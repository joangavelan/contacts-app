package main

import (
	"log"
	"net/http"

	"github.com/joangavelan/todo-app/handler"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))


	mux.HandleFunc("/{$}", handler.Home)

	log.Fatal(http.ListenAndServe(":3000", mux))
}