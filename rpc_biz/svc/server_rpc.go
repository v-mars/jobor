package svc

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	pbapi "jobor/kitex_gen/pbapi"
	"jobor/rpc_biz/registry"
)

// HeartbeatImpl implements the last service interface defined in the IDL.
type HeartbeatImpl struct{}

// RegistryWorker implements the HeartbeatImpl interface.
func (s *HeartbeatImpl) RegistryWorker(ctx context.Context, req *pbapi.RegistryReq) (resp *pbapi.Empty, err error) {
	// TODO: Your code here...

	//p, ok := peer.FromContext(ctx)
	//if !ok {
	//	return &pbapi.Empty{}, fmt.Errorf("registry failed, %s", p)
	//}
	//ip, _, _ := net.SplitHostPort(p.Addr.String())
	hlog.Debugf("`registry host` new worker req %s", req)
	//req.Ip = ip

	if req.Ip == "::1" || req.Ip == "[::]" {
		req.Ip = "127.0.0.1"
	}
	addr := fmt.Sprintf("%s:%d", req.Ip, req.Port)

	subscriptions := registry.Subscriptions{Ip: req.Ip, Hostname: req.Hostname, Port: req.Port, Weight: req.Weight,
		Version: req.Version, RoutingKey: req.RoutingKey, Addr: addr}
	var regCenter = registry.GetRegistryCenter(registry.MysqlCenter, subscriptions)
	if err := regCenter.Keepalive(5); err != nil {
		hlog.Errorf(err.Error())
	}
	hlog.Infof("New Worker SendHeartbeatAndRegistryWorker Success addr %s", addr)
	return &pbapi.Empty{}, nil
}

// SendHeartbeat implements the HeartbeatImpl interface.
func (s *HeartbeatImpl) SendHeartbeat(ctx context.Context, req *pbapi.HeartbeatReq) (resp *pbapi.Empty, err error) {
	// TODO: Your code here...
	//p, ok := peer.FromContext(ctx)
	//if !ok {
	//	return &pbapi.Empty{}, fmt.Errorf("get peer failed")
	//}
	//ip, _, _ := net.SplitHostPort(p.Addr.String())
	hlog.Debugf("heartbeat host keepalive req %s", req)
	if req.Ip == "::1" || req.Ip == "::" {
		req.Ip = "127.0.0.1"
	}
	hlog.Debugf("recv hearbeat addr %s", fmt.Sprintf("%s:%d", req.Ip, req.Port))
	subscriptions := registry.Subscriptions{Ip: req.Ip, Port: req.Port, Addr: fmt.Sprintf("%s:%d", req.Ip, req.Port)}
	var registryStore registry.Center = &registry.Mysql{Subscriptions: subscriptions}

	if err := registryStore.Keepalive(5); err != nil {
		hlog.Errorf(err.Error())
	}

	return &pbapi.Empty{}, nil
}
