package db

import "fmt"

func InsertExpression(expression string) error {
	query := `
		INSERT INTO records (expression)
		VALUES (?)
	`
	_, err := conn.Exec(query, expression)
	return err
}

func AddRecord(expression string, num string, voice string) error {
	if len(num) != 1 {
		return fmt.Errorf("invalid record number: %s", num)
	}

	query := fmt.Sprintf("UPDATE records SET v%s = ? WHERE expression = ?", num)
	_, err := conn.Exec(query, voice, expression)
	return err
}
