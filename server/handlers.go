package server

import (
	"fmt"
	"log"
	"net/http"

	"tts/db"
	"tts/files"
)

func checkTemp(mode string, exp string) bool {
	if mode != "temp" {
		return false
	}

	_, err := db.SelectRecordsVoice(exp, 1)
	if err != nil { // in 99% cases no record in the db
		return true
	}

	return false
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	log.Println(r.URL.Path, q.Encode())
	// log.Println()

	exp := r.PathValue("expression")
	rec := r.PathValue("record")
	isTemp := checkTemp(r.URL.Query().Get("mode"), exp)

	data, err := files.GetOrGenerate(exp, rec, isTemp)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Cache-Control", "public, max-age=604800")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("X-Voice", "Test Voice")
	w.Header().Set("Access-Control-Expose-Headers", "X-Voice")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func delRecord(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	exp := r.PathValue("expression")
	rec := r.PathValue("record")

	err := files.Delete(exp, rec)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "deleted"}`))
}
