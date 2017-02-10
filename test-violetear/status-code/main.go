package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I handle GET requests\n"))
	// do anything here with the Status code
	cw := w.(*violetear.ResponseWriter)
	fmt.Printf("The status code is: %d\n", cw.Status())
}

func main() {
	router := violetear.New()
	router.HandleFunc("/", handleGET, "GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
