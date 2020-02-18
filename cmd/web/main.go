package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}


