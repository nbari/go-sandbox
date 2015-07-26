package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		fn(w, r)
		log.Println("After")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "Hi there, url: %s", r.URL.Path)
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

func main() {
	http.HandleFunc("/", logger(handler))
	http.ListenAndServe(":8080", nil)
}
