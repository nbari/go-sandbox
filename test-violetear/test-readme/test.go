package main

import (
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm catching all\n"))
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I handle GET requests\n"))
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I handle POST requests\n"))
}

func handleUUID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I handle dynamic requests\n"))
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Request_ID = "Request-ID"

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)

	router.HandleFunc("*", catchAll)
	router.HandleFunc("/method", handleGET, "GET")
	router.HandleFunc("/method", handlePOST, "POST")
	router.HandleFunc("/:uuid", handleUUID, "GET,HEAD")

	log.Fatal(http.ListenAndServe(":8080", router))
}
