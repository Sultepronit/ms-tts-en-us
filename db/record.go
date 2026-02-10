package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func CreateRecord(expression string) error {
	query := `
		INSERT INTO records (expression)
		VALUES (?)
	`
	_, err := conn.Exec(query, expression)
	return err
}

func UpdateRecord(expression string, num int, voice string) error {
	query := fmt.Sprintf("UPDATE records SET v%d = ? WHERE expression = ?", num)
	fmt.Println(query)
	_, err := conn.Exec(query, voice, expression)
	return err
}

func SelectRecord(expression string) ([]string, error) {
	query := `SELECT v1, v2, v3, v4, v5, v6 FROM records WHERE expression = ?`
	re := make([]string, 6)
	
	ptrs := make([]any, 6)
	for i := range re {
		ptrs[i] = &re[i]
	}

	err := conn.QueryRow(query, expression).Scan(ptrs...)
	if errors.Is(err, sql.ErrNoRows) {
		return re, CreateRecord(expression)
	}
	if err != nil {
		return nil, err
	}

	return re, nil
}
