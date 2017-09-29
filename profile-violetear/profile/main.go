package main

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/nbari/violetear"
)

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Millisecond)
	w.Write([]byte("hello"))
}

func oneWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1: word"))
}
func twoWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("2: word"))
}

func main() {
	r := violetear.New()
	r.LogRequests = true
	r.AddRegex(":word", `^\w+$`)
	r.HandleFunc("/", hello, "GET,HEAD")
	r.HandleFunc("/:word", oneWord)
	r.HandleFunc("/:word/:word", twoWord)

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/*", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
