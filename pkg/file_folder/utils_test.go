package file_folder

import (
	"fmt"
	"testing"
)

func TestKeep(t *testing.T) {
	err := KeepLastedFolder("./fss", 1, 3)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
