package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/1", panicRecover(handler))
	http.HandleFunc("/2", panicRecover(handler2))
	http.HandleFunc("/3", panicRecover(handler3))
	fmt.Println("Hello, playground")
}

func panicRecover(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC RECOVERED:%s\n", r)
			}
		}()
		f(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Your stuff
}

func handler2(w http.ResponseWriter, r *http.Request) {
	// Your stuff
}

func handler3(w http.ResponseWriter, r *http.Request) {
	// Your stuff
}
