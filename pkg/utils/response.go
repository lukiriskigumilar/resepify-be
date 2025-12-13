package utils

import (
	"github.com/gin-gonic/gin"
)

type ApiResponseSuccess struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
}

type ApiResponseError struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Errors     interface{} `json:"errors,omitempty"`
}

func NewApiResponseSuccess(c *gin.Context, message string, data interface{}, statusCode int) {
	response := ApiResponseSuccess{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
	c.JSON(statusCode, response)
}

func NewApiResponseError(c *gin.Context, message string, statusCode int,
	errors interface{}) {
	response := ApiResponseError{
		Success:    false,
		Message:    message,
		StatusCode: statusCode,
		Errors:     errors,
	}
	c.JSON(statusCode, response)
}
