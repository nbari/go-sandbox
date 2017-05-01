package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	privKey := pem.EncodeToMemory(privateKeyPEM)

	fmt.Printf("privKey = %s\n", privKey)

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	pubBytes := ssh.MarshalAuthorizedKey(pub)
	fmt.Printf("pubBytes = %s\n", pubBytes)
	return nil
}

func main() {
	MakeSSHKeyPair("/tmp/pub", "/tmp/priv")
}
