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
	SuccessStatus = "success"
	StatusError   = "error"
)

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func Success() Response {
	return Response{
		Status: SuccessStatus,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMessages []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMessages = append(errMessages, err.Field()+" is required")
		case "url":
			errMessages = append(errMessages, err.Field()+" is not a valid URL")
		default:
			errMessages = append(errMessages, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMessages, ", "),
	}
}
