package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

func sleep() {
	time.Sleep(3 * time.Second)
	fmt.Println("done...")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	go sleep()
	return
}

func main() {
	router := violetear.New()
	router.HandleFunc("*", index)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
