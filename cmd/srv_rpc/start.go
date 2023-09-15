package srv_rpc

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"jobor/conf"
	"jobor/kitex_gen/pbapi/heartbeat"
	"jobor/rpc_biz/mw"
	"jobor/rpc_biz/svc"
	"net"
)

func StartSrvRpc() error {
	klog.Infof("start srv grpc server service")
	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: conf.AppWorkerName,
		Tags:        make(map[string]string),
	}
	ebi.Tags["idc"] = "cn-sh"
	ebi.Tags["name"] = conf.AppName
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
	addr, _ := net.ResolveTCPAddr("tcp", conf.GetConf().Server.GrpcAddress)
	svr := heartbeat.NewServer(new(svc.HeartbeatImpl), server.WithServerBasicInfo(ebi),
		server.WithServiceAddr(addr),
		server.WithMiddleware(mw.AccessLogMW),
		//server.WithLogger(),
		//server.WithACLRules(),
		//server.WithErrorHandler(),
		server.WithMuxTransport(), // IO多路复用
	)
	if err := svr.Run(); err != nil {
		klog.Info("srv grpc stopped with error:", err)
		return err
	} else {
		klog.Info("srv grpc stopped")
		return nil
	}
}
