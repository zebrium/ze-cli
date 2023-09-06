package batch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var auth = "123456"
var batchId = "batch123"

func TestBatchBegin(t *testing.T) {
	testCases := []struct {
		batchResponse BeginBatchResp
		httpResponse  int
		expectedFail  bool
	}{
		{
			batchResponse: BeginBatchResp{
				Data:    &BatchBeginDataResp{BatchId: batchId},
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: false,
		},
		{
			batchResponse: BeginBatchResp{
				Data:    nil,
				Message: "",
				Code:    502,
				Status:  "502 bad gateway",
			},
			httpResponse: http.StatusBadGateway,
			expectedFail: true,
		},
		{
			batchResponse: BeginBatchResp{
				Data:    nil,
				Message: "",
				Code:    502,
				Status:  "502 bad gateway",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
		{
			batchResponse: BeginBatchResp{
				Data:    nil,
				Message: "",
				Code:    200,
				Status:  "Success",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
		{
			batchResponse: BeginBatchResp{
				Data:    nil,
				Message: "",
				Code:    403,
				Status:  "Unable to create record",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
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
			w.WriteHeader(tc.httpResponse)
			_, err := w.Write(byteResponse)
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}

		}))
		value, err := Begin(server.URL, auth, "")
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
		if value.Code != 200 {
			t.Fatalf("expected response code of 200, instead got: %d in tc: %d ", value.Code, i)
		}
		if value.Data == nil {
			t.Fatalf("expected value.Data to be not nil but was.  tc: %d", i)
		}
		if value.Data.BatchId != batchId {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data.BatchId, i)
		}
		server.Close()
	}
}

func TestCancelBatch(t *testing.T) {

	testCases := []struct {
		batchResponse CancelBatchResp
		httpResponse  int
		expectedFail  bool
	}{
		{
			batchResponse: CancelBatchResp{
				Data: &cancelBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: false,
		},
		{
			batchResponse: CancelBatchResp{
				Data: &cancelBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    502,
				Status:  "Success",
			},
			httpResponse: http.StatusBadGateway,
			expectedFail: true,
		},
		{
			batchResponse: CancelBatchResp{
				Data: &cancelBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    502,
				Status:  "Success",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
		{
			batchResponse: CancelBatchResp{
				Data:    nil,
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
	}

	for i, tc := range testCases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != fmt.Sprintf("/%s/%s", batchURL, batchId) {
				t.Errorf("Expected to request %s, got: %s tc: %d", batchURL, r.URL.Path, i)
			}
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json header, got: %s tc: %d", r.Header.Get("Accept"), i)
			}
			if r.Header.Get("Authorization") != fmt.Sprintf("Token %s", auth) {
				t.Errorf("Expected Authorization: %s, got: %s tc: %d", fmt.Sprintf("Token %s", auth), r.Header.Get("Authorization"), i)
			}

			byteResponse, _ := json.Marshal(tc.batchResponse)
			w.WriteHeader(tc.httpResponse)
			_, err := w.Write(byteResponse)
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}

		}))
		value, err := Cancel(server.URL, auth, batchId)
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
		if value.Code != 200 {
			t.Fatalf("expected response code of 200, instead got: %d in tc: %d ", value.Code, i)
		}
		if value.Data == nil {
			t.Fatalf("expected value.Data to be not nil but was.  tc: %d", i)
		}
		if value.Data.BatchID != batchId {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data.BatchID, i)
		}
		if value.Data.State != tc.batchResponse.Data.State {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data.State, i)
		}
		server.Close()
	}
}

func TestEndBatch(t *testing.T) {

	testCases := []struct {
		batchResponse EndBatchResp
		httpResponse  int
		expectedFail  bool
	}{
		{
			batchResponse: EndBatchResp{
				Data: &endBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: false,
		},
		{
			batchResponse: EndBatchResp{
				Data: &endBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    502,
				Status:  "Success",
			},
			httpResponse: http.StatusBadGateway,
			expectedFail: true,
		},
		{
			batchResponse: EndBatchResp{
				Data: &endBatchData{
					BatchID: batchId,
					State:   "Canceled",
				},
				Message: "",
				Code:    502,
				Status:  "Success",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
		{
			batchResponse: EndBatchResp{
				Data:    nil,
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
	}

	for i, tc := range testCases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != fmt.Sprintf("/%s/%s", batchURL, batchId) {
				t.Errorf("Expected to request %s, got: %s tc: %d", batchURL, r.URL.Path, i)
			}
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json header, got: %s tc: %d", r.Header.Get("Accept"), i)
			}
			if r.Header.Get("Authorization") != fmt.Sprintf("Token %s", auth) {
				t.Errorf("Expected Authorization: %s, got: %s tc: %d", fmt.Sprintf("Token %s", auth), r.Header.Get("Authorization"), i)
			}

			byteResponse, _ := json.Marshal(tc.batchResponse)
			w.WriteHeader(tc.httpResponse)
			_, err := w.Write(byteResponse)
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}

		}))
		value, err := End(server.URL, auth, batchId)
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
		if value.Code != 200 {
			t.Fatalf("expected response code of 200, instead got: %d in tc: %d ", value.Code, i)
		}
		if value.Data == nil {
			t.Fatalf("expected value.Data to be not nil but was.  tc: %d", i)
		}
		if value.Data.BatchID != batchId {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data.BatchID, i)
		}
		if value.Data.State != tc.batchResponse.Data.State {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data.State, i)
		}
		server.Close()
	}
}

func TestShowBatch(t *testing.T) {

	testCases := []struct {
		batchResponse ShowBatchResp
		httpResponse  int
		expectedFail  bool
	}{
		{
			batchResponse: ShowBatchResp{
				Data: []ShowBatchData{
					{
						BatchID:            batchId,
						Account:            "",
						State:              "",
						Lines:              0,
						Bundles:            0,
						BundlesCompleted:   0,
						Created:            "",
						UploadTimeSecs:     0,
						ProcessingTimeSecs: 0,
						CompletionTime:     "",
						RetentionHours:     0,
						ExpirationTime:     "",
						ProcessingMethod:   "",
						Reason:             "",
					},
				},
				Message: "",
				Code:    200,
				Status:  "Success",
			},
			httpResponse: http.StatusOK,
			expectedFail: false,
		},
		{
			batchResponse: ShowBatchResp{
				Data: []ShowBatchData{
					{
						BatchID:            batchId,
						Account:            "",
						State:              "",
						Lines:              0,
						Bundles:            0,
						BundlesCompleted:   0,
						Created:            "",
						UploadTimeSecs:     0,
						ProcessingTimeSecs: 0,
						CompletionTime:     "",
						RetentionHours:     0,
						ExpirationTime:     "",
						ProcessingMethod:   "",
						Reason:             "",
					},
				},
				Message: "",
				Code:    500,
				Status:  "Success",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
		{
			batchResponse: ShowBatchResp{
				Data: []ShowBatchData{
					{
						BatchID:            batchId,
						Account:            "",
						State:              "",
						Lines:              0,
						Bundles:            0,
						BundlesCompleted:   0,
						Created:            "",
						UploadTimeSecs:     0,
						ProcessingTimeSecs: 0,
						CompletionTime:     "",
						RetentionHours:     0,
						ExpirationTime:     "",
						ProcessingMethod:   "",
						Reason:             "",
					},
				},
				Message: "",
				Code:    502,
				Status:  "Success",
			},
			httpResponse: http.StatusBadGateway,
			expectedFail: true,
		},
		{
			batchResponse: ShowBatchResp{
				Data:    nil,
				Message: "",
				Code:    200,
				Status:  "200",
			},
			httpResponse: http.StatusOK,
			expectedFail: true,
		},
	}

	for i, tc := range testCases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != fmt.Sprintf("/%s/%s", batchURL, batchId) {
				t.Errorf("Expected to request %s, got: %s tc: %d", batchURL, r.URL.Path, i)
			}
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json header, got: %s tc: %d", r.Header.Get("Accept"), i)
			}
			if r.Header.Get("Authorization") != fmt.Sprintf("Token %s", auth) {
				t.Errorf("Expected Authorization: %s, got: %s tc: %d", fmt.Sprintf("Token %s", auth), r.Header.Get("Authorization"), i)
			}

			byteResponse, _ := json.Marshal(tc.batchResponse)
			w.WriteHeader(tc.httpResponse)
			_, err := w.Write(byteResponse)
			if err != nil {
				t.Fatalf("encountered error when one wasnt expected in tc: %d, Error: %v", i, err)
			}

		}))
		value, err := Show(server.URL, auth, batchId)
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
		if value.Code != 200 {
			t.Fatalf("expected response code of 200, instead got: %d in tc: %d ", value.Code, i)
		}
		if value.Data == nil {
			t.Fatalf("expected value.Data to be not nil but was.  tc: %d", i)
		}
		if value.Data[0].BatchID != batchId {
			t.Fatalf("Expected %s but instead got: %s for tc: %d", batchId, value.Data[0].BatchID, i)
		}

		server.Close()
	}
}

func TestFormatTime(t *testing.T) {
	testcases := []struct {
		input    int
		expected string
	}{
		{
			input:    416446,
			expected: "115h40m46s",
		},
		{
			input:    60,
			expected: "1m0s",
		},
		{
			input:    3600,
			expected: "1h0m0s",
		},
	}

	for i, tc := range testcases {
		act, err := formatTime(tc.input)
		if err != nil {
			t.Fatalf("Unexpected error for tc: %d, error: %v", i, err)
		}
		if act != tc.expected {
			t.Fatalf("expected: %s, actual: %s  tc: %d", tc.expected, act, i)
		}
	}

}
