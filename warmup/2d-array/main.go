package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	m := [6][6]int{}

	io := bufio.NewReader(os.Stdin)

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Fscan(io, &n)
			m[i][j] = n
		}
	}
	fmt.Printf("m = %+v\n", m)
}
