package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go.com/movies/cmd/web/pkg/models"
)

func (app *application) printHtml(w http.ResponseWriter, filename []string, data interface{}) {
	t, err := template.ParseFiles(filename...)
	if err != nil {
		app.servErr(w, err)
		return
	}
	if err := t.Execute(w, data); err != nil {
		app.servErr(w, err)
		return
	}
}
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	k, err := app.movies.LastTwenty()
	if err != nil {
		app.servErr(w, err)
		return
	}
	data := &movieData{Movies: k}
	files := []string{
		"./templates/home.page.tmpl",
		"./templates/base.layout.tmpl",
	}
	app.printHtml(w, files, data)
}
func (app *application) secondPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		w.Write([]byte("Movie not found"))
		return
	}
	k, err := app.movies.Get(id)
	if err != nil {
		if errors.Is(err, models.NoRecord) {
			app.notFound(w)
		} else {
			app.servErr(w, err)
		}
		return
	}
	data := &movieData{Movie: k}
	files := []string{
		"./templates/show.page.tmpl",
		"./templates/base.layout.tmpl",
	}
	app.printHtml(w, files, data)
}
func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientErr(w, http.StatusMethodNotAllowed)
		return
	}
	name := "Example"
	releaseDate := "2023-01-01"
	rating := "5"
	description := "Example movie"
	lastid, err := app.movies.Insert(name, releaseDate, rating, description)
	if err != nil {
		app.servErr(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/movie?id=%d", lastid), http.StatusSeeOther)
}
