package main

import (
	"log"
	"tts/db"
	"tts/voices"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	// db.Edit()

	// voices.FilterVoices()
	err = db.FillVoices(voices.FilterVoices())
	if err != nil {
		log.Fatal(err)
	}

}
