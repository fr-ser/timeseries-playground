package tools

import "os"

// GetEnvDefault returns the value from the environment for the specified key
// or returns the fallback
func GetEnvDefault(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
