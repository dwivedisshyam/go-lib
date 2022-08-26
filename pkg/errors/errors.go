package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Err        string `json:"error"`
	Code       string `json:"code"`
	StatusCode int    `json:"-"`
}

func (err Error) Error() string {
	return err.Err
}

func EntityNotFound(entity, id string) *Error {
	return &Error{
		Err:        fmt.Sprintf("%s not found with id: %s", entity, id),
		StatusCode: http.StatusNotFound,
	}
}

func InvalidBody() *Error {
	return &Error{
		Err:        "invalid body",
		StatusCode: http.StatusBadRequest,
	}
}

func EntityAlreadyExist(entity string) *Error {
	return &Error{
		Err:        fmt.Sprintf("%s already exist", entity),
		StatusCode: http.StatusConflict,
	}
}
