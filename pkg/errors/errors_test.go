package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	expected := BadRequest("invalid body")

	t.Run("ErrorString", func(t *testing.T) {
		err := Error{Message: "invalid body"}
		actual := err.Error()

		assert.Equal(t, expected.Error(), actual)
	})
}

func Test_NotFound(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusNotFound,
		Code:       http.StatusText(http.StatusNotFound),
		Message:    "user",
	}

	t.Run("NotFound", func(t *testing.T) {
		actual := NotFound("user")
		assert.Equal(t, expected, actual)
	})
}

func Test_Unexpected(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusInternalServerError,
		Code:       http.StatusText(http.StatusInternalServerError),
		Message:    "db error",
	}

	t.Run("UnexpectedErr", func(t *testing.T) {
		actual := Unexpected("db error")
		assert.Equal(t, expected, actual)
	})
}

func Test_BadRequest(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusBadRequest,
		Code:       http.StatusText(http.StatusBadRequest),
		Message:    "invalid body",
	}

	t.Run("BadRequest", func(t *testing.T) {
		actual := BadRequest("invalid body")
		assert.Equal(t, expected, actual)
	})
}

func Test_Validation(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusUnprocessableEntity,
		Code:       http.StatusText(http.StatusUnprocessableEntity),
		Message:    "invalid parameter",
	}

	t.Run("Validation", func(t *testing.T) {
		actual := Validation("invalid parameter")
		assert.Equal(t, expected, actual)
	})
}

func Test_Unauthenticated(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusUnauthorized,
		Code:       http.StatusText(http.StatusUnauthorized),
		Message:    "unauthenticated user",
	}

	t.Run("Unauthenticated", func(t *testing.T) {
		actual := Unauthenticated("unauthenticated user")
		assert.Equal(t, expected, actual)
	})
}

func Test_Unauthorized(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusForbidden,
		Code:       http.StatusText(http.StatusForbidden),
		Message:    "unauthorized user",
	}

	t.Run("Unauthorized", func(t *testing.T) {
		actual := Unauthorized("unauthorized user")
		assert.Equal(t, expected, actual)
	})
}
