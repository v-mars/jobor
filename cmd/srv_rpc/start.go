package srv_rpc

func Start() error {
	//klog.Infof("start grpc server service")
	//ebi := &rpcinfo.EndpointBasicInfo{
	//	ServiceName: conf.AppName,
	//	Tags:        make(map[string]string),
	//}
	//ebi.Tags["idc"] = "cn-sh"
	//ebi.Tags["env"] = conf.GetEnv()
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
	//addr, _ := net.ResolveTCPAddr("tcp", conf.GetConf().Server.GrpcAddress)
	//svr := sso.NewServer(new(rpc_biz.SsoImpl), server.WithRegistry(r), server.WithServerBasicInfo(ebi),
	//	server.WithServiceAddr(addr),
	//	server.WithMiddleware(mw.AccessLogMW),
	//	//server.WithLogger(),
	//	//server.WithACLRules(),
	//	//server.WithErrorHandler(),
	//	//server.WithMuxTransport(), // IO多路复用
	//)
	//if err = svr.Run(); err != nil {
	//	klog.Info("server stopped with error:", err)
	//	return err
	//} else {
	//	klog.Info("server stopped")
	//	return nil
	//}
	return nil
}
