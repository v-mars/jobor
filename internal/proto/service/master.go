package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"jobor/internal/proto"
	"jobor/internal/proto/pb"
	"jobor/internal/proto/registry"
	"jobor/pkg/logger"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)


const (
	MasterGRPCPort = ":50052"
	//fmt.Sprintf("%s:%d", config.WorkerConfig.IP,config.WorkerConfig.Port)
)

// Auth check rpc request valid
type Auth struct {
	SecretToken string
}

// HeartbeatService implementation proto Heartbeat interface
type HeartbeatService struct {
	Auth Auth
}

// RegistryWorker worker registry
func (hs *HeartbeatService) RegistryWorker(ctx context.Context, req *pb.RegistryReq) (*pb.Empty, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.Empty{}, fmt.Errorf("registry failed")
	}
	ip, _, _ := net.SplitHostPort(p.Addr.String())
	logger.Debugf("`registryHost` new worker req %s", req)
	req.Ip = ip

	if req.Ip=="::1"{req.Ip="127.0.0.1"}
	addr := fmt.Sprintf("%s:%d", req.Ip, req.Port)

	subscriptions := registry.Subscriptions{Ip: req.Ip,Hostname: req.Hostname,Port: req.Port,Weight: req.Weight,
		Version: req.Version,RoutingKey: req.RoutingKey,Addr: addr}
	var regCenter = registry.GetRegistryCenter(registry.MysqlCenter, subscriptions)
	if err:= regCenter.Keepalive(5);err!=nil{
		logger.Fatal(err)
	}
	logger.Infof("New Worker SendHeartbeatAndRegistryWorker Success addr %s", addr)
	return &pb.Empty{}, nil
}

// SendHeartbeat recv heatneat from client
func (hs *HeartbeatService) SendHeartbeat(ctx context.Context, hb *pb.HeartbeatReq) (*pb.Empty, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.Empty{}, fmt.Errorf("get peer failed")
	}
	ip, _, _ := net.SplitHostPort(p.Addr.String())
	if ip=="::1"{ip="127.0.0.1"}
	logger.Debugf("recv hearbeat addr %s",fmt.Sprintf("%s:%d", ip, hb.Port))
	subscriptions := registry.Subscriptions{Ip: ip,Port: hb.Port, Addr: fmt.Sprintf("%s:%d", ip, hb.Port)}
	var registryStore registry.Center = &registry.Mysql{Subscriptions: subscriptions}

	if err:= registryStore.Keepalive(5);err!=nil{
		logger.Fatal(err)
	}

	return &pb.Empty{}, nil
}


func MasterGRPC() error {
	//绑定端口
	lis, err := net.Listen("tcp", MasterGRPCPort)
	if err != nil{
		log.Fatal("gRPC fail to listen")
		return err
	}
	// 新建gRPC服务器实例
	// 默认单次接收最大消息长度为`1024*1024*4`bytes(4M)，单次发送消息最大长度为`math.MaxInt32`bytes
	// grpcServer := service.NewServer(service.MaxRecvMsgSize(1024*1024*4), service.MaxSendMsgSize(math.MaxInt32))
	serverOptions := []grpc.ServerOption{
		//grpcMiddleware.WithUnaryServerChain(
			//middleware.RecoveryInterceptor,
			//middleware.LoggerInterceptor,
			//middleware.CheckSecretInterceptor,
		//),
		grpc.MaxRecvMsgSize(1024 * 1024 * 18),
		grpc.MaxSendMsgSize(math.MaxInt32),
		grpc.KeepaliveEnforcementPolicy(proto.Kaep),
		grpc.KeepaliveParams(proto.Kasp),
	}

	gRPCServer := grpc.NewServer(serverOptions...)
	pb.RegisterHeartbeatServer(gRPCServer, &HeartbeatService{})
	reflection.Register(gRPCServer)
	go tryDisConn(gRPCServer)
	return gRPCServer.Serve(lis)
}


// tryDisConn will close service and http conn
// if time rather than 10s, will immediately close
func tryDisConn(gRPCServer *grpc.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGKILL,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT,
	)

	select {
	case sig := <-signals:
		go func() {
			select {
			case <-time.After(time.Second * 10):
				logger.Jobor.Warn("Shutdown gracefully timeout, application will shutdown immediately.")
				os.Exit(0)
			}
		}()
		logger.Jobor.Info(fmt.Sprintf("get signal %s, application service will shutdown.", sig))
		//dispatcher.DoStopConn(mode)

		logger.Jobor.Debug("start stop grpcServer")
		gRPCServer.Stop()
		//if mode == define.Server {
		//	logger.Debug("Start Stop HttpServer")
		//	httpServer.Shutdown(context.Background())
		//}
		os.Exit(0)
	}
}