package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// d, err := time.ParseDuration("1m30s")
	d, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m.Seconds() = %+v\n", d.Seconds())
	fmt.Printf("d = %+T\n", d)
	fmt.Println("Unix:  ", time.Now().Unix())
	t := time.Now().Add(d)
	fmt.Println("Unix:+ ", t.Unix())
}
