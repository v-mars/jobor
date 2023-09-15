package mw

import (
	"context"
	"jobor/conf"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func AccessLog() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
		//spanid1, _ := c.Value("Spanid").(string)
		//spanid2, _ := c.Value("SpanId").(string)
		//spanid3, _ := c.Value("spanid").(string)
		//fmt.Println("spanid:", spanid1, spanid3, spanid2)
		c.Set(conf.RequestIDHeaderValue, reqId)
		//reqId := ""
		defer func() {
			un := c.GetString("username")
			if un == "" {
				un = "Anonymous"
			}
			end := time.Now()
			latency := end.Sub(start).String()
			hlog.CtxDebugf(ctx, "username=%s status=%d cost=%s method=%s full_path=%s client_ip=%s host=%s user_agent=%s request_id=%s",
				un, c.Response.StatusCode(), latency,
				c.Request.Header.Method(), c.Request.URI().RequestURI(), c.ClientIP(), c.Request.Host(),
				c.UserAgent(), reqId)
		}()
		c.Next(ctx)
		if c.Response.Header.Get("Server") == "hertz" {
			c.Response.Header.Del("Server")
		}
	}
}
