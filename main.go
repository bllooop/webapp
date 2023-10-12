package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("домашняя страница"))
}
func secondPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("вторая страница"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/pagetwo", secondPage)
	log.Println("запуск")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
