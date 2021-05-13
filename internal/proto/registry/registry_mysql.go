package registry

import (
	"jobor/internal/app/jobor/worker"
	"jobor/internal/models/tbs"
	"jobor/internal/proto"
	"fmt"
	"time"
)



type Mysql struct {
	Subscriptions
}

func (m *Mysql) Registry(port int32,ttl int64) error {

	//var d = time.Second * time.Duration(ttl)
	ticker := time.NewTicker(proto.DefaultHeartbeatInterval)
	go func() {
		for {
			//_, err = hbClient.SendHb(context.Background(), &pb.HeartbeatReq{Port: port})
			//if err != nil {
			//	time.Sleep(time.Second)
			//	fmt.Printf("send hearbeat failed：%s", err)
			//	return
			//}
			fmt.Println("get ticker", time.Now().Format("2006-01-02 15:04:05"))
			<-ticker.C
		}
	}()

	return nil
}

// Keepalive
//保持服务器与mysql的同步更新
func (m *Mysql) Keepalive(interval int64) error {
	var data = tbs.JoborWorker{Hostname: m.Hostname,Addr: m.Addr,Version: m.Version,RoutingKey: m.RoutingKey,
		Weight: m.Weight,LeaseUpdate: time.Now().Unix(),Status: "running"}
	return worker.CreateOrUpdate(data)
}

func (m *Mysql) UnRegistry() error {
	var data = tbs.JoborWorker{Hostname: m.Hostname,Addr: m.Ip,Version: m.Version,RoutingKey: m.RoutingKey,
		Weight: m.Weight,LeaseUpdate: time.Now().Unix(),Status: "offline"}
	return worker.CreateOrUpdate(data)
}