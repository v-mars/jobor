package mw

import (
	"context"
	"jobor/biz/response"
	"jobor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// ReqAesDec header:  X-Enc-Data = yes
func ReqAesDec() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		encryptTag := ctx.Request.Header.Get("X-Enc-Data")
		if encryptTag == "yes" {
			rawData := ctx.GetRawData() // body 只能读一次，读出来之后需要重置下 Body
			data := utils.DeTxtByAes(string(rawData), response.DefaultAesKey)
			ctx.Request.ResetBody()
			ctx.Request.SetBodyRaw([]byte(data)) // 重置body
			ctx.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
		// 处理请求
		ctx.Next(c)
	}
}
