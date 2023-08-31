package dispatcher

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/conf"
)

const (
	exchange         = "jobor_exchange"
	ExchangeType     = "direct"
	BindingKey       = "jobor_tasks123"
	Queue            = "jobor_queue"
	RegisterTaskName = "TaskWorker"
)

var QSrv *machinery.Server

func InitQSrv(rds *conf.Redis, queue string) (*machinery.Server, error) {
	var qbr = fmt.Sprintf("redis://%s@%s/1", rds.Password, rds.Address)
	var cnf = &config.Config{
		//Broker:        "amqp://guest:guest@localhost:5672/",
		// redis://password@localhost:6379/1 "amqp://guest:guest@localhost:5672/",
		DefaultQueue:  queue,
		Broker:        qbr,
		ResultBackend: qbr,
		//AMQP: &config.AMQPConfig{
		//	Exchange:     exchange,
		//	ExchangeType: ExchangeType,
		//	BindingKey:   BindingKey,
		//},
		ResultsExpireIn: 120,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		// do something with the error
		return nil, err
	}
	QSrv = server
	//server.SendTask()
	err = QSrv.RegisterTask(RegisterTaskName, TaskWorker)
	if err != nil {
		hlog.Errorf("register Task  err: %s", err)
		return nil, err
	}
	return server, nil
}
