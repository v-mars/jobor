package internal

import (
	"context"
	"fmt"
	"jobor/internal/app/jobor/schedule"
	"jobor/internal/config"
	"jobor/internal/models"
	redisCli "jobor/internal/redis"
	"jobor/internal/utils/casbin"
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




