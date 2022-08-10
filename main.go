package main

import (
	"goweb/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/contact", handler.ContactHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	port := os.Getenv("PORT")

	log.Println("Starting web on port" + port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
