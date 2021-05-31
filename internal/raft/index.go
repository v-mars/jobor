package raft

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type stCached struct {
	HttpServer   *httpServer
	Opts         Options
	Log          *log.Logger
	CacheManager *cacheManager
	RaftNode     *raftNodeInfo
}

type stCachedContext struct {
	StCached *stCached
}

var (
	St = stCached{}
)


func Server(opts Options) {
	St.Opts = opts
	St.Log = log.New(os.Stderr, "stCached: ", log.Ldate|log.Ltime)
	St.CacheManager = NewCacheManager()
	ctx := &stCachedContext{&St}
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", St.Opts.HttpAddress)
	if err != nil {
		St.Log.Fatal(fmt.Sprintf("listen %s failed: %s", St.Opts.HttpAddress, err))
	}
	St.Log.Printf("http server listen:%s", l.Addr())

	logger := log.New(os.Stderr, "httpserver: ", log.Ldate|log.Ltime)
	httpServer := NewHttpServer(ctx, logger)
	St.HttpServer = httpServer
	go func() {
		http.Serve(l, httpServer.mux)
	}()

	raft, err := newRaftNode(St.Opts, ctx)
	if err != nil {
		St.Log.Fatal(fmt.Sprintf("new RaftNode node failed:%v", err))
	}
	St.RaftNode = raft

	if St.Opts.JoinAddress != "" {
		err = JoinRaftCluster(St.Opts)
		if err != nil {
			St.Log.Fatal(fmt.Sprintf("join RaftNode cluster failed:%v", err))
		}

	}

	/*go func() {
		for{
			//_ = StCached.RaftNode.RaftNode.
			//fmt.Println("RaftNode:",StCached.RaftNode.RaftNode.String())
			fmt.Println("RaftNode Leader,Id,state:",StCached.RaftNode.raft.Leader(),StCached.RaftNode.raft.State())
			fmt.Println("RaftNode Servers:",StCached.RaftNode.raft.GetConfiguration().Configuration().Servers)
			time.Sleep(30*time.Second)
		}
	}()*/

	// monitor leadership
	for {
		select {
		case leader := <-St.RaftNode.leaderNotifyCh:
			if leader {
				St.Log.Println("become leader, enable write api")
				St.HttpServer.setWriteFlag(true)
				//StCached.HttpServer.Set("ping","pong")
				//StCached.HttpServer.Set(StCached.Opts.RaftTCPAddress,fmt.Sprintf("%s:%s",
				//	strings.Split(StCached.Opts.RaftTCPAddress,":")[0],strings.Split(StCached.Opts.HttpAddress,":")[1]))
			} else {
				St.Log.Println("become follower, close write api")
				St.HttpServer.setWriteFlag(false)
				//leaderHttpAddr := StCached.HttpServer.Get(StCached.Opts.RaftTCPAddress)
				//err = SetRaftCluster(leaderHttpAddr,StCached.Opts.RaftTCPAddress,fmt.Sprintf("%s:%s",
				//	strings.Split(StCached.Opts.RaftTCPAddress,":")[0],strings.Split(StCached.Opts.HttpAddress,":")[1]))
				//if err != nil {
				//	StCached.Log.Fatal(fmt.Sprintf("set key RaftNode cluster failed:%v", err))
				//	return
				//}
			}
		}
	}
}