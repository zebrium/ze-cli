// Package batch Copyright © 2023 ScienceLogic Inc

package batch

type beginBatch struct {
	ProcessingMethod string `json:"processing_method"`
	RetentionHours   int    `json:"retention_hours"`
	BatchId          string `json:"batch_id,omitempty" `
}

type BeginBatchResp struct {
	Data    *batchBeginDataResp `json:"data"`
	Message string              `json:"message"`
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
}
type batchBeginDataResp struct {
	BatchId string `json:"batch_id"`
}

type EndBatchResp struct {
	Data    *endBatchData `json:"data"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Status  string        `json:"status"`
}

type endBatchData struct {
	BatchID string `json:"batch_id"`
	State   string `json:"state"`
}

type ShowBatchResp struct {
	Data    []showBatchData `json:"data"`
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Status  string          `json:"status"`
}

type showBatchData struct {
	BatchID            string `json:"batch_id"`
	Account            string `json:"account"`
	State              string `json:"state"`
	Lines              int    `json:"lines"`
	Bundles            int    `json:"bundles"`
	BundlesCompleted   int    `json:"bundles_completed"`
	Created            string `json:"created"`
	UploadTimeSecs     int    `json:"upload_time_secs"`
	ProcessingTimeSecs int    `json:"processing_time_secs"`
	CompletionTime     string `json:"completion_time"`
	RetentionHours     int    `json:"retention_hours"`
	ExpirationTime     string `json:"expiration_time"`
	ProcessingMethod   string `json:"processing_method"`
	Reason             string `json:"reason"`
}

type CancelBatchResp struct {
	Data    *cancelBatchData `json:"data"`
	Message string           `json:"message"`
	Code    int              `json:"code"`
	Status  string           `json:"status"`
}

type cancelBatchData struct {
	BatchID string `json:"batch_id"`
	State   string `json:"state"`
}
