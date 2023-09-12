package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	"jobor/kitex_gen/worker"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
	"jobor/rpc_biz"
	"time"
)

const (
	NameWorker             = "jobor_worker"
	WorkerModeSsh          = "ssh"
	WorkerModeAgent        = "agent"
	WorkerAuthModePassword = "password"
	WorkerAuthModePubKey   = "pub_key"
)

type JoborWorker struct {
	db.Model
	Hostname    string    `gorm:"type:varchar(128);index:idx_code;not null;comment:worker hostname" json:"hostname" form:"hostname"`
	Description string    `gorm:"type:mediumtext;default:null;comment:描述" json:"description" form:"description"`
	Addr        string    `gorm:"type:varchar(64);unique_index;not null;comment:worker ip:port" json:"addr" form:"addr"`
	Ip          string    `gorm:"type:varchar(128);index:ip;not null;comment:worker ip" json:"ip" form:"ip"`
	Port        int32     `gorm:"default:22;comment:端口" json:"port" form:"port"`
	Mode        string    `gorm:"type:varchar(8);default:agent;comment:模式[agent,ssh]" json:"mode" form:"mode"`
	AuthMode    string    `gorm:"type:varchar(8);default:key;comment:认证模式[password,pub_key]" json:"auth_mode" form:"auth_mode"`
	Username    string    `gorm:"type:varchar(152);comment:认证用户" json:"username" form:"username"`
	Password    db.AesStr `gorm:"type:varchar(255);comment:认证密码" json:"password" form:"password"  keep_data:"yes"`
	Rsa         db.AesStr `gorm:"type:text;comment:SSH认证私钥" json:"rsa" form:"rsa"  keep_data:"yes"`
	Private     db.AesStr `gorm:"type:varchar(255);comment:SSH key秘钥" json:"private" form:"private"  keep_data:"yes"`
	Version     string    `gorm:"type:varchar(32);comment:'版本'" json:"version" form:"version"`
	RoutingKey  string    `gorm:"type:varchar(64);default:default;comment:routing_key" json:"routing_key" form:"routing_key"`
	Weight      int32     `gorm:"comment:权重" json:"weight" form:"weight"`
	LeaseUpdate int64     `gorm:"comment:租约时间更新" json:"lease_update" form:"lease_update"`
	Status      string    `gorm:"type:varchar(32);default:offline;comment:状态：running,stop,offline" json:"status" form:"status"`
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
		GetScopesList(SelectWorker()),
		WhereWorker(req),
		OrderWorker(), GroupWorker()); err != nil {
		return nil, err
	}
	leaseUpdate := time.Now().Unix() - int64(rpc_biz.DefaultHeartbeatInterval.Seconds())
	for i, v := range *u {
		v := v
		i := i
		if v.Mode == WorkerModeAgent && v.LeaseUpdate < leaseUpdate && v.Status == TaskLogStatusRunning {
			(*u)[i].Status = WorkerStatusOffline
		}
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
	if req.Addr == "" {
		row.Addr = fmt.Sprintf("%s:%d", row.Ip, row.Port)
	}
	if err := Db.Table(row.TableName()).Create(&row).Error; err != nil {
		return row, err
	}
	return row, nil
}

func ModWorker(ctx context.Context, Db *gorm.DB, _id interface{}, req *worker.PutWorkerReq) (JoborWorker, error) {
	var mapData map[string]interface{}
	var err error
	var workerObj JoborWorker
	if mapData, err = convert.StructToMap(req); err != nil {
		return workerObj, err
	}
	if err = convert.AnyToAny(req, &workerObj); err != nil {
		return workerObj, err
	}
	if req.Rsa != nil {
		mapData["rsa"] = workerObj.Rsa
	}
	if req.Password != nil {
		mapData["password"] = workerObj.Password
		fmt.Println("mapData[\"password\"]:", mapData["password"], workerObj.Password)
	}
	if req.Private != nil {
		mapData["private"] = workerObj.Private
	}
	fmt.Println("mapData:", mapData)

	tx := Db.Begin()
	defer func() { tx.Rollback() }()

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
		if err := Db.Table(NameWorker).Delete(&JoborWorker{}, _id).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("Jobor任务节点删除成功")
	return us, nil
}

// GetByNameWorker 根据Jobor任务节点名称获取Jobor任务节点信息
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

func GetWorkers(routingKey string) ([]JoborWorker, error) {
	var workers []JoborWorker
	leaseUpdate := time.Now().Unix() - int64(rpc_biz.DefaultHeartbeatInterval.Seconds())
	var whereSql = "status=? and lease_update>=?"
	var whereArgs = []interface{}{WorkerStatusRunning, leaseUpdate}
	if len(routingKey) > 0 {
		whereSql = whereSql + " and routing_key=?"
		whereArgs = append(whereArgs, routingKey)
	}
	err := db.DB.Model(&JoborWorker{}).Where(whereSql, whereArgs...).Find(&workers).Error
	if err != nil {
		return nil, err
	}
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

	return workers, nil
}

func CreateOrUpdate(data JoborWorker) error {
	var exist = JoborWorker{}
	//if err:=db.DB.Model(&tbs.JoborWorker{}).Where(tbs.JoborWorker{Addr: data.Addr}).FirstOrCreate(&data).Error;err!=nil{
	//return err
	//}
	if err := db.DB.Model(&JoborWorker{}).Where(JoborWorker{Addr: data.Addr}).First(&exist).Error; errors.Is(
		err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&JoborWorker{}).Create(&data).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		if err = db.DB.Model(&JoborWorker{}).Where("addr=? and status!=?", data.Addr,
			WorkerStatusStop).Updates(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
