package home

// Request .
type Request struct {
	RequestID string `json:"requestId"`
	Inputs    []struct {
		Intent string `json:"intent"`
	} `json:"inputs"`
}

const (
	SyncIntent = "action.devices.SYNC"
	QueryIntent = "action.devices.QUERY"
	ExecuteIntent = "action.devices.EXECUTE"
	DisconnectIntent = "action.devices.DISCONNECT"
)