package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		return ""
	}

	return os.Getenv(key)
}
