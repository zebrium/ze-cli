// Package up Copyright © 2023 ScienceLogic Inc/*

package up

type MetaData struct {
	Stream             string            `json:"stream,omitempty"`
	LogBaseName        string            `json:"logbasename"`
	LogType            string            `json:"log_type,omitempty"`
	Ids                map[string]string `json:"ids,omitempty"`
	Cfgs               map[string]string `json:"cfgs,omitempty"`
	Tags               map[string]string `json:"tags,omitempty"`
	Tz                 string            `json:"tz,omitempty"`
	ZeLogCollectorVers string            `json:"ze_log_collector_vers,omitempty"`
	TM                 bool              `json:"ze_tm,omitempty"`
}

type TokenResp struct {
	Token string `json:"token"`
}
