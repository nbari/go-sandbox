package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/nbari/violetear"
)

func stream(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("python", "game.py")
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = wPipe
	cmd.Stderr = wPipe
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	go writeOutput(w, rPipe)
	cmd.Wait()
	wPipe.Close()
}

func writeOutput(w http.ResponseWriter, input io.ReadCloser) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Immportant to make it work in browsers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	in := bufio.NewScanner(input)
	for in.Scan() {
		fmt.Fprintf(w, "data: %s\n", in.Text())
		flusher.Flush()
	}
	input.Close()
}

func main() {
	router := violetear.New()
	router.HandleFunc("/", stream, "GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
