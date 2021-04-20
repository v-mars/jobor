package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jobor/pkg/logger"
	"net/http"
	"regexp"
	"strings"
)

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}


type PageDataList struct {
	List       interface{} `json:"list"`
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
	Total      int64       `json:"total"`
}

func InitPageData(c *gin.Context,list interface{}, all bool) PageDataList {
	pageSize := 10
	pageNumber := 1
	if c!=nil{
		pageSize = GetPageLimit(c)
		pageNumber = GetPageIndex(c)
	}
	if all{
		pageSize = 0
	}
	var pageData = PageDataList{PageNumber: pageNumber,PageSize:pageSize,List:list}
	return pageData
}


// 响应成功
func PageSuccess(c *gin.Context, v PageDataList) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 创建响应成功
func CreateSuccess(c *gin.Context, v...interface{}) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: "创建成功.", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 更新响应成功
func UpdateSuccess(c *gin.Context, v...interface{}) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: "更新成功.", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 删除响应成功
func DeleteSuccess(c *gin.Context, v...interface{}) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: "删除成功.", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func Error(c *gin.Context, err error, v...interface{}) {
	ret := Data{Code: FAIL_CODE,Status:"error", Message: err.Error(), Data: v}
	c.Set("errorMsg", err)
	stack:=Stack(2)
	logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

func ParamFailed(c *gin.Context, err error, v...interface{}) {
	errStr := err.Error()
	// Key: 'QueryParam.K8sClusterID' Error:Field validation for 'K8sClusterID' failed on the 'required' tag
	comp := regexp.MustCompile(`Key: '.+' Error:Field validation for '(.+)?' failed on the '(.+)?' .+`)
	subMatches := comp.FindAllStringSubmatch(errStr, -1)
	// 报错格式化"
	if len(subMatches) > 0{
		if len(subMatches[0]) >= 3{
			errStr = fmt.Sprintf("请求参数验证错误：%s参数【%s】", subMatches[0][2],subMatches[0][1])
			errStr = strings.Replace(errStr, "required", "缺少", -1)
		}else {
			errStr = fmt.Sprintf("请求参数验证错误：%s", subMatches[0])
		}
	}
	ret := Data{Code: FAIL_CODE,Status:"error", Message: errStr, Data: v}
	c.Set("errorMsg", err)
	stack:=Stack(2)
	logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

func SqlFailed(c *gin.Context, err error, v...interface{}) {
	errStr := err.Error()
	// Error 1062: Duplicate entry '1.1.1.2' for key 'idx_name_code'
	comp := regexp.MustCompile(`Error 1062: Duplicate entry '(.+)?' for key '(.+)?'`)
	subMatches := comp.FindAllStringSubmatch(errStr, -1)
	// "记录重复。具体报错："
	if len(subMatches) > 0{
		if len(subMatches[0]) >= 2{
			errStr = fmt.Sprintf("记录重复。具体报错：%s", subMatches[0][1])
		}else {
			errStr = fmt.Sprintf("记录重复。具体报错：%s", subMatches[0])
		}
	}
	ret := Data{Code: FAIL_CODE,Status:"error", Message: errStr, Data: v}
	c.Set("errorMsg", err)
	stack:=Stack(2)
	logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

func NoPermission(c *gin.Context, v...interface{}) {
	ret := Data{Code: ErrNoPerm,Status:"error", Message: "无权限访问", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

func FailedMsg(c *gin.Context, msg error, v...interface{}) {
	ret := Data{Code: FAIL_CODE,Status:"error", Message: msg.Error(), Data: v}
	c.Set("errorMsg", msg.Error())
	stack:=Stack(2)
	logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败 code
func FailedCode(c *gin.Context,code int, msg error, v...interface{}) {
	ret := Data{Code: code,Status:"error", Message: msg.Error(), Data: v}
	c.Set("errorMsg", msg.Error())
	stack:=Stack(2)
	logger.Error(string(stack))
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败 code
func FailedCodeRecovery(c *gin.Context,code int, msg error, RecoveryErr error) {
	ret := Data{Code: code,Status:"error", Message: msg.Error(), Data:nil}
	c.Set("errorMsg", msg.Error())
	if RecoveryErr!=nil{
		c.Set("stack", fmt.Errorf("%s",RecoveryErr))
	}
	ResJSON(c, http.StatusInternalServerError, &ret)
}

func Success(c *gin.Context, v interface{}) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

func Res(c *gin.Context, code int,msg string, v interface{}) {
	ret := Data{Code: code,Status:"success", Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

func SuccessMsg(c *gin.Context, msg string, v interface{}) {
	ret := Data{Code: SUCCESS_CODE,Status:"success", Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}


//func ErrLog(c *gin.Context, msg string)  {
//	stack:=middleware.Stack(3)
//	logger.Errorf("method: %s url: %s, msg: %s, \nStack: %s", c.Request.Method, c.FullPath(), msg, stack)
//}