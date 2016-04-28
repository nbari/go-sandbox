package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64

	password = "hello"
)

func main() {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", hash)
}
