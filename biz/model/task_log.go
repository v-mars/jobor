package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	task "jobor/kitex_gen/task"
	"jobor/kitex_gen/task_log"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
	"time"
)

const (
	NameLog = "jobor_log"
)

type JoborLog struct {
	db.Model
	Name          string        `gorm:"type:varchar(128);index:idx_code;comment:任务名" json:"name" form:"name"`
	Lang          string        `gorm:"type:varchar(16);index:idx_code;not null;comment:任务类型:[shell,api,python,golang,e.g.]" json:"lang" form:"lang"`
	TaskId        int           `gorm:"index:task_id;comment:关联任务id" json:"task_id"`
	ByTaskLogId   int           `gorm:"index:by_task_log_id;comment:关联任务log id" json:"by_task_log_id"`
	TriggerMethod string        `gorm:"type:varchar(16);comment:任务触发方式：auto,manual" json:"trigger_method"`
	Expr          string        `gorm:"type:varchar(32);not null;comment:定时任务表达式：0/1 * * ? * * * 秒分时天月星期" json:"expr" form:"expr"`
	Data          task.TaskData `gorm:"type:mediumtext;not null;comment:任务执行详细，格式：json" json:"data" form:"data"`
	Resp          string        `gorm:"type:mediumtext;comment:任务返回结果" json:"resp"`
	Result        string        `gorm:"type:varchar(16);comment:任务结果: success,failed" json:"result" form:"result"`
	ErrCode       int           `gorm:"comment:'错误码'" json:"err_code" form:"err_code"`
	ErrMsg        string        `gorm:"type:mediumtext;comment:错误信息" json:"err_msg" form:"err_msg"`
	Addr          string        `gorm:"type:varchar(64);not null;comment:任务运行的log节点" json:"addr" form:"addr"`
	Idempotent    string        `gorm:"type:varchar(156);comment:幂等标志" json:"idempotent" form:"idempotent"`
	ExpectContent string        `gorm:"type:varchar(500);default:null;comment:期望任务返回结果" json:"expect_content" form:"expect_content"`
	ExpectCode    int           `gorm:"default:0;comment:期望任务状态码" json:"expect_code" form:"expect_code"`
	StartTime     db.JSONTime   `gorm:"default: null;type:datetime;comment:开始时间" json:"start_time" form:"start_time"`
	EndTime       db.JSONTime   `gorm:"default: null;type:datetime;comment:结束时间" json:"end_time" form:"end_time"`
	CostTime      db.MillisTime `gorm:"type:float;default:0;comment:执行耗时(毫秒)" json:"cost_time" form:"cost_time"`                                                          // 单位：毫秒
	Parents       []*JoborLog   `gorm:"many2many:jobor_log_parents;association_autoupdate:false;association_autocreate:false;constraint:OnDelete:CASCADE" json:"parents"` // Many-To-Many
	Childs        []*JoborLog   `gorm:"many2many:jobor_log_childs;association_autoupdate:false;association_autocreate:false;constraint:OnDelete:CASCADE" json:"childs"`   // Many-To-Many
	//CostTime      float32       `gorm:"comment:'任务耗时'" json:"cost_time" form:"cost_time"`
}

func (u *JoborLog) TableName() string {
	return NameLog
}

func (u *JoborLog) GetLogRpc() *task_log.LogResp {
	rpcObj := task_log.LogResp{}
	err := utils.AnyToAny(u, &rpcObj)
	if err != nil {
		panic(err)
	}
	return &rpcObj
}

type Logs []JoborLog

func (u *Logs) List(req *task_log.LogQuery, resp *response.PageDataList) (Logs, error) {
	resp.List = u
	if err := PageDataWithScopes(db.DB.Model(&JoborLog{}), NameLog, Find, resp,
		GetScopesList(SelectLog()), PreloadLog("Parents"), PreloadLog("Childs"),
		WhereLog(req),
		OrderLog(), GroupLog()); err != nil {
		return nil, err
	}
	for i, v := range *u {
		if v.Result == TaskLogStatusRunning || v.Result == TaskLogStatusWait {
			var totalTime = time.Now().Unix() - v.StartTime.Unix()
			(*u)[i].CostTime = db.MillisTime((time.Second * time.Duration(totalTime)).String())
		}
	}
	return *u, nil
}
func (u *Logs) ListLog() (uis []*task_log.LogResp) {
	if u != nil {
		err := utils.AnyToAny(u, &uis)
		if err != nil {
			panic(err)
		}
	}
	return uis
}
func NewLogModel(Db *gorm.DB) *JoborLog {
	return &JoborLog{Model: db.Model{GormDB: db.DB}}
}

func SelectLog() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `jobor_log.*`
		return Db.Select(sql)
	}
}
func SelectAllLog() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,name")
	}
}

func WhereLog(req *task_log.LogQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "(by_task_log_id < 1 or by_task_log_id is null) and name like ?"
		var sqlArgs = []interface{}{"%" + req.Name + "%"}
		if len(req.Lang) > 0 {
			sql = sql + " and lang=?"
			sqlArgs = append(sqlArgs, req.Lang)
		}
		if len(req.Result) > 0 {
			sql = sql + " and result = ?"
			sqlArgs = append(sqlArgs, req.Result)
		}
		if len(req.TriggerMethod) > 0 {
			sql = sql + " and trigger_method = ?"
			sqlArgs = append(sqlArgs, req.TriggerMethod)
		}
		if len(req.Addr) > 0 {
			sql = sql + " and addr like ?"
			sqlArgs = append(sqlArgs, req.Addr)
		}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsLog() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		sql := ``
		return Db.Joins(sql)
	}
}
func PreloadLog(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderLog() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("jobor_log.id desc")
	}
}

func GroupLog() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("jobor_log.id")
	}
}

func AddLog(ctx context.Context, Db *gorm.DB, req *task_log.PostLogReq) (JoborLog, error) {
	var row JoborLog
	if err := utils.AnyToAny(req, &row); err != nil {
		return row, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := Db.Table(row.TableName()).Create(&row).Error; err != nil {
		return row, err
	}
	return row, nil
}

func ModLog(ctx context.Context, Db *gorm.DB, _id interface{}, req *task_log.PutLogReq) (JoborLog, error) {
	var mapData map[string]interface{}
	var err error
	if mapData, err = convert.StructToMap(req); err != nil {
		return JoborLog{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var logObj JoborLog
	if err = tx.First(&logObj, _id).Error; err != nil {
		return logObj, err
	}
	if err = tx.Table(NameLog).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return logObj, err
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return logObj, err
	}
	hlog.Debug("事务提交成功")
	if err = Db.First(&logObj, _id).Error; err != nil {
		return logObj, err
	}
	return logObj, nil
}

func DelLog(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]JoborLog, error) {
	var us []JoborLog
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&JoborLog{}).Find(&us, "id in (?)", _ids).Error; err != nil {
		return us, err
	}
	for _, _id := range _ids {
		if err := Db.Table(NameLog).Where("id!=?", _id).Update("deleted", true).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("Jobor任务删除成功")
	return us, nil
}

// GetByNameLog 根据Jobor任务名称获取Jobor任务信息
func (u *JoborLog) GetByNameLog(name string) (JoborLog, error) {
	err := db.DB.Table(u.TableName()).Where("name = ?", name).Take(&u).Error
	if err != nil {
		return JoborLog{}, err
	}
	return *u, nil
}

func GetLogByName(ctx context.Context, Db *gorm.DB, name string) (*JoborLog, error) {
	var row = JoborLog{}
	err := Db.Table(NameLog).Where("name=?", name).First(&row).Error
	return &row, err
}

func GetLogInfoById(id interface{}, isPanic bool) (*JoborLog, error) {
	var err error
	var u JoborLog
	err = db.DB.Table(NameLog).Where("id= ?", id).Scopes(PreloadLog("Parents"), PreloadLog("Childs")).First(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the log information does not exist")
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	return &u, nil
}

func GetLogInfoByName(name string, isPanic bool) (task_log.LogResp, error) {
	var err error
	var u task_log.LogResp
	err = db.DB.Table(NameLog).Where("name = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the log information does not exist")
		if isPanic {
			panic(err)
		}
		return task_log.LogResp{}, err
	}
	return u, nil
}
