package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func hello_world_all(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "catch %s!", r.URL.Path)
}
func not_found(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found %s!", r.URL.Path)
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-app-Version", "1.0")
		fn(w, r)
	}
}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		w.Header().Set("X-app-Version", "1.0")
		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := violetear.NewResponseWriter(w, "")
		w.Header().Set("request-id", "naranjas")
		next.ServeHTTP(lw, r)
		log.Printf("---- %s [%s] %d %d",
			r.RemoteAddr,
			r.URL,
			lw.Status(),
			lw.Size())
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

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "request-id"
	//	router.Verbose = false

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	router.AddRegex(":sopas", `sopas`)

	stdChain := middleware.New(logger, middlewareOne, middlewareTwo)

	router.Handle("/", stdChain.ThenFunc(final), "GET")
	router.Handle("/foo", stdChain.ThenFunc(final), "GET")
	router.HandleFunc("/hello/world/:sopas", hello_world_all, "GET")
	router.HandleFunc(":uuid", hello_world)
	router.HandleFunc("*", hello_world)
	router.HandleFunc("/hola/*", hello_world)
	router.HandleFunc("/hola/:sopas", hello_world)
	router.HandleFunc("/hola/epazote", hello_world)
	router.HandleFunc("/hola/epazote/*", hello_world)
	router.HandleFunc("/hola/epazote/a/papas", hello_world)

	log.Fatal(http.ListenAndServe(":8080", router))
}
