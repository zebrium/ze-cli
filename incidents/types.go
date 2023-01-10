// Package incidents Copyright © 2023 ScienceLogic Inc/*

package incidents

type readIncident struct {
	TimeFrom           int    `json:"time_from"`
	TimeTo             int    `json:"time_to"`
	Timezone           string `json:"timezone"`
	RepeatingIncidents string `json:"repeating_incidents"`
	Occurrences        string `json:"occurrences"`
	TimeBuckets        string `json:"time_buckets"`
	InciId             string `json:"inci_id,omitempty"`
	ITypeId            string `json:"itype_id,omitempty"`
	InciSignal         string `json:"inci_signal,omitempty"`
	BatchIds           string `json:"batch_ids,omitempty"`
}

type ReadIncidentResponse struct {
	Data            []readIncidentData `json:"data"`
	Error           IncidentError      `json:"error"`
	Op              string             `json:"op"`
	SoftwareRelease string             `json:"softwareRelease"`
}

type readIncidentData struct {
	InciBadLvl       int                `json:"inci_bad_lvl"`
	InciCode         string             `json:"inci_code"`
	InciEvents       readIncidentEvents `json:"inci_events"`
	InciEventsStr    string             `json:"inci_events_str"`
	InciHasSignal    bool               `json:"inci_has_signal"`
	InciHosts        string             `json:"inci_hosts"`
	InciId           string             `json:"inci_id"`
	InciItypeOcc     int                `json:"inci_itype_occ"`
	InciItypeTtl     string             `json:"inci_itype_ttl"`
	InciLogs         string             `json:"inci_logs"`
	InciRareLvl      string             `json:"inci_rare_lvl"`
	InciSignificance string             `json:"inci_significance"`
	InciSources      string             `json:"inci_sources"`
	InciSvcGroups    string             `json:"inci_svc_grps"`
	InciTagIds       string             `json:"inci_tag_ids"`
	InciTs           string             `json:"inci_ts"`
}

type readIncidentEvents struct {
	IevtEtext string `json:"ievt_etext"`
	IevtTs    string `json:"ievt_ts"`
	IevtLevel int    `json:"ievt_level"`
}

type IncidentError struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
