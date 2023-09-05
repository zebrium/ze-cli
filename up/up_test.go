package up

import (
	"encoding/json"
	"fmt"
	"github.com/zebrium/ze-cli/batch"
	"net/http"
	"net/http/httptest"
	"testing"
)

var url = "https://test.local"
var auth = "xxxxxxxxxxxxxxxxxxxxxxx"
var batchURL = "log/api/v2/batch"

func TestCreateMap(t *testing.T) {
	testCases := []struct {
		set string
		len int
	}{
		{
			set: "",
			len: 0,
		},
		{
			set: "key1=value1",
			len: 1,
		},
		{
			set: "key1=value1,key2=value2,key3=value3",
			len: 3,
		},
	}
	for i, tc := range testCases {
		actual := createMap(tc.set)
		if len(actual) != tc.len {
			t.Fatalf("Subtest: %d failed. Incorrect number of objects created in map.  Expected: %d, Got: %d", i, tc.len, len(actual))
		}
	}
}

func TestMetadataBatchPassed(t *testing.T) {
	batchId := "test123456"
	version := "1.0.0"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "test.log", "", "", "", "", "", "", "", batchId, false, version)
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to false when it should of been true")
	}
	if metadata.Cfgs["ze_batch_id"] != batchId {
		t.Fatalf("Batchids did not match what was expected. Expected: %s, Actual %s", batchId, metadata.Cfgs["ze_batch_id"])
	}
	if batchId != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batchId, updatedBatch)
	}

}
func TestMetadataBatchConfig(t *testing.T) {
	batchId := "test123456"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "test.log", "", "",
		"", "", "", "ze_batch_id="+batchId, "", "", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to false when it should of been true")
	}
	if metadata.Cfgs["ze_batch_id"] != batchId {
		t.Fatalf("Batchids did not match what was expected. Expected: %s, Actual %s", batchId, metadata.Cfgs["ze_batch_id"])
	}
	if batchId != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batchId, updatedBatch)
	}
}

func TestMetadataFileWithNoLogType(t *testing.T) {
	batchId := "test123456"
	filename := "test_one-1234.log"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, filename, "", "", "",
		"", "", "", "", batchId, false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.LogBaseName != "test_one-1234" {
		t.Fatalf("LogBaseName incorrectly set.  Expected: %s, Actual: %s", "test_one", metadata.LogBaseName)
	}
	if batchId != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batchId, updatedBatch)
	}
	metadata, existingBatch, _, err = generateMetadata(url, auth, "test_one-1234", "", "",
		"", "", "", "", "", batchId, false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.LogBaseName != "test_one-1234" {
		t.Fatalf("LogBaseName incorrectly set.  Expected: %s, Actual: %s", "test_one", metadata.LogBaseName)
	}
	if metadata.Stream != "zefile" {
		t.Fatalf("Incorrect setting of Stream.  For  files, it should be zefile. Actual: %s", metadata.Stream)

	}
}

func TestMetadataStreaming(t *testing.T) {
	logtype := "peanutbutter"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, "", "",
		"", "", "", "", "", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to True when it should of been False")
	}
	if len(metadata.Cfgs["ze_batch_id"]) != 0 {
		t.Fatalf("BatchId was set even though streaming was enabled.")
	}
	if metadata.LogBaseName != logtype {
		t.Fatalf("Incorrect Log Type Set.  Expected: %s Actual: %s", logtype, metadata.LogBaseName)
	}
	if metadata.Stream != "native" {
		t.Fatalf("Incorrect setting of Stream.  For non files, it should be native. Actual: %s", metadata.Stream)

	}
	if updatedBatch != "" {
		t.Fatalf("Got %s for batchId when should of been empty", updatedBatch)
	}
}

func TestMetadataGeneral(t *testing.T) {
	logtype := "peanutbutter"
	host := "countlogula"
	svcgrp := "jelly"
	tz := "EST"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, host, svcgrp, tz,
		"", "", "", "123", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.Tz != tz {
		t.Fatalf("Incorrect setting of TZ.  Expected: %s Actual: %s", tz, metadata.Tz)

	}
	if metadata.TM != false {
		t.Fatalf("TM is set to true and should always be set to false")

	}
	if metadata.Ids["zid_host"] != host {
		t.Fatalf("Expected: %s Actual: %s", host, metadata.Ids["zid_host"])

	}
	if metadata.Ids["ze_deployment_name"] != svcgrp {
		t.Fatalf("Expected: %s Actual: %s", host, metadata.Ids["ze_deployment_name"])

	}
	if updatedBatch != "123" {
		t.Fatalf("Expected %s Actual %s", "123", updatedBatch)
	}
}

func TestMetadataFilename(t *testing.T) {
	testCases := []struct {
		filename    string
		expectedLBN string
	}{
		{
			filename:    "test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "../../test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "/this/is/a/test/test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "test_one.1.log",
			expectedLBN: "test_one",
		},
		{
			filename:    "test_one",
			expectedLBN: "test_one",
		},
		{
			filename:    "../../test_one",
			expectedLBN: "test_one",
		},
		{
			filename:    "/this/is/a/test/test_one",
			expectedLBN: "test_one",
		},
	}
	for i, tc := range testCases {
		metadata, _, _, err := generateMetadata(url, auth, tc.filename, "", "", "", "", "", "", "", "test", false, "")
		if err != nil {
			t.Fatal(err)
		}
		if metadata.LogBaseName != tc.expectedLBN {
			t.Fatalf("Subtest: %d failed. Expected: %s, Actual: %s", i, tc.expectedLBN, metadata.LogBaseName)
		}
	}
}

func TestParseBatchIdFromConfig(t *testing.T) {
	testCases := []struct {
		cfgs    string
		batchId string
	}{
		{
			cfgs:    "",
			batchId: "",
		},
		{
			cfgs:    "key1=value,key2=value",
			batchId: "",
		},
		{
			cfgs:    "key1=value,key2=value,ze_batch_id=test123",
			batchId: "test123",
		},
		{
			cfgs:    "ze_batch_id=test123",
			batchId: "test123",
		},
	}
	for i, tc := range testCases {
		batchId := parseBatchIdFromConfigs(tc.cfgs)
		if batchId != tc.batchId {
			t.Fatalf("Subtest: %d failed. Expected: %s, Actual: %s", i, tc.batchId, batchId)
		}
	}
}

func TestGenerateMetadataBatching(t *testing.T) {
	testCases := []struct {
		cfgs          string
		batchId       string
		disableBatch  bool
		batchResponse batch.BeginBatchResp
		expectedFail  bool
		existingBatch bool
	}{
		{
			batchResponse: batch.BeginBatchResp{
				Data:    &batch.BatchBeginDataResp{BatchId: "batch123456"},
				Message: "",
				Code:    200,
				Status:  "200",
			},
			disableBatch:  false,
			cfgs:          "",
			expectedFail:  false,
			existingBatch: false,
		},
		{
			batchResponse: batch.BeginBatchResp{
				Data:    &batch.BatchBeginDataResp{BatchId: "batch123456"},
				Message: "",
				Code:    500,
				Status:  "200",
			},
			disableBatch:  false,
			cfgs:          "",
			expectedFail:  true,
			existingBatch: false,
		},
	}
	for i, tc := range testCases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != fmt.Sprintf("/%s", batchURL) {
				t.Errorf("Expected to request %s, got: %s tc: %d", batchURL, r.URL.Path, i)
			}
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json header, got: %s tc: %d", r.Header.Get("Accept"), i)
			}
			if r.Header.Get("Authorization") != fmt.Sprintf("Token %s", auth) {
				t.Errorf("Expected Authorization: %s, got: %s tc: %d", fmt.Sprintf("Token %s", auth), r.Header.Get("Authorization"), i)
			}
			byteResponse, _ := json.Marshal(tc.batchResponse)
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(byteResponse)
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}

		}))
		metadata, existing, updatedBatch, err := generateMetadata(server.URL, auth, "potato", "", "", "", "", "", tc.cfgs, "", tc.batchId, tc.disableBatch, "")

		if tc.expectedFail {
			if err == nil {
				t.Fatalf("failed to encounter error when one was expected in tc: %d, Error: %v", i, err)
			} else {
				server.Close()
				continue
			}
		} else {
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}
		}
		if existing != tc.existingBatch {
			t.Fatalf("expected: %t, actual: %t, tc: %d", tc.existingBatch, existing, i)
		}
		if updatedBatch != tc.batchResponse.Data.BatchId {
			t.Fatalf("expected: %s, actual: %s, tc: %d", tc.batchResponse.Data.BatchId, updatedBatch, i)
		}
		if !tc.existingBatch {
			if metadata.Cfgs["ze_batch_id"] != tc.batchResponse.Data.BatchId {
				t.Fatalf("expected: %s, actual: %s, tc: %d", tc.batchResponse.Data.BatchId, metadata.Cfgs["ze_batch_id"], i)

			}
		}
	}
}
