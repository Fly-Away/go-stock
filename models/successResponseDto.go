package models

type SuccessResponseDto struct {
	Success   bool        `json:"success"`
	Message   interface{} `json:"message"`
	Data      interface{}   `json:"data"`
}
