package env

import (
	"os"
	"strconv"
)

// GetString retrieves the value of the environment variable named by key.
func GetString(key string, defaultValue string) string {
	// If the variable is not set, it returns the provided default value.
	if value := os.Getenv(key); value != "" {
		return value
	}
	// If the variable is set, it returns the value.
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	// Retrieves the value of the environment variable named by key.
	if value := os.Getenv(key); value != "" {
		// If the variable is set, it attempts to convert the value to an integer.
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	// If the variable is not set or conversion fails, it returns the provided default value.
	return defaultValue
}
