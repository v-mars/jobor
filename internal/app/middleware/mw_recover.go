package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jobor/internal/app/response"
)

// RecoveryMiddleware 崩溃恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var data string
		//if c.Request.Method != http.MethodGet { // 如果不是GET请求，则读取body
		//	body, err := c.GetRawData() 		// body 只能读一次，读出来之后需要重置下 Body
		//	if err != nil {logger.Error(err)}
		//	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重置body
		//	data = string(body)
		//}
		//start := time.Now()
		defer func() {
			if err := recover(); err != nil {
				//cost := time.Since(start)
				//userName, ok := c.Get("nickname")
				//if !ok {userName = "nil"}
				stack := response.Stack(3)
				//errMsg :=fmt.Sprintf(
				//	"用户: %s, 方法: %s, URL: %s, CODE: %d, 耗时: %dms, Body数据: %s, \nERROR: %s, \n堆栈信息: \n%s",
				//	userName,c.Request.Method,c.Request.URL.Path,c.Writer.Status(),cost.Milliseconds(),data, err,stack)
				// 这里会打印出错栈信息 time.Now().Format("2006-01-02 15:04:05")
				//logger.Error(errMsg)
				response.FailedCodeRecovery(c,5009,fmt.Errorf("[Recovery]: %s", err), fmt.Errorf("%s", stack))
			c.Abort()
			return
			}
		}()
		c.Next()
	}
}

