// Package up Copyright © 2023 ScienceLogic Inc

package up

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/batch"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tokenPath = "log/api/v2/token"
var logstashPath = "log/api/v2/ingest?log_source=logstash&log_format=json_batch"
var postPath = "log/api/v2/post"

// UploadFile Main processing function for uploading a file to Zebrium's backend.
func UploadFile(url string, auth string, file string, logtype string, host string, svcgrp string,
	dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool, logstash bool,
	version string) (err error) {

	defer func() {
		if err != nil {
			cleanUpBatchOnExit(batchId)
		}
	}()

	route := postPath
	tr := &http.Transport{
		MaxIdleConns:    10,
		WriteBufferSize: 32768,
	}
	client := &http.Client{Transport: tr}
	metadata, existingBatch, updatedBatchId, err := generateMetadata(url, auth, file, logtype, host, svcgrp, dtz, ids, cfgs, tags, batchId, disableBatch, version)
	if err != nil {
		return err
	}
	body, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	// request stream token
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, tokenPath), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	tokenResp := &TokenResp{}
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}
	uploadAuth := tokenResp.Token
	if logstash {
		route = logstashPath
		uploadAuth = auth
	}
	req, err = sendRequestBuilder(url, route, file)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", uploadAuth))
	req.Close = true
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	// Close batch
	if !disableBatch && !existingBatch {
		_, err = batch.End(url, auth, updatedBatchId)
		if err != nil {
			return err
		}
	}
	return nil
}

// sendRequestBuilder Builds the request for sending log files.
// Will switch between Readers based on if a file is present or if Stdin is the intended target
func sendRequestBuilder(url string, route string, filename string) (*http.Request, error) {
	if len(filename) != 0 {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, route), file)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/octet-stream")
		return req, nil
	} else {
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, route), os.Stdin)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/octet-stream")
		return req, nil

	}
}

// createMap Helper function for creating mapping passed in metadata in a key=value csv form to a string map
func createMap(in string) (result map[string]string) {
	m := make(map[string]string)
	if len(in) != 0 {
		args := strings.Split(in, ",")
		for _, e := range args {
			kvp := strings.Split(e, "=")
			m[kvp[0]] = kvp[1]
		}
	}
	return m
}

// cleanUpBatchOnExit Helper function to clean up autogenerate batchId's on error of exit
func cleanUpBatchOnExit(batchId string) {
	if len(batchId) != 0 {
		_, err := batch.Cancel(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Unable to cleanly exit on Error.  Will need to clean up Batch ID: %s\n", batchId)
	}
}

// parseBatchIdFromConfigs Helper function that will retrieve a batchId from a config map
func parseBatchIdFromConfigs(cfgs string) (batchId string) {
	cfgMap := createMap(cfgs)
	val, ok := cfgMap["ze_batch_id"]
	if ok {
		return val
	}
	return ""
}

// generateMetadata Helper function that generates the metadata struct that is needed to request a stream token
func generateMetadata(url string, auth string, file string, logtype string, host string, svcgrp string,
	dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool,
	version string) (metadata MetaData, existingBatch bool, updatedBatchId string, err error) {
	existingBatch = true
	updatedBatchId = batchId
	metadata = MetaData{
		Stream:             "native",
		ZeLogCollectorVers: version,
		TM:                 false,
	}

	if len(file) != 0 {
		metadata.Stream = "zefile"
		if len(logtype) == 0 {
			fileName := filepath.Base(file)
			metadata.LogBaseName = strings.ToLower(strings.Split(strings.TrimSpace(fileName), ".")[0])
		}
	} else {
		disableBatch = true
	}

	if len(cfgs) != 0 {
		cfgBatchId := parseBatchIdFromConfigs(cfgs)
		if len(cfgBatchId) != 0 {
			updatedBatchId = cfgBatchId
		}
	}
	metadata.Tz = dtz
	metadata.Ids = createMap(ids)
	metadata.Cfgs = createMap(cfgs)
	metadata.Tags = createMap(tags)
	if len(logtype) != 0 {
		metadata.LogBaseName = logtype
	}
	if len(host) != 0 {
		metadata.Ids["zid_host"] = host
	}

	if len(svcgrp) != 0 {
		metadata.Ids["ze_deployment_name"] = svcgrp
	}

	if len(metadata.Ids["ze_deployment_name"]) == 0 {
		metadata.Ids["ze_deployment_name"] = "default"
	}
	// Generate a new batch upload if needed
	if !disableBatch && len(updatedBatchId) == 0 && !strings.Contains(cfgs, "ze_batch_id") {
		batchResp, err := batch.Begin(url, auth, "")
		if err != nil {
			return metadata, false, updatedBatchId, err
		}
		existingBatch = false
		updatedBatchId = batchResp.Data.BatchId
	}
	// Add new batch id to configs
	if len(updatedBatchId) != 0 && !disableBatch && !strings.Contains(cfgs, "ze_batch_id") {
		metadata.Cfgs["ze_batch_id"] = updatedBatchId
	}
	return metadata, existingBatch, updatedBatchId, nil
}
