package main

import (
	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
	"log"
	"net/http"
)

func my404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ne ne ne", 404)
	})
}

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-app-Version", "1.0")
		next.ServeHTTP(w, r)
	})
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("I catch all"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("foo"))
}

func main() {
	router := violetear.New()
	router.NotFoundHandler = my404()

	stdChain := middleware.New(commonHeaders, middlewareOne, middlewareTwo)

	router.Handle("/", stdChain.ThenFunc(catchAll), "GET")
	router.Handle("/foo", stdChain.ThenFunc(foo), "GET")
	router.HandleFunc("/bar", foo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
