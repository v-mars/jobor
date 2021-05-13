package proto

import (
	"jobor/internal/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

const (
	// DefaultRPCTimeout RPC connect timeout
	DefaultRPCTimeout = time.Second * 3
	// DefaultHeartbeatInterval worker send heartbeat ttl
	DefaultHeartbeatInterval        = time.Second * 15 // maxWorkerTTL int64 = 20
	DefaultLastFailHearBeatInterval = time.Second * 3
	// DefaultMaxRetryGetWorker max retry get host time for func Next
	DefaultMaxRetryGetWorker = 3
)

var Kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var Kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var Kacp = keepalive.ClientParameters{
	Time:                5 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}




func getgRPCClientConn(addr string) *grpc.ClientConn {
	//cg.Lock()
	//conn, exist := cg.conn[addr]
	//cg.Unlock()
	//if exist && conn.GetState() == connectivity.Ready {
	//	return conn
	//}
	//if conn != nil {
	//	conn.Close()
	//}

	return nil
}

type GrpcSession struct {
	exit chan bool
}

func DoWithTimeout(f func() (*pb.StreamResponse, error), ctx context.Context) (*pb.StreamResponse, error) {
	errChan := make(chan error, 1)
	Message := make(chan *pb.StreamResponse, 1)
	go func() {
		message, err := f()
		if err != nil {
			errChan <- err
		} else {
			Message <- message
		}
		//close(errChan)
	}()
	//t := time.NewTimer(5 * time.Second)
	select {
	case <-ctx.Done():
	//case <-t.C:
	//	return nil, fmt.Errorf("slow")
	case message := <-Message:
		return message, nil

	}
	return nil, fmt.Errorf("slow")
}




func ClientWithTimeout() {
	conn ,err := grpc.Dial(":20052",grpc.WithInsecure(), grpc.WithKeepaliveParams(Kacp))
	if err != nil{
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	tc := pb.NewTaskServiceClient(conn)
	taskCtx, ctxCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(10))
	defer ctxCancel()
	resStream, err := tc.RunTask(taskCtx,&pb.TaskRequest{TaskId: int32(1),TaskLang: "python3",TaskData: []byte("")})
	if err != nil {
		//logger.Jobor.Errorf("run task %s[%d] RunTasks err: %s", t.Name,t.ID, err)
		return
	}
	streamChan := func() (chan *pb.StreamResponse, chan error) {
		errChan := make(chan error, 1)
		Message := make(chan *pb.StreamResponse, 1)
		rec,errRecv := resStream.Recv()
		Message <- rec
		errChan <- errRecv
		return Message, errChan
	}
	var abort = make(chan bool)
	for {
		//rec,errRecv := resStream.Recv()
		rec,errChan:=streamChan()
		select {
		case <-abort:
			return
		case <-taskCtx.Done():
			return
		case d:=<-rec:
			_=d
		case e:=<-errChan:
			_=e
		}
	}
}