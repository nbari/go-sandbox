package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	fmt.Println("Location:", t.Location(), ":Time:", t.Format(time.RFC3339Nano))
	fmt.Println("UnixNano:", t.UnixNano())
}
