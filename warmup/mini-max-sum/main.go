package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	values := make([]int, 5)
	for i := range values {
		fmt.Fscan(io, &values[i])
	}
	var x, min, max int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j != i {
				x = x + values[j]
			}
		}
		if i == 0 {
			max = x
			min = x
		}
		if x > max {
			max = x
		} else if x < min {
			min = x
		}
		x = 0
	}
	fmt.Printf("%d %d\n", min, max)
}
