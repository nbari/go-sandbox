package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type scheduler struct {
	t    <-chan time.Time
	quit chan struct{}
	f    func()
}

type schedulerPool struct {
	schedulers map[string]scheduler
}

func NewScheduler() *schedulerPool {
	return &schedulerPool{
		schedulers: make(map[string]scheduler),
	}
}

func (s *schedulerPool) Add(name string, duration time.Duration) {

	scheduler := scheduler{
		t:    time.NewTicker(duration).C,
		quit: make(chan struct{}),
		f:    func() { fmt.Println(name) },
	}

	s.schedulers[name] = scheduler
	go func() {
		for {
			select {
			case <-scheduler.t:
				scheduler.f()
			case <-scheduler.quit:
				return
			}
		}
	}()
}

func (s *schedulerPool) Stop(name string) error {
	scheduler := s.schedulers[name]
	close(scheduler.quit)
	return nil
}

func New(s *schedulerPool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := r.URL.Query().Get("duration")
		n := r.URL.Query().Get("name")
		seconds, _ := time.ParseDuration(d)
		s.Add(n, seconds)
		w.Write([]byte("Starting scheduler"))
	})
}

func Stop(s *schedulerPool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := r.URL.Query().Get("name")
		s.Stop(n)
		w.Write([]byte("Stoping scheduler"))
	})
}

func main() {
	s := NewScheduler()
	http.Handle("/new", New(s))
	http.Handle("/stop", Stop(s))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
