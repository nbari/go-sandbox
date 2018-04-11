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
	ctx := r.Context()
	ch := make(chan struct{})

	cmd := exec.CommandContext(ctx, "python", "game.py")
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

	go func(ch chan struct{}) {
		cmd.Wait()
		wPipe.Close()
		ch <- struct{}{}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("Client disconnected: %s\n", err)
	}
}

func writeOutput(w http.ResponseWriter, input io.ReadCloser) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Important to make it work in browsers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	in := bufio.NewScanner(input)
	for in.Scan() {
		data := in.Text()
		log.Printf("data: %s\n", data)
		fmt.Fprintf(w, "data: %s\n", data)
		flusher.Flush()
	}
	input.Close()
}

func main() {
	router := violetear.New()
	router.HandleFunc("/", stream, "GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
