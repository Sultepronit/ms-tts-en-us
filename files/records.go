package files

import (
	"fmt"
	"os"
	"strings"
)

func genFnDn(expression string, record string, isTemp bool) (string, string) {
	prefix := ""
	if isTemp {
		prefix = "temp-"
	}

	esc := strings.ReplaceAll(expression, " ", "_")
	dn := fmt.Sprintf("%srecords/%s", prefix, esc)
	fn := fmt.Sprintf("%s/%s", dn, record)
	return fn, dn
}

func writeRecord(expression string, record string, data []byte, isTemp bool) error {
	fn, dn := genFnDn(expression, record, isTemp)
	err := os.MkdirAll(dn, 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(fn, data, 0644)
}

func readRecord(expression string, record string, isTemp bool) ([]byte, error) {
	fn, _ := genFnDn(expression, record, isTemp)

	return os.ReadFile(fn)
}

func delRecord(expression string, record string) error {
	fn, _ := genFnDn(expression, record, false)
	return os.Remove(fn)
}
