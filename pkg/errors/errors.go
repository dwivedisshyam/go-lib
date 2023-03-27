package errors

import (
	"net/http"
)

type Error struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	StatusCode int    `json:"-"`
}

func (err Error) Error() string {
	return err.Message
}

func NotFound(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Code:       http.StatusText(http.StatusNotFound),
	}
}

func Unexpected(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Code:       http.StatusText(http.StatusInternalServerError),
	}
}

func BadRequest(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Code:       http.StatusText(http.StatusBadRequest),
	}
}

func Validation(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		Code:       http.StatusText(http.StatusUnprocessableEntity),
	}
}

func Unauthenticated(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Code:       http.StatusText(http.StatusUnauthorized),
	}
}

func Unauthorized(message string) error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Code:       http.StatusText(http.StatusForbidden),
	}
}
