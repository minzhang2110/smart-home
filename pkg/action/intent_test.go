package action

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIntent(t *testing.T) {
	payload := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.QUERY",
		  "payload": {
			"devices": [{
			  "id": "123",
			  "customData": {
				"fooValue": 74,
				"barValue": true,
				"bazValue": "foo"
			  }
			}]
		  }
		}]
	}`

	req := &Request{}
	json.NewDecoder(strings.NewReader(payload)).Decode(req)
	if req.RequestID != "ff36a3cc-ec34-11e6-b1a0-64510650abcf" ||
		req.Inputs[0].Intent != "action.devices.QUERY" {
		t.Error()
	}

	v, err := json.Marshal(&Response{
		RequestID: req.RequestID,
		Payload: Payload{
			ErrorCode: CodeUnknownError,
		},
	})
	if err != nil {
		t.Error(err)
	}
	if string(v) != `{"requestId":"ff36a3cc-ec34-11e6-b1a0-64510650abcf","payload":{"errorCode":"unknownError"}}` {
		t.Error()
	}
}
