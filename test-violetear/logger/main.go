package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("hello world"))
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.HandleFunc("*", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", router))
}
