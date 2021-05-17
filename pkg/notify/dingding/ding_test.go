package dingding

import (
	"fmt"
	"testing"
)

func TestNewDing(t *testing.T) {
	notify := NewDing(
		"https://oapi.dingtalk.com/robot/send?access_token=xxx",
		2,
		"xxx")
	err := notify.Send([]string{}, "dingtitle", "test dingding")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
