package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"jobor/pkg/logger"
	"time"
)

func ErrorLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data string
		if c.Request.Method != http.MethodGet { // 如果是post请求，则读取body
			body, err := c.GetRawData() // body 只能读一次，读出来之后需要重置下 Body
			if err != nil {logger.Error(err)}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重置body
			data = string(body)
		}
		start := time.Now()
		c.Next()
		//`用户: 张东海, 方法: PUT, URL: /api/v1/sys/usergroup, CODE: 200, 耗时: 2ms, Body数据: {"name":"g12","description":"","users":[1,2,3s":[6,8],"id":16,"owner_id":1},`
		errMsg, ok := c.Get("errorMsg")
		if ok {
			cost := time.Since(start)
			//stack := response.Stack(0) // 3
			//fmt.Println("stack:", string(stack))
			userName, ok := c.Get("nickname")
			if !ok {userName = "nil"}
			logMsg :=fmt.Sprintf(
				"用户: %s, 方法: %s, URL: %s, CODE: %d, 耗时: %dms, Body数据: %s, \nERROR: %s",
				userName,c.Request.Method,c.Request.URL.Path,c.Writer.Status(),cost.Milliseconds(),data, errMsg)
			//stack, stackOk := c.Get("stack")
			//if stackOk{
			//	logMsg = fmt.Sprintf("%s,\n堆栈信息: \n%s", logMsg, stack)
			//}
			logger.Error(logMsg)
		}


		//logMsg := fmt.Sprintf("方法: %s, URL: %s, CODE: %d, 用时: %dms",
		//	c.Request.Method, c.Request.URL, c.Writer.Status(), cost.Milliseconds())
		//if c.Request.Method == http.MethodGet {
		//	logger.Infof("%s",logMsg)
		//} else {
		//	logger.Infof("%s, body数据: %s",logMsg, data)
		//}
	}
}
