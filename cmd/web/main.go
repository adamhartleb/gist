package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

type myFileSystem string

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

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
	address := flag.String("address", ":4000", "HTTP Network Address")
	flag.Parse()

	app := application{
		infoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()


	fileServer := http.FileServer(myFileSystem("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

	// Made custom http.Server to inject error logger.
	srv := http.Server{
		Addr: *address,
		ErrorLog: app.errorLog,
		Handler: mux,
	}

	app.infoLog.Printf("Starting server on %s", *address)
	if err := srv.ListenAndServe(); err != nil {
		app.errorLog.Fatal(err)
	}
}


