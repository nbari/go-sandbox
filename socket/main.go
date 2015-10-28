package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
	"log"
	//	"net"
	"net/http"
)

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-app-Version", "1.0")
		fn(w, r)
	}
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := violetear.NewResponseWriter(w)
		w.Header().Set("request-id", "naranjas")
		next.ServeHTTP(lw, r)
		log.Printf("%s [%s] %d %d",
			r.RemoteAddr,
			r.URL,
			lw.Status(),
			lw.Size())
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s - %s\n", k, v)
	}
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	stdChain := middleware.New(logger)

	router.Handle("/", stdChain.ThenFunc(final), "GET")
	router.Handle("/foo", stdChain.ThenFunc(final), "GET")

	//l, err := net.Listen("unix", "/tmp/violetear.sock")
	//if err != nil {
	//fmt.Printf("%s\n", err)
	//} else {
	//err := http.Serve(l, router)
	//if err != nil {
	//panic(err)
	//}
	//}

	log.Fatal(http.ListenAndServe(":8080", router))
}
