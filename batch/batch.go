// Package batch Copyright © 2023 ScienceLogic Inc

package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var batchURL = "log/api/v2/batch"

// End Ends a batchId so that it can be updated to begin processing
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error processing request, http code: %d, http status: %s", resp.StatusCode, resp.Status)
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
	if respMap.Code != 200 {
		return nil, fmt.Errorf("batch end failed with error code: %d. error message: %s", respMap.Code, respMap.Message)
	}
	if respMap.Data == nil {
		return nil, fmt.Errorf("error processing the response from the server, response: %v", respMap.Data)
	}
	return respMap, nil
}

// Begin Begins a batchId
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error processing request, http code: %d, http status: %s", resp.StatusCode, resp.Status)
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
	if respMap.Code != 200 {
		return nil, fmt.Errorf("batch begin failed with error code: %d, error message: %s, error status: %s", respMap.Code, respMap.Message, respMap.Status)
	}
	if respMap.Data == nil || len(respMap.Data.BatchId) == 0 {
		return nil, fmt.Errorf("batch begin returned an invalid batch id.  response from server: %v", respMap.Data)
	}
	return respMap, nil

}

// Show Gets the current status of a batch ID
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error processing request, http code: %d, http status: %s", resp.StatusCode, resp.Status)
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
	if respMap.Code != 200 {
		return nil, fmt.Errorf("batch show failed with error code: %d. error message: %s", respMap.Code, respMap.Message)
	}
	if respMap.Data == nil {
		return nil, fmt.Errorf("error processing the response from the server, response: %v", respMap.Data)
	}
	return respMap, nil
}

// Cancel Cancels a batch Id
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error processing request, http code: %d, http status: %s", resp.StatusCode, resp.Status)
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
	if respMap.Code != 200 {
		return nil, fmt.Errorf("batch cancel failed with error code: %d. error message: %s", respMap.Code, respMap.Message)
	}
	if respMap.Data == nil {
		return nil, fmt.Errorf("error processing the response from the server, response: %v", respMap.Data)
	}
	return respMap, nil
}

// String Overrides the toString func for showBatchData to pretty print data for user consumption
func (d ShowBatchData) String() (string, error) {
	fmtUploadTime, err := formatTime(d.UploadTimeSecs)
	if err != nil {
		return "", err
	}
	fmtProcessingTime, err := formatTime(d.ProcessingTimeSecs)
	if err != nil {
		return "", err
	}
	response := fmt.Sprintf(
		"         Batch ID: %s\n"+
			"            State: %s\n", d.BatchID, d.State)
	if len(d.Reason) > 0 {
		response = response + fmt.Sprintf("   Failure Reason: %s\n", d.Reason)
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
		d.Created, d.CompletionTime, d.ExpirationTime, d.RetentionHours, d.Lines, d.Bundles, d.BundlesCompleted, fmtUploadTime, fmtProcessingTime), nil
}

// formatTime Formats time objects into a duration for pretty printing
func formatTime(duration int) (string, error) {
	seconds, err := time.ParseDuration(fmt.Sprintf("%ds", duration))
	if err != nil {
		return "", err
	}
	return seconds.String(), nil
}
