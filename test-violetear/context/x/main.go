package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-app-Version", "1.0")
		next.ServeHTTP(w, r)
	})
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		ctx := context.WithValue(r.Context(), "m1", "m1")
		ctx = context.WithValue(ctx, "key", 1)
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		ctx := context.WithValue(r.Context(), "m2", "m2")
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Println("Executing middlewareTwo again")
	})
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Executing finalHandler\nm1:%s\nkey:%d\nm2:%s\n",
		r.Context().Value("m1"),
		r.Context().Value("key"),
		r.Context().Value("m2"),
	)
	w.Write([]byte("I catch all"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	panic("this will never happen, because of the return")
}

func main() {
	router := violetear.New()

	stdChain := middleware.New(commonHeaders, middlewareOne, middlewareTwo)

	router.Handle("/", stdChain.ThenFunc(catchAll), "GET,HEAD")
	router.Handle("/foo", stdChain.ThenFunc(foo), "GET,HEAD")
	router.HandleFunc("/bar", foo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
