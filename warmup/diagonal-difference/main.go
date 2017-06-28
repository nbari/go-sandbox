package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	values := readInput()
	var a, b int
	for k, v := range values {
		a = a + v[k]
		b = b + v[(len(v)-1)-k]
	}
	fmt.Println(math.Abs(float64(a - b)))
}

func readInput() [][]int {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	values := [][]int{}
	for r := 0; r < n; r++ {
		row := make([]int, n)
		for i := range row {
			fmt.Fscan(io, &row[i])
		}
		values = append(values, row)
	}
	return values
}
