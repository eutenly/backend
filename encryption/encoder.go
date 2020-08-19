package encryption

import (
	"encoding/base64"
)

func encodeBytes(byteCode []byte) (text string) {
	data := byteCode
	str := base64.StdEncoding.EncodeToString(data)
	//fmt.Println(str)
	return str
}

func decodeBytes(text string) (byteCode []byte) {
	bts, _ := base64.StdEncoding.DecodeString(text)
	return bts
}
