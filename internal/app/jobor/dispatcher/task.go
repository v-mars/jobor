package dispatcher

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app"
	"jobor/internal/app/sys/user"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/proto"
	"jobor/internal/redis"
	"jobor/internal/response"
	"jobor/pkg/convert"
	"strings"
	"time"
)

type ITask interface {
	app.CommonInterfaces
	RunOrStopStatus(c *gin.Context)
	RunTask(c *gin.Context)
	Dashboard(c *gin.Context)
}
type Task struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) ITask {
	if DB==nil{DB= db.DB
	}
	return Task{DB: DB}
}

// Query
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
		Name       string `form:"name"`
		RoutingKey string `form:"routing_key"`
		Status     string `form:"status"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = r.Option()
	o.Where = "name like ? and routing_key like ?"
	o.Value = append(o.Value, "%"+param.Name+"%","%"+param.RoutingKey+"%")
	if len(param.Status)>0{
		o.Where = o.Where + " and status=?"
		o.Value = append(o.Value,param.Status)
	}
	o.Order = "id desc"
	//o.Scan = true
	err := models.Query(r.DB,&tbs.JoborTask{}, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		response.Success(c, pageData)
	}
}

// Create
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
	if err = c.ShouldBindJSON(&obj); err!=nil{
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

	if err := redis.Publish(c, PubSubChannel, Event{TaskID: obj.ID, TE: AddEvent});err!=nil{
		response.Error(c, err)
		return
	}
	response.CreateSuccess(c, obj)
}

// Update
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
	if obj.Data!=nil{
		bytes,err:=json.Marshal(obj.Data)
		if err!=nil{
			response.Error(c, err)
			return
		}
		MapData["data"]=string(bytes)
	}
	if obj.Notify!=nil{
		bytes,err:=json.Marshal(obj.Notify)
		if err!=nil{
			response.Error(c, err)
			return
		}
		MapData["notify"]=string(bytes)
	}

	var res tbs.JoborTask
	if err:= models.UpdateById(tx, &res,_id,MapData,ass, true);err!=nil{
		response.Error(c, err)
		return
	}

	if err := redis.Publish(c, PubSubChannel, Event{TaskID: res.ID, TE: ChangeEvent});err!=nil{
		response.Error(c, err)
		return
	}
	//err = EventFunc(Event{TaskID: res.ID, TE: ChangeEvent},res)
	//if err != nil {
	//	response.Error(c, err)
	//	return
	//}
	response.UpdateSuccess(c, res)
}

// Delete
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
	if err := redis.Publish(c, PubSubChannel, Event{TaskID: convert.ToUint(_id), TE: DeleteEvent});err!=nil{
		response.Error(c, err)
		return
	}
	//err := EventFunc(Event{TaskID: convert.ToUint(_id), TE: DeleteEvent},res)
	//if err != nil {
	//	response.Error(c, err)
	//	return
	//}
	if err:= r.DB.Where("id in (?)", _id).Delete(&obj).Error;err!=nil{
		response.Error(c, err)
		return
	}

	response.DeleteSuccess(c)
}

func (r Task) Get(c *gin.Context) {
	var o models.Option
	o.Select = "id,name,deleted,by_update,ctime,mtime"
	//o.Select = "*"
	//o.Joins = ""
	o.Order = "ID DESC"

}

// Dashboard
// @Tags Jobor管理
// @Summary Jobor大盘
// @Description Jobor大盘
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/dashboard [get]
func (r Task)Dashboard(c *gin.Context) {
	type Param struct {
		EndTime   time.Time `form:"end_time"`
		StartTime time.Time `form:"start_time"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	//fmt.Println("StartTime:", param.StartTime.Local().String())

	type res struct {
		Tasks                int64 `json:"tasks"`
		TaskLogs             int64 `json:"task_logs"`
		Workers              int64 `json:"workers"`
		TodayTaskLogs        int64 `json:"today_task_logs"`
		TaskLogStatusDay     []struct{
			Name         string `json:"name"`
			Count        int64  `json:"count"`
			FailedCount  int64  `json:"failed_count"`
			TimeoutCount int64  `json:"timeout_count"`
			SuccessCount int64  `json:"success_count"`
			Time         string `json:"time"`
		} `json:"task_log_status_day"`
		TaskRun []struct{
			OrdNum  int    `json:"ord_num"`
			Task    string `json:"task"`
			Count   int64  `json:"count"`
		} `json:"task_run"`
	}
	var data = res{}
	if err:=r.DB.Model(&tbs.JoborTask{}).Count(&data.Tasks).Error;err!=nil{
		response.Error(c, err)
		return
	}
	if err:=r.DB.Model(&tbs.JoborLog{}).Count(&data.TaskLogs).Error;err!=nil{
		response.Error(c, err)
		return
	}
	if err:=r.DB.Model(&tbs.JoborWorker{}).Count(&data.Workers).Error;err!=nil{
		response.Error(c, err)
		return
	}
	if err:=r.DB.Model(&tbs.JoborLog{}).Where("to_days(ctime) = to_days(now())").Count(&data.TodayTaskLogs).Error;err!=nil{
		response.Error(c, err)
		return
	}


	taskLogStatusSql :=fmt.Sprintf(`
SELECT
  DATE_FORMAT(d.ctime,'%%Y-%%m-%%d') as time,
	d.name,
	SUM( CASE d.result WHEN 'failed' THEN 1 ELSE 0 END ) AS failed_count,
	SUM( CASE d.result WHEN 'success' THEN 1 ELSE 0 END ) AS success_count,
	SUM( CASE d.result WHEN 'timeout' THEN 1 ELSE 0 END ) AS timeout_count,
	count(d.id) as count
FROM
	jobor_log AS d
WHERE
	d.ctime BETWEEN '%s' AND '%s'
GROUP BY
	time
	ORDER BY
	time asc
`,param.StartTime.Local().String(), param.EndTime.Local().String())
	if err:=r.DB.Raw(taskLogStatusSql).Scan(&data.TaskLogStatusDay).Error;err!=nil{
		response.Error(c, err)
		return
	}

	taskRunSql := fmt.Sprintf(`
SELECT
	t.NAME as task,
	count(log.id) AS count,
    ( @i := @i + 1 ) AS ord_num
FROM
	jobor_task as t
	LEFT JOIN jobor_log AS log ON t.id = log.task_id
	left join ( SELECT @i := 0 ) as b on 1=1
WHERE
	log.ctime BETWEEN '%s' AND '%s'
	GROUP BY t.NAME
	ORDER BY count DESC
`,param.StartTime.Local().String(), param.EndTime.Local().String())
	if err:=r.DB.Raw(taskRunSql).Scan(&data.TaskRun).Error;err!=nil{
		response.Error(c, err)
		return
	}
	response.Success(c, data)
}

// RunTask
// @Tags Jobor任务管理
// @Summary Jobor任务运行
// @Description Jobor任务
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Jobor任务id"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/jobor/task/{id}/run [post]
func (r Task) RunTask(c *gin.Context) {
	_id := c.Params.ByName("id")
	t, err := GetById(_id)
	if err!=nil{
		response.Error(c, err)
		return
	}
	go RunTasks(TriggerAuto,TriggerManual, t)
	response.SuccessMsg(c, "开始手动执行",_id)

}

// RunOrStopStatus
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
func (r Task) RunOrStopStatus(c *gin.Context) {
	_id := c.Params.ByName("id")
	status := c.Params.ByName("status")
	te:= AddEvent
	if status=="stop"{te= KillEvent
	}
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
	if err := redis.Publish(c, PubSubChannel, Event{TaskID: res.ID, TE: te});err!=nil{
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

func GetWorkers(routingKey string) ([]tbs.JoborWorker, error) {
	var workers []tbs.JoborWorker
	leaseUpdate :=time.Now().Unix() - int64(proto.DefaultHeartbeatInterval.Seconds())
	//fmt.Println("lease_update:", leaseUpdate)
	var o = models.Option{Where: "status=? and lease_update>=?",
		Value: []interface{}{tbs.WorkerStatusRunning, leaseUpdate}}
	if len(routingKey)>0{
		o.Where = o.Where + " and routing_key=?"
		o.Value = append(o.Value, routingKey)
	}

	err := models.Get(db.DB, &tbs.JoborWorker{}, o, &workers)
	//var onlineWorkers []*tbs.JoborWorker
	//for _, worker := range workers {
	//	if worker.Status==tbs.WorkerStatusRunning {
	//		onlineWorkers = append(onlineWorkers, worker)
	//	}
	//}
	if len(workers) == 0 {
		err = fmt.Errorf("can not get valid worker from routing_key[%s]", routingKey)
		return nil, err
	}

	return workers,err
}