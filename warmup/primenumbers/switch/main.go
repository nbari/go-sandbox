package main

import (
	"fmt"
	"time"
)

func isPrime(number int) bool {
	switch {
	case number < 2:
		return false
	case number == 2:
		return true
	case number%2 == 0:
		return false
	default:
		for i := 3; (i * i) <= number; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

func main() {
	t0 := time.Now()
	prime := 0
	for i := 1; i < 10000; i++ {
		if isPrime(i) {
			//	fmt.Println(i)
			prime += i
		}
	}
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Printf("prime = %+v\n", prime)
}
