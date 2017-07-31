package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func GenerateSalt(size int) ([]byte, error) {
	nonce := make([]byte, size)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func main() {
	p := "secret"
	salt, _ := GenerateSalt(32)
	dk, err := scrypt.Key([]byte(p), salt, 16384, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dk = %x\n", dk)
}
