package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home"))
}

func showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("Showing gist with id %d", id)))
}

func createGist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create Gist"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}


