package models

type ErrorResponse struct {
	Success   bool        `json:"success"`
	Message   interface{} `json:"message"`
	ErrorCode interface{} `json:"error_code"`
	Url       string      `json:"url"`
}