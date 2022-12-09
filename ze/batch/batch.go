// Copyright (c) 2023 ScienceLogic Inc

package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var batchURL = "/log/api/v2/batch"

func End(url string, auth string, batchId string) (response *EndBatchResp, err error) {
	client := &http.Client{}
	var jsonStr = []byte(`{"uploads_complete": true}`)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", url, batchURL, batchId), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respMap := &EndBatchResp{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, err
	}
	return respMap, nil
}

func Begin(url string, auth string, batchId string) (response *BeginBatchResp, err error) {
	client := &http.Client{}
	beginBatch := beginBatch{ProcessingMethod: "opportunistic", RetentionHours: 48, BatchId: batchId}
	body, _ := json.Marshal(beginBatch)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, batchURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respMap := &BeginBatchResp{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, err
	}
	return respMap, nil

}
func Show(url string, auth string, batchId string) (response *ShowBatchResp, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", url, batchURL, batchId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respMap := &ShowBatchResp{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, err
	}
	return respMap, nil
}
func Cancel(url string, auth string, batchId string) (response *CancelBatchResp, err error) {
	client := &http.Client{}
	var jsonStr = []byte(`{"cancel": true}`)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", url, batchURL, batchId), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", auth))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respMap := &CancelBatchResp{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, err
	}
	return respMap, nil
}

func (this showBatchData) String() string {
	fmtUploadTime := formatTime(this.UploadTimeSecs)
	fmtProcessingTime := formatTime(this.ProcessingTimeSecs)
	response := fmt.Sprintf(
		"         Batch ID: %s\n"+
			"            State: %s\n", this.BatchID, this.State)
	if len(this.Reason) > 0 {
		response = response + fmt.Sprintf("   Failure Reason: %s\n", this.Reason)
	}
	return response + fmt.Sprintf(
		"          Created: %s\n"+
			"  Completion Time: %s\n"+
			"  Expiration Time: %s\n"+
			"Retention (hours): %d\n"+
			"            Lines: %d\n"+
			"  Bundles Created: %d\n"+
			"Bundles Completed: %d\n"+
			"      Upload time: %s\n"+
			"  Processing time: %s\n",
		this.Created, this.CompletionTime, this.ExpirationTime, this.RetentionHours, this.Lines, this.Bundles, this.BundlesCompleted, fmtUploadTime, fmtProcessingTime)
}

func formatTime(duration int) string {
	seconds, err := time.ParseDuration(fmt.Sprintf("%ds", duration))
	if err != nil {
		log.Fatal(err.Error())
	}
	return seconds.String()
}
func ValidateId(batchId string) {

}
