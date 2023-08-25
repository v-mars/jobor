package kitex_gen

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"jobor/biz/dal/db"
)

type JoborTask struct {
	Name          string      `gorm:"type:varchar(128);unique_index;not null;comment:任务名" json:"name" form:"name"`
	Description   string      `gorm:"type:varchar(512);default:null;comment:'任务描述" json:"description" form:"description"`
	Lang          string      `gorm:"type:varchar(16);index:idx_code;not null;comment:任务类型:[shell,api,python,golang,e.g.]" json:"lang" form:"lang"`
	Data          TaskData    `gorm:"type:text;not null;comment:任务执行详细，格式：json" json:"data" form:"data"`
	Notify        Notify      `gorm:"type:text;null;comment:告警通知，格式：json" json:"notify" form:"notify"`
	UserId        int         `gorm:"index:user_id;comment:关联用户id" json:"user_id"`
	User          string      `gorm:"index:user;comment:关联用户" json:"user"`
	Count         int         `gorm:"comment:执行次数" json:"count" form:"count"`
	Expr          string      `gorm:"type:varchar(32);not null;comment:定时任务表达式：0/1 * * ? * * * 秒分时天月星期" json:"expr" form:"expr"`
	Timeout       int         `gorm:"default:-1;comment:超时时间" json:"timeout" form:"timeout"`
	RoutePolicy   int         `gorm:"default:1;comment:路由策略 1:Random 2:RoundRobin 3:Weight 4:LeastTask" json:"route_policy" form:"route_policy"`
	RoutingKey    string      `gorm:"type:varchar(32);default:'default';comment:执行worker路由标识" json:"routing_key" form:"routing_key"`
	Status        string      `gorm:"type:varchar(32);default:'running';comment:定时任务状态: running,stop" json:"status" form:"status"`
	AlarmPolicy   int         `gorm:"default:2;comment:告警策略：{0:never,1:always,2:failed,3:success}" json:"alarm_policy" form:"alarm_policy"`
	ExpectContent string      `gorm:"type:varchar(500);default:null;comment:期望任务返回结果" json:"expect_content" form:"expect_content"`
	ExpectCode    int         `gorm:"default:0;comment:期望任务状态码" json:"expect_code" form:"expect_code"`
	Retry         int         `gorm:"default:0;comment:重试次数" json:"retry" form:"retry"`
	Prev          db.JSONTime `gorm:"default: null;type:datetime;comment:上次执行时间" json:"prev" form:"prev"`
	Next          db.JSONTime `gorm:"default: null;type:datetime;comment:'下次执行时间'" json:"next" form:"next"`
	Updater       string      `gorm:"type:varchar(156);" json:"updater" form:"updater"`
	db.Model
	//D         TestD    `gorm:"type:text;comment:'任务执行详细，格式：json'" json:"d" form:"d"`
	//Deleted  bool   `gorm:"default:false;comment:'逻辑删除'" json:"deleted" form:"deleted"`
}

type JoborLog struct {
	Name          string      `gorm:"type:varchar(128);index:idx_code;comment:'任务名'" json:"name" form:"name"`
	Lang          string      `gorm:"type:varchar(16);index:idx_code;not null;comment:'任务类型:[shell,api,python,golang,e.g.]'" json:"lang" form:"lang"`
	TaskID        *int        `gorm:"index:task_id;comment:'关联任务id'" json:"task_id"`
	TriggerMethod string      `gorm:"type:varchar(16);comment:'任务触发方式：auto,manual'" json:"trigger_method"`
	Expr          string      `gorm:"type:varchar(32);not null;comment:'定时任务表达式：0/1 * * ? * * * 秒分时天月星期'" json:"expr" form:"expr"`
	Data          TaskData    `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	Resp          string      `gorm:"type:text;comment:'任务返回结果'" json:"resp"`
	CostTime      float32     `gorm:"comment:'任务耗时'" json:"cost_time" form:"cost_time"`
	Result        string      `gorm:"type:varchar(16);comment:'任务结果: success,failed'" json:"result" form:"result"`
	ErrCode       int         `gorm:"comment:'错误码'" json:"err_code" form:"err_code"`
	ErrMsg        string      `gorm:"type:text;comment:'错误信息'" json:"err_msg" form:"err_msg"`
	Addr          string      `gorm:"type:varchar(64);not null;comment:'任务运行的worker节点'" json:"addr" form:"addr"`
	StartTime     db.JSONTime `gorm:"default: null;type:datetime;comment:'开始时间'" json:"start_time" form:"start_time"`
	EndTime       db.JSONTime `gorm:"default: null;type:datetime;comment:'结束时间'" json:"end_time" form:"end_time"`
	db.Model
	//ParentTaskIds     []string    `json:"parent_taskids"`
	//ParentTaskIdsDesc []string    `json:"parent_taskidsdesc" comment:"父任务"`
	//ParentRunParallel bool        `json:"parent_runparallel" comment:"父任务运行策略"`
	//ChildTaskIds      []string    `json:"child_taskids"`
	//ChildTaskIdsDesc  []string    `json:"child_taskidsdesc"  comment:"子任务"`
	//ChildRunParallel  bool        `json:"child_runparallel" comment:"子任务运行策略"`
}

type JoborWorker struct {
	Hostname    string `gorm:"type:varchar(128);index:idx_code;not null;comment:'worker hostname'" json:"hostname" form:"hostname"`
	Ip          string `gorm:"type:varchar(128);index:ip;not null;comment:'worker ip'" json:"ip" form:"ip"`
	Addr        string `gorm:"type:varchar(64);unique_index;not null;comment:'worker ip:port'" json:"addr" form:"addr"`
	Version     string `gorm:"type:varchar(32);comment:'版本'" json:"version" form:"version"`
	RoutingKey  string `gorm:"type:varchar(64);default:'default';comment:'routing_key'" json:"routing_key" form:"routing_key"`
	Weight      int32  `gorm:"comment:'权重'" json:"weight" form:"weight"`
	LeaseUpdate int64  `gorm:"comment:'租约时间更新'" json:"lease_update" form:"lease_update"`
	Status      string `gorm:"type:varchar(32);default:'offline';comment:'状态：running,stop,offline'" json:"status" form:"status"`
	db.Model
}

const (
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

type TestD struct {
	Api1 string `json:"api_1"`
	Api2 string `json:"api_2"`
	Api3 int    `json:"api_3"`
	Api4 []struct {
		Aa string `json:"aa"`
		Bb string `json:"bb"`
	} `json:"api_4"`
}

func (j *TestD) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := TestD{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j TestD) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type TaskData struct {
	Code string `json:"data"`
	Api  struct {
		Url         string              `json:"url"`
		Method      string              `json:"method"`
		ContentType string              `json:"content_type"`
		Payload     string              `json:"payload"`
		Header      []map[string]string `json:"header"`
		Forms       []map[string]string `json:"forms"`
	} `json:"api"`
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *TaskData) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := TaskData{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j TaskData) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type Notify struct {
	Email struct {
		Enable    bool     `json:"enable"`
		Receivers []string `json:"receivers"` // 多个邮箱地址以逗号分割
	} `json:"email"`
	Webhook struct {
		Enable bool     `json:"enable"`
		Urls   []string `json:"urls"` // 多个api url以逗号分割
	} `json:"webhook"`
	// 多个飞书webhook以逗号分割
	Lark struct {
		Enable   bool `json:"enable"`
		Webhooks []struct {
			WebhookUrl string `json:"webhook_url"`
			Secret     string `json:"secret"`
		} `json:"webhooks"`
	} `json:"lark"`
	// 多个钉钉webhook以逗号分割
	DingDing struct {
		Enable   bool `json:"enable"`
		Webhooks []struct {
			WebhookUrl string `json:"webhook_url"`
			Secret     string `json:"secret"`
		} `json:"webhooks"`
	} `json:"dingding"`
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *Notify) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Notify{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j Notify) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}
