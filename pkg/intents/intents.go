package intents

import (
	"encoding/json"
	"github.com/minzhang2110/smart-home/pkg/devices"

	"github.com/minzhang2110/smart-home/pkg/actions"
)

// Execute .
func Execute(v []byte, dvcs *devices.Mgr) []byte {
	req := &actions.Request{}
	err := json.Unmarshal(v, req)
	if err != nil {
		ret, _ := json.Marshal(&actions.Response{
			Payload: actions.Payload{
				ErrorCode: actions.CodeProtocolError,
			},
		})
		return ret
	}
	switch req.Inputs[0].Intent {
	case actions.SyncIntent:
		return sync(v, dvcs)
	case actions.QueryIntent:
		return query(v, dvcs)
	case actions.ExecuteIntent:
		return sync(v, dvcs)
	case actions.DisconnectIntent:
		return query(v, dvcs)
	}
	return nil
}

func sync(v []byte, dvcs *devices.Mgr) []byte {
	return nil
}

func query(v []byte, dvcs *devices.Mgr) []byte {
	return nil
}

func execute(v []byte, dvcs *devices.Mgr) []byte {
	return nil
}

func disconnect(v []byte, dvcs *devices.Mgr) []byte {
	return nil
}
