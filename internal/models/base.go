package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"jobor/internal/config"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"log"
	"os"
	"time"
)

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error
	var dialector gorm.Dialector
	if config.Gorm.DBType=="mysql"{
		//dialector = mysql.New(mysql.Config{
		//	DSN: config.MySQL.DSN(), // DSN data source name
		//	DefaultStringSize: 256, // string 类型字段的默认长度
		//	DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		//})
		dialector=mysql.Open(config.MySQL.DSN())
	} else if config.Gorm.DBType=="sqlite3"{
		dialector=mysql.Open(config.Sqlite3.DSN())
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	var gormConfig = gorm.Config{
		//DryRun: false,
		Logger: newLogger,
		QueryFields: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	//gormConfig.Logger.Info(context.TODO(),log.New(os.Stdout, "\r\n", 0))
	gdb, err = gorm.Open(dialector,&gormConfig)
	if err != nil {
		log.Fatal("SQL Connect Failed: ", err)
	}
	//sqlDB, err := gdb.DB()
	//if err != nil {
	//	log.Fatal("SQL Failed: ", err)
	//}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConn)
	//gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConn)
	//gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second)
	db.DB =gdb
}


func Migration() {
	err := db.DB.AutoMigrate(&tbs.JoborTask{},&tbs.JoborLog{},&tbs.JoborWorker{},&tbs.User{},
	&tbs.Usergroup{},&tbs.Role{},&tbs.Permission{})
	log.Printf("db migration err: %s", err)
}

type DataBases struct {
	Engine     string            `json:"engine" default:"mysql"`
	Name       string            `json:"name" default:"ops"`
	User       string            `json:"user" default:"root"`
	Password   string            `json:"password" default:"123456"`
	Host       string            `json:"host" default:"127.0.0.1"`
	Port       string            `json:"port" default:"3306"`
	Charset    string            `json:"charset" default:"utf8mb4"`
	Parameters string            `json:"parameters"`
	Options    map[string]string `json:"options"`
}

func (a DataBases) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.Name, a.Parameters)
}