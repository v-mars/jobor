package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jobor/internal/app/response"
)

// PermissionMiddleware 权限中间件
func PermissionMiddleware(skipper ...SkipperFunc) gin.HandlerFunc  {
	return func(c *gin.Context) {
		// 白名单
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}else {
			response.FailedMsg(c, errors.New("没有访问权限"))
			return
		}
	}
}

