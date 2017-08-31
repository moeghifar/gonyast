package sekolah

type (
	// DataSekolah ...
	DataSekolah struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	// ResponseTime ...
	ResponseTime struct {
		Duration string `json:"duration"`
		Unit     string `json:"unit"`
	}
	// Error ...
	Error struct {
		ErrorMessage string `json:"error_message"`
		ErrorCode    string `json:"error_code"`
	}
	// OutputResponse ...
	OutputResponse struct {
		Error        Error         `json:"error,omitempty"`
		Data         []DataSekolah `json:"data,omitempty"`
		ResponseTime ResponseTime  `json:"response_time"`
	}
)
