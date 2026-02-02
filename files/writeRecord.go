package files

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func WriteRecord(expression string, data []byte) error {
	esc := strings.ReplaceAll(expression, " ", "_")
	ri := rand.IntN(10)
	fmt.Println(ri)
	err := os.MkdirAll("records/"+esc, 0755)
	if err != nil {
		return err
	}
	fn := fmt.Sprintf("records/%s/%d.mp3", esc, ri)

	return os.WriteFile(fn, data, 0644)
}
