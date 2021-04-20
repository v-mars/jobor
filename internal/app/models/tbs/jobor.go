package tbs


type JoborTask struct {
	BaseID
	Name          string `gorm:"type:varchar(128);unique_index;comment:'任务名'" json:"name" form:"name"`
	Description   string `gorm:"type:varchar(512);default:null;comment:'任务描述" json:"description" form:"description"`
	Lang          string `gorm:"type:varchar(16);index:idx_code;not null;comment:'任务类型:[shell,api,python,golang,e.g.]'" json:"lang" form:"lang"`
	Data          string `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	UserID        *uint  `gorm:"index:user_id;comment:'关联用户id'" json:"user_id"`
	Count         int    `gorm:"comment:'执行次数'" json:"count" form:"count"`
	Expr          string `gorm:"type:varchar(32);not null;comment:'定时任务表达式：0/1 * * ? * * * 秒分时天月星期'" json:"expr" form:"expr"`
	Timeout       int    `gorm:"default:-1;comment:'超时时间'" json:"timeout" form:"timeout"`
	Status        string `gorm:"type:varchar(32);default:'running';comment:'定时任务状态: running,stop'" json:"status" form:"status"`
	AlarmPolicy   int    `gorm:"default:2;comment:'告警策略：{0:never,1:always,2:failed,3:success}'" json:"alarm_policy" form:"alarm_policy"`
	ExpectContent string `gorm:"type:varchar(500);default:null;comment:'期望任务返回结果'" json:"expect_content" form:"expect_content"`
	ExpectCode    int    `gorm:"default:0;comment:'期望任务状态码'" json:"expect_code" form:"expect_code"`
	Retry         int    `gorm:"default:0;comment:'重试次数'" json:"retry" form:"retry"`
	//Deleted  bool   `gorm:"default:false;comment:'逻辑删除'" json:"deleted" form:"deleted"`
	Prev JSONTime `gorm:"default: null;type:datetime;comment:'上次执行时间'" json:"prev" form:"prev"`
	Next JSONTime `gorm:"default: null;type:datetime;comment:'下次执行时间'" json:"next" form:"next"`
	BaseByUpdate
	BaseTime
}

type JoborLog struct {
	BaseID
	Name          string   `gorm:"type:varchar(128);index:idx_code;comment:'任务名'" json:"name" form:"name"`
	Lang          string   `gorm:"type:varchar(16);index:idx_code;not null;comment:'任务类型:[shell,api,python,golang,e.g.]'" json:"lang" form:"lang"`
	TaskID        *uint    `gorm:"index:task_id;comment:'关联任务id'" json:"task_id"`
	TriggerMethod string   `gorm:"type:varchar(16);comment:'任务触发方式：auto,manual'" json:"trigger_method"`
	Expr          string   `gorm:"type:varchar(32);not null;comment:'定时任务表达式：0/1 * * ? * * * 秒分时天月星期'" json:"expr" form:"expr"`
	Data          string   `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	Resp          string   `gorm:"type:text;comment:'任务返回结果'" json:"resp"`
	CostTime      float32  `gorm:"comment:'任务耗时'" json:"cost_time" form:"cost_time"`
	Result        string   `gorm:"type:varchar(16);comment:'任务结果: success,failed'" json:"result" form:"result"`
	ErrCode       int      `gorm:"comment:'错误码'" json:"err_code" form:"err_code"`
	ErrMsg        string   `gorm:"type:text;comment:'错误信息'" json:"err_msg" form:"err_msg"`
	StartTime     JSONTime `gorm:"default: null;type:datetime;comment:'开始时间'" json:"start_time" form:"start_time"`
	EndTime       JSONTime `gorm:"default: null;type:datetime;comment:'结束时间'" json:"end_time" form:"end_time"`
	BaseTime
}