package actions

// Request .
type Request struct {
	RequestID string `json:"requestId"`
	Inputs    []struct {
		Intent string `json:"intent"`
	} `json:"inputs"`
}

const (
	SyncIntent       = "action.devices.SYNC"
	QueryIntent      = "action.devices.QUERY"
	ExecuteIntent    = "action.devices.EXECUTE"
	DisconnectIntent = "action.devices.DISCONNECT"
)

// Response .
type Response struct {
	RequestID string  `json:"requestId"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	ErrorCode string `json:"errorCode"`
}

const (
	CodeNotSupported  = "notSupported"
	CodeProtocolError = "protocolError"
	CodeUnknownError  = "unknownError"
)
