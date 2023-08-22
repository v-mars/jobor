package response

import (
	"fmt"
	"jobor/pkg/convert"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	page        = "page"
	pageSize    = "pageSize"
	exportExcel = "export_excel"
)

// ResJSON 响应JSON数据
func ResJSON[T int | uint32 | int64](c *app.RequestContext, status T, v interface{}) {
	c.JSON(int(status), v)
	c.Abort()
}

// GetPageIndex 获取页码
func GetPageIndex(c *app.RequestContext) int {
	return GetQueryToInt(c, page, 1)
}

// GetPageLimit 获取每页记录数
func GetPageLimit(c *app.RequestContext) int {
	limit := GetQueryToInt(c, pageSize, 10)
	//if limit > 500 {
	//	limit = 500
	//}
	return limit
}

// GetPageExport 获取导出参数
func GetPageExport(c *app.RequestContext) string {
	return GetQueryToStr(c, exportExcel)
}

// GetPageSort 获取排序信息
func GetPageSort(c *app.RequestContext) string {
	return GetQueryToStr(c, "sort")
}

// GetPageKey 获取搜索关键词信息
func GetPageKey(c *app.RequestContext) string {
	return GetQueryToStr(c, "key")
}

func GetQueryToStrE(c *app.RequestContext, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", fmt.Errorf("没有这个值[%s]传入", key)
	}
	return str, nil
}

func GetQueryToStr(c *app.RequestContext, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		return defaultValue
	}
	return str
}

// GetQueryToUintE QueryToUintE
func GetQueryToUintE(c *app.RequestContext, key string) (uint, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUintE(str)
}

// GetQueryToUint QueryToUint
func GetQueryToUint(c *app.RequestContext, key string, defaultValues ...uint) uint {
	var defaultValue uint
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUintE(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetQueryToUint64E QueryToUintE
func GetQueryToUint64E(c *app.RequestContext, key string) (uint64, error) {
	str, err := GetQueryToStrE(c, key)
	//fmt.Println(str)
	if err != nil {
		return 0, err
	}
	return convert.ToUint64E(str)
}

// GetQueryToIntE QueryToUintE
func GetQueryToIntE(c *app.RequestContext, key string) (int, error) {
	str, err := GetQueryToStrE(c, key)
	//fmt.Println(str)
	if err != nil {
		return 0, err
	}
	return convert.ToInt(str), nil
}

// GetQueryToUint64 QueryToUint
func GetQueryToUint64(c *app.RequestContext, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUint64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetQueryToInt QueryToInt
func GetQueryToInt(c *app.RequestContext, key string, defaultValues ...int) int {
	var defaultValue int
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToIntE(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}
