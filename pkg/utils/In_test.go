package utils

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {
	r, err := In([]interface{}{1, "two", 3}, "two")
	//r, err := In([]interface{}{1, "two", 3}, "two")
	fmt.Println(r)
	fmt.Println(err)
	//var arr = []string{"aaa", "bbb", "ccc"}
	//fmt.Println("a")
}
