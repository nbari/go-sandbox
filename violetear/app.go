package main

import (
	"github.com/nbari/go-sandbox/violetear/app"
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

func main() {
	router := violetear.New()
	router.Get("/user/:name/profile", http.HandlerFunc(resource.Profile))
	router.Get("/user/:name", http.HandlerFunc(resource.Test))

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}
