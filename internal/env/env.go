package env

import (
	"os"
	"strconv"
)

// GetString retrieves the value of the environment variable named by key.
func GetString(key, fallback string) string {
	// os.LookupEnv returns the value of the environment variable if it exists, otherwise it returns false.
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

// GetBool retrieves the value of the environment variable named by key as a boolean.
func getInt(key string, fallback int) int {
	// os.LookupEnv returns the value of the environment variable if it exists, otherwise it returns false.
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	// Convert the string value to an integer.
	valAsInt, err := strconv.Atoi(val)
	// If conversion fails, return the fallback value.
	if err != nil {
		return fallback
	}

	return valAsInt
}
