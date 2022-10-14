package utils

import "os"

func GetEnvOrDefault(key string, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else {
		return def
	}
}
