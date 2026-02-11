package server

import (
	"fmt"
	"log"
	"net/http"
	"tts/files"

	"github.com/rs/cors"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{expression}/{record}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here we go!")
		log.Println(r.URL.Path)
	})

	port := "8080"
	log.Printf("Listening on port %s\n", port)

	handler := cors.Default().Handler(mux)
	// log.Fatal(http.ListenAndServe(":"+port, mux))
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
