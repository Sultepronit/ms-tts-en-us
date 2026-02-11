package main

import (
	"log"
	"tts/db"
	"tts/server"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("v 0.5.1")

	godotenv.Load()

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	// db.Edit()
	// db.FillParsedVoices()
	// v, err := voices.GetRandomVoice(false, []string{"en-US-Ava:DragonHDLatestNeural"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(v)
	// exp := "lest we forget"
	// data, err := service.Generate(exp, v.CodeName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// files.WriteRecord(exp, data)

	// err = db.InsertExpression("test")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.AddRecord("test", 1, v.CodeName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// re, err := db.SelectRecord("tests")
	// log.Println(re)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server.Start()
}
