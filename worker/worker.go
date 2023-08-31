package worker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"jobor/biz/pack/dispatcher"
	"jobor/conf"
	"jobor/kitex_gen/pbapi/taskservice"
	"jobor/rpc_biz/mw"
	"jobor/rpc_biz/registry"
	"jobor/rpc_biz/svc"
	"net"
	"strconv"
)

func MQWorker() {
	wc := conf.GetWorkerConf()
	srv, err := dispatcher.InitQSrv(&wc.Redis, dispatcher.Queue)
	if err != nil {
		hlog.Error(err)
		return
	}
	//worker := srv.NewCustomQueueWorker("worker_name", 10, "default")
	worker := srv.NewWorker("worker_name", 10)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		hlog.Errorf("I am an error handler:", err)
	}

	pretaskhandler := func(signature *tasks.Signature) {
		hlog.Infof("I am a start of task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		hlog.Infof("I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)
	err = worker.Launch()
	fmt.Println("worker is start")
	if err != nil {
		// do something with the error
		hlog.Error(err)
	}
}

func StartWorkerRpc() error {
	klog.Infof("start worker grpc server service")
	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: conf.AppWorkerName,
		Tags:        make(map[string]string),
	}
	ebi.Tags["idc"] = "cn-sh"
	ebi.Tags["type"] = "worker"
	ebi.Tags["name"] = conf.GetWorkerConf().WorkerName
	ebi.Tags["routing_key"] = conf.GetWorkerConf().RoutingKey
	//ebi.Tags["dyeing"] = conf.GetDyeing()
	//consulConfig := consulapi.Config{
	//	Address: conf.GetConf().Consul.Address,
	//	//Scheme:  "https",
	//	//Token:   "TEST-MY-TOKEN",
	//}
	//r, err := consul.NewConsulRegisterWithConfig(&consulConfig,
	//	consul.WithCheck(&consulapi.AgentServiceCheck{
	//		Interval:                       "7s",
	//		Timeout:                        "5s",
	//		DeregisterCriticalServiceAfter: "30s",
	//	},
	//	))
	//if err != nil {
	//	klog.Fatal(err)
	//}
	addr, _ := net.ResolveTCPAddr("tcp", conf.GetWorkerConf().GRpcAddr)
	svr := taskservice.NewServer(new(svc.TaskServiceImpl), server.WithServerBasicInfo(ebi),
		server.WithServiceAddr(addr),
		server.WithMiddleware(mw.AccessLogMW),
		//server.WithLogger(),
		//server.WithACLRules(),
		//server.WithErrorHandler(),
		server.WithMuxTransport(), // IO多路复用
	)
	_, port, _ := net.SplitHostPort(conf.GetWorkerConf().GRpcAddr)
	intPort, _ := strconv.Atoi(port)
	go registry.SendHeartbeatAndRegistryWorker(int32(intPort), 5)
	if err := svr.Run(); err != nil {
		klog.Info("worker grpc stopped with error:", err)
		return err
	} else {
		klog.Info("worker grpc stopped")
		return nil
	}
	//return nil
}
