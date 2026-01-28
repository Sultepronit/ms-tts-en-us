package rawvoices

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

var dir string

func getDir() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("no path!")
	}
	dir = filepath.Dir(filename)
}

func prepPath(end string) string {
	if dir == "" {
		getDir()
	}
	return filepath.Join(dir, end)
}

type RawVoice struct {
	CodeName string `json:"ShortName"`
	Name     string `json:"DisplayName"`
	Gender   string `json:"Gender"`
	Locale   string `json:"Locale"`
	Rate     int    `json:"WordsPerMinute,string"`
}

func Parse() []RawVoice {
	path := prepPath("voices.json")
	fmt.Println(path)
	file, err := os.ReadFile(path)
	handleErr(err)

	var all []RawVoice
	err = json.Unmarshal(file, &all)
	handleErr(err)
	// fmt.Println(all)

	re := make([]RawVoice, 0, 100)
	for _, v := range all {
		if v.Locale != "en-US" || v.Gender == "Neutral" {
			continue
		}
		re = append(re, v)
		// fmt.Println(v)
	}
	// fmt.Println(enUS)
	return re
}
