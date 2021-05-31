package worker

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/proto"
	"jobor/internal/response"
	"jobor/pkg/convert"
	"time"
)

type IJoborWorker interface {
	GetRoutingKey(c *gin.Context)
	Query(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type JoborWorker struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) IJoborWorker {
	if DB==nil{DB= db.DB}
	return JoborWorker{DB: DB}
}

// GetRoutingKey
// @Tags Jobor Worker管理
// @Summary Jobor routing_key列表
// @Description JoborWorker
// @Produce  json
// @Security ApiKeyAuth
// //@Param Name query string false ""
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/routing_key [get]
func (r JoborWorker)GetRoutingKey(c *gin.Context)  {
	var obj []ShowRoutingKey
	var pageData = response.PageDataList{PageNumber: 1,PageSize:0,List:&obj}
	o := models.Option{}
	o.Select = "distinct routing_key"
	o.Order = "ID DESC"
	o.Scan = true
	err := models.Query(r.DB,&tbs.JoborWorker{}, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		response.Success(c, pageData)
	}

}

// Query
// @Tags Jobor Worker管理
// @Summary Jobor worker列表
// @Description JoborWorker
// @Produce  json
// @Security ApiKeyAuth
// @Param Name query string false "JoborHostname"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/worker [get]
func (r JoborWorker) Query(c *gin.Context) {
	var obj []tbs.JoborWorker
	var pageData = response.InitPageData(c, &obj, false)
	type Param struct {
		Hostname       string `form:"hostname"` // `form:"Name" binding:"required"`
		Addr           string `form:"addr"` // `form:"Name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = r.Option()
	o.Where = "hostname like ? and addr like ?"
	o.Value = append(o.Value, "%"+param.Hostname+"%","%"+param.Addr+"%")
	o.Order = "id desc"
	err := models.Query(r.DB,&tbs.JoborWorker{}, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		leaseUpdate :=time.Now().Unix() - int64(proto.DefaultHeartbeatInterval.Seconds())
		for i,v:=range obj{
			if v.LeaseUpdate < leaseUpdate && v.Status==tbs.TaskLogStatusRunning{
				obj[i].Status = tbs.WorkerStatusOffline
			}
		}
		response.Success(c, pageData)
	}
}

// Update
// @Tags Jobor Worker管理
// @Summary 删除JoborWorker
// @Description JoborWorker
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/worker/{id} [put]
func (r JoborWorker) Update(c *gin.Context) {
	var obj update
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	var ass []models.Association
	var MapData map[string]interface{}
	var err error
	if MapData,err=convert.StructToMap(obj); err!=nil{
		response.Error(c, err)
		return
	}
	_id := c.Params.ByName("id")
	var res tbs.JoborWorker
	if err:= models.UpdateById(tx, &res,_id,MapData,ass, true);err!=nil{
		response.Error(c, err)
		return
	}

	response.UpdateSuccess(c, res)
}

// Delete
// @Tags Jobor Worker管理
// @Summary 删除JoborWorker
// @Description JoborWorker
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/worker/{id} [delete]
func (r JoborWorker) Delete(c *gin.Context) {
	var obj tbs.JoborWorker
	_id := c.Params.ByName("id")
	if err:= r.DB.Where("id in (?)", _id).Delete(&obj).Error;err!=nil{
		response.Error(c, err)
		return
	}

	response.DeleteSuccess(c)
}

func (r JoborWorker) Option() models.Option {
	var o models.Option
	//o.Select = "id,name,lang,data,user_id,count,expr,timeout,status,prev,next,by_update,ctime,mtime"
	//o.Select = "*"
	//o.Joins = ""
	o.Order = "ID DESC"
	return o
}

func CreateOrUpdate(data tbs.JoborWorker) error {
	var exist tbs.JoborWorker
	//if err:=db.DB.Model(&tbs.JoborWorker{}).Where(tbs.JoborWorker{Addr: data.Addr}).FirstOrCreate(&data).Error;err!=nil{
		//return err
	//}
	if err:=db.DB.Model(&tbs.JoborWorker{}).Where(tbs.JoborWorker{Addr: data.Addr}).First(&exist).Error;errors.Is(
		err, gorm.ErrRecordNotFound){
		if err:=db.DB.Model(&tbs.JoborWorker{}).Create(&data).Error;err!=nil{
			return err
		}
	}else if err!=nil{
		return err
	}else {
		if err:=db.DB.Model(&tbs.JoborWorker{}).Where("addr=? and status!=?",data.Addr,tbs.WorkerStatusStop).Updates(&data).Error;err!=nil{
			return err
		}
	}

	return nil
}
