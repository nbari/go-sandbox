package main

import "fmt"

func main() {
	println(ones(3 * 7))
}

func ones(x int) int {
	n := 0
	for x != 0 {
		fmt.Printf("x & 1: %v     x >> 1: %v    x: %d\n", x&1, x>>1, x)
		n += x & 1
		x = x >> 1
	}
	return n
}
