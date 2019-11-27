package home

// SyncRequest .
type SyncRequest Request

// SyncResponse .
type SyncResponse struct {
	RequestID string `json:"requestId"`
	Payload   struct {
		AgentUserID string `json:"agentUserId"`
	} `json:"payload"`
}
