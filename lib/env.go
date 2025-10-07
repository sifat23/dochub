package lib

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
}
