package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func SieveOfEratosthenes(value int) int {
	sum := 0
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			//	fmt.Printf("%v ", i)
			sum += i
		}
	}
	//	fmt.Println("")
	return sum
}

func main() {
	sum := SieveOfEratosthenes(20000000)
	fmt.Printf("sum = %+v\n", sum)
}
