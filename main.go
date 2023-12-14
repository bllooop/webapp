package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Movies struct {
	Name        string
	ReleaseDate string
	Rating      int
	Description string
}

type application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errolog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		ErrorLog: errolog,
		InfoLog:  infolog,
	}

	serv := &http.Server{
		Addr:     *addr,
		ErrorLog: errolog,
		Handler:  app.routes(),
	}
	infolog.Printf("Запуск веб-сервиса на %s", *addr)
	err := serv.ListenAndServe()
	errolog.Fatal(err)
}
