package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Println("stderr")
	log.SetOutput(os.Stdout)
	log.Println("stdout")
}
