package dispatcher

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"jobor/internal/models/tbs"
	"jobor/internal/proto"
	"jobor/pkg/logger"
	"time"
)

// TryGetGrpcClientConn try Get service client conn
// Scheduling Algorithm
// - random
// - LeastTask
// - Weight
// - roundRobin
// get rpc conn
func TryGetGrpcClientConn(ctx context.Context, worker JoborWorker) (*grpc.ClientConn,*tbs.JoborWorker, error) {
	var (
		err  error
		conn *grpc.ClientConn
	)
	for i := 0; i < proto.DefaultMaxRetryGetWorker; i++ {
		w := worker()
		if w == nil {
			err = fmt.Errorf("can't get valid worker")
			continue
		}
		conn, err = getGrpcClientConn(ctx, w.Addr)
		if err != nil {
			logger.Jobor.Errorf("GetGrpcClientConn addr[%s] failed: %s", w.Addr,err)
			continue
		}
		// when only conn is Ready, direct return this conn,otherse
		if conn.GetState() == connectivity.Ready || conn.GetState() == connectivity.Idle {
			return conn,w, nil
		}
		err = conn.Close()
		if err != nil {
			return nil,w, err
		}
	}
	return nil,nil, fmt.Errorf("can't get valid workers")
}


// getGrpcClientConn Get Grpc Client Conn
func getGrpcClientConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	//conn := cachegRPCConnM.getgRPCClientConn(addr)
	//conn ,errDial := service.Dial(addr,service.WithInsecure()) // service.WithKeepaliveParams(proto.Kacp)
	//if errDial != nil{
	//	logger.Jobor.Errorf("%s",errDial)
	//	return nil, errDial
	//}

	var (
		conn *grpc.ClientConn
		c   credentials.TransportCredentials
		err error
	)

	dialOptions := []grpc.DialOption{
		//service.WithPerRPCCredentials(
		//	&Auth{SecretToken: "token"},
		//),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(18 * 1024 * 1024)), // 16M
		grpc.WithBlock(),
		//service.WithKeepaliveParams(keepalive.ClientParameters{}),
		//grpc.WithKeepaliveParams(proto.Kacp),
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoff.Config{MaxDelay: time.Second * 2}, MinConnectTimeout: time.Second * 2}),
	}

	if false {
		c, err = credentials.NewClientTLSFromFile("", "")
		if err != nil {
			logger.Jobor.Errorf("credentials.NewClientTLSFromFile failed: %s", err)
			return nil, err
		}
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(c))
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	//rpcCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	//defer cancel()

	conn ,err = grpc.Dial(addr,grpc.WithInsecure())
	//conn, err = grpc.DialContext(rpcCtx, addr, dialOptions...)
	if err != nil {return nil, err}
	//fmt.Println("conn state:", conn.GetState())
	//cachegRPCConnM.addgRPCClientConn(addr, conn)
	return conn, nil
}
