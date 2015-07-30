package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

func handle_Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my PID: %d\n%s", os.Getpid(), time.Now())
}

func handle_Panic(w http.ResponseWriter, r *http.Request) {
	panic("just panicked and continue living...")
}

func handle_Panic2(w http.ResponseWriter, r *http.Request) {
	panic("woops")
}

func main() {
	log.Printf("Starting with PID: %d", os.Getpid())
	http.HandleFunc("/", handle_Root)
	http.HandleFunc("/panic", handle_Panic)
	http.Handle("/die", PanicHandler(handle_Panic2))
	http.ListenAndServe(":8080", nil)
}
