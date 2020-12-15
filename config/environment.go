package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//LoadENV ... Load .env file from the root of project
func LoadENV() {
	env := os.Getenv("GO_ENV")
	var file string
	if "production" != env {
		file = ".env.development.local"
	} else {
		file = ".env.local"
	}
	log.Printf("Load %s", file)
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading %s", file)
	}
}
