package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	FS "jobor/fs"
	"jobor/internal/app/sys/user"
	"jobor/internal/config"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/internal/utils/casbin"
	"jobor/pkg/logger"
	"strings"
)

var tables = []string{
	models.User,
	models.Usergroup,
	models.Role,
	models.Permission,
	models.JoborTask,
	models.JoborLog,
	models.JoborWorker,
}

// QueryIsInstall check table is create
func QueryIsInstall(ctx context.Context,DB *gorm.DB) (bool, error) {
	var queryTable string
	var needTables []interface{}

	var queryName string
	var params []string
	driveName := config.Configs.Gorm.DBType
	if driveName == "sqlite3" {
		queryTable = `SELECT count(*) FROM sqlite_master WHERE type="table" AND (`
		queryName = "name"
	} else if driveName == "mysql" {
		dbname := strings.Split(strings.Split(config.Configs.MySQL.DSN(), "?")[0], "/")[1]
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
	logger.Info(queryTable, params)

	sqlDB, err := DB.DB()
	if err != nil {
		err = fmt.Errorf("SQL Failed: %s",err)
		logger.Error(err)

	}
	//defer sqlDB.Close()
	err = sqlDB.QueryRowContext(ctx, queryTable, needTables...).Scan(&count)
	if err != nil {
		logger.Errorf("msg string err: %s", err)
		return false, fmt.Errorf("scan failed: %w", err)
	}

	if count != len(tables) {
		return false, nil
	}
	return true, nil
}

// StartInstall start install system
func StartInstall(ctx context.Context,DB *gorm.DB, sqlFileName string) error {
	// create table
	sqlDB, err := DB.DB()
	if err != nil {
		err = fmt.Errorf("SQL Failed: %s", err)
		logger.Error(err)
		return err
	}

	fs := FS.SqlFs
	//sqlFileName:= "jobor.sql"
	file, err := fs.Open(sqlFileName)
	if err != nil {
		err = fmt.Errorf("fs.Open failed, filename %s err: %s", sqlFileName,err)
		logger.Error(err)
		return err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll failed: %s", err)
		logger.Error(err)
		return err
	}
	var execSql string
	if config.Configs.Gorm.DBType == "sqlite3" {
		// sqlite3 TODO 的自增字段为AUTOINCREMENT
		execSql = strings.Replace(string(content), "AUTO_INCREMENT", "AUTOINCREMENT", -1)
		execSql = strings.Replace(string(content), "COMMENT", "--", -1)
	} else {
		execSql = string(content)
	}

	sqlList:=strings.Split(execSql,";")
	for _,sqlRow:= range sqlList{
		sqlRow=strings.TrimPrefix(sqlRow,"\n")
		sqlRow=strings.TrimSuffix(sqlRow,"\n")
		if len(sqlRow)>0 && strings.HasPrefix(sqlRow, "-- ")==false {
			_, err = sqlDB.ExecContext(ctx, sqlRow+";")
			if err != nil {
				logger.Errorf("sqlDB.ExecContext failed: %s", err)
				return fmt.Errorf("sqlDB.ExecContext failed: %w", err)
			}
		}

	}

	logger.Debug("Success Run All jobor Sql")

	err = InitUpdatePermissionByGinRoutes()
	if err != nil {
		return err
	}
	var permissions []tbs.Permission
	if err=DB.Model(&tbs.Permission{}).Where(
		"name like ? or name=?","sys:%","jobor:dashboard:get").Find(&permissions).Error; err != nil {
		return fmt.Errorf("get permissions err: %s", err)
	}

	var roles []tbs.Role
	if err=DB.Find(&roles).Error; err != nil {
		return fmt.Errorf("get roles err: %s", err)
	} else if len(roles)==0{
		roles = []tbs.Role{
			{Name: "admin", Description: "超级管理员"},
			{Name: "system", Description: "系统管理",Permissions: permissions},
			{Name: "normal", Description: "普通用户"},
		}
		if err=DB.Create(&roles).Error; err != nil {
			return fmt.Errorf("add roles err: %s", err)
		}
	}
	for _,role:=range roles{
		for _,v:=range role.Permissions{
			_, err = casbin.Enforcer.AddPolicy(v.Name, "sys",v.Path, v.Method)
			if err!=nil{
				return fmt.Errorf("casbin.Enforcer.AddPolicy err: %s", err)
			}
		}
	}

	var u tbs.User
	if err=DB.First(&u).Error; errors.Is(err,gorm.ErrRecordNotFound) {
		u = tbs.User{Nickname: "admin",Username: "admin",Password: "admin",Email: "admin@example.com",
			Status: true,UserType: "local", Roles: roles}
		if err=user.AddUser(DB,u); err != nil {
			return fmt.Errorf("add user err: %s", err)
		}
	} else if err!=nil{
		return fmt.Errorf("first user err: %s", err)
	}

	// create admin user
	/*err = user.AddUser(User)
	if err != nil {
		return fmt.Errorf("AddUser failed: %w", err)
	}*/
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	logger.Errorf("enforcer.LoadPolicy failed: %s", err)
	//	return fmt.Errorf("enforcer.LoadPolicy failed: %w", err)
	//}

	logger.Debug("Success Install Jobor")
	return nil
}


func CheckInstall(c *gin.Context)  {
	install, err := QueryIsInstall(c,db.DB)
	if err != nil {
		response.Error(c,err)
		return
	}
	response.Success(c,install)
}

func Install(c *gin.Context)  {
	err := StartInstall(c, db.DB, "jobor.sql")
	if err != nil {
		response.Error(c,err)
		return
	}
	response.Success(c,"")
}