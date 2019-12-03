// Package config is a generic helper library for configurations used by other
// packages and projects at Lakesite.Net.
package config

import (
	"log"
	"os"
)

// Getenv takes a key and fallback value, then returns either the corresponding
// os.LookupEnv value for the key if it exists, otherwise it returns the 
// fallback value.
func Getenv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// GetWorkingDirectory returns the current working directory (os.Getwd()).
// It will log a fatal error if os.Getwd() returns one.
func GetWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}
