package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config
/**
This function is to create a configuration
file with the file name .env
*/
func Config(key string) string {
	// Load .env file
	err := godotenv.Load(".env")

	// Check if failed to load file
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Take env file parameters
	return os.Getenv(key)
}
