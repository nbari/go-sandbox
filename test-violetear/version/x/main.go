package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi,  %s!", r.URL.Path[1:])
}
func handleHelloV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version 2")
}

func main() {
	router := violetear.New()

	router.HandleFunc("/hello", handleHello, "GET, HEAD")
	router.HandleFunc("/hello#violetear.v2", handleHelloV2, "GET, HEAD, POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
