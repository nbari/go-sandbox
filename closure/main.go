package main

import "fmt"

func main() {
	echo := func(x int) {
		fmt.Printf("x = %+v\n", x)
	}
	x := 0
	if x <= 0 {
		echo(x)
	} else if x == 1 {
		echo(x)
	}
}
