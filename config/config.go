package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value from key ---

func Config(key string) string {
	//Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file %s\n", err.Error())
	}
	return os.Getenv(key)
}
