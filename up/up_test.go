// Package up Copyright © 2024 ScienceLogic Inc

package up

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zebrium/ze-cli/batch"
)

var url = "https://test.local"
var auth = "xxxxxxxxxxxxxxxxxxxxxxx"
var batchURL = "log/api/v2/batch"
var testFileLocation = "../testfiles/test.log"

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
				t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
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
				t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
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

func TestGenerateMetadataSVG(t *testing.T) {
	testCases := []struct {
		ids      string
		svggrp   string
		expected string
	}{
		{
			ids:      "",
			svggrp:   "",
			expected: "default",
		},
		{
			ids:      "",
			svggrp:   "test123",
			expected: "test123",
		},
		{
			ids:      "ze_deployment_name=test123",
			svggrp:   "",
			expected: "test123",
		},
		{
			ids:      "ze_deployment_name=test123",
			svggrp:   "test1234",
			expected: "test1234",
		},
	}
	for i, tc := range testCases {
		metadata, _, _, err := generateMetadata(url, auth, "", "", "", tc.svggrp, "", tc.ids, "", "", "", false, "")
		if err != nil {
			t.Fatalf("test case %d generated unexpected error: %v", i, err)
		}
		if metadata.Ids["ze_deployment_name"] != tc.expected {
			t.Fatalf("test case %d failed to match expected result expected: %s  actual: %s", i, tc.expected, metadata.Ids["ze_deployment_name"])
		}
	}
}

func TestUpErrorHandling(t *testing.T) {
	host := "testMachine"
	svg := "test"
	dtz := ""
	ids := ""
	tags := ""
	testCases := []struct {
		file                string //Test files can be found in ../testfiles
		logtype             string
		cfgs                string
		batchId             string //Existing Batch ID to use instead of generating one
		generatedBatchId    string //Autogenerated Batch Id to use when no batch passed in
		logstash            bool
		disableBatch        bool
		uploadTokenResponse TokenResp
		uploadTokenStatus   int
		uploadPostStatus    int
		expectedBatchCancel bool
		expectedFail        bool
		existingBatch       bool
		batchEndStatus      int
		expectingBatchEnd   bool
		batchCancelStatus   int
	}{
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: false,
			expectedFail:        false,
			existingBatch:       true,
			batchEndStatus:      http.StatusOK,
			expectingBatchEnd:   false,
		},
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "batch-1234",
			logstash:            false,
			disableBatch:        true,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: false,
			expectedFail:        false,
			existingBatch:       true,
			batchEndStatus:      http.StatusOK,
			expectingBatchEnd:   false,
		},
		{
			file:                "",
			logtype:             "",
			cfgs:                "",
			batchId:             "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusServiceUnavailable,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: false,
			expectedFail:        true,
			existingBatch:       true,
			batchEndStatus:      http.StatusOK,
			expectingBatchEnd:   false,
		},
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusServiceUnavailable,
			expectedBatchCancel: false,
			expectedFail:        true,
			existingBatch:       true,
			batchEndStatus:      http.StatusOK,
			expectingBatchEnd:   false,
		},
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "",
			generatedBatchId:    "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: false,
			expectedFail:        false,
			existingBatch:       true,
			batchEndStatus:      http.StatusOK,
			expectingBatchEnd:   true,
		},
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "",
			generatedBatchId:    "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: true,
			expectedFail:        true,
			existingBatch:       true,
			batchEndStatus:      http.StatusServiceUnavailable,
			expectingBatchEnd:   true,
			batchCancelStatus:   http.StatusOK,
		},
		{
			file:                testFileLocation,
			logtype:             "",
			cfgs:                "",
			batchId:             "",
			generatedBatchId:    "batch-1234",
			logstash:            false,
			disableBatch:        false,
			uploadTokenResponse: TokenResp{Token: "token1234"},
			uploadTokenStatus:   http.StatusOK,
			uploadPostStatus:    http.StatusOK,
			expectedBatchCancel: true,
			expectedFail:        true,
			existingBatch:       true,
			batchEndStatus:      http.StatusServiceUnavailable,
			expectingBatchEnd:   true,
			batchCancelStatus:   http.StatusServiceUnavailable,
		},
	}

	for i, tc := range testCases {
		server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			tcBatchID := tc.batchId
			if len(tc.batchId) == 0 {
				tcBatchID = tc.generatedBatchId
			}
			switch req.URL.Path {
			// Batch End or Chancel Request
			case fmt.Sprintf("/%s/%s", batchURL, tcBatchID):
				switch req.Method {
				//Generate Batch ID Request
				case http.MethodPost:
					resp.WriteHeader(http.StatusOK)
					batchResponse := batch.BeginBatchResp{
						Message: "Ok", Code: http.StatusOK, Status: "Ok", Data: &batch.BatchBeginDataResp{BatchId: tc.generatedBatchId},
					}
					byteResponse, _ := json.Marshal(batchResponse)
					_, err := resp.Write(byteResponse)
					if err != nil {
						t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
					}
				// Batch End or Batch Cancel
				case http.MethodPut:
					reqBody, err := io.ReadAll(req.Body)
					if err != nil {
						t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
						t.Fail()
					}
					if strings.Contains(string(reqBody), "cancel") {
						if !tc.expectedBatchCancel {
							t.Fatalf("called batch cancel when it was not supposed too")
							t.Fail()
						}
						resp.WriteHeader(http.StatusOK)
						batchResponse := batch.CancelBatchResp{
							Code: tc.batchCancelStatus, Status: http.StatusText(tc.batchCancelStatus), Message: "", Data: &batch.CancelBatchData{BatchID: tcBatchID, State: "cancelled"},
						}
						byteResponse, _ := json.Marshal(batchResponse)
						_, err := resp.Write(byteResponse)
						if err != nil {
							t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
						}

					} else {
						if !tc.expectingBatchEnd {
							t.Fatalf("called batch end when it was not supposed too")
							t.Fail()
						} else {
							resp.WriteHeader(tc.batchEndStatus)
							batchResponse := batch.EndBatchResp{
								Message: http.StatusText(tc.batchEndStatus), Code: tc.batchEndStatus, Data: &batch.EndBatchData{BatchID: tcBatchID, State: "Processing"},
							}
							byteResponse, _ := json.Marshal(batchResponse)
							_, err := resp.Write(byteResponse)
							if err != nil {
								t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
							}
						}
					}
				default:
					t.Fatalf("unsupported method %s called", req.Method)
					t.Fail()
				}
			// Batch Generation
			case fmt.Sprintf("/%s", batchURL):
				t.Log("Called generic Batch URL")
				switch req.Method {
				case http.MethodPost:
					resp.WriteHeader(http.StatusOK)
					batchResponse := batch.BeginBatchResp{
						Message: "Ok", Code: http.StatusOK, Status: "Ok", Data: &batch.BatchBeginDataResp{BatchId: tc.generatedBatchId},
					}
					byteResponse, _ := json.Marshal(batchResponse)
					_, err := resp.Write(byteResponse)
					if err != nil {
						t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
					}
				default:
					t.Fatalf("unsupported method %s called", req.Method)
					t.Fail()
				}
			// Call Token Request
			case fmt.Sprintf("/%s", tokenPath):
				byteResponse, _ := json.Marshal(tc.uploadTokenResponse)
				resp.WriteHeader(tc.uploadTokenStatus)
				_, err := resp.Write(byteResponse)
				if err != nil {
					t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
					t.Fail()
				}

			//Upload Logstash
			case fmt.Sprintf("/%s", logstashPath):
				t.Log("called logstash post")
				resp.WriteHeader(tc.uploadPostStatus)
			//Upload Post
			case fmt.Sprintf("/%s", postPath):
				resp.WriteHeader(tc.uploadPostStatus)
			default:
				t.Fatalf("called unimplemented method %s", req.RequestURI)
				t.Fatal()
			}
		}))
		err := UploadFile(server.URL, auth, tc.file, tc.logtype, host, svg, dtz, ids, tc.cfgs, tags, tc.batchId, tc.disableBatch, tc.logstash, "test")
		if tc.expectedFail {
			if err == nil {
				t.Fatalf("failed to encounter error when one was expected in tc: %d, Error: %v", i, err)
				t.Fatal()
			} else {
				server.Close()
				continue
			}
		} else {
			if err != nil {
				t.Fatalf("encountered error when one wasn't expected in tc: %d, Error: %v", i, err)
				t.Fail()
			}
		}
		server.Close()
	}
}
