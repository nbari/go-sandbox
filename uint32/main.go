package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x uint32 = 4294967290
	for i := 0; i < 10; i++ {
		atomic.AddUint32(&x, 1)
		fmt.Printf("x = %+v\n", x)
	}

}
