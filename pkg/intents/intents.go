package intents

import (
	"log"
	"encoding/json"
	"github.com/minzhang2110/smart-home/pkg/actions/disconnect"
	"github.com/minzhang2110/smart-home/pkg/actions/execute"
	"github.com/minzhang2110/smart-home/pkg/actions/query"
	"github.com/minzhang2110/smart-home/pkg/actions/sync"
	"github.com/minzhang2110/smart-home/pkg/devices"

	"github.com/minzhang2110/smart-home/pkg/actions"
)

// Execute .
func Execute(v []byte, dvcs *devices.Mgr) []byte {
	req := &actions.Request{}
	err := json.Unmarshal(v, req)
	if err != nil {
		return errorCode(nil, actions.CodeProtocolError)
	}

	var resp interface{}
	switch req.Inputs[0].Intent {

	case actions.SyncIntent:
		resp, err = ExecuteSync((*sync.Request)(req), dvcs)

	case actions.QueryIntent:
		request := &query.Request{}
		err = json.Unmarshal(v, request)
		if err != nil {
			return errorCode(nil, actions.CodeProtocolError)
		}
		resp, err = ExecuteQuery(request, dvcs)

	case actions.ExecuteIntent:
		request := &execute.Request{}
		err = json.Unmarshal(v, request)
		if err != nil {
			return errorCode(nil, actions.CodeProtocolError)
		}
		resp, err = ExecuteExecute(request, dvcs)

	case actions.DisconnectIntent:
		resp, err = ExecuteDisconnect((*disconnect.Request)(req), dvcs)

	default:
		return errorCode(req, actions.CodeNotSupported)
	}

	if err != nil {
		return errorCode(req, actions.CodeUnknownError)
	}

	ret, err := json.Marshal(resp)
	if err != nil {
		return errorCode(req, actions.CodeUnknownError)
	}
	return ret
}

func errorCode(req *actions.Request, code string) []byte {
	var id string
	if req != nil {
		id = req.RequestID
	}
	ret, _ := json.Marshal(&actions.Response{
		RequestID: id,
		Payload: actions.Payload{
			ErrorCode: code,
		},
	})
	return ret
}

func ExecuteSync(r *sync.Request, dvcs *devices.Mgr) (resp *sync.Response, err error) {
	defer func() {
		log.Printf("[sync] req: %+v, resp: %+v, error: %v", r, resp, err)
	}()

	o := dvcs.Outlet
	resp = &sync.Response{
		RequestID: r.RequestID,
		Payload: sync.Payload{
			AgentUserID: dvcs.AgentUserID,
			Devices: []sync.Device{
				sync.Device{
					ID:     o.ID,
					Type:   o.Type,
					Traits: o.Traits,
					Name: sync.Name{
						Name: o.Name,
					},
					WillReportState: false,
				},
			},
		},
	}

	return resp, nil
}

func ExecuteQuery(r *query.Request, dvcs *devices.Mgr) (resp *query.Response, err error) {
	defer func() {
		log.Printf("[query] req: %+v, resp: %+v, error: %v", r, resp, err)
	}()

	o := dvcs.Outlet
	resp = &query.Response{
		RequestID: r.RequestID,
		Payload:   query.Payload {
			Devices: map[string]map[string]interface{}{
				o.ID: map[string]interface{}{"on": o.On, "online": o.Online},
			},
		},
	}
	return resp, nil
}

func ExecuteExecute(r *execute.Request, dvcs *devices.Mgr) (resp *execute.Response, err error) {
	defer func() {
		log.Printf("[execute] req: %+v, resp: %+v, error: %v", r, resp, err)
	}()

	o := dvcs.Outlet
	switch r.Inputs[0].Payload.Commands[0].Execution[0].Command {
	case execute.CommandOnOff:
		o.Turn(r.Inputs[0].Payload.Commands[0].Execution[0].Params["on"].(bool))
	}
	return &execute.Response{
		RequestID: r.RequestID,
		Payload: execute.Payload{
			Commands: []execute.Command{
				execute.Command{
					IDs:    []string{o.ID},
					Status: "SUCCESS",
					State:  map[string]interface{}{"on": o.On, "online": o.Online},
				},
			},
		},
	}, nil
}

func ExecuteDisconnect(r *disconnect.Request, dvcs *devices.Mgr) (resp *disconnect.Response, err error) {
	defer func() {
		log.Printf("[disconnect] req: %+v, resp: %+v, error: %v", r, resp, err)
	}()
	return &disconnect.Response{}, nil
}
