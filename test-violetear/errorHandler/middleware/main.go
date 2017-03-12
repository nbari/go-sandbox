package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nbari/violetear"
	"github.com/nbari/violetear/middleware"
)

type myError interface {
	error
	Status() int
}

// Error represents an error with an associated HTTP status code.
type Error struct {
	Code int
	Err  error
}

// Error return error message
func (e Error) Error() string {
	return e.Err.Error()
}

// Status return  HTTP status code.
func (e Error) Status() int {
	return e.Code
}

type myErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h myErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		switch e := err.(type) {
		case myError:
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) error {
	t := time.Now()
	if t.Second()%2 != 0 {
		return Error{500, fmt.Errorf("Test error %d", t.Unix())}
	}
	fmt.Fprintf(w, "Hi,  %s!", t)
	return nil
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

func foo(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("foo"))
}

func main() {
	router := violetear.New()
	stdChain := middleware.New(commonHeaders, middlewareOne, middlewareTwo)
	router.Handle("/", stdChain.ThenFunc(foo))
	router.Handle(
		"/hello",
		myErrorHandler(handleHello),
		"GET, HEAD",
	)
	log.Fatal(http.ListenAndServe(":8000", router))
}
