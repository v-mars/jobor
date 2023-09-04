package utils

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var literals = strings.NewReplacer(
	"SUN", "0", "MON", "1", "TUE", "2", "WED", "3", "THU", "4", "FRI", "5", "SAT", "6",
	"JAN", "1", "FEB", "2", "MAR", "3", "APR", "4", "MAY", "5", "JUN", "6", "JUL", "7",
	"AUG", "8", "SEP", "9", "OCT", "10", "NOV", "11", "DEC", "12",
)

var expressions = map[string]string{
	"@yearly":    "0 0 1 1 *",
	"@annually":  "0 0 1 1 *",
	"@monthly":   "0 0 1 * *",
	"@weekly":    "0 0 * * 0",
	"@daily":     "0 0 * * *",
	"@hourly":    "0 * * * *",
	"@always":    "* * * * *",
	"@5minutes":  "*/5 * * * *",
	"@10minutes": "*/10 * * * *",
	"@15minutes": "*/15 * * * *",
	"@30minutes": "0,30 * * * *",

	"@everysecond": "* * * * * *",
}

// SpaceRe is regex for whitespace.
var SpaceRe = regexp.MustCompile(`\s+`)
var yearRe = regexp.MustCompile(`\d{4}`)

func normalize(expr string) []string {
	expr = strings.Trim(expr, " \t")
	if e, ok := expressions[strings.ToLower(expr)]; ok {
		expr = e
	}

	expr = SpaceRe.ReplaceAllString(expr, " ")
	expr = literals.Replace(strings.ToUpper(expr))

	return strings.Split(strings.ReplaceAll(expr, "  ", " "), " ")
}

// IsDue checks if cron expression is due for given reference time (or now).
// It returns bool or error if any.
func IsDue(expr string, ref ...time.Time) (bool, error) {
	ref = append(ref, time.Now())

	segs, err := Segments(expr)
	if err != nil {
		return false, err
	}

	return SegmentsDue(segs)
}

// Segments splits expr into array array of cron parts.
// If expression contains 5 parts or 6th part is year like, it prepends a second.
// It returns array or error.
func Segments(expr string) ([]string, error) {
	segs := normalize(expr)
	slen := len(segs)
	if slen < 5 || slen > 7 {
		return []string{}, errors.New("expr should contain 5-7 segments separated by space")
	}

	// Prepend second if required
	prepend := slen == 5 || (slen == 6 && yearRe.MatchString(segs[5]))
	if prepend {
		segs = append([]string{"0"}, segs...)
	}

	return segs, nil
}

// SegmentsDue checks if all cron parts are due.
// It returns bool. You should use IsDue(expr) instead.
func SegmentsDue(segs []string) (bool, error) {
	for pos, seg := range segs {
		if seg == "*" || seg == "?" {
			continue
		}

		if due, err := CheckDue(seg, pos); !due {
			return due, err
		}
	}

	return true, nil
}

// IsValid checks if cron expression is valid.
// It returns bool.
func IsValidExpression(expr string) bool {
	segs, err := Segments(expr)
	if err != nil {
		return false
	}

	for pos, seg := range segs {
		if _, err := CheckDue(seg, pos); err != nil {
			return false
		}
	}

	return true
}

func valueByPos(ref time.Time, pos int) (val int) {
	switch pos {
	case 0:
		val = ref.Second()
	case 1:
		val = ref.Minute()
	case 2:
		val = ref.Hour()
	case 3:
		val = ref.Day()
	case 4:
		val = int(ref.Month())
	case 5:
		val = int(ref.Weekday())
	case 6:
		val = ref.Year()
	}
	return
}

func isOffsetDue(offset string, val, pos int) (bool, error) {
	if offset == "*" || offset == "?" {
		return true, nil
	}

	bounds, isWeekDay := boundsByPos(pos), pos == 5
	if strings.Contains(offset, "/") {
		return inStep(val, offset, bounds)
	}
	if strings.Contains(offset, "-") {
		if isWeekDay {
			offset = strings.Replace(offset, "7-", "0-", 1)
		}
		return inRange(val, offset, bounds)
	}

	if !isWeekDay && (val == 0 || offset == "0") {
		return offset == "0" && val == 0, nil
	}

	nval, err := strconv.Atoi(offset)
	if err != nil {
		return false, err
	}
	if nval < bounds[0] || nval > bounds[1] {
		return false, fmt.Errorf("segment#%d: '%s' out of bounds(%d, %d)", pos, offset, bounds[0], bounds[1])
	}

	return nval == val || (isWeekDay && nval == 7 && val == 0), nil
}

func CheckDue(segment string, pos int) (due bool, err error) {
	ref, last := time.Now(), -1
	val, loc := valueByPos(ref, pos), ref.Location()
	isMonth, isWeekDay := pos == 3, pos == 5

	for _, offset := range strings.Split(segment, ",") {
		mod := (isMonth || isWeekDay) && strings.ContainsAny(offset, "LW#")
		if due, err = isOffsetDue(offset, val, pos); due || (!mod && err != nil) {
			return
		}
		if !mod {
			continue
		}
		if last == -1 {
			last = time.Date(ref.Year(), ref.Month(), 1, 0, 0, 0, 0, loc).AddDate(0, 1, 0).Add(-time.Nanosecond).Day()
		}
		if isMonth {
			due, err = isValidMonthDay(offset, last, ref)
		} else if isWeekDay {
			due, err = isValidWeekDay(offset, last, ref)
		}
		if due || err != nil {
			return due, err
		}
	}

	return false, nil
}

func boundsByPos(pos int) (bounds []int) {
	bounds = []int{0, 0}
	switch pos {
	case 0, 1:
		bounds = []int{0, 59}
	case 2:
		bounds = []int{0, 23}
	case 3:
		bounds = []int{1, 31}
	case 4:
		bounds = []int{1, 12}
	case 5:
		bounds = []int{0, 7}
	case 6:
		bounds = []int{1, 9999}
	}
	return
}

func inStep(val int, s string, bounds []int) (bool, error) {
	parts := strings.Split(s, "/")
	step, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}
	if step == 0 {
		return false, errors.New("step can't be 0")
	}

	if strings.Index(s, "*/") == 0 || strings.Index(s, "0/") == 0 {
		return val%step == 0, nil
	}

	sub, end := strings.Split(parts[0], "-"), val
	start, err := strconv.Atoi(sub[0])
	if err != nil {
		return false, err
	}

	if len(sub) > 1 {
		end, err = strconv.Atoi(sub[1])
		if err != nil {
			return false, err
		}
	}

	if (len(sub) > 1 && end < start) || start < bounds[0] || end > bounds[1] {
		return false, fmt.Errorf("step '%s' out of bounds(%d, %d)", parts[0], bounds[0], bounds[1])
	}

	return inStepRange(val, start, end, step), nil
}

func inRange(val int, s string, bounds []int) (bool, error) {
	parts := strings.Split(s, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}

	if end < start || start < bounds[0] || end > bounds[1] {
		return false, fmt.Errorf("range '%s' out of bounds(%d, %d)", s, bounds[0], bounds[1])
	}

	return start <= val && val <= end, nil
}

func inStepRange(val, start, end, step int) bool {
	for i := start; i <= end && i <= val; i += step {
		if i == val {
			return true
		}
	}
	return false
}

func isValidMonthDay(val string, last int, ref time.Time) (valid bool, err error) {
	day, loc := ref.Day(), ref.Location()
	if val == "L" {
		return day == last, nil
	}

	pos := strings.Index(val, "W")
	if pos < 1 {
		return false, errors.New("invalid offset value: " + val)
	}

	nval, err := strconv.Atoi(val[0:pos])
	if err != nil {
		return false, err
	}

	for _, i := range []int{0, -1, 1, -2, 2} {
		incr := i + nval
		if incr > 0 && incr <= last {
			iref := time.Date(ref.Year(), ref.Month(), incr, ref.Hour(), ref.Minute(), ref.Second(), 0, loc)
			week := int(iref.Weekday())

			if week > 0 && week < 6 && iref.Month() == ref.Month() {
				valid = day == iref.Day()
				break
			}
		}
	}

	return valid, nil
}

func isValidWeekDay(val string, last int, ref time.Time) (bool, error) {
	loc := ref.Location()
	if pos := strings.Index(strings.ReplaceAll(val, "7L", "0L"), "L"); pos > 0 {
		nval, err := strconv.Atoi(val[0:pos])
		if err != nil {
			return false, err
		}

		for i := 0; i < 7; i++ {
			decr := last - i
			dref := time.Date(ref.Year(), ref.Month(), decr, ref.Hour(), ref.Minute(), ref.Second(), ref.Nanosecond(), loc)

			if int(dref.Weekday()) == nval {
				return ref.Day() == decr, nil
			}
		}

		return false, nil
	}

	pos := strings.Index(val, "#")
	parts := strings.Split(strings.ReplaceAll(val, "7#", "0#"), "#")
	if pos < 1 || len(parts) < 2 {
		return false, errors.New("invalid offset value: " + val)
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, err
	}

	nth, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}

	if day < 0 || day > 7 || nth < 1 || nth > 5 || int(ref.Weekday()) != day {
		return false, nil
	}

	return ref.Day()/7 == nth-1, nil
}

var CRON_TIME_CN = []string{"秒", "分", "点", "天", "月", "周", "年"}
var HOURS = 24
var TIMESCALE = 60

// TranslateToChinese 识别cron表达式转成中文
func TranslateToChinese(cronExp string) string {
	return translateToChinese(cronExp, CRON_TIME_CN)
}

func translateToChinese(cronExp string, cronTimeArr []string) string {
	//入参判断
	if cronExp == "" || len(cronTimeArr) < 1 {
		return ""
	}
	tmpCorns := strings.Split(cronExp, " ")
	var sBuffer bytes.Buffer //声明字节缓冲
	//stringBuilder.WriteString("哈喽")   //把字符串写入缓冲区
	//stringBuilder.WriteString("沃德")
	if len(tmpCorns) == 6 || len(tmpCorns) == 7 {
		if len(tmpCorns) == 7 {
			// 解析年 Year
			year := tmpCorns[6]
			//0 15 10 L * ? 切割后 len==7 且最后一个是空串。go和java不一样，java length==6
			if year != "*" && year != "?" && year != "" {
				sBuffer.WriteString(year)
				sBuffer.WriteString(cronTimeArr[6])
			}
		}
		// 解析月 Month 主要解析 /
		months := tmpCorns[4]
		if months != "*" && months != "?" {
			if strings.Contains(months, "/") {
				sBuffer.WriteString("从")
				sBuffer.WriteString(strings.Split(months, "/")[0])
				sBuffer.WriteString("号开始")
				sBuffer.WriteString(",每")
				sBuffer.WriteString(strings.Split(months, "/")[1])
				sBuffer.WriteString(cronTimeArr[4])
			} else {
				if strings.Contains(sBuffer.String(), "-") {
					sBuffer.WriteString("每年")
				}
				sBuffer.WriteString(months)
				sBuffer.WriteString(cronTimeArr[4])
			}
		}
		// 解析周 DayofWeek 主要解析 , - 1~7/L 1L~7L
		dayofWeek := tmpCorns[5]
		if dayofWeek != "*" && dayofWeek != "?" {
			if strings.Contains(dayofWeek, ",") {
				sBuffer.WriteString("每周的第")
				sBuffer.WriteString(dayofWeek)
				sBuffer.WriteString(cronTimeArr[3])
			} else if strings.Contains(dayofWeek, "L") && len(dayofWeek) > 1 { // 1L-7L
				weekNum := strings.Split(dayofWeek, "L")[0]
				weekName := judgeWeek(weekNum)
				sBuffer.WriteString("每月的最后一周的")
				sBuffer.WriteString(weekName)
			} else if strings.Contains(dayofWeek, "-") {
				splitWeek := strings.Split(dayofWeek, "-")
				weekOne := judgeWeek(splitWeek[0])
				weekTwo := judgeWeek(splitWeek[1])
				sBuffer.WriteString("每周的")
				sBuffer.WriteString(weekOne)
				sBuffer.WriteString("到")
				sBuffer.WriteString(weekTwo)
			} else {                    // 1-7/L
				if "L" == (dayofWeek) { // L 转换为7，便于识别
					dayofWeek = "7"
				}
				sBuffer.WriteString("每周的")
				weekName := judgeWeek(dayofWeek)
				sBuffer.WriteString(weekName)
			}
		}
		// 解析日 days -- DayofMonth 需要解析的 / # L W
		// * “6#3”表示本月第三周的星期五（6表示星期五，3表示第三周）;
		// * “2#1”表示本月第一周的星期一;
		// * “4#5”表示第五周的星期三。
		days := tmpCorns[3]
		if days != "?" && days != "*" {
			if strings.Contains(days, "/") {
				sBuffer.WriteString("每周从第")
				sBuffer.WriteString(strings.Split(days, "/")[0])
				sBuffer.WriteString("天开始")
				sBuffer.WriteString(",每")
				sBuffer.WriteString(strings.Split(days, "/")[0])
				sBuffer.WriteString(cronTimeArr[3])
			} else if "L" == days { // 处理L 每月的最后一天
				sBuffer.WriteString("每月最后一天")
			} else if "W" == days { // 处理W 暂时不懂怎么处理
				sBuffer.WriteString(days)
			} else if strings.Contains(days, "#") {
				splitDays := strings.Split(days, "#")
				weekNum := splitDays[0]     // 前面数字表示周几
				weekOfMonth := splitDays[1] // 后面的数字表示第几周 范围1-4 一个月最多4周
				weekName := judgeWeek(weekNum)
				sBuffer.WriteString("每月第")
				sBuffer.WriteString(weekOfMonth)
				sBuffer.WriteString(cronTimeArr[5])
				sBuffer.WriteString(weekName)
			} else {
				sBuffer.WriteString("每月")
				sBuffer.WriteString(days)
				sBuffer.WriteString("号")
			}
		} else {
			sBufferStr := sBuffer.String()
			if (len(sBufferStr) == 0 || len(tmpCorns) == 7) && !strings.Contains(sBufferStr, "星期") {
				// 前面没有内容的话，拼接下
				sBuffer.WriteString("每")
				sBuffer.WriteString(cronTimeArr[3])
			}
		}

		// 解析时 Hours 主要解析 /
		hours := tmpCorns[2]
		if hours != "*" {
			if strings.Contains(hours, "/") {
				sBuffer.WriteString(appendGapInfo(hours, HOURS, 2))
			} else {
				if !(len(sBuffer.String()) > 0) { // 对于 , 的情况，直接拼接
					sBuffer.WriteString("每天")
					sBuffer.WriteString(hours)
					sBuffer.WriteString(cronTimeArr[2])
				} else {
					sBuffer.WriteString(hours)
					sBuffer.WriteString(cronTimeArr[2])
				}
			}
		}

		// 解析分 Minutes 主要解析 /
		minutes := tmpCorns[1]
		if minutes != "*" {
			if strings.Contains(minutes, "/") {
				sBuffer.WriteString(appendGapInfo(minutes, TIMESCALE, 1))
			} else if minutes == "0" {
			} else if strings.Contains(minutes, "-") {
				splitMinute := strings.Split(months, "-")
				sBuffer.WriteString(splitMinute[0])
				sBuffer.WriteString(cronTimeArr[1])
				sBuffer.WriteString("到")
				sBuffer.WriteString(splitMinute[1])
				sBuffer.WriteString(cronTimeArr[1])
				sBuffer.WriteString("每分钟")
			} else {
				sBuffer.WriteString(minutes)
				sBuffer.WriteString(cronTimeArr[1])
			}
		}

		// 解析秒 Seconds 主要解析 /
		seconds := tmpCorns[0]
		if seconds != "*" {
			if strings.Contains(seconds, "/") {
				sBuffer.WriteString(appendGapInfo(seconds, TIMESCALE, 0))
			} else if seconds != "0" {
				sBuffer.WriteString(seconds)
				sBuffer.WriteString(cronTimeArr[0])
			}
		}

		if len(sBuffer.String()) > 0 {
			sBuffer.WriteString("执行")
		} else {
			sBuffer.WriteString("表达式中文转换异常")
		}

	}
	return sBuffer.String()
}

var WeekEnumMap = map[string]string{
	"0": "星期天", "1": "星期一",
	"2": "星期二", "3": "星期三",
	"4": "星期四", "5": "星期五",
	"6": "星期六", "SUN": "星期天",
	"MON": "星期一", "TUE": "星期二",
	"WED": "星期三", "THU": "星期四",
	"FRI": "星期五", "SAT": "星期六",
}

func judgeWeek(weekNum string) string {
	weekName := WeekEnumMap[weekNum]
	if weekName != "" {
		return weekName
	}
	return weekNum
}

func appendGapInfo(time string, rangeNum int, index int) string {
	var sBufferTemp bytes.Buffer
	splitTime := strings.Split(time, "/")
	//startNum := splitTime[0]
	gapNum := splitTime[1]
	splitTimeUnit := CRON_TIME_CN[index]
	/*
		atoi, err := strconv.Atoi(startNum)
		if err != nil {
			atoi = 0
		}
		atoi2, err := strconv.Atoi(gapNum)
		if err != nil {
			atoi2 = 0
		}
		endNum := rangeNum + atoi - atoi2
		endStr := strconv.Itoa(endNum)
		timeUnit := CRON_TIME_CN[index]

		if index == 1 {
			splitTimeUnit = "分钟"
		} else if index == 2 {
			splitTimeUnit = "小时"
		}

			sBufferTemp.WriteString("从")
			sBufferTemp.WriteString(startNum)
			sBufferTemp.WriteString(timeUnit)
			sBufferTemp.WriteString("开始")
			sBufferTemp.WriteString("到")
			sBufferTemp.WriteString(endStr)
			sBufferTemp.WriteString(timeUnit)
			sBufferTemp.WriteString("范围内")

	*/
	sBufferTemp.WriteString(",每隔")
	sBufferTemp.WriteString(gapNum)
	sBufferTemp.WriteString(splitTimeUnit)
	return sBufferTemp.String()
}
