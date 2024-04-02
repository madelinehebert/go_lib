package MDH

import (
	"crypto/aes"
	"crypto/cipher"
)

// Function to decrypt
func decrypt_text(key []byte, cypherText []byte) string {
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

	//Grab nonce size
	nonceSize := gcm.NonceSize()
	if len(cypherText) < nonceSize {
		panic(err)
	}

	//do a bunch of fancy stuff
	nonce, cypherText := cypherText[:nonceSize], cypherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cypherText, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
