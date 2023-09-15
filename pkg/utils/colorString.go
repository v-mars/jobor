package utils

import (
	"fmt"
	_ "fmt"
	"jobor/pkg/convert"
	"strings"
	_ "strings"
)

// 绿色字体，modifier里，第一个控制闪烁，第二个控制下划线
func Green(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 0, modifier...)
}

// 淡绿
func LightGreen(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 1, modifier...)
}

// 青色/蓝绿色
func Cyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 0, modifier...)
}

// 淡青色
func LightCyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 1, modifier...)
}

// 红字体
func Red(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 0, modifier...)
}

// 淡红色
func LightRed(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 1, modifier...)
}

// 黄色字体
func Yellow(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

// 黑色
func Black(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 0, modifier...)
}

// 深灰色
func DarkGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 1, modifier...)
}

// 浅灰色
func LightGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 0, modifier...)
}

// 白色
func White(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 1, modifier...)
}

// 蓝色
func Blue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 0, modifier...)
}

// 淡蓝
func LightBlue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 1, modifier...)
}

// 紫色
func Purple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 0, modifier...)
}

// 淡紫色
func LightPurple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 1, modifier...)
}

// 棕色
func Brown(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

func cliColorRender(str string, color int, weight int, extraArgs ...interface{}) string {
	//闪烁效果
	isBlink := 0
	if len(extraArgs) > 0 {
		isBlink = convert.ToInt(extraArgs[0])
	}
	//下划线效果
	isUnderLine := 0
	if len(extraArgs) > 1 {
		isUnderLine = convert.ToInt(extraArgs[1])
	}
	var mo []string
	if isBlink > 0 {
		mo = append(mo, "05")
	}
	if isUnderLine > 0 {
		mo = append(mo, "04")
	}
	if weight > 0 {
		mo = append(mo, fmt.Sprintf("%d", weight))
	}
	if len(mo) <= 0 {
		mo = append(mo, "0")
	}
	return fmt.Sprintf("\033[%s;%dm"+str+"\033[0m", strings.Join(mo, ";"), color)
}

//fmt.Println(Green("闪烁字体：Green", 1, 1))
//fmt.Println(LightGreen("闪烁字体：LightGreen", 1))
//fmt.Println(Cyan("闪烁字体：Cyan", 1))
//fmt.Println(LightCyan("闪烁字体：LightCyan", 1))
//fmt.Println(Red("闪烁字体：Red", 1))
//fmt.Println(LightRed("闪烁字体：LightRed", 1))
//fmt.Println(Yellow("闪烁字体：Yellow", 1))
//fmt.Println(Black("闪烁字体：Black", 1))
//fmt.Println(DarkGray("闪烁字体：DarkGray", 1))
//fmt.Println(LightGray("闪烁字体：LightGray", 1))
//fmt.Println(White("闪烁字体：White", 1))
//fmt.Println(Blue("闪烁字体：Blue", 1))
//fmt.Println(LightBlue("闪烁字体：LightBlue", 1))
//fmt.Println(Purple("闪烁字体：Purple", 1))
//fmt.Println(LightPurple("闪烁字体：LightPurple", 1))
//fmt.Println(Brown("闪烁字体：Brown", 1))
//fmt.Println(Green("字体：Green"))
//fmt.Println(LightGreen("字体：LightGreen"))
//fmt.Println(Cyan("字体：Cyan"))
//fmt.Println(LightCyan("字体：LightCyan"))
//fmt.Println(Red("字体：Red"))
//fmt.Println(LightRed("字体：LightRed"))
//fmt.Println(Yellow("字体：Yellow"))
//fmt.Println(Black("字体：Black"))
//fmt.Println(DarkGray("字体：DarkGray"))
//fmt.Println(LightGray("字体：LightGray"))
//fmt.Println(White("字体：White"))
//fmt.Println(Blue("字体：Blue"))
//fmt.Println(LightBlue("字体：LightBlue"))
//fmt.Println(Purple("字体：Purple"))
//fmt.Println(LightPurple("字体：LightPurple"))
//fmt.Println(Brown("字体：Brown"))
//fmt.Println(Blue("字体：Blue", 1, 1))
