package errors

import (
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (err Error) Error() string {
	return err.Message
}

func (err Error) StatusCode() int {
	return err.Code
}

func NotFound(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func Unexpected(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func BadRequest(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func Validation(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func Unauthenticated(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func Unauthorized(message string) error {
	return &Error{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
