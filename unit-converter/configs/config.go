package config

import "os"

// GetPort returns the server port
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080" // Default port
	}
	return ":" + port
}
