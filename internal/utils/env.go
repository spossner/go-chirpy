package utils

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func MustGetEnvString(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("missing configuration value for: %s", key)
	return ""
}

func GetEnvString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
		log.Fatalf("not a int configuration value: %s=%v\n", key, value)
	}
	return defaultValue
}

var TRUE_VALUES = []string{"yes", "y", "true", "enable", "enabled", "active", "activate", "activated", "on"}
var FALSE_VALUES = []string{"no", "n", "false", "disable", "disabled", "inactive", "deactivate", "deactivated", "off"}

func GetEnvBool(key string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if slices.ContainsFunc(TRUE_VALUES, func(tag string) bool {
			return strings.EqualFold(value, tag)
		}) {
			return true
		}

		if slices.ContainsFunc(FALSE_VALUES, func(tag string) bool {
			return strings.EqualFold(value, tag)
		}) {
			return false
		}
		log.Fatalf("not a boolean configuration value: %s=%v\n", key, value)
	}
	return defaultValue
}
