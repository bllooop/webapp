package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) printHtml(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		return
	}
	if err := t.Execute(w, data); err != nil {
		app.ErrorLog.Println(err.Error())
		return
	}
}
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	movie := Movies{"Example", "2023-1-1", 5, "Example movie"}
	app.printHtml(w, "./templates/main_page.html", movie)
}
func (app *application) secondPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		w.Write([]byte("Movie not found"))
		return
	}
	w.Write([]byte("Детали фильма"))
}
func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET method not allowed"))
		return
	}
	w.Write([]byte("Добавление нового фильма"))
}
