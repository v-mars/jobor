package conf

import (
	"fmt"
	"testing"
	"time"
)

func TestGetConf(t *testing.T) {
	var aa, err = time.ParseDuration("4d")
	fmt.Println("time:", aa.Seconds(), err)
}
