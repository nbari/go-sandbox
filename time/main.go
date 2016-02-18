package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location:", t.Location(), ":Time:", t)
}
