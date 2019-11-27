package query

import (
	"encoding/json"
	"testing"
)

func TestQuery(t *testing.T) {
	payload := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.QUERY",
		  "payload": {
			"devices": [{
			  "id": "123"
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
		req.Inputs[0].Intent != "action.devices.QUERY" ||
		req.Inputs[0].Payload.Devices[0].ID != "123" {
		t.Error()
	}

	v, err := json.Marshal(&Response{
		RequestID: "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		Payload:   map[string]map[string]interface{}{"123": map[string]interface{}{"on": true, "online": true}},
	})
	if err != nil {
		t.Fatal(err)
	}
	e := `{"requestId":"ff36a3cc-ec34-11e6-b1a0-64510650abcf","payload":{"123":{"on":true,"online":true}}}`
	if string(v) != e {
		t.Error()
	}
}
