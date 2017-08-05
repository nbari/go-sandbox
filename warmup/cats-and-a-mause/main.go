package main

import (
	"bufio"
	"fmt"
	"os"
)

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	var n, input int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	for i := 0; i < n; i++ {
		query := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(io, &input)
			query[j] = input
		}
		a2m := Abs(query[2] - query[0])
		b2m := Abs(query[2] - query[1])
		if a2m == b2m {
			fmt.Println("Mouse C")
		} else if a2m > b2m {
			fmt.Println("Cat B")
		} else {
			fmt.Println("Cat A")
		}
	}
}
