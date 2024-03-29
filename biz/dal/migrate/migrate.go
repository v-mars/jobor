package migrate

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/db"
	"jobor/biz/model"
	"jobor/kitex_gen/audit"
)

func Migrate() {

	err := db.DB.AutoMigrate(
		&model.JoborTask{},
		&model.JoborWorker{},
		&model.JoborLog{},
		&audit.AuditLog{},
		&model.Api{},
		&model.Role{},
		&model.User{},
	)
	if err != nil {
		hlog.Fatal("AutoMigrate err:", err)
	}
	//fmt.Println(db.DB.AutoMigrate(new(tbs.Key)).Error)
	//MigSys()
}

func MigSys() {
	err := db.DB.AutoMigrate(
	//&api.Api{},
	)
	if err != nil {
		hlog.Fatal("auto migrate sys err:", err)
	}
}
