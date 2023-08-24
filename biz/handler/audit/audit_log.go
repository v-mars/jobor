// Code generated by hertz generator.

package audit

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	audit "jobor/kitex_gen/audit"
)

// GetLoginHistory .
// @router /api/v1/jobor/audit-log [GET]
func GetLoginHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req audit.AuditReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(audit.AuditResp)

	c.JSON(consts.StatusOK, resp)
}