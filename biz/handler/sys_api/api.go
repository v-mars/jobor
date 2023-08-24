// Code generated by hertz generator.

package sys_api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	sys_api "jobor/kitex_gen/sys_api"
)

// TriggerUpdateApi .
// @router /api/v1/sys/api-auto-update [GET]
func TriggerUpdateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiQuery
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}

// GetApiAll .
// @router /api/v1/sys/apis [GET]
func GetApiAll(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiQuery
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}

// GetApi .
// @router /api/v1/sys/api [GET]
func GetApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiQuery
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}

// PostApi .
// @router /api/v1/sys/api [POST]
func PostApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiPost
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}

// PutApi .
// @router /api/v1/sys/api/:id [PUT]
func PutApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiPut
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteApi .
// @router /api/v1/sys/api/:id [DELETE]
func DeleteApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys_api.ApiDel
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys_api.ApiResp)

	c.JSON(consts.StatusOK, resp)
}