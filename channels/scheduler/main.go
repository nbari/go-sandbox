package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var timer *time.Ticker

func scheduler(seconds time.Duration) *time.Ticker {
	ticker := time.NewTicker(seconds * time.Second)
	go func() {
		for t := range ticker.C {
			// do stuff
			fmt.Println(t)
		}
	}()
	return ticker
}

func Start(timer *time.Ticker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer = scheduler(1)
		w.Write([]byte("Starting scheduler"))
	})
}

func Stop(timer *time.Ticker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer.Stop()
		w.Write([]byte("Stoping scheduler"))
	})
}

func main() {
	timer = scheduler(1)
	http.Handle("/start", Start(timer))
	http.Handle("/stop", Stop(timer))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
