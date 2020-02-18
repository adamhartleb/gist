package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

type myFileSystem string

func (s myFileSystem) Open(name string) (http.File, error) {
	fileInfo, err := os.Stat(path.Join(string(s), name))
	if err != nil || fileInfo.IsDir()  {
		return nil, os.ErrNotExist
	}

	file, err := os.Open(path.Join(string(s), name))
	if err != nil {
		return nil, os.ErrNotExist
	}

	return file, err
}

func main() {
	mux := http.NewServeMux()


	fileServer := http.FileServer(myFileSystem("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}


