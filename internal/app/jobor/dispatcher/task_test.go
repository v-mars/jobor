package dispatcher

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGetById(t *testing.T) {
	s:="fadasfk:    a"
	taskRespCode, err := strconv.Atoi(strings.TrimLeft(s[len(s)-5:], " "))
	if err != nil {
		err = fmt.Errorf("change str to int failed: %s",err)
		fmt.Println(err)
	}
	fmt.Println("taskRespCode:", taskRespCode)

}

