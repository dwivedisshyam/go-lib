package config

// MockConfig is the mock of Config
type MockConfig struct {
	Env map[string]string
}

// Get returns the value of given environment vairable
func (c *MockConfig) Get(key string) string {
	return c.Env[key]
}

// GetDefault returns the value of given environment vairable, returns the value otherwise
func (c *MockConfig) GetDefault(key, value string) string {
	val := c.Env[key]
	if val == "" {
		val = value
	}

	return val
}
