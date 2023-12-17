package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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
	psqlInfo := "postgres://person:12345@localhost:5432/movies"
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errolog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dbpool, err := pgxpool.New(context.Background(), os.Getenv(psqlInfo))
	if err != nil {
		errolog.Fatal(err)
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		errolog.Fatal(err)
	}
	defer dbpool.Close()
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
	err = serv.ListenAndServe()
	errolog.Fatal(err)
}
