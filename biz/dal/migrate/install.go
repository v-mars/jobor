package migrate

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"io/ioutil"
	"jobor/biz/dal/db"
	"jobor/biz/model"
	"jobor/biz/response"
	"jobor/conf"
	FS "jobor/fs"
	"jobor/kitex_gen/audit"
	"strings"
)

var tables = []string{
	model.NameUser,
	model.NameRole,
	model.NameTask,
	model.NameLog,
	model.NameWorker,
	model.NameApi,
	audit.NameAuditLog,
}

// QueryIsInstall check table is create
func QueryIsInstall(ctx context.Context, DB *gorm.DB) (bool, error) {
	var queryTable string
	var needTables []interface{}

	var queryName string
	var params []string
	driveName := "mysql"
	if driveName == "sqlite3" {
		queryTable = `SELECT count(*) FROM sqlite_master WHERE type="table" AND (`
		queryName = "name"
	} else if driveName == "mysql" {
		dbname := strings.Split(strings.Split(conf.GetConf().MySQL.DSN, "?")[0], "/")[1]
		needTables = append(needTables, dbname)
		queryTable = `SELECT count(*) FROM information_schema.TABLES WHERE TABLE_SCHEMA=? AND (`
		queryName = "table_name"
	} else {
		return false, fmt.Errorf("unsupport drive type %s, only support sqlite3 or mysql", driveName)
	}

	for _, tbName := range tables {
		needTables = append(needTables, tbName)
	}

	for i := 0; i < len(tables); i++ {
		params = append(params, queryName+"=?")
	}
	queryTable += strings.Join(params, " OR ")
	queryTable += ")"
	var count int
	hlog.Info(queryTable, params)

	sqlDB, err := DB.DB()
	if err != nil {
		err = fmt.Errorf("SQL Failed: %s", err)
		hlog.Error(err)

	}
	//defer sqlDB.Close()
	err = sqlDB.QueryRowContext(ctx, queryTable, needTables...).Scan(&count)
	if err != nil {
		hlog.Errorf("msg string err: %s", err)
		return false, fmt.Errorf("scan failed: %w", err)
	}

	if count != len(tables) {
		return false, nil
	}
	return true, nil
}

// StartInstall start install system
func StartInstall(ctx context.Context, DB *gorm.DB, sqlFileName string) error {
	// create table
	sqlDB, err := DB.DB()
	if err != nil {
		err = fmt.Errorf("SQL Failed: %s", err)
		hlog.Error(err)
		return err
	}

	fs := FS.SqlFs
	//sqlFileName:= "jobor.sql"
	file, err := fs.Open(sqlFileName)
	if err != nil {
		err = fmt.Errorf("fs.Open failed, filename %s err: %s", sqlFileName, err)
		hlog.Error(err)
		return err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll failed: %s", err)
		hlog.Error(err)
		return err
	}
	var execSql string
	driveName := "mysql"
	if driveName == "sqlite3" {
		// sqlite3 TODO 的自增字段为AUTOINCREMENT
		execSql = strings.Replace(string(content), "AUTO_INCREMENT", "AUTOINCREMENT", -1)
		execSql = strings.Replace(string(content), "COMMENT", "--", -1)
	} else {
		execSql = string(content)
	}

	//sqlList := strings.Split(execSql, ";")
	//for _, sqlRow := range sqlList {
	//	sqlRow = strings.TrimPrefix(sqlRow, "\n")
	//	sqlRow = strings.TrimSuffix(sqlRow, "\n")
	//	if len(sqlRow) > 0 && strings.HasPrefix(sqlRow, "-- ") == false {
	//		_, err = sqlDB.ExecContext(ctx, sqlRow+";")
	//		if err != nil {
	//			hlog.Errorf("sqlDB.ExecContext failed: %s", err)
	//			return fmt.Errorf("sqlDB.ExecContext failed: %w", err)
	//		}
	//	}
	//
	//}
	if len(execSql) > 0 {
		if _, err = sqlDB.ExecContext(ctx, execSql); err != nil {
			hlog.Error(err)
			return err
		}
	}

	hlog.Debug("Success Run All jobor Sql")

	//var sysApi []model.Api
	//if err = DB.Model(&model.Api{}).Where(
	//	"name like ? or name=?", "sys:%", "jobor:dashboard:get").Find(&sysApi).Error; err != nil {
	//	return fmt.Errorf("get sysApi err: %s", err)
	//}
	//
	//var normalPermissions []model.Api
	//if err = DB.Model(&model.Api{}).Where(
	//	"name like ? or name like ? or name in (?)",
	//	"jobor:task:%", "jobor:log:%", []string{"jobor:dashboard:get", "sys:user-profile:put", "sys:user-pass:put"},
	//).Find(&normalPermissions).Error; err != nil {
	//	return fmt.Errorf("get normalPermissions err: %s", err)
	//}
	//
	//var roles []model.Role
	//if err = DB.Find(&roles).Error; err != nil {
	//	return fmt.Errorf("get roles err: %s", err)
	//} else if len(roles) == 0 {
	//	roles = []model.Role{
	//		{Name: "admin", Description: "超级管理员"},
	//		{Name: "system", Description: "系统管理", Apis: sysApi},
	//		{Name: "normal", Description: "普通用户"},
	//	}
	//	if err = DB.Create(&roles).Error; err != nil {
	//		return fmt.Errorf("add roles err: %s", err)
	//	}
	//}
	//for _, role := range roles {
	//	for _, v := range role.Apis {
	//		_, err = casbin.Enforcer.AddPolicy(v.Name, conf.Dom, v.Path, v.Method)
	//		if err != nil {
	//			return fmt.Errorf("casbin.Enforcer.AddPolicy err: %s", err)
	//		}
	//	}
	//}
	//
	//var u model.User
	//if err = DB.First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	//	ur := user.PostUserReq{Nickname: "admin", Username: "admin", Password: "admin", Email: "admin@example.com",
	//		Status: true, UserType: "local"}
	//	if _, err = model.AddUser(ctx, DB, &ur); err != nil {
	//		return fmt.Errorf("add user err: %s", err)
	//	}
	//} else if err != nil {
	//	return fmt.Errorf("first user err: %s", err)
	//}

	// create admin user
	/*err = user.AddUser(User)
	if err != nil {
		return fmt.Errorf("AddUser failed: %w", err)
	}*/
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	hlog.Errorf("enforcer.LoadPolicy failed: %s", err)
	//	return fmt.Errorf("enforcer.LoadPolicy failed: %w", err)
	//}

	hlog.Debug("Success Install Jobor")
	return nil
}

func CheckInstall(ctx context.Context, c *app.RequestContext) {
	install, err := QueryIsInstall(ctx, db.DB)
	if err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}
	response.Success(ctx, c, install)
}

func Install(ctx context.Context, c *app.RequestContext) {
	err := StartInstall(ctx, db.DB, "jobor.sql")
	if err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}
	response.Success(ctx, c, "")
}
