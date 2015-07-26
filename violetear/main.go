package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

func main() {
	router := violetear.New()

	log.Fatal(http.ListenAndServe(":8080", router))
}
