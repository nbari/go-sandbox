package main

import (
	"fmt"

	"github.com/nbari/base58"
)

func encode(i uint64) {
	x := base58.Encode(i)
	fmt.Printf("%d = %+v\n", i, x)
}

func main() {
	for i, val := uint64(0), uint64(1<<24); i <= val; i++ {
		encode(i)
	}
}
