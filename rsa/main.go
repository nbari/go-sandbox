package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	private_key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Precompute some calculations -- Calculations that speed up private key operations in the future
	private_key.Precompute()

	//Validate Private Key -- Sanity checks on the key
	if err = private_key.Validate(); err != nil {
		panic(err)
	}

	private_key_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(private_key),
		},
	)

	public_key, err := x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	if err != nil {
		panic(err)
	}

	public_key_pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: public_key,
	})

	fmt.Printf("%s%s", private_key_pem, public_key_pem)

	key := [][]byte{private_key_pem, public_key_pem}
	ioutil.WriteFile("key.pem", bytes.Join(key, []byte("")), 0644)
}
