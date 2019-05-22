package main

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			println("fizzBuzz")
		} else if i%3 == 0 {
			println("fizz")
		} else if i%5 == 0 {
			println("buzz")
		}
	}
}
