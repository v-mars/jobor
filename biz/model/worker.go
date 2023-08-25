package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	"jobor/kitex_gen/worker"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
)

const (
	NameWorker = "jobor_worker"
)

type JoborWorker struct {
	db.Model
	Hostname    string `gorm:"type:varchar(128);index:idx_code;not null;comment:'worker hostname'" json:"hostname" form:"hostname"`
	Ip          string `gorm:"type:varchar(128);index:ip;not null;comment:'worker ip'" json:"ip" form:"ip"`
	Addr        string `gorm:"type:varchar(64);unique_index;not null;comment:'worker ip:port'" json:"addr" form:"addr"`
	Version     string `gorm:"type:varchar(32);comment:'版本'" json:"version" form:"version"`
	RoutingKey  string `gorm:"type:varchar(64);default:'default';comment:'routing_key'" json:"routing_key" form:"routing_key"`
	Weight      int32  `gorm:"comment:'权重'" json:"weight" form:"weight"`
	LeaseUpdate int64  `gorm:"comment:'租约时间更新'" json:"lease_update" form:"lease_update"`
	Status      string `gorm:"type:varchar(32);default:'offline';comment:'状态：running,stop,offline'" json:"status" form:"status"`
}

func (u *JoborWorker) TableName() string {
	return NameWorker
}

func (u *JoborWorker) GetWorkerRpc() *worker.WorkerResp {
	rpcObj := worker.WorkerResp{}
	err := utils.AnyToAny(u, &rpcObj)
	if err != nil {
		panic(err)
	}
	return &rpcObj
}

type Workers []JoborWorker

func (u *Workers) List(req *worker.WorkerQuery, resp *response.PageDataList) (Workers, error) {
	resp.List = u
	if err := PageDataWithScopes(db.DB.Model(&JoborWorker{}), NameWorker, Find, resp,
		GetScopesList(), SelectWorker(),
		WhereWorker(req),
		OrderWorker(), GroupWorker()); err != nil {
		return nil, err
	}
	return *u, nil
}
func (u *Workers) ListWorker() (uis []*worker.WorkerResp) {
	if u != nil {
		err := utils.AnyToAny(u, &uis)
		if err != nil {
			panic(err)
		}
	}
	return uis
}
func NewWorkerModel(Db *gorm.DB) *JoborWorker {
	return &JoborWorker{Model: db.Model{GormDB: db.DB}}
}

func SelectWorker() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `jobor_worker.*`
		return Db.Select(sql)
	}
}
func SelectAllWorker() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,name")
	}
}

func WhereWorker(req *worker.WorkerQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "(hostname like ? or ip like ?)"
		var sqlArgs = []interface{}{"%" + req.Hostname + "%", "%" + req.Hostname + "%"}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsWorker() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		sql := ``
		return Db.Joins(sql)
	}
}
func PreloadWorker(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderWorker() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("jobor_worker.id desc")
	}
}

func GroupWorker() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("jobor_worker.id")
	}
}

func AddWorker(ctx context.Context, Db *gorm.DB, req *worker.PostWorkerReq) (JoborWorker, error) {
	var row JoborWorker
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

func ModWorker(ctx context.Context, Db *gorm.DB, _id interface{}, req *worker.PutWorkerReq) (JoborWorker, error) {
	var mapData map[string]interface{}
	var err error
	if mapData, err = convert.StructToMap(req); err != nil {
		return JoborWorker{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var workerObj JoborWorker
	if err = tx.First(&workerObj, _id).Error; err != nil {
		return workerObj, err
	}
	if err = tx.Table(NameWorker).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return workerObj, err
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return workerObj, err
	}
	hlog.Debug("事务提交成功")
	if err = Db.First(&workerObj, _id).Error; err != nil {
		return workerObj, err
	}
	return workerObj, nil
}

func DelWorker(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]JoborWorker, error) {
	var us []JoborWorker
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&JoborWorker{}).Find(&us, "id in (?)", _ids).Error; err != nil {
		return us, err
	}
	for _, _id := range _ids {
		if err := Db.Table(NameWorker).Where("id!=?", _id).Update("deleted", true).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("Jobor任务删除成功")
	return us, nil
}

// GetByNameWorker 根据Jobor任务名称获取Jobor任务信息
func (u *JoborWorker) GetByNameWorker(name string) (JoborWorker, error) {
	err := db.DB.Table(u.TableName()).Where("name = ?", name).Take(&u).Error
	if err != nil {
		return JoborWorker{}, err
	}
	return *u, nil
}

func GetWorkerByName(ctx context.Context, Db *gorm.DB, name string) (*JoborWorker, error) {
	var row = JoborWorker{}
	err := Db.Table(NameWorker).Where("name=?", name).First(&row).Error
	return &row, err
}

func GetWorkerInfoById(id interface{}, isPanic bool) (*worker.WorkerResp, error) {
	var err error
	var u worker.WorkerResp
	err = db.DB.Table(NameWorker).Where("id= ?", id).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Hostname == "" && u.Id == 0 {
		err = fmt.Errorf("the worker information does not exist")
		if isPanic {
			panic(err)
		}
		return &worker.WorkerResp{}, err
	}
	return &u, nil
}

func GetWorkerInfoByName(name string, isPanic bool) (worker.WorkerResp, error) {
	var err error
	var u worker.WorkerResp
	err = db.DB.Table(NameWorker).Where("name = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Hostname == "" && u.Id == 0 {
		err = fmt.Errorf("the worker information does not exist")
		if isPanic {
			panic(err)
		}
		return worker.WorkerResp{}, err
	}
	return u, nil
}
