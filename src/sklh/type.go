package sklh

type (
	// DataSklh ...
	DataSklh struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	// ResponseTime ...
	ResponseTime struct {
		Duration string `json:"duration"`
		Unit     string `json:"unit"`
	}
	// OutputResponse ...
	OutputResponse struct {
		Data         []DataSklh   `json:"data"`
		ResponseTime ResponseTime `json:"response_time"`
	}
)
