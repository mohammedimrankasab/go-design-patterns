package singleton

import "sync"

// ConfigManager represents a shared configuration manager.
type ConfigManager struct {
	values map[string]string
}

var (
	instance *ConfigManager
	once     sync.Once
)

// GetInstance returns the singleton instance.
//
// sync.Once guarantees that initialization happens only once
// even when accessed concurrently.
func GetInstance() *ConfigManager {

	once.Do(func() {
		instance = &ConfigManager{
			values: map[string]string{
				"environment": "development",
			},
		}
	})

	return instance
}

// Set stores a configuration value.
func (c *ConfigManager) Set(
	key string,
	value string,
) {
	c.values[key] = value
}

// Get retrieves a configuration value.
func (c *ConfigManager) Get(
	key string,
) (string, bool) {

	value, exists := c.values[key]

	return value, exists
}
