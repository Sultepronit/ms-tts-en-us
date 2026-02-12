package files

import (
	"fmt"
	"os"
	"strings"
)

func genFnEsc(expression string, record string) (string, string) {
	esc := strings.ReplaceAll(expression, " ", "_")
	fn := fmt.Sprintf("records/%s/%s", esc, record)
	return fn, esc
}

func writeRecord(expression string, record string, data []byte) error {
	// esc := strings.ReplaceAll(expression, " ", "_")
	fn, esc := genFnEsc(expression, record)
	err := os.MkdirAll("records/"+esc, 0755)
	if err != nil {
		return err
	}
	// fn := fmt.Sprintf("records/%s/%s", esc, record)

	return os.WriteFile(fn, data, 0644)
}

func readRecord(expression string, record string) ([]byte, error) {
	// esc := strings.ReplaceAll(expression, " ", "_")
	// fn := fmt.Sprintf("records/%s/%s", esc, record)
	fn, _ := genFnEsc(expression, record)

	return os.ReadFile(fn)
}

func delRecord(expression string, record string) error {
	fn, _ := genFnEsc(expression, record)
	return os.Remove(fn)
}
