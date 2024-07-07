package utils

import "os"

// Getenv retrieves the value of the environment variable named by the key.
// If the variable is not present in the environment or is empty, it returns the defaultValue.
func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
