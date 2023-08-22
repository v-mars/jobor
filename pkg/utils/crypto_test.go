package utils

import (
	"fmt"
	"testing"
)

func TestDecryptDES(t *testing.T) {
	//TestAll()
	//AesDemo2()
	aa := DeTxtByAes("15bbddabdf3113087b72336ac2b3063c", "xxx")
	fmt.Println("aa:", aa)
}
