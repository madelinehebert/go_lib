package go_lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Function to encrypt data
func encrypt_text(key []byte, text []byte) []byte {
	//generate new aes cypher
	cypher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//Make new GCM
	gcm, err := cipher.NewGCM(cypher)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	return gcm.Seal(nonce, nonce, text, nil)
}
