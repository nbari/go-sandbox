package main

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/nbari/violetear"
)

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond)
	w.Write([]byte("hello"))
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("*"))
}

func main() {
	r := violetear.New()
	r.LogRequests = true
	r.AddRegex(":word", `^\w+$`)
	r.HandleFunc("/hello", hello, "GET,HEAD")
	r.HandleFunc("/hello/:word/", helloWord, "GET,HEAD")
	r.HandleFunc("/*", catchAll, "GET,HEAD")

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/*", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
