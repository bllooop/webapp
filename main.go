package main

import (
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	w.Write([]byte("Список фильмов"))
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
