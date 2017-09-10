package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	Status, Size int
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	if w.Status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	Size, err := w.ResponseWriter.Write(data)
	w.Size += Size
	return Size, err
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func www() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := &ResponseWriter{w, 0, 0}
		http.StripPrefix("/",
			http.FileServer(http.Dir("/tmp/x")),
		).ServeHTTP(lw, r)
		log.Printf("%s [%s] %d %d %s",
			r.RemoteAddr,
			r.URL,
			lw.Status,
			lw.Size,
			time.Since(start))
	})
}

func main() {
	certPEMBlock, keyPEMBlock, err := CreateSSL()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0], err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:      ":8080",
		Handler:   www(),
		TLSConfig: tlsConfig,
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

// CreateSSL creates self signed certificate
func CreateSSL() ([]byte, []byte, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	host, err := os.Hostname()
	if err != nil {
		return nil, nil, err
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"localhost"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost", host},
	}
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	crt, err := x509.CreateCertificate(rand.Reader, &template, &template, &privatekey.PublicKey, privatekey)
	if err != nil {
		return nil, nil, err
	}
	var certOut, keyOut bytes.Buffer
	pem.Encode(&certOut, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crt},
	)
	pem.Encode(&keyOut, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)},
	)
	return certOut.Bytes(), keyOut.Bytes(), nil
}
