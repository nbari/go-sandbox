package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func main() {
	// 4096

	private_key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// Precompute some calculations -- Calculations that speed up private key operations in the future
	private_key.Precompute()

	//Validate Private Key -- Sanity checks on the key
	if err = private_key.Validate(); err != nil {
		log.Fatal(err)
	}

	private_key_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(private_key),
		},
	)

	public_key, err := x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	public_key_pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: public_key,
	})

	log.Printf("\n%s\n%s", private_key_pem, public_key_pem)

	ioutil.WriteFile("public_key.pem", public_key_pem, 0644)
	ioutil.WriteFile("private_key.pem", private_key_pem, 0644)
}
