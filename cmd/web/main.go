package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
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

func openDB(dsn string) (*sql.DB, error) {
	db, err :=  sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// The sql.Open method does not actually create a database connection so in order to test if everything
	// was set up correctly, we ping the database.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func main() {
	address := flag.String("address", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL Data Source Name")
	flag.Parse()

	app := application{
		infoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}


	db, err := openDB(*dsn)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	// A bit superfluous because the main function doesn't exit unless we terminate it so
	// this never runs.
	defer db.Close()

	// Made custom http.Server to inject error logger.
	srv := http.Server{
		Addr: *address,
		ErrorLog: app.errorLog,
		Handler: app.routes(),
	}

	app.infoLog.Printf("Starting server on %s", *address)
	if err := srv.ListenAndServe(); err != nil {
		app.errorLog.Fatal(err)
	}
}


