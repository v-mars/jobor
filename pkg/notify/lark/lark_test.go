package lark

import (
	"fmt"
	"testing"
)

func TestNewLark(t *testing.T) {
	l:=NewLark("https://open.feishu.cn/open-apis/bot/v2/hook/xxx",
		1,"xxxx")
	err := l.Send([]string{}, "lark title", "test lark text example")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
