package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*", "POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Content-Length", "X-CSRF-Token",
			"Token,session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding",
			"Accept-Language,DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With",
			"If-Modified-Since", "Cache-Control", "Content-TypeV1", "Pragma",
			"Access-Control-Request-Private-Network", "access-control-request-private-network"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Cache-Control", "Content-Language", "Content-TypeV1", "Expires", "Last-Modified", "Pragma", "FooBar",
			"Access-Control-Request-Private-Network"},
		AllowCredentials: true,
		MaxAge:           24 * 7 * time.Hour,
	})
}

func Cors2() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		method := c.Request.Method()
		origin := c.Request.Header.Get("Origin")
		//var headerKeys []string
		//for k, _ := range ctx.Request.Header{
		//	headerKeys = append(headerKeys, k)
		//}
		//headerStr := strings.Join(headerKeys, ", ")
		//if headerStr != "" {
		//	headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		//} else {
		//	headerStr = "access-control-allow-origin, access-control-allow-headers"
		//}
		if origin != "" {
			//ctx.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-TypeV1, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-TypeV1,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                             // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                    //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                                // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//    c.JSON(http.StatusOK, "Options Request!")
		//}
		if string(method) == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next(ctx) //  处理请求
	}
}
