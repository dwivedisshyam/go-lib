package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockGet(t *testing.T) {
	c := NewMockConfig()

	c = &MockConfig{Env: map[string]string{"MOCK_APP_NAME": "mock_app"}}

	tcs := []struct {
		name  string
		key   string
		value string
	}{
		{"Success", "MOCK_APP_NAME", "mock_app"},
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

func TestMockGetDefault(t *testing.T) {
	c := NewMockConfig()

	c = &MockConfig{Env: map[string]string{"MOCK_APP_NAME": "mock_app"}}

	tcs := []struct {
		name  string
		key   string
		value string
	}{
		{"Success", "APP_NAME", "mock_app"},
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
