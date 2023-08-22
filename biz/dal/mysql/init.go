package mysql

import (
	"jobor/biz/dal/db"
	"jobor/conf"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//var DB *gorm.DB

func Init() {
	db.DB = ConnectDB(conf.GetConf().MySQL.DSN)
	//db.DB = DB
}

func ConnectDB(dsn string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // TaskLog level
			Colorful:      false,        // 禁用彩色打印
		},
	)
	var gconf = gorm.Config{
		//DryRun: false,
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}

	tdb, err := gorm.Open(mysql.Open(dsn), &gconf)
	if err != nil {
		panic(err)
	}
	return tdb
}
