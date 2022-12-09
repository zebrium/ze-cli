// Copyright (c) 2022 ScienceLogic Inc

package up

import "github.com/zebrium/ze-cli/ze/batch"

func UploadFile(url string, auth string, file string, logtype string, host string, svcgrp string, dtz string, ids string, cfgs string, tags string, batchId string, disableBatch bool, logstash bool) (err error) {
	//Generate a new batch upload
	if !disableBatch && batchId == "" {
		batch, err := batch.Begin(url, auth, "")
		if err != nil {
			return err
		}
		batchId = batch.Data.BatchId
	}

	getFileToken()
	sendFile()

	//Close batch
	if !disableBatch {
		_, err = batch.End(url, auth, batchId)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileToken() {

}

func sendFile() {

}
