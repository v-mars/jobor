package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
)

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

const Base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
const Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const HexChars = "0123456789abcdef"
const DecChars = "0123456789"

// Base64 generates a random base64 string with length of n
func Base64(n int) string { return RandStr(n, Base64Chars) }

// Base62 generates a random base62 string with length of n
func Base62(s int) string { return RandStr(s, Base62Chars) }

// Dec generates a random decimal number string with length of n
func Dec(n int) string { return RandStr(n, DecChars) }

// Hex generates a random hex string with length of n
// e.g: 67aab2d956bd7cc621af22cfb169cba8
func Hex(n int) string { return RandStr(n, HexChars) }

// list of default letters that can be used to make a random string when calling String
// function with no letters provided
var defLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// String generates a random string using only letters provided in the letters parameter
// if user ommit letters parameters, this function will use defLetters instead
func RandStr(n int, letters ...string) string {
	var letterRunes []rune
	if len(letters) == 0 {
		letterRunes = defLetters
	} else {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))
	// on each loop, generate one random rune and append to output
	for i := 0; i < n; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%l])
	}
	return bb.String()
}
