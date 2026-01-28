package voices

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

type Voice struct {
	CodeName string `json:"ShortName"`
	Name     string `json:"DisplayName"`
	Gender   string `json:"Gender"`
	Locale   string `json:"Locale"`
	Rate     string `json:"WordsPerMinute"`
}

func FilterVoices() {
	// dir, err := os.ReadDir(".")
	// handleErr(err)
	// fmt.Println(dir)

	path := prepPath("voices.json")
	fmt.Println(path)
	file, err := os.ReadFile(path)
	handleErr(err)

	var all []Voice
	err = json.Unmarshal(file, &all)
	handleErr(err)
	// fmt.Println(all)

	// var enUS []Voice
	enUS := make([]Voice, 0, 100)
	for _, v := range all {
		if v.Locale != "en-US" || v.Gender == "Neutral" {
			continue
		}
		enUS = append(enUS, v)
		fmt.Println(v)
	}
	// fmt.Println(enUS)
}
