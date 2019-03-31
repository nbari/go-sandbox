package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

func showId(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, violetear.GetParam("*", r))
}

func showEmpty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("/", showEmpty, "HEAD, GET")
	router.HandleFunc("/user/*", showId, "HEAD, GET")
	router.HandleFunc("/user", showEmpty, "POST")

	srv := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   7 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
