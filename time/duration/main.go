package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	d, err := time.ParseDuration("1m30s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m.Seconds() = %+v\n", d.Seconds())
	fmt.Printf("d = %+T\n", d)
}
