package log

import (
	"jobor/internal/app/jobor/dispatcher"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/pkg/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ITaskLog interface {
	Query(c *gin.Context)
	Abort(c *gin.Context)
}
type TaskLog struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) ITaskLog {
	if DB==nil{DB= db.DB
	}
	return TaskLog{DB: DB}
}

// Query
// @Tags Jobor任务Log
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
		for i,v:=range obj{
			if v.Result == tbs.TaskLogStatusRunning || v.Result == tbs.TaskLogStatusWait{
				obj[i].CostTime = float32(time.Now().Unix() - v.StartTime.Unix())
			}
		}
		response.Success(c, pageData)
	}
}

// Abort
// @Tags Jobor任务Log
// @Summary Abort Jobor task
// @Description JoborLog列表
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务log id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/log/{id}/abort [post]
func (r TaskLog) Abort(c *gin.Context) {
	_id := c.Params.ByName("id")
	_,ok:=dispatcher.CacheTask.TaskLogs[convert.ToUint(_id)]
	if !ok{
		response.Error(c, fmt.Errorf("任务[%s]已经完成或不存在",_id), "")
		return
	}
	dispatcher.CacheTask.TaskLogs[convert.ToUint(_id)].TaskCancel()
	response.SuccessMsg(c, "任务终止成功", "")
}