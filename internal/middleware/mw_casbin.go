package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jobor/internal/app/sys/user"
	"jobor/internal/response"
	"jobor/internal/utils/casbin"
	"jobor/pkg/utils"
)

// 权限检查中间件
// PermissionMiddleware 权限中间件
func CasbinMiddlewareV2(allowSkipper, allowNoSkipper SkipperFunc) gin.HandlerFunc  {
	return func(c *gin.Context) {
		// 白名单
		if allowSkipper(c) || allowNoSkipper(c) {
			c.Next()
			return
		}else {
			userValue, err:= user.GetUserValue(c)
			if err!=nil{
				response.Error(c, err)
				return
			}
			isRoot := utils.Intersect([]string{"admin", "root"}, userValue.Roles)
			//fmt.Println("intersects:", isRoot)
			if len(isRoot)>0{
				c.Next()
				return
			}
			var dom = "sys"
			var obj=c.Request.URL.Path // path
			//var obj=c.Request.URL.RequestURI() // path
			var act=c.Request.Method   // method
			var sub= userValue.Username
			isPass, err := casbin.Enforcer.Enforce(sub, dom, obj, act)
			if err!=nil{
				response.NoPermission(c, err)
				return
			}
			//fmt.Println("isPass:", isPass, userValue, sub, obj, act)
			if isPass{
				c.Next()
			}else {
				//c.Next()
				response.NoPermission(c, fmt.Errorf("没有访问权限"))
			}
			return
		}
	}
}

// 权限检查中间件
// PermissionMiddleware 权限中间件
func CasbinMiddleware(skippers ...SkipperFunc) gin.HandlerFunc  {
	return func(c *gin.Context) {
		// 白名单
		for _,skipper :=range skippers{
			if skipper(c){
				c.Next()
				return
			}
		}
		userValue, err:= user.GetUserValue(c)
		if err!=nil{
			response.Error(c, err)
			return
		}
		isRoot := utils.Intersect([]string{"admin", "root"}, userValue.Roles)
		//fmt.Println("intersects:", isRoot)
		if len(isRoot)>0{
			c.Next()
			return
		}
		var dom = "sys"
		var obj=c.Request.URL.Path // path
		//var obj=c.Request.URL.RequestURI() // path
		var act=c.Request.Method   // method
		var sub= userValue.Username
		isPass, err := casbin.Enforcer.Enforce(sub, dom, obj, act)
		if err!=nil{
			response.NoPermission(c, err)
			return
		}
		//fmt.Println("isPass:", isPass, userValue, sub, obj, act)
		if isPass{
			c.Next()
		}else {
			//c.Next()
			response.NoPermission(c, fmt.Errorf("没有访问权限"))
		}
		return

	}
}