package intents

import (
	"testing"
	"github.com/minzhang2110/smart-home/pkg/devices/outlet"
	"github.com/minzhang2110/smart-home/pkg/devices"
)

func TestT(t *testing.T) {
	dvcs := devices.New(outlet.New("1001", "light"))

	req := `{
		"requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
		"inputs": [{
		  "intent": "action.devices.SYNC"
		}]
	}`
	t.Error(string(Execute([]byte(req), dvcs)))
}
