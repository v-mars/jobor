package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/biz/dal/db"
	"jobor/biz/dal/redis"
	"jobor/biz/response"
	task2 "jobor/kitex_gen/task"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
)

const (
	NameTask          = "jobor_task"
	TaskStatusRunning = "running"
	TaskStatusStop    = "stop"

	WorkerStatusRunning = "running"
	WorkerStatusStop    = "stop"
	WorkerStatusOffline = "offline"

	TaskLogStatusSuccess = "success"
	TaskLogStatusRunning = "running"
	TaskLogStatusFailed  = "failed"
	TaskLogStatusUnknown = "unknown"
	TaskLogStatusAbort   = "abort"
	TaskLogStatusTimeout = "timeout"
	TaskLogStatusWait    = "wait"
	TaskLogStatusCancel  = "cancel"
)

type AlarmPolicy int

type TaskEvent uint8

const (
	// AddEvent recv add event
	AddEvent TaskEvent = iota + 1
	// ChangeEvent recv delete task
	ChangeEvent
	// DeleteEvent recv delete task
	DeleteEvent
	// RunEvent run a task2
	RunEvent
	// KillEvent recv stop task
	KillEvent

	Never   AlarmPolicy = 0
	Always  AlarmPolicy = 1
	Failed  AlarmPolicy = 2
	Success AlarmPolicy = 3

	TriggerAuto   = "auto"
	TriggerManual = "manual"

	PubSubChannel = "jobor.event"
)

type Event struct {
	TaskID int       // task id
	TE     TaskEvent // task event: add change delete stop task
}

func (ap AlarmPolicy) String() string {
	switch ap {
	case 0:
		return "从不"
	case 1:
		return "总是"
	case 2:
		return "失败"
	case 3:
		return "成功"
	default:
		return "未知告警策略"
	}
}

type JoborTask struct {
	db.Model
	Name          string         `gorm:"type:varchar(128);unique_index;not null;comment:任务名" json:"name" form:"name"`
	Description   string         `gorm:"type:mediumtext;default:null;comment:任务描述" json:"description" form:"description"`
	Lang          string         `gorm:"type:varchar(16);index:idx_code;not null;comment:任务类型:[shell,api,python,golang,e.g.]" json:"lang" form:"lang"`
	Data          task2.TaskData `gorm:"type:mediumtext;not null;comment:任务执行详细，格式：json" json:"data" form:"data"`
	Notify        task2.Notify   `gorm:"type:mediumtext;null;comment:告警通知，格式：json" json:"notify" form:"notify"`
	UserId        int            `gorm:"index:user_id;comment:关联用户id" json:"user_id"`
	User          string         `gorm:"index:user;comment:关联用户" json:"user"`
	Count         int            `gorm:"comment:执行次数" json:"count" form:"count"`
	Expr          string         `gorm:"type:varchar(32);not null;comment:定时任务表达式：0/1 * * ? * * * 秒分时天月星期" json:"expr" form:"expr"`
	Timeout       int            `gorm:"default:-1;comment:超时时间" json:"timeout" form:"timeout"`
	RoutePolicy   int            `gorm:"default:1;comment:路由策略 1:Random 2:RoundRobin 3:Weight 4:LeastTask" json:"route_policy" form:"route_policy"`
	RoutingKey    string         `gorm:"type:varchar(32);default:'default';comment:执行worker路由标识" json:"routing_key" form:"routing_key"`
	Status        string         `gorm:"type:varchar(32);default:'running';comment:定时任务状态: running,stop" json:"status" form:"status"`
	AlarmPolicy   int            `gorm:"default:2;comment:告警策略：{0:never,1:always,2:failed,3:success}" json:"alarm_policy" form:"alarm_policy"`
	ExpectContent string         `gorm:"type:varchar(500);default:null;comment:期望任务返回结果" json:"expect_content" form:"expect_content"`
	ExpectCode    int            `gorm:"default:0;comment:期望任务状态码" json:"expect_code" form:"expect_code"`
	Retry         int            `gorm:"default:0;comment:重试次数" json:"retry" form:"retry"`
	Prev          db.JSONTime    `gorm:"default: null;type:datetime;comment:上次执行时间" json:"prev" form:"prev"`
	Next          db.JSONTime    `gorm:"default: null;type:datetime;comment:'下次执行时间'" json:"next" form:"next"`
	Updater       string         `gorm:"type:varchar(156);" json:"updater" form:"updater"`
	Deleted       bool           `gorm:"default:false;comment:逻辑删除" json:"deleted" form:"deleted"`
	ExprZh        string         `gorm:"-" json:"expr_zh" form:"expr_zh"`
	//D         TestD    `gorm:"type:text;comment:'任务执行详细，格式：json'" json:"d" form:"d"`

}

func (u *JoborTask) TableName() string {
	return NameTask
}

func (u *JoborTask) GetTaskRpc() *task2.TaskResp {
	rpcObj := task2.TaskResp{}
	err := utils.AnyToAny(u, &rpcObj)
	if err != nil {
		panic(err)
	}
	return &rpcObj
}

type Tasks []JoborTask

func (u *Tasks) List(req *task2.TaskQuery, resp *response.PageDataList) (Tasks, error) {
	resp.List = u
	if err := PageDataWithScopes(db.DB.Model(&JoborTask{}), NameTask, Scan, resp,
		GetScopesList(SelectTask()),
		WhereTask(req),
		OrderTask(), GroupTask()); err != nil {
		return nil, err
	}
	for i, v := range *u {
		i := i
		v := v
		(*u)[i].ExprZh = utils.TranslateToChinese(v.Expr)
	}
	return *u, nil
}
func (u *Tasks) ListTask() (uis []*task2.TaskResp) {
	if u != nil {
		err := utils.AnyToAny(u, &uis)
		if err != nil {
			panic(err)
		}
	}
	return uis
}
func NewTaskModel(Db *gorm.DB) *JoborTask {
	return &JoborTask{Model: db.Model{GormDB: db.DB}}
}

func SelectTask() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `jobor_task.*`
		return Db.Select(sql)
	}
}
func SelectAllTask() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,name")
	}
}

func WhereTask(req *task2.TaskQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "name like ?"
		var sqlArgs = []interface{}{"%" + req.Name + "%"}
		if len(req.Lang) > 0 {
			sql = sql + " and lang=?"
			sqlArgs = append(sqlArgs, req.Lang)
		}
		if len(req.RoutingKey) > 0 {
			sql = sql + " and routing_key like ?"
			sqlArgs = append(sqlArgs, req.RoutingKey)
		}
		if len(req.Status) > 0 {
			sql = sql + " and status = ?"
			sqlArgs = append(sqlArgs, req.Status)
		}
		if req.RoutePolicy > 0 {
			sql = sql + " and route_policy = ?"
			sqlArgs = append(sqlArgs, req.RoutePolicy)
		}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsTask() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		sql := ``
		return Db.Joins(sql)
	}
}
func PreloadTask(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderTask() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("jobor_task.id desc")
	}
}

func GroupTask() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("jobor_task.id")
	}
}

func AddTask(ctx context.Context, Db *gorm.DB, req *task2.PostTaskReq) (JoborTask, error) {
	var row JoborTask
	if err := utils.AnyToAny(req, &row); err != nil {
		return row, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := Db.Table(row.TableName()).Omit([]string{"Prev", "Next"}...).Create(&row).Error; err != nil {
		return row, err
	}
	if err := redis.Publish(ctx, PubSubChannel, Event{TaskID: row.ID,
		TE: ChangeEvent}); err != nil {
		return row, err
	}
	return row, nil
}

func ModTask(ctx context.Context, Db *gorm.DB, _id interface{}, req *task2.PutTaskReq) (JoborTask, error) {
	var mapData map[string]interface{}
	var err error
	if mapData, err = convert.StructToMap(req); err != nil {
		return JoborTask{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var taskObj JoborTask
	if err = tx.First(&taskObj, _id).Error; err != nil {
		return taskObj, err
	}
	if err = tx.Table(NameTask).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return taskObj, err
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return taskObj, err
	}
	hlog.Debug("事务提交成功")
	if err = redis.Publish(ctx, PubSubChannel, Event{TaskID: taskObj.ID,
		TE: ChangeEvent}); err != nil {
		return taskObj, err
	}
	if err = Db.First(&taskObj, _id).Error; err != nil {
		return taskObj, err
	}
	return taskObj, nil
}

func DelTask(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]JoborTask, error) {
	var us []JoborTask
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&JoborTask{}).Find(&us, "id in (?)", _ids).Error; err != nil {
		return us, err
	}
	for _, _id := range _ids {
		if err := redis.Publish(ctx, PubSubChannel,
			Event{TaskID: convert.ToInt(_id), TE: DeleteEvent}); err != nil {
			return nil, err
		}
		if err := Db.Table(NameTask).Where("id=?", _id).Updates(
			map[string]interface{}{"deleted": true, "status": TaskStatusStop}).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("Jobor任务删除成功")
	return us, nil
}

// GetByNameTask 根据Jobor任务名称获取Jobor任务信息
func (u *JoborTask) GetByNameTask(name string) (JoborTask, error) {
	err := db.DB.Table(u.TableName()).Where("name = ?", name).Take(&u).Error
	if err != nil {
		return JoborTask{}, err
	}
	return *u, nil
}

func GetTaskByName(ctx context.Context, Db *gorm.DB, name string) (*JoborTask, error) {
	var row = JoborTask{}
	err := Db.Table(NameTask).Where("name=?", name).First(&row).Error
	return &row, err
}

func GetTaskInfoById(id interface{}, isPanic bool) (*JoborTask, error) {
	var err error
	var u JoborTask
	err = db.DB.Table(NameTask).Where("id= ?", id).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Name == "" && u.ID == 0 {
		err = fmt.Errorf("the task information does not exist")
		if isPanic {
			panic(err)
		}
		return &JoborTask{}, err
	}
	return &u, nil
}

func GetTaskInfoByName(name string, isPanic bool) (task2.TaskResp, error) {
	var err error
	var u task2.TaskResp
	err = db.DB.Table(NameTask).Where("name = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the task information does not exist")
		if isPanic {
			panic(err)
		}
		return u, err
	}
	return u, nil
}

func GetAllRunningTask() ([]JoborTask, error) {
	var resList []JoborTask
	err := db.DB.Where("status=?", TaskStatusRunning).Find(&resList).Error
	return resList, err
}
