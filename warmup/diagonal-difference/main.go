package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	//fmt.Println(math.Abs(float64(a - b)))
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	var a, b int
	for r := 0; r < n; r++ {
		row := make([]int, n)
		for i := range row {
			fmt.Fscan(io, &row[i])
		}
		a = a + row[r]
		b = b + row[(n-1)-r]
	}
	fmt.Println(math.Abs(float64(a - b)))
}
