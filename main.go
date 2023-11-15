package main

import (
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("page not found"))
		return
	}
	w.Write([]byte("домашняя страница"))
}
func secondPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("вторая страница"))
}
func thirdPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		w.Write([]byte("page not found"))
		return
	}
	w.Write([]byte("третья страница"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/pagetwo", secondPage)
	mux.HandleFunc("/pagethree", thirdPage)
	log.Println("запуск")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
