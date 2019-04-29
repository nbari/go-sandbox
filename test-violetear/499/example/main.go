package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

func catch499(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Println("Starting middleware")
		defer log.Println("End of middleware")

		ch := make(chan struct{})

		go func() {
			time.Sleep(3 * time.Second)
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			next.ServeHTTP(w, r)
		case <-ctx.Done():
			http.Error(w, ctx.Err().Error(), 499)
			return
		}
	})
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ch := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "Hello World!")
		ch <- struct{}{}
	}()

	select {
	case <-ch:
	case <-ctx.Done():
		http.Error(w, ctx.Err().Error(), http.StatusPartialContent)
	}
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	chain := middleware.New(catch499)

	router.Handle("/", chain.ThenFunc(helloWorld), "GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
