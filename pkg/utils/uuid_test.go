package utils

import (
	"fmt"
	"testing"
)

func TestNewMD5(t *testing.T) {
	fmt.Println(NewMD5(UUID{16}, []byte("1233")))
}
