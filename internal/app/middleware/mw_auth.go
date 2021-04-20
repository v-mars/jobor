package middleware

import (
	"github.com/gin-gonic/gin"
	"jobor/internal/app/response"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//path := c.Request.URL.Path
		//fmt.Println("UserAuthMiddleware:", path)
		//fmt.Println("skipper:", skipper, len(skipper), skipper[0](c))

		// 白名单
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		userInfo, tokenErr := TokenData(c, "Authorization") // AUTHORIZATION Authorization
		//sData, sessionErr := SessionData(c) // sessionErr
		//fmt.Println("UserAuthMiddleware:", tokenErr, sData, sessionErr)
		if tokenErr!=nil {
			response.FailedCode(c,999,tokenErr)
			return
		}else {
			c.Set("userInfo", userInfo)
			c.Set("nickname", userInfo["nickname"])
			c.Set("username", userInfo["username"])
			c.Set("email", userInfo["email"])
			c.Set("roles", userInfo["roles"])
			c.Set("dom", "sys")
			c.Next()
			return
		}
	}
}

