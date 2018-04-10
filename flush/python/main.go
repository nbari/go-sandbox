package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/nbari/violetear"
)

func export(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("python", "game.py")
	rPipe, wPipe := io.Pipe()
	cmd.Stdout = wPipe
	cmd.Stderr = wPipe
	go writeCmdOutput(w, rPipe)
	cmd.Run()
	wPipe.Close()
}

func writeCmdOutput(w http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, 1024)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		w.Write(data)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}

func main() {
	router := violetear.New()
	router.HandleFunc("/export", export, "GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
