package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"time"
)

func main() {
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject:      pkix.Name{Organization: []string{"localhost"}},
		NotBefore:    time.Now().AddDate(0, 0, 7),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:     []string{"localhost"},
	}
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	crt, err := x509.CreateCertificate(rand.Reader,
		&template,
		&template,
		&privatekey.PublicKey,
		privatekey)
	if err != nil {
		panic(err)
	}
	var certOut, keyOut bytes.Buffer
	pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: crt})
	pem.Encode(&keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)})
	ioutil.WriteFile("/tmp/key", keyOut.Bytes(), 0644)
	ioutil.WriteFile("/tmp/pem", certOut.Bytes(), 0644)
}
