package disconnect

import (
	"encoding/json"
	"testing"
)

func TestExecute(t *testing.T) {
	payload := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.DISCONNECT"
		}]
	}`

	req := &Request{}
	err := json.Unmarshal([]byte(payload), req)
	if err != nil {
		t.Fatal(err)
	}
	if req.RequestID != "ff36a3cc-ec34-11e6-b1a0-64510650abcf" ||
		req.Inputs[0].Intent != "action.devices.DISCONNECT" {
		t.Errorf("%+v", req)
	}

	v, err := json.Marshal(&Response{})
	if err != nil {
		t.Fatal(err)
	}
	e := `{}`
	if string(v) != e {
		t.Error(string(v))
	}
}
