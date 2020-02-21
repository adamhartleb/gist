package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()


	fileServer := http.FileServer(myFileSystem("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

	return mux
}
