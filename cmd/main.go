package main

import (
	"log"
	"net/http"

	"github.com/joangavelan/todo-app/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)

	log.Fatal(http.ListenAndServe(":3000", mux))
}