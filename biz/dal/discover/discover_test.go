package discover

import (
	"fmt"
	utils2 "jobor/pkg/utils"
	"testing"
)

func TestRegistry(t *testing.T) {
	fmt.Println(utils2.GetLocalIPv4Address())
	fmt.Println(utils2.GetOutBoundIP("tcp", "172.20.106.5:8500"))
}
