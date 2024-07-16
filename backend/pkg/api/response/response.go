package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	statusOK    = "OK"
	statusError = "ERROR"
)

func OK() *Response {
	return &Response{
		Status: statusOK,
	}
}

func Error(message string) *Response {
	return &Response{
		Status: statusError,
		Error:  message,
	}
}

func ValidationError(errs validator.ValidationErrors) *Response {
	var errList []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errList = append(errList, fmt.Sprintf("field %s is required", err.Field()))
		case "url":
			errList = append(errList, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return &Response{
		Status: statusError,
		Error:  strings.Join(errList, ","),
	}
}
