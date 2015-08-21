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

func main() {
	router := violetear.New(true)

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	router.AddRegex(":catchall", `.*`)

	router.AddPath("/hello/world", hello_world)
	router.AddPath("/hello/world/:catchall", hello_world_all)
	router.AddPath(":uuid", hello_world)

	router.Run(":8080")

}
