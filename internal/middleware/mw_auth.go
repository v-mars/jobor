package middleware

import (
	"github.com/gin-gonic/gin"
	"jobor/internal/response"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单
		for _,skipper :=range skippers{
			if skipper(c){
				c.Next()
				return
			}
		}

		userInfo, tokenErr := TokenData(c, "Authorization") // AUTHORIZATION Authorization
		//sData, sessionErr := SessionData(c) // sessionErr
		//fmt.Println("UserAuthMiddleware:", userInfo, tokenErr)
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