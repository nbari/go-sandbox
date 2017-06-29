package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	array := make([]int, n)
	for a := range array {
		fmt.Fscan(io, &array[(n-1)-a])
	}
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
	println()
}
