package main

import (
	"fmt"
	"github.com/nbari/violetear"
	//	"log"
	"net/http"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func hello_world_all(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "catch %s!", r.URL.Path)
}
func not_found(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found %s!", r.URL.Path)
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-app-Version", "1.0")
		fn(w, r)
	}
}

func main() {
	router := violetear.New(true)

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	router.AddRegex(":sopas", `sopas`)

	router.HandleFunc("/hello", hello_world)
	router.HandleFunc("/hello/world", commonHeaders(hello_world), "GET,POST")
	router.HandleFunc("/hello/world/:sopas", hello_world_all, "GET")
	router.HandleFunc(":uuid", hello_world)

	//	router.NotAllowedHandler = http.HandlerFunc(not_found)
	//	router.NotFoundHandler = http.HandlerFunc(not_found)

	router.Run(":8080")

}
