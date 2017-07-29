package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("1234567890123456")
	plaintext := []byte("text can be a random lenght")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	// BTW (only for test purpose) I don't include it

	ciphertext := make([]byte, len(plaintext))

	iv := []byte{'\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f', '\x0f'}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewCTR.
	base := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("encodedHEX: %x\n", ciphertext)
	fmt.Printf("encodedBASE: %s\n", base)

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext)

	fmt.Printf("decoded: %s\n", plaintext2)
}
