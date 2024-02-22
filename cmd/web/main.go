package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.com/movies/cmd/web/pkg/models/psql"
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
	movies   *psql.MovieModel
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	password := os.Getenv("DB_PASSWORD")
	databas := flag.String("databas", "postgres://db:"+password+"@localhost:5432/movies", "Подключение к PSQL")
	flag.Parse()
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dbpool, err := OpenDB(*databas)
	if err != nil {
		errorlog.Fatal(err)
	}
	/*dbpool, err := pgxpool.New(context.Background(), os.Getenv(*databas))
	if err != nil {
		errorlog.Fatal(err)
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		errorlog.Fatal(err)
	} */
	defer dbpool.Close()
	app := &application{
		ErrorLog: errorlog,
		InfoLog:  infolog,
		movies:   &psql.MovieModel{DB: dbpool},
	}

	serv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}
	infolog.Printf("Запуск веб-сервиса на %s", *addr)
	err = serv.ListenAndServe()
	errorlog.Fatal(err)
}

func OpenDB(databas string) (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv(databas))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
