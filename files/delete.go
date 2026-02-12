package files

import (
	"fmt"
	"log"
	"tts/db"
)

func Delete(expression string, record string) error {
	d := getNum(record)
	if d < 0 {
		return fmt.Errorf("invalid record: %s", record)
	}

	err := delRecord(expression, record)
	if err != nil {
		return err
	}
	
	v, err := db.SelectRecordsVoice(expression, d)
	if err != nil {
		return err
	}
	log.Println(v)

	return db.DownrateVoice(v)
}
