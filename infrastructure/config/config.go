package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// ConfigEnv function takes key and returns the equivalent of that key from the env file.
func ConfigEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
