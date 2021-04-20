package schedule

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app/controllers"
	"jobor/internal/app/controllers/sys/user"
	"jobor/internal/app/models"
	"jobor/internal/app/models/db"
	"jobor/internal/app/models/tbs"
	"jobor/internal/app/response"
	"jobor/internal/redis"
	"jobor/pkg/convert"
	"strings"
)

type ITask interface {
	controllers.CommonInterfaces
	RunOrStop(c *gin.Context)
}
type Task struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) ITask {
	if DB==nil{DB= db.DB
	}
	return Task{DB: DB}
}


// @Tags Jobor任务管理
// @Summary Jobor任务列表
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// @Param Name query string false "Jobor任务名"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/task [get]
func (r Task) Query(c *gin.Context) {
	var obj []tbs.JoborTask
	var pageData = response.InitPageData(c, &obj, false)
	type Param struct {
		Name       string `form:"Name"` // `form:"Name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = r.Option()
	o.Where = "Name like ?"
	o.Value = append(o.Value, "%"+param.Name+"%")
	o.Order = "id desc"
	//o.Scan = true
	err := models.Query(r.DB,&obj, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		response.PageSuccess(c, pageData)
	}
}

// @Tags Jobor任务管理
// @Summary 创建Jobor任务
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body  tbs.Task true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/task [post]
func (r Task) Create(c *gin.Context) {
	u, err:= user.GetUserValue(c)
	if err!=nil{
		response.Error(c, err)
		return
	}
	var obj PostSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Name = strings.ToLower(obj.Name)
	obj.ByUpdate = u.Nickname
	obj.Status = "running"
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()

	if err:= models.Create(tx, &obj,true);err!=nil{
		response.Error(c, err)
		return
	}

	if err := redis.Publish(c,PubSubChannel, Event{TaskID: obj.ID, TE: AddEvent});err!=nil{
		response.Error(c, err)
		return
	}
	response.CreateSuccess(c, obj)
}

// @Tags Jobor任务管理
// @Summary 更新Jobor任务
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务id"
// @Param payload body PutSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/task/{id} [put]
func (r Task) Update(c *gin.Context) {
	u, err:= user.GetUserValue(c)
	if err!=nil{
		response.Error(c, err)
		return
	}
	var obj PutSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	if obj.Name!=nil{
		*obj.Name = strings.ToLower(*obj.Name)
	}
	obj.ByUpdate = &u.Nickname
	_id := c.Params.ByName("id")
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	var ass []models.Association
	var MapData map[string]interface{}
	if MapData,err=convert.StructToMap(obj); err!=nil{
		response.Error(c, err)
		return
	}
	var res tbs.JoborTask
	if err:= models.UpdateById(tx, &res,_id,MapData,ass, true);err!=nil{
		response.Error(c, err)
		return
	}

	if err := redis.Publish(c,PubSubChannel, Event{TaskID: res.ID, TE: ChangeEvent});err!=nil{
		response.Error(c, err)
		return
	}
	response.UpdateSuccess(c, res)
}

// @Tags Jobor任务管理
// @Summary 删除Jobor任务
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// //@Param payload body [] true "id list"
// @Param id path int true "Jobor任务id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// //@Router /api/v1/jobor/task [delete]
// @Router /api/v1/jobor/task/{id} [delete]
func (r Task) Delete(c *gin.Context) {
	var obj tbs.JoborTask
	_id := c.Params.ByName("id")
	if err := redis.Publish(c,PubSubChannel, Event{TaskID: convert.ToUint(_id), TE: DeleteEvent});err!=nil{
		response.Error(c, err)
		return
	}
	if err:= r.DB.Where("id in (?)", _id).Delete(&obj).Error;err!=nil{
		response.Error(c, err)
		return
	}

	response.DeleteSuccess(c)
}

func (r Task) Get(c *gin.Context) {
	var o models.Option
	o.Select = "id,Name,deleted,by_update,ctime,mtime"
	//o.Select = "*"
	//o.Joins = ""
	o.Order = "ID DESC"

}

// @Tags Jobor任务管理
// @Summary 停止Jobor任务
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务id"
// @Param status path string true "Jobor任务状态"
// //@Param payload body PutSchema true "参数信息[stop, run]"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/task/{id}/{status} [put]
func (r Task) RunOrStop(c *gin.Context) {
	_id := c.Params.ByName("id")
	status := c.Params.ByName("status")
	te:=AddEvent
	if status=="stop"{te=KillEvent}
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	var ass []models.Association
	var MapData = map[string]interface{}{"status": status}
	var res tbs.JoborTask
	if err:= models.UpdateById(tx, &res,_id,MapData,ass, true);err!=nil{
		response.Error(c, err)
		return
	}
	if err := redis.Publish(c,PubSubChannel, Event{TaskID: res.ID, TE: te});err!=nil{
		response.Error(c, err)
		return
	}
	response.UpdateSuccess(c, res)

}

func (r Task) Option() models.Option {
	var o models.Option
	//o.Select = "id,name,lang,data,user_id,count,expr,timeout,status,prev,next,by_update,ctime,mtime"
	//o.Select = "*"
	//o.Joins = ""
	o.Order = "ID DESC"
	return o
}

func GetById(id interface{}) (tbs.JoborTask, error) {
	var res tbs.JoborTask
	err:= db.DB.First(&res,id).Error
	return res,err
}

func GetAllRunningTask() ([]tbs.JoborTask, error) {
	var resList []tbs.JoborTask
	err:= db.DB.Where("status='running'").Find(&resList).Error
	return resList,err
}