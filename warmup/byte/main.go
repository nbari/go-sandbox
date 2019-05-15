package main

import (
	"fmt"
)

func main() {
	var b byte
	//After the b == 255 iteration, b++ is executed. This overflows (since the
	//maximum value for a byte is 255) and results in b == 0. Therefore b <= 255
	//still holds and the loop restarts from 0.
	for b = 250; b <= 255; b++ {
		fmt.Printf("%d %c\n", b, b)
	}

	//for b := byte(250); ; b++ {
	//fmt.Printf("%d %c\n", b, b)
	//if b == 255 {
	//break
	//}
	//}
}
