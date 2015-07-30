package main

import (
	"log"
	"net/http"
)

type PanicHandler func(w http.ResponseWriter, r *http.Request)

func (h PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	h(w, r)
}

func PanickyHandler(w http.ResponseWriter, r *http.Request) {
	panic("woops")
}

func main() {

	http.Handle("/", PanicHandler(PanickyHandler))
	http.ListenAndServe(":8080", nil)
}
