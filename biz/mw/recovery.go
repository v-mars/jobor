package mw

import (
	"context"
	"fmt"
	"jobor/biz/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Recovery() app.HandlerFunc {
	return recovery.Recovery(
		recovery.WithRecoveryHandler(func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
			response.FailedCodeRecovery(c, ctx, int(response.Err_ServerInternalErr), fmt.Errorf("[Recovery]: %v", err), fmt.Errorf("%s", string(stack)))
			ctx.Next(c)
		}))
}
