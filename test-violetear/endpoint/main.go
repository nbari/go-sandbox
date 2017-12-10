package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path: %s", r.URL.Path)
}

func hello_world_all(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path: %s", r.URL.Path)
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "request-id"
	//	router.Verbose = false

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	router.AddRegex(":sopas", `sopas`)

	router.HandleFunc("/", hello_world)
	router.HandleFunc("/hello/world/:sopas", hello_world_all, "GET")
	router.HandleFunc(":uuid", hello_world)
	router.HandleFunc("*", hello_world)
	router.HandleFunc("/hola/*", hello_world)
	router.HandleFunc("/hola/:sopas", hello_world)
	router.HandleFunc("/hola/epazote", hello_world)
	router.HandleFunc("/hola/epazote/*", hello_world)
	router.HandleFunc("/hola/epazote/a/papas", hello_world)

	log.Fatal(http.ListenAndServe(":8080", router))
}
