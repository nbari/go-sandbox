/*
I was able to achieve these speeds on my laptop: Intel Core i7-4500U @ 1.80GHZ
2015/08/20 09:28:05 Test 1: 664579 	took 47.0074ms
2015/08/20 09:28:05 Test 2: 5761455 	took 626.9954ms
2015/08/20 09:28:13 Test 3: 50847534 	took 8.0860009s
*/

package main

import (
	"log"
	"math"
	"time"
)

func main() {
	for i, v := range []int{100, 10000000, 100000000, 1000000000} {
		start := time.Now()
		num := countPrimes(v)
		elapsed := time.Since(start)
		log.Printf("Test %d: %v \t  took %s", i+1, num, elapsed)
	}
}

func countPrimes(limit int) int {
	// Return if less than 1
	if limit <= 1 {
		return 0
	}

	// Get the sqrt of the limit
	sqrtLimit := int(math.Sqrt(float64(limit)))

	// Create array
	numbers := make([]bool, limit)

	// Set 1 to prime
	numbers[0] = true
	numPrimes := 0

	// Count the number of olds
	if limit%2 == 0 {
		numPrimes = limit / 2
	} else {
		numPrimes = (limit + 1) / 2
	}

	// Loop through odd numbers
	for i := 3; i <= sqrtLimit; i += 2 {
		if !numbers[i] {
			for j := i * i; j < limit; j += i * 2 {
				if !numbers[j] {
					numbers[j] = true
					numPrimes -= 1
				}
			}
		}
	}

	return numPrimes
}
