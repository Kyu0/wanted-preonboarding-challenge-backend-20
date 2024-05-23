package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

var FILES = []string{".env"}

func LoadEnvironmentVariables() {
	err := godotenv.Load(FILES...)

	if err != nil {
		log.Fatal("Failed to Loading '.env' file.")
	}
}
