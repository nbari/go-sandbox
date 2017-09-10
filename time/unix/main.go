package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Printf("now = %+v\n", now)
	fmt.Printf("now = %T\n", now)
}
