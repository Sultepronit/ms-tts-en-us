package files

import (
	"fmt"
	"os"
	"strings"
)

func writeRecord(expression string, record string, data []byte) error {
	esc := strings.ReplaceAll(expression, " ", "_")
	err := os.MkdirAll("records/"+esc, 0755)
	if err != nil {
		return err
	}
	fn := fmt.Sprintf("records/%s/%s", esc, record)

	return os.WriteFile(fn, data, 0644)
}

func readRecord(expression string, record string) ([]byte, error) {
	esc := strings.ReplaceAll(expression, " ", "_")
	fn := fmt.Sprintf("records/%s/%s", esc, record)

	return os.ReadFile(fn)
}
