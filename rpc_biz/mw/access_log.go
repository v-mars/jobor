package mw

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func AccessLogMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		start := time.Now()
		//fmt.Printf("request: %v\n", request)
		defer func() {
			un := ""
			if un == "" {
				un = "Anonymous"
			}
			end := time.Now()
			latency := end.Sub(start).String()
			//fmt.Println("ctx:", ctx)
			klog.CtxDebugf(ctx, "cost=%s", latency)
		}()
		err := next(ctx, request, response)
		//fmt.Printf("response: %v", response)
		return err
	}
}
