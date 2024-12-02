package common

import (
	"encoding/json"
	"time"
)

type ApiResponse[T any] struct {
	Success   bool     `json:"success,omitempty"`
	Timestamp int64    `json:"timestamp,omitempty"`
	Message   string   `json:"message"`
	Errors    []string `json:"errors,omitempty"`
	Data      T        `json:"data,omitempty"`
}

func NewApiResponse() *ApiResponse[any] {
	return &ApiResponse[any]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Errors:    []string{},
		Data:      nil,
	}
}

func (t ApiResponse[T]) String() string {
	jsonBytes, _ := json.Marshal(t)
	return string(jsonBytes)
}
