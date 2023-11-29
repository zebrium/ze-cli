// Package batch Copyright © 2023 ScienceLogic Inc

package batch

type beginBatch struct {
	ProcessingMethod string `json:"processing_method"`
	RetentionHours   int    `json:"retention_hours"`
	BatchId          string `json:"batch_id,omitempty" `
}

type BeginBatchResp struct {
	Data    *BatchBeginDataResp `json:"data"`
	Message string              `json:"message"`
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
}
type BatchBeginDataResp struct {
	BatchId string `json:"batch_id"`
}

type EndBatchResp struct {
	Data    *EndBatchData `json:"data"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Status  string        `json:"status"`
}

type EndBatchData struct {
	BatchID string `json:"batch_id"`
	State   string `json:"state"`
}

type ShowBatchResp struct {
	Data    []ShowBatchData `json:"data"`
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Status  string          `json:"status"`
}

type ShowBatchData struct {
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
	Data    *CancelBatchData `json:"data"`
	Message string           `json:"message"`
	Code    int              `json:"code"`
	Status  string           `json:"status"`
}

type CancelBatchData struct {
	BatchID string `json:"batch_id"`
	State   string `json:"state"`
}
