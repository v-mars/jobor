package registry

import (
	"jobor/internal/config"
	"jobor/internal/proto"
	"jobor/internal/proto/pb"
	"jobor/internal/response"
	"jobor/pkg/logger"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"time"
)

const (
	MysqlCenter = "mysql"
	RedisCenter = "redis"
	EtcdCenter = "etcd"

)


type Subscriptions struct {
	Addr       string `json:"addr,omitempty"`
	Ip         string `json:"ip,omitempty"`
	Port       int32  `json:"port,omitempty"`
	Weight     int32  `json:"weight,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	Version    string `json:"version,omitempty"`
	HostGroup  string `json:"hostGroup,omitempty"`
	Remark     string `json:"remark,omitempty"`
	RoutingKey string `json:"routingKey,omitempty"`
}

type Center interface {
	Registry (port int32,ttl int64) error
	Keepalive (ttl int64) error
	UnRegistry () error
}

func GetRegistryCenter(regCenter string, s Subscriptions) Center {
	switch regCenter {
	case MysqlCenter:
		return &Mysql{Subscriptions: s}
	case RedisCenter:
		return &Redis{Subscriptions: s}
	case EtcdCenter:
		return &Etcd{Subscriptions: s}
	}
	return &Mysql{}
}


func SendHeartbeatAndRegistryWorker(port int32,ttl int64) {
	rand.Seed(time.Now().UnixNano())
	var (
		lastAddr string
	)
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	for {
		var Servers = config.WorkerConfig.Servers
		if len(Servers) == 0 {
			log.Fatal("server address is empty")
			// cancel()
			return
		}
		// do not get last addr
		for {
			getAddr := Servers[rand.Int()%len(Servers)]
			// do not get failed addr
			if getAddr != lastAddr || len(Servers) == 1 {
				lastAddr = getAddr
				break
			}
		}

		conn, err := grpc.DialContext(ctx,lastAddr,grpc.WithInsecure())
		if err != nil {
			log.Fatal("gRPC client dial error: ", err)
			return
		}
		heartbeatClient := pb.NewHeartbeatClient(conn)

		regHost := pb.RegistryReq{
			Port:      port,
			Hostname:  config.WorkerConfig.Name,
			Version:   config.WorkerConfig.Version,
			Weight:    config.WorkerConfig.Weight,
			Remark:    "",
			RoutingKey:    config.WorkerConfig.RoutingKey,
		}
		_, err = heartbeatClient.RegistryWorker(context.Background(), &regHost)
		if err != nil {
			time.Sleep(time.Second)
			logger.Infof("worker registry failed, err: %s", err)
			return

		}
		logger.Infof("worker registry success, server: %s", lastAddr)

		//var d = time.Second * time.Duration(ttl)
		ticker := time.NewTicker(proto.DefaultHeartbeatInterval)
		canNotConn := 0
		for {
			select {
			case <-ticker.C:
				ctxBeat, cancel := context.WithTimeout(context.Background(), proto.DefaultRPCTimeout)
				hbReq := &pb.HeartbeatReq{
					Port: port,
				}

				_, err := heartbeatClient.SendHeartbeat(ctxBeat, hbReq)
				if err != nil {
					cancel()
					//err := DealRPCErr(err)
					statusErr, ok := status.FromError(err)
					if ok && statusErr.Code() == codes.Unavailable && canNotConn > 1 {
						if len(config.WorkerConfig.Servers) >= 2 {
							logger.Debug("can not connect server,change other server")
							conn.Close()
							goto Next
						}
					} else {
						canNotConn++
					}
					logger.Errorf("client.SendHeartbeat failed, err: %s", err)
					ticker.Reset(proto.DefaultLastFailHearBeatInterval)
					continue
				}
				canNotConn = 0
				cancel()
				logger.Debugf("send heartbeat success, server %s", lastAddr)
				ticker.Reset(proto.DefaultHeartbeatInterval)
				//case <-stopHeatBeat:
				//	log.Info("Stop Send HearBeat")
				//	timer.Stop()
			}

		}
		Next:
			time.Sleep(time.Second*time.Duration(1))
	}

}

// DealRPCErr change rpc error to err code
func DealRPCErr(err error) error {
	statusErr, ok := status.FromError(err)
	if ok {
		switch statusErr.Code() {
		case codes.DeadlineExceeded:
			return response.GetMsgErr(response.ErrCtxDeadlineExceeded)
		case codes.Canceled:
			return response.GetMsgErr(response.ErrCtxCanceled)
		case codes.Unauthenticated:
			return response.GetMsgErr(response.ErrRPCUnauthenticated)
		case codes.Unavailable:
			return response.GetMsgErr(response.ErrRPCUnavailable)
		}
	}
	return err
}