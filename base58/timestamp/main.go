package main

import (
	"fmt"
	"time"

	"github.com/nbari/base58"
)

func main() {
	var f string
	var sum uint64
	for len(f) < 3 {
		sum++
		f = base58.Encode(sum)
	}
	sum--
	fmt.Printf("%d --> %s\n", sum, base58.Encode(sum))

	for i := 0; i <= 60; i++ {
		now := uint64(time.Now().Unix())
		x := base58.Encode(now)
		fmt.Printf("x = %s now: %d\n", x, now)
		time.Sleep(time.Second)
	}
}
