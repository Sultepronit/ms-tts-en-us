package files

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"tts/db"
	"tts/service"
	"tts/voices"

	"github.com/bogem/id3v2/v2"
)

func setTag(data []byte, voice string, exp string) ([]byte, error) {
	tag := id3v2.NewEmptyTag()
	tag.SetArtist(voice)
	tag.SetTitle(exp)

	buff := new(bytes.Buffer)
	if _, err := tag.WriteTo(buff); err != nil {
		return nil, err
	}

	buff.Write(data)

	return buff.Bytes(), nil
}

var tempRecords = make(map[string][]string)

func getUsedVoices(exp string, isTemp bool) ([]string, error) {
	if !isTemp {
		return db.SelectOrCreateRecord(exp)
	}

	if val, ok := tempRecords[exp]; ok {
		return val, nil
	}

	tempRecords[exp] = make([]string, 6)
	return tempRecords[exp], nil
}

func updateRecord(exp string, num int, vName string, isTemp bool) error {
	if !isTemp {
		return db.UpdateRecord(exp, num, vName)
	}

	tempRecords[exp][num-1] = vName
	return nil
}

func generate(expression string, record string, isTemp bool) ([]byte, error) {
	d := getNum(record)
	if d < 0 {
		return nil, fmt.Errorf("invalid record: %s", record)
	}

	isMale := d%2 == 1
	// usedVoices, err := db.SelectOrCreateRecord(expression)
	usedVoices, err := getUsedVoices(expression, isTemp)
	if err != nil {
		return nil, err
	}

	v, err := voices.GetRandomVoice(isMale, usedVoices)
	if err != nil {
		return nil, err
	}
	log.Println(v)

	// err = db.UpdateRecord(expression, d, v.CodeName)
	err = updateRecord(expression, d, v.CodeName, isTemp)
	if err != nil {
		return nil, err
	}

	data, err := service.Generate(expression, v.CodeName)
	if err != nil {
		return nil, err
	}

	data, err = setTag(data, v.Name, expression)
	if err != nil {
		return nil, err
	}

	go func() {
		err := writeRecord(expression, record, data, isTemp)
		if err != nil {
			log.Println("Error writing record:", err)
		}
	}()

	return data, err
}

func GetOrGenerate(expression string, record string, isTemp bool) ([]byte, error) {
	data, err := readRecord(expression, record, isTemp)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Generating: %s - %s", expression, record)
			return generate(expression, record, isTemp)
		}

		return nil, err
	}

	log.Printf("Found: %s - %s", expression, record)
	return data, nil
}
