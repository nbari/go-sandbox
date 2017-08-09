package main

import "fmt"

func main() {
	sum := 0
	for num := 1; num <= 100; num++ {
		isPrime := true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime && num != 1 {
			sum += num
		}
	}
	fmt.Printf("Sum of prime numbers is = %d\n", sum)
}
