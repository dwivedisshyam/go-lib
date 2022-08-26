package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	expected := InvalidBody()

	t.Run("ErrorString", func(t *testing.T) {
		err := Error{Err: "invalid body"}
		actual := err.Error()

		assert.Equal(t, expected.Error(), actual)
	})
}

func Test_EntityNotFound(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusNotFound,
		Err:        "user not found with id: 123",
	}

	t.Run("EntityNotFound", func(t *testing.T) {
		actual := EntityNotFound("user", "123")
		assert.Equal(t, expected, actual)
	})
}

func Test_EntityAlreadyExist(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusConflict,
		Err:        "user already exist",
	}

	t.Run("EntityAlreadyExist", func(t *testing.T) {
		actual := EntityAlreadyExist("user")
		assert.Equal(t, expected, actual)
	})
}

func Test_InvalidBody(t *testing.T) {
	expected := &Error{
		StatusCode: http.StatusBadRequest,
		Err:        "invalid body",
	}

	t.Run("InvalidBody", func(t *testing.T) {
		actual := InvalidBody()
		assert.Equal(t, expected, actual)
	})
}
