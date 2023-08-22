package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"jobor/biz/response"
	"jobor/kitex_gen/user"
	"jobor/pkg/utils"
)

func HandlerFuncWithRole(allowRoles []string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userValue, err := user.GetUserSession(c, false)
		if err != nil {
			response.SendBaseResp(ctx, c, err)
			return
		}

		if len(utils.Intersect(utils.Union([]string{admin, root, JoborAdmin, JoborRW}, allowRoles), userValue.Roles)) > 0 {
			c.Next(ctx)
			return
		} else {
			response.SendBaseResp(ctx, c, response.UnauthorizedErr)
			c.Abort()
			return
		}
	}
}

var (
	//AdminRoles = []string{"jobor_admin", "jobor_rw"}
	JoborAdmin       = "jobor_admin"
	JoborRW          = "jobor_rw"
	JoborR           = "jobor_r"
	AmsRw            = "jobor_ams_rw"
	AmsR             = "jobor_ams_r"
	AssetRW          = "jobor_asset_rw"
	AssetR           = "jobor_asset_r"
	JoborInfoRw      = "jobor_info_rw"
	JoborInfoR       = "jobor_info_r"
	SysRw            = "jobor_sys_rw"
	SysR             = "jobor_sys_r"
	AmsRolesRW       = []string{AmsRw}
	AmsRolesR        = append([]string{AmsR, JoborR}, AmsRolesRW...)
	AssetRolesRW     = []string{AssetRW}
	AssetRolesR      = append([]string{AssetR, JoborR}, AssetRolesRW...)
	JoborInfoRolesRW = []string{JoborInfoRw}
	JoborInfoRolesR  = append([]string{JoborInfoR, JoborR}, JoborInfoRolesRW...)
	SysRolesRW       = []string{SysRw}
	SysRolesR        = append([]string{SysR, JoborR}, SysRolesRW...)
)
