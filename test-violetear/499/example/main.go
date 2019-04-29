package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "Hello World!")
		ch <- struct{}{}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		http.Error(w, ctx.Err().Error(), http.StatusPartialContent)
	}
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("/", helloWorld, "GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
