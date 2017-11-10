package main

import "fmt"

func main() {
	elements := [...]string{2: "foo", 4: "bar"}

	fmt.Printf("elements = %+v\n", elements)
	fmt.Printf("len(elements) = %+v\n", len(elements))
	for k, v := range elements {
		fmt.Printf("k: %v v: %v\n", k, v)
	}
}
