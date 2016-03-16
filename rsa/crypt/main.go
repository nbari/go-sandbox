package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/ascii85"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var pem_file_path string
	var err error
	var block *pem.Block
	var private_key *rsa.PrivateKey
	var public_key *rsa.PublicKey
	var pem_data, plain_text, encrypted, decrypted, label []byte

	plain_text = []byte("Plain text message to be encrypted")

	// A PEM file can contain a Private key among others (Public certificate, Intermidiate Certificate, Root certificate, ...)
	pem_file_path = "../key.pem"
	if pem_data, err = ioutil.ReadFile(pem_file_path); err != nil {
		log.Fatalf("Error reading pem file: %s", err)
	}

	//Package pem implements the PEM data encoding, most commonly used in TLS keys and certificates.
	//Decode will find the next PEM formatted block (certificate, private key etc) in the input.
	//Expected Block type "RSA PRIVATE KEY"
	//http://golang.org/pkg/encoding/pem/
	if block, _ = pem.Decode(pem_data); block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("No valid PEM data found")
	}

	//x509 parses X.509-encoded keys and certificates.
	//ParsePKCS1PrivateKey returns an RSA private key from its ASN.1 PKCS#1 DER encoded form.
	if private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		log.Fatalf("Private key can't be decoded: %s", err)
	}

	public_key = &private_key.PublicKey

	private_key.Precompute()

	encrypted = encrypt_oaep(public_key, plain_text, label)
	decrypted = decrypt_oaep(private_key, encrypted, label)

	fmt.Printf("OAEP Encrypted [%s] to \n[%x]\n", string(plain_text), encrypted)
	fmt.Printf("OAEP Decrypted [%x] to \n[%s]\n", encrypted, decrypted)

	fmt.Printf("PLAIN TEXT: %s\n", string(plain_text))
	fmt.Printf("OAEP Encrypted HEX len: %d\n%x\n", len(fmt.Sprintf("%x", encrypted)), encrypted)
	base64_out_URL := base64.URLEncoding.EncodeToString(encrypted)
	base64_out_STD := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Printf("OAEP Encrypted Base64 Std len: %d\n%s\n", len(base64_out_URL), base64_out_URL)
	fmt.Printf("OAEP Encrypted Base64 URL len: %d\n%s\n", len(base64_out_STD), base64_out_STD)
	base85_buffer := make([]byte, ascii85.MaxEncodedLen(len(encrypted)))
	encodedbytes := ascii85.Encode(base85_buffer, encrypted)
	fmt.Printf("OAEP Encrypted Base85 len: %d\n%s\n", encodedbytes, base85_buffer)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("OAEP Decrypted [%x] to \n[%s]\n", encrypted, decrypted)

	return
}

//OAEP Encrypt
func encrypt_oaep(public_key *rsa.PublicKey, plain_text, label []byte) (encrypted []byte) {
	var err error
	if encrypted, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, public_key, plain_text, label); err != nil {
		log.Fatal(err)
	}
	return
}

func decrypt_oaep(private_key *rsa.PrivateKey, encrypted, label []byte) (decrypted []byte) {
	var err error
	if decrypted, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, private_key, encrypted, label); err != nil {
		log.Fatal(err)
	}
	return
}
