package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}
