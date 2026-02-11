package server

// import (
// 	"log"
// 	"regexp"
// 	"strings"
// )

// var rgxMp3 = regexp.MustCompile(`^([1-6])\.mp3$`)

// func getNum(fn string) string {
// 	matches := rgxMp3.FindStringSubmatch(fn)
// 	if len(matches) > 1 {
// 		return matches[1]
// 	}

// 	return ""
// }

// func parseReq(path string) []string {
// 	parts := strings.Split(path, "/")
// 	if len(parts) < 3 {
// 		return nil
// 	}
// 	log.Println(parts)

// 	num := getNum(parts[2])
// 	// log.Println(num)
// 	if num == "" {
// 		return nil
// 	}

// 	return []string{parts[1], parts[2], num}
// }
