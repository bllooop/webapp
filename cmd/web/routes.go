package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", app.home)
	r.HandleFunc("/movie", app.secondPage)
	r.HandleFunc("/movie/createMovie", app.createMovie)
	fileServer := http.FileServer(http.Dir("./templates/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	return r
}
