// Copyright (c) 2022 ScienceLogic Inc

package up

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/ze/batch"
	"io"
	"net/http"
	"strings"
)

var tokenPath = "/log/api/v2/token"
var logstashPath = "/log/api/v2/ingest?log_source=logstash&log_format=json_batch"
var postPath = "/log/api/v2/post"

func UploadFile(url string, auth string, file string, logtype string, host string, svcgrp string,
	dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool, logstash bool,
	version string) (err error) {
	client := &http.Client{}
	metadata := MetaData{
		Stream:             "native",
		ZeLogCollectorVers: version,
		TM:                 false,
	}
	if file != "" {
		metadata.Stream = "zefile"
		if logtype == "" {
			if strings.Contains(file, "(?=[^\\/]++$)([a-zA-Z]{3,})") {
				metadata.LogBaseName = strings.ToLower(file)
			} else {
				metadata.LogBaseName = "stream"
			}
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
			return err
		}
		batchId = batchResp.Data.BatchId
	}
	// Add new batch id to configs
	if batchId != "" && !disableBatch {
		metadata.Cfgs["ze_batch_id"] = batchId
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

	//Send Data using an IO pipe https://stackoverflow.com/questions/42261421/send-os-stdin-via-http-post-without-loading-file-into-memory

	//Close batch
	if !disableBatch {
		_, err = batch.End(url, auth, batchId)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendFile() {

}

func createMap(in string) (result map[string]string) {
	m := make(map[string]string)
	args := strings.Split(in, ",")
	for _, e := range args {
		kvp := strings.Split(e, "=")
		m[kvp[0]] = kvp[1]
	}
	return m
}

func cleanUpBatchOnExit(batchId string) {
	if batchId != "" {
		batch.Cancel(viper.GetString("url"), viper.GetString("auth"), batchId)
		fmt.Sprintf("Unable to cleanly exit on Error.  Will need to clean up Batch ID: %s\n", batchId)
	}
}
