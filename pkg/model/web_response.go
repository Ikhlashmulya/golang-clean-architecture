package model

type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   []any  `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
