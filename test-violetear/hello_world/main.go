package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	router := violetear.New()
	//	router.LogRequests = false

	router.HandleFunc("/", hello_world)
	log.Fatal(http.ListenAndServe(":8080", router))
}
