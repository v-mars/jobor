package convert

import (
	"fmt"
	"time"
)

func TimeToString(t time.Time) (string, error) {
	//loc, _ := time.LoadLocation("PRC")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-05-31 09:22:19", loc)
	//fmt.Println(t)
	//fmt.Sprintln(t.Format("2006-01-02 15:04:05"))
	return fmt.Sprintln(t.Format("2006-01-02 15:04:05")), nil
}

// 输出:
// 2018-05-31 09:22:19 +0800 CST
// 2018-05-31 09:22:19

//var cstSh, _ = time.LoadLocation("Asia/Shanghai")
//var a = pods.Items[i].PodStatus.StartTime.In(cstSh)

func NowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowTimeNoFormatStr() string {
	return time.Now().Format("20060102150405")
}
