package main

import (
	"fmt"
	"time"

	"github.com/nbari/base58"
)

func encode(i uint64) {
	x := base58.Encode(i)
	fmt.Printf("%d = %s\n", i, x)
	time.Sleep(time.Second)
}

func main() {
	concurrency := 4
	sem := make(chan struct{}, concurrency)
	//for i, val := uint64(0), uint64(1<<16); i <= val; i++ {
	for i, val := uint64(0), uint64(1<<4); i <= val; i++ {
		sem <- struct{}{}
		go func(i uint64) {
			defer func() { <-sem }()
			encode(i)
		}(i)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}
