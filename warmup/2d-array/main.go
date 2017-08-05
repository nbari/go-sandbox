package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func main() {
	var n int
	m := [6][6]int{}

	io := bufio.NewReader(os.Stdin)

	// fill the matrix
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Fscan(io, &n)
			m[i][j] = n
		}
	}

	sums := []int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s := 0
			for l := 0; l < 3; l++ {
				if l == 0 {
					//	fmt.Printf("-- m[i][j] = %+v\n", m[i][j:3+j])
					s += sum(m[i][j : 3+j])
				} else if l == 1 {
					//	fmt.Printf("-- m[i][j] = %+v\n", m[i+l][j+1:2+j])
					s += sum(m[i+l][j+1 : 2+j])
				} else if l == 2 {
					//	fmt.Printf("-- m[i][j] = %+v\n", m[i+l][j:3+j])
					s += sum(m[i+l][j : 3+j])
				}
			}
			sums = append(sums, s)
		}
	}

	max := sums[0]
	for _, v := range sums {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%d\n", max)
}
