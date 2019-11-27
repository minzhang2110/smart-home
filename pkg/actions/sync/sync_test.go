package sync

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSync(t *testing.T) {
	payload := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.SYNC"
		}]
	}`

	req := &Request{}
	json.NewDecoder(strings.NewReader(payload)).Decode(req)
	if req.RequestID != "ff36a3cc-ec34-11e6-b1a0-64510650abcf" ||
		req.Inputs[0].Intent != "action.devices.SYNC" {
		t.Error()
	}

	v, err := json.Marshal(&Response{
		RequestID: "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		Payload: Payload{
			AgentUserID: "1836.15267389",
			Devices: []Device{
				Device{
					ID:     "123",
					Type:   TypeOutlet,
					Traits: []string{TraitOnOff},
					Name: Name{
						Name: "Night light",
					},
					WillReportState: false,
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	e := `{"requestId":"ff36a3cc-ec34-11e6-b1a0-64510650abcf","payload":{"agentUserId":"1836.15267389","devices":[{"id":"123","type":"action.devices.types.OUTLET","traits":["action.devices.traits.OnOff"],"name":{"name":"Night light"},"willReportState":false}]}}`
	if string(v) != e {
		t.Error()
	}
}
