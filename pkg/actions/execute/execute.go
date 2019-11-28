package execute

// Request .
type Request struct {
	RequestID string `json:"requestId"`
	Inputs    []struct {
		Intent  string `json:"intent"`
		Payload struct {
			Commands []struct {
				Devices []struct {
					ID string `json:"id"`
				} `json:"devices"`
				Execution []struct {
					Command string                 `json:"command"`
					Params  map[string]interface{} `json:"params"`
				} `json:"execution"`
			} `json:"commands"`
		} `json:"payload"`
	} `json:"inputs"`
}

const (
	CommandOnOff = "action.devices.commands.OnOff"
)

// Response .
type Response struct {
	RequestID string  `json:"requestId"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	IDs    []string               `json:"ids"`
	Status string                 `json:"status"`
	State  map[string]interface{} `json:"state"`
}
