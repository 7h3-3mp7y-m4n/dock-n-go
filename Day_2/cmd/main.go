package main

import (
	"day2/handler"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8080 ...")
	http.Handle("/", http.HandlerFunc(handler.Hello))
	http.Handle("/bye", http.HandlerFunc(handler.Bye))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
