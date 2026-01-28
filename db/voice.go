package db

import (
	"fmt"
	"log"
	"tts/rawvoices"
	"tts/voices"
)

func SelectVoices(is_male bool) ([]voices.Voice, error) {
	query := `
		SELECT name, code_name, rate, rating, comment
		FROM voices
		WHERE excluded = false
			AND is_male = ?
	`
	rows, err := conn.Query(query, is_male)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	re := make([]voices.Voice, 0, 40)
	for rows.Next() {
		var v voices.Voice
		err = rows.Scan(&v.Name, &v.CodeName, &v.Rate, &v.Rating, &v.Comment)
		if err != nil {
			return nil, err
		}
		re = append(re, v)
	}

	return re, nil
}

func fillVoices(list []rawvoices.RawVoice) error {
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

func FillParsedVoices() {
	err := fillVoices(rawvoices.Parse())
	if err != nil {
		log.Fatal(err)
	}
}
