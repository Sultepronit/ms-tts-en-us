package main

import (
	"fmt"
	"log"
	"tts/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	// db.Edit()
	// db.FillParsedVoices()
	v, err := db.SelectVoices(true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

}
