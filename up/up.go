// Package up Copyright © 2023 ScienceLogic Inc/*

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
	"strings"
)

var tokenPath = "log/api/v2/token"
var logstashPath = "log/api/v2/ingest?log_source=logstash&log_format=json_batch"
var postPath = "log/api/v2/post"

func UploadFile(url string, auth string, file string, logtype string, host string, svcgrp string,
	dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool, logstash bool,
	version string) (err error) {
	route := postPath
	client := &http.Client{}

	metadata, existingBatch, err := generateMetadata(url, auth, file, logtype, host, svcgrp, dtz, ids, cfgs, tags, batchId, disableBatch, version)
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}
	body, err := json.Marshal(metadata)
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}
	//Request Stream Token
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, tokenPath), bytes.NewBuffer(body))
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		cleanUpBatchOnExit(batchId)
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
		cleanUpBatchOnExit(batchId)
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", uploadAuth))
	resp, err = client.Do(req)
	if err != nil {
		cleanUpBatchOnExit(batchId)
		return err
	}

	//Close batch
	if !disableBatch && !existingBatch {
		_, err = batch.End(url, auth, batchId)
		if err != nil {
			return err
		}
	}
	return nil
}

func sendRequestBuilder(url string, route string, filename string) (*http.Request, error) {
	if filename != "" {
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

func createMap(in string) (result map[string]string) {
	m := make(map[string]string)
	if len(in) > 0 {
		args := strings.Split(in, ",")
		for _, e := range args {
			kvp := strings.Split(e, "=")
			m[kvp[0]] = kvp[1]
		}
	}
	return m
}

func cleanUpBatchOnExit(batchId string) {
	if batchId != "" {
		_, err := batch.Cancel(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Unable to cleanly exit on Error.  Will need to clean up Batch ID: %s\n", batchId)
	}
}

func generateMetadata(url string, auth string, file string, logtype string, host string, svcgrp string,
	dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool,
	version string) (metadata MetaData, existingBatch bool, err error) {
	existingBatch = true
	metadata = MetaData{
		Stream:             "native",
		ZeLogCollectorVers: version,
		TM:                 false,
	}

	if file != "" {
		metadata.Stream = "zefile"
		if logtype == "" {
			metadata.LogBaseName = strings.ToLower(strings.Split(strings.TrimSpace(file), ".")[0])
		}
	} else {
		disableBatch = true
	}
	metadata.Tz = dtz
	metadata.Ids = createMap(ids)
	metadata.Cfgs = createMap(cfgs)
	metadata.Tags = createMap(tags)
	if logtype != "" {
		metadata.LogBaseName = logtype
	}
	if host != "" {
		metadata.Ids["zid_host"] = host
	}
	if svcgrp != "" {
		metadata.Ids["ze_deployment_name"] = svcgrp
	}
	//Generate a new batch upload if needed
	if !disableBatch && batchId == "" && !strings.Contains(cfgs, "ze_batch_id") {
		batchResp, err := batch.Begin(url, auth, "")
		if err != nil {
			return metadata, false, err
		}
		existingBatch = false
		batchId = batchResp.Data.BatchId
	}
	// Add new batch id to configs
	if batchId != "" && !disableBatch && !strings.Contains(cfgs, "ze_batch_id") {
		metadata.Cfgs["ze_batch_id"] = batchId
	}
	return metadata, existingBatch, nil
}
