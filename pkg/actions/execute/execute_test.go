package execute

import (
	"encoding/json"
	"testing"
)

func TestExecute(t *testing.T) {
	payload := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.EXECUTE",
		  "payload": {
			"commands": [{
			  "devices": [{
				"id": "123"
			  }],
			  "execution": [{
				"command": "action.devices.commands.OnOff",
				"params": {
				  "on": true
				}
			  }]
			}]
		  }
		}]
	}`

	req := &Request{}
	err := json.Unmarshal([]byte(payload), req)
	if err != nil {
		t.Fatal(err)
	}
	if req.RequestID != "ff36a3cc-ec34-11e6-b1a0-64510650abcf" ||
		req.Inputs[0].Intent != "action.devices.EXECUTE" ||
		req.Inputs[0].Payload.Commands[0].Execution[0].Command != "action.devices.commands.OnOff" ||
		req.Inputs[0].Payload.Commands[0].Execution[0].Params["on"] != true ||
		req.Inputs[0].Payload.Commands[0].Devices[0].ID != "123" {
		t.Errorf("%+v", req)
	}

	v, err := json.Marshal(&Response{
		RequestID: "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		Payload: Payload{
			Commands: []Command{
				Command{
					IDs:    []string{"123"},
					Status: "SUCCESS",
					State:  map[string]interface{}{"on": true, "online": true},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	e := `{"requestId":"ff36a3cc-ec34-11e6-b1a0-64510650abcf","payload":{"Commands":[{"ids":["123"],"status":"SUCCESS","state":{"on":true,"online":true}}]}}`
	if string(v) != e {
		t.Error(string(v))
	}
}
