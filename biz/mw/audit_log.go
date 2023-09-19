// Package mw
// 审计日志
// 目前支持POST、PUT、DELETE
// POST 使用 c.Set("post_id",id)
// PUT 支持在body中id参数的自动获取，如没有id时，可以使用c.Set("put_id",id)
// DELETE 支持在body中id参数的自动获取，如没有id时，可以使用c.Set("delete_id",id)
package mw

import (
	"context"
	"fmt"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/audit"
	"jobor/pkg/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tidwall/gjson"
)

var (
	noAuditList = []string{
		"/api/v1.0/user/get_token_by_param",
		"/api/v1/login", "/api/v1/logout",
	}
)

// AuditLog 审计日志对象
type AuditLog struct {
	Action   string `json:"action"`   //操作
	Business string `json:"business"` //业务
}

// AuditLogHandler 审计日志处理函数
func AuditLogHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		reqBody := c.Request.Body()
		definePath := c.FullPath()
		defer func() {
			end := time.Now()
			//latency := end.Sub(startTime).String()
			requestMethod := string(c.Request.Method())
			if requestMethod == "" {
				hlog.CtxErrorf(ctx, "request method is null")
				return
			}
			for _, v := range noAuditList {
				_ = v
			}
			isAudit, _, title, group := GetInfoFromPath(definePath, requestMethod)
			if !isAudit || utils.InOfStr(string(c.Request.Path()), noAuditList) {
				return
			}
			handler := fmt.Sprintf("%s%s", string(c.Request.Host()), string(c.Request.URI().RequestURI()))
			auditModel := new(audit.AuditLog)
			auditModel.UserIp = GetReqClientIp(ctx, c)
			auditModel.Action = title
			auditModel.Business = group
			auditModel.Dom = conf.Dom
			auditModel.Method = requestMethod
			auditModel.Handler = handler
			auditModel.CostTime = db.MillisTime(end.Sub(start).String())
			// get userid 方法暂未封装 使用 auditModel.UserId进行赋值
			auditModel.User = c.GetString("username")
			auditModel.Nickname = c.GetString("nickname")
			auditModel.UserId = c.GetInt64("uid")
			if len(string(reqBody)) > 0 {
				auditModel.Body = string(reqBody)
			}
			if len(string(c.Response.Body())) > 0 {
				auditModel.RespBody = string(c.Response.Body())
			}
			auditModel.HttpCode = c.Response.StatusCode()

			switch requestMethod {
			case http.MethodPut:
				putAuditHandler(ctx, c, auditModel)
			case http.MethodPost:
				postAuditHandler(ctx, c, auditModel)
			case http.MethodDelete:
				deleteAuditHandler(ctx, c, auditModel)
			}
			statusAuditHandler(ctx, c, auditModel)
			err := auditModel.CreateAuditLog()
			if err != nil {
				hlog.CtxErrorf(ctx, "create audit log err, %s", err)
				return
			}

		}()
		c.Next(ctx)
	}
}

// put处理函数
func putAuditHandler(ctx context.Context, c *app.RequestContext, auditModel *audit.AuditLog) {
	_id := c.Params.ByName("id")
	objId := gjson.Parse(auditModel.Body).Get("id")
	if len(_id) > 0 {
		auditModel.ObjId, _ = strconv.ParseInt(_id, 0, 0)
	} else if objId.Exists() {
		auditModel.ObjId = objId.Int()
	} else {
		objId, ok := c.Get("id")
		if ok {
			auditModel.ObjId = objId.(int64)
		}
	}
	return
}

// post处理函数
func postAuditHandler(c context.Context, ctx *app.RequestContext, auditModel *audit.AuditLog) {
	if len(auditModel.RespBody) > 0 {
		objId := gjson.Parse(auditModel.RespBody).Get("data.id")
		if objId.Exists() {
			auditModel.ObjId = objId.Int()
		} else {
			objId2 := gjson.Parse(auditModel.RespBody).Get("id")
			if objId2.Exists() {
				auditModel.ObjId = objId2.Int()
			}
		}
	}
}

// delete处理函数
func deleteAuditHandler(ctx context.Context, c *app.RequestContext, auditModel *audit.AuditLog) {
	_id := c.Params.ByName("id")
	objId := gjson.Parse(auditModel.Body).Get("id")
	if len(_id) > 0 {
		auditModel.ObjId, _ = strconv.ParseInt(_id, 0, 0)
	} else if objId.Exists() {
		auditModel.ObjId = objId.Int()
	} else {
		objId, ok := c.Get("id")
		if ok {
			auditModel.ObjId = objId.(int64)
		}
	}
}

// status处理函数
func statusAuditHandler(ctx context.Context, c *app.RequestContext, auditModel *audit.AuditLog) {
	code := gjson.Parse(auditModel.RespBody).Get("code")
	if code.Exists() {
		if response.Code(code.Int()) == response.Err_Success {
			auditModel.Status = 1
		} else if response.Code(code.Int()) == response.Err_Unauthenticated || response.Code(code.Int()) == response.Err_Unauthorized {
			auditModel.Status = 2
		} else if int64(response.Err_ServerInternalErr) == code.Int() {
			auditModel.Status = 3
		} else {
			auditModel.Status = 0
		}
	} else {
		if 200 <= auditModel.HttpCode && auditModel.HttpCode < 300 {
			auditModel.Status = 1
		} else if 500 <= auditModel.HttpCode && auditModel.HttpCode < 600 {
			auditModel.Status = 3
		} else {
			auditModel.Status = 0
		}
	}
	return
}

func GetReqClientIp(ctx context.Context, c *app.RequestContext) string {
	remoteIP := c.ClientIP()
	xff := c.GetHeader("X-Forwarded-For")
	if string(xff) != "" {
		remoteIP = string(xff)
	}
	return remoteIP
}

func GetInfoFromPath(path string, method string) (audit bool, name, title, group string) {
	pathArray := strings.Split(strings.Trim(path, "/"), "/")
	title = ""
	group = conf.AppName
	audit = false
	if strings.ToLower(method) != "get" {
		audit = true
	}
	if len(pathArray) >= 5 {
		group = pathArray[2]
		title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(method)], pathArray[3])
		name = fmt.Sprintf("%s:%s:%s:%s", strings.ToLower(method), pathArray[2], pathArray[3], pathArray[4])
	} else if len(pathArray) >= 4 {
		group = pathArray[2]
		title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(method)], pathArray[3])
		name = fmt.Sprintf("%s:%s:%s", strings.ToLower(method), pathArray[2], pathArray[3])
	} else if len(pathArray) == 3 {
		group = pathArray[2]
		title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(method)], pathArray[2])
		name = fmt.Sprintf("%s:%s", strings.ToLower(method), pathArray[2])
	} else if len(pathArray) == 2 {
		title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(method)], pathArray[1])
		name = fmt.Sprintf("%s:%s", strings.ToLower(method), pathArray[1])
	} else if len(pathArray) == 1 {
		title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(method)], pathArray[0])
		name = fmt.Sprintf("%s:%s", strings.ToLower(method), pathArray[0])
	}
	return audit, name, title, group
}

var ActionTitle = map[string]string{
	"get":    "查看",
	"post":   "创建",
	"put":    "修改",
	"patch":  "修补",
	"delete": "删除",
}
