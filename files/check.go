package files

import (
	"regexp"
	"strconv"
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
