package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

func index(w http.ResponseWriter, r *http.Request) {
	delay := rand.Intn(100)
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Fprintf(w, "delayed: %d", delay)
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo %s!", r.URL.Path)
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar %s!", r.URL.Path)
}

func BasicAuth(next http.Handler, username, password, realm string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func counterMW(c prometheus.Summary) middleware.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Println("Executing counter middleware")
			next.ServeHTTP(w, r)
			log.Println("End of request")
			// call prometheus
			c.Observe(float64(time.Since(start).Seconds()))
		})
	}
}

func main() {
	counter := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "myAPI",
		Name:      "response_duration_seconds",
		Help:      "Time taken to response",
	})
	prometheus.MustRegister(counter)

	router := violetear.New()

	// midleware
	stdChain := middleware.New(counterMW(counter))

	router.Handle("/", stdChain.ThenFunc(index))
	router.Handle("/foo", stdChain.ThenFunc(foo), "GET")
	router.Handle("/bar", stdChain.ThenFunc(bar), "POST")
	router.Handle("/metrics", BasicAuth(prometheus.Handler(),
		"user",
		"password",
		"Please enter username and password"),
	)

	// configure server
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   7 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
