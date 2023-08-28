package q

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"jobor/conf"
)

const (
	exchange     = "jobor_exchange"
	ExchangeType = "direct"
	BindingKey   = "jobor_tasks123"
	queue        = "jobor_queue"
)

var QSrv *machinery.Server

func InitQSrv() (*machinery.Server, error) {
	rds := conf.GetConf().Redis
	var qbr = fmt.Sprintf("redis://%s@%s", rds.Password, rds.Address)
	var cnf = &config.Config{
		// redis://password@localhost:6379 "amqp://guest:guest@localhost:5672/",
		Broker: qbr,
		//Broker:        "amqp://guest:guest@localhost:5672/",
		DefaultQueue:  queue,
		ResultBackend: qbr,
		AMQP: &config.AMQPConfig{
			Exchange:     exchange,
			ExchangeType: ExchangeType,
			BindingKey:   BindingKey,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		// do something with the error
		return nil, err
	}
	QSrv = server
	//server.SendTask()
	return server, nil
}
