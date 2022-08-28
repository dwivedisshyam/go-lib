package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	c := New()
	tcs := []struct {
		name  string
		key   string
		value string
	}{
		{"Success", "APP_NAME", "go-lib"},
		{"EnvNotExist", "ABC_ENV", ""},
	}

	for i := range tcs {
		tc := tcs[i]

		t.Run(tc.name, func(t *testing.T) {
			actual := c.Get(tc.key)

			assert.Equal(t, tc.value, actual)
		})
	}
}

func TestGetDefault(t *testing.T) {
	c := New()
	tcs := []struct {
		name  string
		key   string
		value string
	}{
		{"Success", "APP_NAME", "go-lib"},
		{"EnvNotExist", "DB_VERSION", "2.0"},
	}

	for i := range tcs {
		tc := tcs[i]

		t.Run(tc.name, func(t *testing.T) {
			actual := c.GetDefault(tc.key, tc.value)

			assert.Equal(t, tc.value, actual)
		})
	}
}
