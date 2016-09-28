package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var exitCh chan int = make(chan int)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		i := rand.Int()
		fmt.Println(i)
		if i%3 == 0 {
			exit(0)
		}
		if i%2 == 0 {
			fmt.Println("status 1")
			exit(1)
		}
		time.Sleep(time.Second)
	}
}

func exit(code int) {
	if flag.Lookup("test.coverprofile") != nil {
		exitCh <- code
		runtime.Goexit()
	} else {
		os.Exit(code)
	}
}
