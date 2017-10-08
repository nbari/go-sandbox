package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/nbari/violetear"
)

func main() {
	router := violetear.New()
	router.Verbose = false
	router.LogRequests = true
	router.Logger = func(w *violetear.ResponseWriter, r *http.Request) {
		fmt.Printf("%d\n", w.Status())
	}
	router.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("hello world"))
	})
	ts := httptest.NewServer(router)
	defer ts.Close()
	client := ts.Client()
	client.Timeout = time.Duration(time.Millisecond)
	client.Get(ts.URL)
}
