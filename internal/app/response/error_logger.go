package response

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"runtime"
	"strings"
)



func ErrLog(c *gin.Context, msg string)  {
	//start := time.Now()
	//cost := time.Since(start)
	//userName, ok := c.Get("nickname")
	//if !ok {userName = "nil"}
	//stack := Stack(3)
	//errMsg :=fmt.Sprintf("用户: %s, 方法: %s, URL: %s, CODE: %d, 耗时: %dms, Body数据: %s, \nERROR: %s, \n堆栈信息: \n%s",
	//	userName,c.Request.Method,c.Request.URL.Path,c.Writer.PodStatus(),cost.Milliseconds(),data, err,stack)

	//logger.Errorf("method: %s url: %s, msg: %s, \nStack: %s", c.Request.Method, c.FullPath(), msg, stack)
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)


// Stack returns a nicely formatted Stack frame, skipping skip frames.
func Stack(skip int) []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, _ = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in Stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}



//   err stack
func caller() (file string, line int) {
	skipCaller, skipPath := 4, 4
	_, file, line, ok := runtime.Caller(skipCaller)
	if !ok {
		return "<???>", 0
	}
	separator := "/"
	subFile := strings.Split(file, separator)
	if len(subFile) > skipPath {
		subFile = subFile[len(subFile)-skipPath:]
	}
	return strings.Join(subFile, separator), line
}