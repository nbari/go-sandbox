package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func v1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func v3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Header.Get("Accept"))
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Verbose = false
	//router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	//router.AddRegex(":sopas", `sopas`)
	//router.HandleFunc("/hello/world/:sopas", hello_world_all, "GET")
	//router.HandleFunc(":uuid", hello_world)
	//router.HandleFunc("*", hello_world)
	router.HandleFunc("/hola/epazote/*", v1)
	router.HandleFunc("/hola/epazote/*#violetear.v3", v3)
	router.HandleFunc("/hola/epazote/a/papas", v1)
	router.HandleFunc("/hola/epazote/a/papas#violetear.v3", v3)
	//router.HandleFunc("/hola/epazote/a/papas", hello_world)

	log.Fatal(http.ListenAndServe(":8000", router))
}
