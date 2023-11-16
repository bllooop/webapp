package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Movies struct {
	Name        string
	ReleaseDate string
	Rating      int
	Description string
}

func printHtml(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "500 Server error", 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "500 Server error", 500)
		return
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	movie := Movies{"Example", "2023-1-1", 5, "Example movie"}
	//w.Write([]byte(jsonMovie))
	printHtml(w, "./templates/main_page.html", movie)
}
func secondPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		w.Write([]byte("Movie not found"))
		return
	}
	w.Write([]byte("Детали фильма"))
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET method not allowed"))
		return
	}
	w.Write([]byte("Добавление нового фильма"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/movie", secondPage)
	mux.HandleFunc("/movie/createMovie", createMovie)
	log.Println("Запуск веб-сервиса")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
