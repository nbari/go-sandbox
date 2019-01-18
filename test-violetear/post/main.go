package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	x := r.Form.Get("foo")
	fmt.Println(x)
}

func main() {
	router := violetear.New()
	router.HandleFunc("*", handler, "POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
