package main

import (
	"fmt"
	"sync"

	"github.com/nbari/base58"
)

func encode(i uint64) {
	x := base58.Encode(i)
	fmt.Printf("%d = %s\n", i, x)
}

func main() {
	var wg sync.WaitGroup
	for i, val := uint64(0), uint64(1<<24); i <= val; i++ {
		wg.Add(1)
		go func(j uint64, wg *sync.WaitGroup) {
			encode(j)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
