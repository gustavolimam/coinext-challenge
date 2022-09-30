package model

type CustomError struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
