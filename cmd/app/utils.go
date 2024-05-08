package main

import "os"

// os package library doesnot provide an option to set fallback value
// this function will take a key, return the corresponding env value
// if not present return the fallback value
func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
