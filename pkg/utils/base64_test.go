package utils

import (
	"fmt"
	"testing"
)

func TestBase64Enc(t *testing.T) {
	fmt.Println(Base64Enc([]byte("127.0.0.1:8080")))
}

func TestBase64Des(t *testing.T) {
	fmt.Println(Base64Dec("xx"))
}
