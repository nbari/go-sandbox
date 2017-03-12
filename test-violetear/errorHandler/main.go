package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	v "github.com/nbari/violetear"
)

func handleHello(w http.ResponseWriter, r *http.Request) error {
	t := time.Now()
	if t.Second()%2 != 0 {
		return v.Error{500, fmt.Errorf("Test error %d", t.Unix())}
	}
	fmt.Fprintf(w, "Hi,  %s!", t)
	return nil
}

func main() {
	router := v.New()
	router.Handle(
		"/hello",
		v.ErrorHandler(handleHello),
		"GET, HEAD",
	)
	log.Fatal(http.ListenAndServe(":8000", router))
}
