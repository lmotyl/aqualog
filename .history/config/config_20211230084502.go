package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading %s file", filePath)
	}
}

func GetEnv(key string) string {

	return os.Getenv(key)
}
