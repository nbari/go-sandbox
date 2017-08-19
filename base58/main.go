package main

import (
	"fmt"
	"time"

	"github.com/nbari/base58"
)

func encode(i uint64) {
	x := base58.Encode(i)
	fmt.Printf("%d = %s\n", i, x)
	time.Sleep(time.Millisecond)
}

func main() {
	for i, val := uint64(0), uint64(100); i < val; i++ {
		encode(i)
	}
}
