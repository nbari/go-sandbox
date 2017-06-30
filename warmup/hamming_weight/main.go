package main

import "fmt"

func main() {
	println(ones(13))
}

func ones(x int) int {
	n := 0
	for x != 0 {
		fmt.Printf("%08b\n", x)
		n += x & 1
		x = x >> 1
	}
	return n
}
