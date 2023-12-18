package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/movie", app.secondPage)
	mux.HandleFunc("/movie/createMovie", app.createMovie)
	fileServer := http.FileServer(http.Dir("./templates/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
