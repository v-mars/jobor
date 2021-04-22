package casbin

import (
	"fmt"
	"testing"
)

func TestNewCasbin(t *testing.T) {
	t.Helper()
	//t.Logf("%s < %s: %t", key1, key2, myRes)
	_, err := NewCasbin(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "ocean", "123456","127.0.0.1","3306", "ops"))
	fmt.Println("err:", err)
}