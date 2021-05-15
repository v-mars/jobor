package jobor

import (
	"errors"
	"fmt"
	"strings"
)

var (
	xin    = "*"
	wenhao = "?"
	dao    = "-"
	mei    = "/"
	huo    = ","
)

func main() {
	cron := "* 5-8 1,5,7 * * ?"
	s, _ := descCorn(cron)
	fmt.Println(s)
}

func descCorn(cronExp string) (str string, err error) {

	tmpCorns := strings.Split(cronExp, " ")
	if len(tmpCorns) != 6 {
		err = errors.New("长度必须为6")
		return
	}
	var sBuffer strings.Builder
	// 解析月
	descMonth(tmpCorns[4], &sBuffer)
	// 解析周
	descWeek(tmpCorns[5], &sBuffer)
	// 解析日
	descDay(tmpCorns[3], &sBuffer)
	// 解析时
	descHour(tmpCorns[2], &sBuffer)
	// 解析分
	descMintue(tmpCorns[1], &sBuffer)
	// 解析秒
	descSecond(tmpCorns[0], &sBuffer)
	sBuffer.WriteString("执行")
	str = sBuffer.String()
	return
}

func descSecond(s string, sb *strings.Builder) {
	_ = desc(s, sb, "秒")
}

func descMintue(s string, sb *strings.Builder) {
	_ = desc(s, sb, "分")
}

func descHour(s string, sb *strings.Builder) {
	_ = desc(s, sb, "时")
}

func descDay(s string, sb *strings.Builder) {
	_ = desc(s, sb, "天")
}

func descWeek(s string, sb *strings.Builder) {

}

func descMonth(s string, sb *strings.Builder) {
	_ = desc(s, sb, "月")
}

func desc(s string, sb *strings.Builder, danwei string) (e error) {
	if s == "1/1" {
		s = "*"
	}
	if s == "0/0" {
		s = "0"
	}
	if xin == s {
		sb.WriteString("每" + danwei)
		return
	}
	if wenhao == s {
		return
	}
	if strings.Contains(s, huo) {
		arr := strings.Split(s, huo)
		for _, x := range arr {
			if len(x) != 0 {
				sb.WriteString("第" + x + danwei + "和")
			}
		}
		sb.WriteString("的")
		return
	}
	if strings.Contains(s, dao) {
		arr := strings.Split(s, huo)
		if len(arr) != 2 {
			e = errors.New("长度不能小于2")
			return
		}
		sb.WriteString("从第" + arr[0] + danwei + "到第" + arr[1] + danwei + "每" + danwei)
		sb.WriteString("的")
		return
	}
	if strings.Contains(s, mei) {
		arr := strings.Split(s, huo)
		if len(arr) != 2 {
			e = errors.New("长度不能小于2")
			return
		}
		if arr[0] == arr[1] || arr[0] == "0" {
			sb.WriteString("每" + arr[1] + danwei)
		} else {
			sb.WriteString("每" + arr[1] + danwei + "的第" + arr[0] + danwei)
		}
		return
	}
	sb.WriteString("第" + s + danwei)
	return
}
