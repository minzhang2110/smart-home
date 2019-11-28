package query

// Request .
type Request struct {
	RequestID string `json:"requestId"`
	Inputs    []struct {
		Intent  string `json:"intent"`
		Payload struct {
			Devices []struct {
				ID string `json:"id"`
			} `json:"devices"`
		} `json:"payload"`
	} `json:"inputs"`
}

// Response .
type Response struct {
	RequestID string  `json:"requestId"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	Devices map[string]map[string]interface{} `json:"devices"` // [id][type]status, "123":"on":true
}
