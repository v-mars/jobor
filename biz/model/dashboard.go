package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/kitex_gen/dashboard"
)

const (
	NameDash = "dashboard"
)

type DashParam struct {
	EndTime   string `query:"endDate" form:"endDate" json:"endDate"`
	StartTime string `query:"startDate" form:"startDate" json:"startDate"`
	//EndTime   time.Time `form:"endDate" json:"endDate"`
	//StartTime time.Time `form:"startDate" json:"startDate"`
}
type Dash struct {
	Tasks            int64 `json:"tasks"`
	TaskLogs         int64 `json:"task_logs"`
	Workers          int64 `json:"workers"`
	TodayTaskLogs    int64 `json:"today_task_logs"`
	TaskLogStatusDay []struct {
		Name         string `json:"name"`
		Count        int64  `json:"count"`
		FailedCount  int64  `json:"failed_count"`
		TimeoutCount int64  `json:"timeout_count"`
		SuccessCount int64  `json:"success_count"`
		Time         string `json:"time"`
	} `json:"task_log_status_day"`
	TaskRun []struct {
		OrdNum int    `json:"ord_num"`
		Task   string `json:"task"`
		Count  int64  `json:"count"`
	} `json:"task_run"`
}

func GetDash(ctx context.Context, DB *gorm.DB, param DashParam) (Dash, error) {
	var err error
	defer func() {
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
		}
	}()
	var data Dash

	if err = DB.Model(&JoborTask{}).Where("status=?", TaskStatusRunning).Count(&data.Tasks).Error; err != nil {
		return data, err
	}
	if err = DB.Model(&JoborLog{}).Count(&data.TaskLogs).Error; err != nil {
		return data, err
	}
	if err = DB.Model(&JoborWorker{}).Where("status=?", WorkerStatusRunning).Count(&data.Workers).Error; err != nil {
		return data, err
	}
	if err = DB.Model(&JoborLog{}).Where("to_days(created_at) = to_days(now())").Count(&data.TodayTaskLogs).Error; err != nil {
		return data, err
	}

	taskLogStatusSql := fmt.Sprintf(`
SELECT
  DATE_FORMAT(d.created_at,'%%Y-%%m-%%d') as time,
	d.name,
	SUM( CASE d.result WHEN 'failed' THEN 1 ELSE 0 END ) AS failed_count,
	SUM( CASE d.result WHEN 'success' THEN 1 ELSE 0 END ) AS success_count,
	SUM( CASE d.result WHEN 'timeout' THEN 1 ELSE 0 END ) AS timeout_count,
	count(d.id) as count
FROM
	jobor_log AS d
WHERE
	d.created_at BETWEEN '%s' AND '%s'
GROUP BY
	time
	ORDER BY
	time asc
`, param.StartTime, param.EndTime)
	if err = DB.Model(&JoborLog{}).Raw(taskLogStatusSql).Scan(&data.TaskLogStatusDay).Error; err != nil {
		return data, err
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
	log.created_at BETWEEN '%s' AND '%s'
	GROUP BY t.NAME
	ORDER BY count DESC
`, param.StartTime, param.EndTime)
	if err = DB.Model(&JoborTask{}).Raw(taskRunSql).Scan(&data.TaskRun).Error; err != nil {
		return data, err
	}

	return data, nil
}

func SelectScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `dashboard.*`
		return Db.Select(sql)
	}
}
func SelectAllScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,dashboardname,nickname,status")
	}
}

func WhereScopes(req *dashboard.DashboardQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "name like ?"
		var sqlArgs = []interface{}{"%" + req.Name + "%"}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsScopesDash() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Joins("")
	}
}
func PreloadScopesDash(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderScopesDash() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("id desc")
	}
}

func GroupScopesDash() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("dashboard.id")
	}
}
