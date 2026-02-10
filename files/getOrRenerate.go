package files

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"tts/db"
	"tts/service"
	"tts/voices"

	"github.com/bogem/id3v2/v2"
)

var rgxMp3 = regexp.MustCompile(`^([1-6])\.mp3$`)

func getNum(record string) int {
	matches := rgxMp3.FindStringSubmatch(record)
	if len(matches) > 1 {
		d, err := strconv.Atoi(matches[1])
		if err != nil {
			return -1
		}
		return d
	}

	return -1
}

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

func generate(expression string, record string) ([]byte, error) {
	d := getNum(record)
	if d < 0 {
		return nil, fmt.Errorf("invalid record: %s", record)
	}

	isMale := d%2 == 1
	usedVoices, err := db.SelectRecord(expression)
	if err != nil {
		return nil, err
	}

	v, err := voices.GetRandomVoice(isMale, usedVoices)
	if err != nil {
		return nil, err
	}
	log.Println(v)

	err = db.UpdateRecord(expression, d, v.CodeName)
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
