package server

import (
	"fmt"
	"log"
	"net/http"
	"tts/files"
)

func getRecord(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	exp := r.PathValue("expression")
	rec := r.PathValue("record")
	data, err := files.GetOrGenerate(exp, rec)
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
