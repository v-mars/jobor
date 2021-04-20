package app

import (
	"context"
	"fmt"
	"jobor/internal/app/controllers/jobor/schedule"
	"jobor/internal/app/models"
	"jobor/internal/app/utils/casbin"
	"jobor/internal/config"
	redisCli "jobor/internal/redis"
	"jobor/pkg/logger"
	"log"
)


func Run(configPath string)  {
	// 加载配置
	err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	conf :=config.GetConf()
	//fmt.Println(conf)


	logger.Initial()
	initDB(&conf)

	// redis
	redisCli.Init(&conf, redisCli.Client)
	go func() {
		log.Fatal(redisCli.Subscribe(context.TODO(), schedule.Fn, schedule.PubSubChannel))
	}()

	_, err = casbin.NewCasbin(conf.MySQL.CasbinDataSourceName())
	if err != nil {log.Fatal(fmt.Errorf("casbin err: %s",err))}

	schedule.InitCron()

}

func initDB(config *config.Config){
	models.InitDB(config)
	models.Migration()
}




