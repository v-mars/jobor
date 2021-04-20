package log

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app/models"
	"jobor/internal/app/models/db"
	"jobor/internal/app/models/tbs"
	"jobor/internal/app/response"
)

type ITaskLog interface {
	Query(c *gin.Context)
}
type TaskLog struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) ITaskLog {
	if DB==nil{DB= db.DB
	}
	return TaskLog{DB: DB}
}


// @Tags Jobor任务管理
// @Summary JoborLog列表
// @Description JoborLog列表
// @Produce  json
// @Security ApiKeyAuth
// @Param Name query string false "Jobor任务名"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/log [get]
func (r TaskLog) Query(c *gin.Context) {
	var obj []tbs.JoborLog
	var pageData = response.InitPageData(c, &obj, false)
	type Param struct {
		Name       string `form:"Name"` // `form:"Name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = models.Option{}
	o.Where = "Name like ?"
	o.Value = append(o.Value, "%"+param.Name+"%")
	o.Order = "id desc"
	o.Scan = true
	err := models.Query(r.DB,&obj, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		response.PageSuccess(c, pageData)
	}
}
