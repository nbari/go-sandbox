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

func main() {
	router := violetear.New(true)

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)

	router.AddPath("/hello/world", hello_world)
	router.AddPath(":uuid", hello_world)

	router.Run(":8080")

}
