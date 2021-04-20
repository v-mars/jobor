package utils

import "encoding/base64"

// base64编码
func Base64Enc(input []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

// base64解码
func Base64Dec(encodeString string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		panic(err)
	}
	return string(decodeBytes)
}

// base64 URL编码
func Base64EncUrl(input []byte) string {
	encodeString := base64.URLEncoding.EncodeToString(input)
	return encodeString
}

// base64 URL解码
func Base64DecUrl(encodeString string) string {
	decodeBytes, err := base64.URLEncoding.DecodeString(encodeString)
	if err != nil {
		panic(err)
	}
	return string(decodeBytes)
}