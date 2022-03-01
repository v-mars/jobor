package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app/jobor/dispatcher"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/pkg/convert"
	"time"
)

type ITaskLog interface {
	Query(c *gin.Context)
	Abort(c *gin.Context)
	Delete(c *gin.Context)
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
		Name       string `form:"name"` // `form:"name" binding:"required"`
		Lang       string `form:"lang"`
		Trigger    string `form:"trigger"`
		Worker     string `form:"worker"`
		Status     string `form:"status"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = models.Option{}
	o.Where = "name like ? and addr like ?"
	o.Value = append(o.Value, "%"+param.Name+"%","%"+param.Worker+"%")
	if len(param.Lang)>0{
		o.Where = o.Where + " and lang=?"
		o.Value = append(o.Value,param.Lang)
	}
	if len(param.Trigger)>0{
		o.Where = o.Where + " and trigger_method=?"
		o.Value = append(o.Value,param.Trigger)
	}
	if len(param.Status)>0{
		o.Where = o.Where + " and result=?"
		o.Value = append(o.Value,param.Status)
	}
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

// Delete
// @Tags Jobor任务Log
// @Summary 删除 Jobor task log
// @Description 删除 Jobor task log
// @Produce  json
// @Security ApiKeyAuth
// //@Param id path int true "Jobor任务log id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/log [delete]
func (r TaskLog) Delete(c *gin.Context) {
	_id := c.Params.ByName("id")
	//rows["rows"]
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	if err:= models.DeleteById(tx, &tbs.JoborLog{}, _id, []string{}, false); err!=nil{
		response.Error(c, err)
		return
	}
	err := tx.Commit().Error
	if err!=nil{
		response.Error(c, err)
		return
	}
	response.DeleteSuccess(c)
}