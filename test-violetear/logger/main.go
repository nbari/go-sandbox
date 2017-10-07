package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nbari/violetear"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("hander ended")

	ch := make(chan struct{})
	go func(ch chan struct{}) {
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, "Hello world!")
		ch <- struct{}{}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusPartialContent)
	}
}

func myLogger(w *violetear.ResponseWriter, r *http.Request) {
	j := map[string]interface{}{
		"Time":        time.Now().UTC().Format(time.RFC3339),
		"RemoteAddr":  r.RemoteAddr,
		"URL":         r.URL.String(),
		"Status":      w.Status(),
		"Size":        w.Size(),
		"RequestTime": w.RequestTime(),
		"RequestID":   w.RequestID(),
	}
	if err := json.NewEncoder(os.Stdout).Encode(j); err != nil {
		log.Println(err)
	}
}

func main() {
	router := violetear.New()
	router.LogRequests = true
	router.Logger = myLogger
	router.Verbose = false
	router.HandleFunc("*", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", router))
}
