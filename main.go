package main

import (
	"tts/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// voices.FilterVoices()
	db.Open()
}
