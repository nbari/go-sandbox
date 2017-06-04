package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	body := "Hello World\n"
	w.Header().Set("X-Server", "go-http")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(body)))
	fmt.Fprint(w, body)
}

func main() {
	router := violetear.New()
	router.HandleFunc("/", helloWorld)

	log.Fatal(http.ListenAndServe(":8000", router))
}
