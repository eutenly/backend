package encryption

import (
	ecies "github.com/ecies/go"
)

func Decrypt(textToDecrypt string, key *ecies.PrivateKey) string {
	plaintext, err := ecies.Decrypt(key, decodeBytes(textToDecrypt))
	if err != nil {
		panic(err)
	}
	return bytesToString(plaintext)
}

func bytesToString(data []byte) string {
	return string(data[:])
}
