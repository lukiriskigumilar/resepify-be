package util

import (
	"encoding/json"
	"net/http"
)

type ApiResponseSuccess struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
}

type ApiResponseError struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Reason     interface{} `json:"reason,omitempty"`
}

func NewApiResponseSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(ApiResponseSuccess{
		Status:     true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	})
}

func NewApiResponseError(w http.ResponseWriter, message string, statusCode int,
	reason interface{}) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(ApiResponseError{
		Status:     false,
		Message:    message,
		StatusCode: statusCode,
		Reason:     reason,
	})
}
