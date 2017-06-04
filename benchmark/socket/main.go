package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/nbari/violetear"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	body := "Hello World\n"
	w.Header().Set("X-Server", "go-socket")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(body)))
	fmt.Fprint(w, body)
}

const SOCK = "/tmp/benchmark.sock"

func main() {
	os.Remove(SOCK)
	router := violetear.New()
	router.HandleFunc("/", helloWorld)
	l, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatal(http.Serve(l, router))
}
