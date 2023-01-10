// Package incidents Copyright © 2023 ScienceLogic Inc/*

package incidents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/perimeterx/marshmallow"
	"io"
	"net/http"
)

var incidentURL = "/mwsd/v1/incident/read/list"

func List(url string, auth string, startTs int, endTs int, rptInci string, timezone string, batchId string) (response *ReadIncidentResponse, err error) {
	marshmallow.EnableCache()
	client := &http.Client{}
	readIncidentReq := readIncident{TimeFrom: startTs, TimeTo: endTs, Timezone: timezone, BatchIds: batchId, RepeatingIncidents: rptInci}
	body, err := json.Marshal(readIncidentReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", url, incidentURL), bytes.NewBuffer(body))
	req.Header.Add("Authentication", fmt.Sprintf("Bearer %s", auth))
	req.Header.Add("Content-Type", "application/json")
	println("Sending Request")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respMap := &ReadIncidentResponse{}
	println("unmarshling request")
	_, err = marshmallow.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, err
	}
	println("returning ResMap")
	return respMap, nil
}
