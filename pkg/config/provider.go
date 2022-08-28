package config

import "os"

type configProvider struct {
}

// Get returns the value of given environment vairable
func (c *configProvider) Get(key string) string {
	return os.Getenv(key)
}

// GetDefault returns the value of given environment vairable, returns the value otherwise
func (c *configProvider) GetDefault(key, value string) string {
	val := os.Getenv(key)
	if val == "" {
		val = value
	}

	return val
}
