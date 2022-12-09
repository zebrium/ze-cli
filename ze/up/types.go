// Copyright (c) 2022 ScienceLogic Inc

package up

type MetaData struct {
	Stream             string            `json:"stream"`
	Logbasename        string            `json:"logbasename"`
	LogType            string            `json:"log_type"`
	Ids                map[string]string `json:"ids"`
	Cfgs               map[string]string `json:"cfgs"`
	Tags               map[string]string `json:"tags"`
	Tz                 string            `json:"tz"`
	ZeLogCollectorVers string            `json:"ze_log_collector_vers"`
}
