package sync

import "github.com/minzhang2110/smart-home/pkg/actions"

// Request .
type Request actions.Request

// Response .
type Response struct {
	RequestID string  `json:"requestId"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	AgentUserID string   `json:"agentUserId"`
	Devices     []Device `json:"devices"`
}

const (
	TypeOutlet = "action.devices.types.OUTLET"
	TypeLight  = "action.devices.types.LIGHT"
)

const (
	TraitOnOff      = "action.devices.traits.OnOff"
	TraitBrightness = "action.devices.traits.Brightness"
)

type Device struct {
	ID              string   `json:"id"`
	Type            string   `json:"type"`
	Traits          []string `json:"traits"`
	Name            Name     `json:"name"`
	WillReportState bool     `json:"willReportState"`
}

type Name struct {
	// DefaultNames []string `json:"defaultNames"` // Optional
	Name string `json:"name"`
}
