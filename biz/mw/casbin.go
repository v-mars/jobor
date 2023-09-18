package mw

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/casbin"
	"jobor/biz/model"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	hcasbin "github.com/hertz-contrib/casbin"
)

const (
	admin = "admin"
	root  = "root"
	dom   = conf.Dom
)

// CasbinMw
// 权限检查中间件
// PermissionMiddleware 权限中间件
func CasbinMw(skipper ...SkipperFunc) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 白名单
		if len(skipper) > 0 && skipper[0](ctx, c) {
			c.Next(ctx)
			return
		} else {
			//c.Next(ctx)
			userValue, err := model.GetUserValue(c, false)
			if err != nil {
				response.SendBaseResp(ctx, c, err)
				return
			}
			isRoot := utils.Intersect([]string{admin, root}, userValue.Roles)
			//fmt.Println("intersects:", isRoot)
			if len(isRoot) > 0 {
				c.Next(ctx)
				return
			}
			var obj = c.Request.URI().Path() // path
			//var obj=c.Request.Request.URI().RequestURI() // path
			var act = c.Request.Method() // method
			var sub = userValue.Username
			isPass, err := casbin.Enforcer.Enforce(sub, dom, obj, act)
			if err != nil {
				response.SendBaseResp(ctx, c, err)
				return
			}
			//fmt.Println("isPass:", isPass, userValue, sub, obj, act)
			if isPass {
				c.Next(ctx)
			} else {
				//c.Next()
				response.SendBaseResp(ctx, c, fmt.Errorf("没有访问权限"))
			}
			return
		}
	}
}

func PermissionMwWithRole(skipper ...SkipperFunc) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 白名单
		if len(skipper) > 0 && skipper[0](ctx, c) {
			c.Next(ctx)
			return
		} else {
			//c.Next(ctx)
			userValue, err := model.GetUserSession(c, false)
			if err != nil {
				response.SendBaseResp(ctx, c, err)
				return
			}
			isRoot := utils.Intersect([]string{admin, root}, userValue.Roles)
			//fmt.Println("intersects:", isRoot)
			if len(isRoot) > 0 {
				c.Next(ctx)
				return
			}
			//var obj = c.Request.URI().Path() // path
			////var obj=c.Request.Request.URI().RequestURI() // path
			//var act = c.Request.Method() // method
			//var sub = userValue.Username
			//isPass, err := casbin.Enforcer.Enforce(sub, dom, obj, act)
			//if err != nil {
			//	response.NoPermission(ctx, c, err)
			//	return
			//}
			////fmt.Println("isPass:", isPass, userValue, sub, obj, act)
			//if isPass {
			//	c.Next(ctx)
			//} else {
			//	//c.Next()
			//	response.NoPermission(ctx, c, fmt.Errorf("没有访问权限"))
			//}
			return
		}
	}
}

func getLookupHandler(ctx context.Context, c *app.RequestContext) string {
	// 获取访问实体
	//session := sessions.Default(c)
	val, exists := c.Get("name")
	if subject, ok := val.(string); !ok || !exists {
		return ""
	} else {
		return subject
	}
}

func aa() {
	casbinMiddleware, err := hcasbin.NewCasbinMiddlewareFromEnforcer(casbin.Enforcer, getLookupHandler)
	if err != nil {
		hlog.Fatal(err)
	}
	_ = casbinMiddleware
}
