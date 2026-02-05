package files

import (
	"log"
	"os"
	"tts/service"
	"tts/voices"
)

func generate(expression string, record string) ([]byte, error) {
	v, err := voices.GetRandomVoice(false)
	if err != nil {
		return nil, err
	}

	data, err := service.Generate(expression, v.CodeName)
	if err != nil {
		return nil, err
	}

	go writeRecord(expression, record, data)

	return data, err
}

func GetOrGenerate(expression string, record string) ([]byte, error) {
	data, err := readRecord(expression, record)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Generating: %s - %s", expression, record)
			return generate(expression, record)
		}

		return nil, err
	}

	log.Printf("Found: %s - %s", expression, record)
	return data, nil
}
