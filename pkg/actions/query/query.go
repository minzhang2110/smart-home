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
	RequestID string                            `json:"requestId"`
	Payload   map[string]map[string]interface{} `json:"payload"` // [id][type]status, "123":"on":true
}
