// Code generated by hertz generator.

package tool

import (
	"context"
	"github.com/hertz-contrib/jwt"
	"jobor/biz/dal/db"
	"jobor/biz/model"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/audit"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
)

// GetRoutes .
//
//	@router	/routes [GET]
func GetRoutes(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req tool.RoutesReq
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	response.SendBaseResp(ctx, c, err)
	//	return
	//}

	response.SuccessMsg(ctx, c, "db migrate is success", "")
}

// GetMigrate .
//
//	@router	/api/v1/jobor/migrate [GET]
func GetMigrate(ctx context.Context, c *app.RequestContext) {
	var err error
	//var req tool.RoutesReq
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	response.SendBaseResp(ctx, c, err)
	//	return
	//}
	err = db.DB.AutoMigrate(
		&model.JoborTask{},
		&model.JoborWorker{},
		&model.JoborLog{},
		&audit.AuditLog{},
		&model.Api{},
		&model.Role{},
		&model.User{},
	)
	if err != nil {
		//hlog.Fatal("auto migrate sys err:", err)
		response.SendBaseResp(ctx, c, err)
		return
	}
	hlog.CtxInfof(ctx, "%s db migrate is success", conf.AppName)
	response.SuccessMsg(ctx, c, "db migrate is success", "")
}

// GetStateCode .
//
//	@router	/api/v1/jobor/state-code [GET]
func GetStateCode(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req tool.RoutesReq
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	response.SendBaseResp(ctx, c, err)
	//	return
	//}
	//hlog.CtxDebugf(ctx, "not add users: %v", notUser)
	response.SendDataResp(ctx, c, response.Succeed, response.StateCodeList)
}

// GenJwtToken .
//
//	@router	/api/v1/jobor/gen-token [POST]
func GenJwtToken(ctx context.Context, c *app.RequestContext) {
	userinfo, err := model.GetUserSession(c, false)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	if !userinfo.IsAdmin() {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, response.UnauthorizedErr)
		return
	}
	hlog.CtxDebugf(ctx, "用户权限效验成功")
	type param struct {
		Uid     int    `json:"uid,required"`
		Key     string `json:"key,required"`
		Timeout uint32 `json:"timeout,required"`
	}
	var req param
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	hlog.CtxDebugf(ctx, "解析请求参数成功")
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "jwt",
		Key:     []byte(req.Key),
		Timeout: time.Duration(req.Timeout) * time.Hour, // 用于设置 token 过期时间，默认为一小时
		// 用于设置最大 token 刷新时间，允许客户端在 TokenTime + MaxRefresh 内刷新 token 的有效时间，追加一个 Timeout 的时长
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		//TokenLookup:   "", // 用于设置 token 的获取源，可以选择 header、query、cookie、param、form，默认为 header:Authorization
		//TokenHeadName: "", // 用于设置从 header 中获取 token 时的前缀，默认为 Bearer
		// 用于设置登陆成功后为向 token 中添加自定义负载信息的函数
	})
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	type Data struct {
		Token  string    `json:"token"`
		User   string    `json:"user"`
		Role   []string  `json:"role"`
		Expiry time.Time `json:"expiry"`
	}
	var data Data
	info, err := model.GetUserinfoById(req.Uid, false)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	info.Roles, err = model.GetUserRoles(info.Username)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	hlog.CtxDebugf(ctx, "获取生成token的用户信息成功")
	data.Token, data.Expiry, err = authMiddleware.TokenGenerator(info)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		response.SendBaseResp(ctx, c, err)
		return
	}
	data.User = info.Username
	data.Role = info.Roles
	hlog.CtxDebugf(ctx, "生成token成功")
	response.SendDataResp(ctx, c, response.Succeed, data)
}
