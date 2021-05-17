package service

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"jobor/internal/app/jobor/dispatcher"
	"jobor/internal/config"
	"jobor/internal/middleware"
	"jobor/internal/proto"
	"jobor/internal/proto/pb"
	"jobor/internal/proto/registry"
	"jobor/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"sync"
	"time"
)


type myServer struct{}

func (s *myServer) RunTask(request *pb.TaskRequest, server pb.TaskService_RunTaskServer) error {
	//fmt.Printf("request: %++v\n", request)
	defer func() {logger.Jobor.Debugf("task=[%d] lang=%s worker run task finish.", request.TaskId, request.TaskLang)}()
	taskCtx, taskCancel := context.WithCancel(server.Context())
	defer taskCancel()
	runner, err := dispatcher.GetDataRun(request)
	if err != nil {
		logger.Jobor.Errorf("dispatcher getDataRun err: %s", err.Error())
		return err
	}
	out:=runner.Run(taskCtx)
	logger.Jobor.Debugf("task=[%d] lang=%s runner run start", request.TaskId, request.TaskLang)
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			logger.Jobor.Errorf("runner run close failed: %s", err)
		}
	}(out)
	var buf = make([]byte, 1024)
	for {
		n, err := out.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			// if read failed please send default err code -1
			logger.Jobor.Errorf("read failed from %s", err)
			err = server.Send(&pb.StreamResponse{Resp: []byte(err.Error() + fmt.Sprintf("%3d", dispatcher.DefaultExitCode))})
			if err != nil {
				logger.Jobor.Errorf("send failed: %s",err)
			}
			return nil
		}
		if n > 0 {
			resp := pb.StreamResponse{Resp: buf[:n]}
			err = server.Send(&resp)
			if err != nil {
				logger.Jobor.Errorf("stream send failed: %s", err)
				return nil
			}
		}
	}
	//return nil
}

// GetStream 服务端 单向流
func (s *myServer) GetStream(data *pb.StreamReqData, server pb.TaskService_GetStreamServer) error {
	i:= 0
	for{
		i++
		//if len(payload) > data.maxSendMessageSize {
		//	return status.Errorf(codes.ResourceExhausted, "trying to send message larger than max (%d vs. %d)", len(payload), ss.maxSendMessageSize)
		//}
		server.Send(&pb.StreamResData{Data:fmt.Sprintf("%v",time.Now().Unix())})
		time.Sleep(10*time.Second)
		if i >10 {
			break
		}
	}

	return nil
}

// PutStream 客户端 单向流
func (s *myServer) PutStream(cliStr pb.TaskService_PutStreamServer) error {
	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println("server PutStream",tem)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

// AllStream 客户端服务端 双向流
func (s *myServer) AllStream(allStr pb.TaskService_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for {
		data, err := allStr.Recv()
		if err!=nil{
			log.Println("server AllStream err:",err)
			break
		}else {
			log.Println("server AllStream:",data)
		}
	}
	//wg.Done()

	//go func() {
	//	for {
	//		data, _ := allStr.Recv()
	//		log.Println("server AllStream:",data)
	//	}
	//	wg.Done()
	//}()


	//go func() {
	//	for {
	//		allStr.Send(&pb.StreamResData{Data:"ssss"})
	//		time.Sleep(time.Second)
	//	}
	//	wg.Done()
	//}()

	//wg.Wait()
	return nil
}

// RunRPC 这里myServer实现了SayHello
func (s *myServer) RunRPC(ctx context.Context, in *pb.TaskRequest) (*pb.Response, error) {
	fmt.Println("receive " + in.Name)
	return &pb.Response{
		Code:pb.Response_SUCCESS,
		Msg:&pb.Reply{Message:"receive " + in.Name},
	}, nil
}



func WorkerGRPC() error {
	//绑定端口
	var (
		WorkerGRPCPort = fmt.Sprintf("%s:%d", config.WorkerConfig.IP,config.WorkerConfig.Port)
	)

	lis, err := net.Listen("tcp", WorkerGRPCPort)
	if err != nil{
		log.Fatal("gRPC fail to listen")
		return err
	}
	serverOptions := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			middleware.RecoveryInterceptor,
			middleware.LoggerInterceptor,
			//middleware.CheckSecretInterceptor,
		),
		grpc.MaxRecvMsgSize(1024 * 1024 * 18),
		grpc.MaxSendMsgSize(math.MaxInt32),
		grpc.KeepaliveEnforcementPolicy(proto.Kaep),
		grpc.KeepaliveParams(proto.Kasp),
	}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterTaskServiceServer(grpcServer, &myServer{})
	reflection.Register(grpcServer)
	_, port, _ := net.SplitHostPort(lis.Addr().String())
	intPort, _ := strconv.Atoi(port)

	go registry.SendHeartbeatAndRegistryWorker(int32(intPort),5)

	//select {}
	return grpcServer.Serve(lis)
}

