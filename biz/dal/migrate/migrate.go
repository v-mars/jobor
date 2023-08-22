package migrate

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/db"
)

func Migrate() {

	//err := db.DB.AutoMigrate(&tbs.K8sResource{})
	//if err != nil {
	//	log.Fatal("AutoMigrate err:", err)
	//}
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
