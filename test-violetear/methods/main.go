package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func methodGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "methodGET request method: %s", r.Method)
	log.Printf("method = %s\n", r.Method)
}

func methodPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "methodPost request method: %s", r.Method)
}

func methodALL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "methodAll request method: %s", r.Method)
}

func methodRX(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "methodRX request method: %s", r.Method)
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Verbose = true

	router.AddRegex(":any", `\w+`)

	router.HandleFunc("/method", methodGET, "GET, HEAD") // will handle only GET and HEAD
	router.HandleFunc("/method", methodPOST, "POST")     // will handle only POST
	router.HandleFunc("/method", methodALL)              // all but GET, HEAD, POST
	router.HandleFunc("/:any", methodRX)                 // will handle any method

	log.Fatal(http.ListenAndServe(":8080", router))
}
