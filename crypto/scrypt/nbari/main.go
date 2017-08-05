package main

import (
	"fmt"
	"log"

	"github.com/nbari/crypto/scrypt"
)

func main() {
	key, err := scrypt.Create("The quick brown fox jumps over the lazy dog", 64)
	if err != nil {
		log.Fatal(err)
	}

	ok, err := scrypt.Verify("The quick brown fox jumps over the lazy dog", key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ok = %+v\n", ok)
}
