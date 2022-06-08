package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(dotEnvPath string, key string) string {
	err := godotenv.Load(dotEnvPath)
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
