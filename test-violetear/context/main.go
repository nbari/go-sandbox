package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("hander ended")

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "CatchAll")
		ch <- struct{}{}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusPartialContent)
	}
}

func main() {
	router := violetear.New()
	router.HandleFunc("*", catchAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
