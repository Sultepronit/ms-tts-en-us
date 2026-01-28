package db

import (
	"fmt"
	"tts/voices"
)

func FillVoices(list []voices.Voice) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO voices (name, code_name, is_male, rate)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range list {
		rate := -1
		if v.Rate != 0 {
			rate = v.Rate
		}
		fmt.Println(rate)
		_, err = stmt.Exec(v.Name, v.CodeName, v.Gender == "Male", rate)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
