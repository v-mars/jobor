package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"jobor/conf"
	"jobor/kitex_gen/base"
	"jobor/pkg/utils"
	"net/http"
	"regexp"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/render"
)

var DefaultAesKey = "AirkJobor1qaz*xxx"

const (
	username = "username"
)

//type Response struct {
//	Code      Code        `json:"code"`
//	Message   string      `json:"message"`
//	RequestId string      `json:"request_id"`
//	Response      interface{} `json:"data"`
//}

type Response struct {
	Code      uint32      `json:"code"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

type CmdData struct {
	Response
	ExitCode Code   `json:"exit_code"`
	Stderr   string `json:"stderr"`
	Stdout   string `json:"stdout"`
}

type PageDataList struct {
	List     interface{} `json:"list"`
	PageSize int         `json:"pageSize"`
	Page     int         `json:"page"`
	Total    int64       `json:"total"`
}

func (e Response) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.Code, e.Message)
}

func (e Response) WithMessage(msg string) Response {
	e.Message = msg
	return e
}

func InitPageData(ctx context.Context, c *app.RequestContext, list interface{}, all bool) PageDataList {
	pageSize := 10
	page := 1
	exportExcel := ""
	if c != nil {
		pageSize = GetPageLimit(c)
		page = GetPageIndex(c)
		exportExcel = GetPageExport(c)
	}
	if exportExcel == "xlsx" || exportExcel == "csv" {
		all = true
	}
	if all {
		pageSize = 0
		page = 1
	}
	var pageData = PageDataList{Page: page, PageSize: pageSize, List: list}
	return pageData
}

// EncryptSuccess 加密数据响应成功
func EncryptSuccess(ctx context.Context, c *app.RequestContext, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	if v != nil && conf.GetConf().Ext.EnableWebEnc {
		marshal, err := json.Marshal(v)
		if err != nil {
			return
		}
		v = utils.EnTxtByAes(string(marshal), DefaultAesKey)
	}
	ret := Response{Code: Err_Success, Status: "success", Message: "成功.", Data: v, RequestId: reqId}
	c.Header("x-enc-data", "yes")
	ResJSON(c, http.StatusOK, &ret)
}

func ParamFailed(ctx context.Context, c *app.RequestContext, err error, v ...interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	errStr := err.Error()
	// Key: 'QueryParam.K8sClusterID' Error:Field validation for 'K8sClusterID' failed on the 'required' tag
	comp := regexp.MustCompile(`Key: '.+' Error:Field validation for '(.+)?' failed on the '(.+)?' .+`)
	subMatches := comp.FindAllStringSubmatch(errStr, -1)
	// 报错格式化"
	if len(subMatches) > 0 {
		if len(subMatches[0]) >= 3 {
			errStr = fmt.Sprintf("请求参数验证错误：%s参数【%s】", subMatches[0][2], subMatches[0][1])
			errStr = strings.Replace(errStr, "required", "缺少", -1)
		} else {
			errStr = fmt.Sprintf("请求参数验证错误：%s", subMatches[0])
		}
	}
	ret := Response{Code: Err_ParamsErr, Status: "error", Message: errStr, Data: v, RequestId: reqId}
	c.Set("errorMsg", err)
	//stack := Stack(2)
	//logger.Error(string(stack))
	hlog.CtxErrorf(ctx, err.Error())
	ResJSON(c, http.StatusOK, &ret)
}

func SqlFailed(ctx context.Context, c *app.RequestContext, err error, v ...interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	errStr := err.Error()
	// Error 1062: Duplicate entry '1.1.1.2' for key 'idx_name_code'
	comp := regexp.MustCompile(`Error 1062: Duplicate entry '(.+)?' for key '(.+)?'`)
	subMatches := comp.FindAllStringSubmatch(errStr, -1)
	// "记录重复。具体报错："
	if len(subMatches) > 0 {
		if len(subMatches[0]) >= 2 {
			errStr = fmt.Sprintf("记录重复。具体报错：%s", subMatches[0][1])
		} else {
			errStr = fmt.Sprintf("记录重复。具体报错：%s", subMatches[0])
		}
	}
	ret := Response{Code: Err_ParamsErr, Status: "error", Message: errStr, Data: v, RequestId: reqId}
	c.Set("errorMsg", err)
	//stack := Stack(2)
	//logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

func FailedMsg(ctx context.Context, c *app.RequestContext, msg error, v ...interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: Err_BadRequest, Status: "error", Message: msg.Error(), Data: v, RequestId: reqId}
	c.Set("errorMsg", msg.Error())
	//stack := Stack(2)
	ResJSON(c, http.StatusOK, &ret)
}

// FailedHttpCode 响应失败 code
func FailedHttpCode[T int | uint32 | int64](ctx context.Context, c *app.RequestContext, code T, msg error, v ...interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: uint32(code), Status: "error", Message: msg.Error(), Data: v, RequestId: reqId}
	c.Set("errorMsg", msg.Error())
	//stack := Stack(2)
	ResJSON(c, int(code), &ret)
}

// FailedCode 响应失败 code
func FailedCode[T comparable](ctx context.Context, c *app.RequestContext, code int, msg error, v ...interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: uint32(code), Status: "error", Message: msg.Error(), Data: v, RequestId: reqId}
	c.Set("errorMsg", msg.Error())
	//stack := Stack(2)
	ResJSON(c, http.StatusOK, &ret)
}

// FailedCodeRecovery 响应失败 code
func FailedCodeRecovery[T int | uint32 | int64](ctx context.Context, c *app.RequestContext, code T, msg error, RecoveryErr error) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: uint32(code), Status: "error", Message: msg.Error(), Data: nil, RequestId: reqId}
	c.Set("errorMsg", msg.Error())
	if RecoveryErr != nil {
		c.Set("stack", fmt.Errorf("%s", RecoveryErr))
	}
	ResJSON(c, http.StatusInternalServerError, &ret)
}

func Success(ctx context.Context, c *app.RequestContext, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: Err_Success, Status: "success", Message: "ok", Data: v, RequestId: reqId}
	ResJSON(c, http.StatusOK, &ret)
}

func Download(ctx context.Context, c *app.RequestContext, fileName, filePath string) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func Res[T int | uint32 | int64](ctx context.Context, c *app.RequestContext, code T, msg string, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: uint32(code), Status: "success", Message: msg, Data: v, RequestId: reqId}
	ResJSON(c, http.StatusOK, &ret)
}

func ResCode[T int | uint32 | int64](ctx context.Context, c *app.RequestContext, code T, msg string, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: uint32(code), Status: "success", Message: msg, Data: v, RequestId: reqId}
	ResJSON(c, code, &ret)
}

func Send(ctx context.Context, c *app.RequestContext, d Response, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	d.Data = v
	d.RequestId = reqId
	ResJSON(c, http.StatusOK, d)
}

func SendMsg(ctx context.Context, c *app.RequestContext, d Response, msg string, data interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	d.Data = data
	d.Message = msg
	d.RequestId = reqId
	ResJSON(c, http.StatusOK, d)
}

func SendWithCode(ctx context.Context, c *app.RequestContext, d Response, data interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	d.Data = data
	d.RequestId = reqId
	ResJSON(c, int(d.Code), d)
}

func SendWithCodeMsg(ctx context.Context, c *app.RequestContext, d Response, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	d.Data = v
	d.RequestId = reqId
	ResJSON(c, int(d.Code), d)
}

func Raw(ctx context.Context, c *app.RequestContext, v interface{}) {
	ResJSON(c, http.StatusOK, &v)
}

func SuccessMsg(ctx context.Context, c *app.RequestContext, msg string, v interface{}) {
	reqId, _ := ctx.Value(conf.RequestIDHeaderValue).(string)
	ret := Response{Code: Err_Success, Status: "success", Message: msg, Data: v, RequestId: reqId}
	ResJSON(c, http.StatusOK, &ret)
}

func SuccessCmd(ctx context.Context, c *app.RequestContext, stdout, stderr string, exitCode int, v interface{}) {
	ret := CmdData{Response: Response{Code: Err_Success, Status: "success", Message: "ok", Data: v},
		ExitCode: Code(exitCode), Stdout: stdout, Stderr: stderr}
	ResJSON(c, http.StatusOK, &ret)
}

func Html(ctx context.Context, c *app.RequestContext, htmlContent string) {
	var templateName = "text_html"
	var pageInstance render.Render
	pageContent := template.Must(template.New(templateName).Parse(htmlContent))
	htmlRender := render.HTMLProduction{Template: pageContent}
	pageInstance = htmlRender.Instance(templateName, map[string]interface{}{})
	c.Render(http.StatusOK, pageInstance)
}

/* new response */

// SendBaseResp build baseResp from error
func SendBaseResp(ctx context.Context, c *app.RequestContext, err error) {
	if err == nil {
		c.JSON(http.StatusOK, baseResp(Succeed))
		return
	}
	//errStr := err.Error()
	//// Key: 'QueryParam.K8sClusterID' Error:Field validation for 'K8sClusterID' failed on the 'required' tag
	//comp := regexp.MustCompile(`Key: '.+' Error:Field validation for '(.+)?' failed on the '(.+)?' .+`)
	//subMatches := comp.FindAllStringSubmatch(errStr, -1)
	//// 报错格式化"
	//if len(subMatches) > 0 {
	//	if len(subMatches[0]) >= 3 {
	//		errStr = fmt.Sprintf("请求参数验证错误：%s参数【%s】", subMatches[0][2], subMatches[0][1])
	//		errStr = strings.Replace(errStr, "required", "缺少", -1)
	//	} else {
	//		errStr = fmt.Sprintf("请求参数验证错误：%s", subMatches[0])
	//	}
	//}
	//err = fmt.Errorf(errStr)
	defer func() {
		hlog.CtxErrorf(ctx, err.Error())
	}()
	c.Set("errorMsg", err)
	e := ErrNo{}
	if errors.As(err, &e) {
		c.JSON(http.StatusOK, baseResp(e))
		return
	}
	s := ServiceErr.WithMessage(err.Error())
	c.JSON(http.StatusOK, baseResp(s))
	return
}

// SendDataResp pack response
func SendDataResp(ctx context.Context, c *app.RequestContext, err error, data interface{}) {
	Err := ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
	return
}

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{Code: err.GetErrCode(), Message: err.GetErrMsg()}
}
