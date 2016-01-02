package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Stat struct {
	sync.Mutex
	counters map[string]int64
}

type StatAverage Stat

func (s *Stat) increaseCounter(key string) {
	s.Lock()
	defer s.Unlock()

	if c, exists := s.counters[key]; !exists {
		s.counters[key] = 1
	} else {
		s.counters[key] = c + 1
	}
}

func main() {
	s := &Stat{
		counters: make(map[string]int64),
	}
	fmt.Println(runtime.GOMAXPROCS(0))

	go s.increaseCounter("test")
	go s.increaseCounter("test")
	go s.increaseCounter("test")
	s.increaseCounter("test")
	time.Sleep(time.Millisecond)
	fmt.Printf("%#v\n", s.counters)
}
