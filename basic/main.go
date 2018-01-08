package main

import "fmt"

func foo(n int, a ...string) {
	fmt.Printf("len(a): %+v ", len(a))
	fmt.Printf("a: %+v", a)
	println()
}

func main() {
	foo(0)
	foo(0, "foo")
	foo(0, "foo", "*")
}
