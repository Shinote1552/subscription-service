package summary

type Request struct {
	UserID      string `json:"user_id,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	StartPeriod string `json:"start_period"`
	EndPeriod   string `json:"end_period"`
}

type Response struct {
	TotalPrice int64 `json:"total_price"`
}
