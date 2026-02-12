package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{expression}/{record}", getRecord)
	mux.HandleFunc("DELETE /{expression}/{record}", delRecord)

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
