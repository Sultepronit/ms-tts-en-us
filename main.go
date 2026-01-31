package main

import (
	"fmt"
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
	// db.FillParsedVoices()
	// v, err := db.SelectVoices(true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(v)
	v, err := voices.GetRandomVoice(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
