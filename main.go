package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go!!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("form for creating snippet"))
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
