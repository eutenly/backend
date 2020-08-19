package encryption

import (
	ecies "github.com/ecies/go"
)

func Encrypt(textToEncrypt string, key *ecies.PrivateKey) string {
	ciphertext, err := ecies.Encrypt(key.PublicKey, []byte(textToEncrypt))
	if err != nil {
		panic(err)
	}
	return encodeBytes(ciphertext)
}
