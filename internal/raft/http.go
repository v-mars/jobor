package raft

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/hashicorp/raft"
)

const (
	EnableWriteTrue  = int32(1)
	EnableWriteFalse = int32(0)
)

type httpServer struct {
	ctx         *stCachedContext
	log         *log.Logger
	mux         *http.ServeMux
	enableWrite int32
}

func NewHttpServer(ctx *stCachedContext, log *log.Logger) *httpServer {
	mux := http.NewServeMux()
	s := &httpServer{
		ctx:         ctx,
		log:         log,
		mux:         mux,
		enableWrite: EnableWriteFalse,
	}

	mux.HandleFunc("/set", s.doSet)
	mux.HandleFunc("/get", s.doGet)
	mux.HandleFunc("/delete", s.doDelete)
	mux.HandleFunc("/join", s.doJoin)
	mux.HandleFunc("/remove", s.doRemove)
	mux.HandleFunc("/member", s.Member)
	mux.HandleFunc("/stats", s.Stats)
	return s
}

func (h *httpServer) checkWritePermission() bool {
	return atomic.LoadInt32(&h.enableWrite) == EnableWriteTrue
}

func (h *httpServer) setWriteFlag(flag bool) {
	if flag {
		atomic.StoreInt32(&h.enableWrite, EnableWriteTrue)
	} else {
		atomic.StoreInt32(&h.enableWrite, EnableWriteFalse)
	}
}

func (h *httpServer) doGet(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	key := vars.Get("key")

	ret := h.Get(key)

	fmt.Fprintf(w, "%s\n", ret)
}

// doSet saves data to cache, only RaftNode master node provides this api
func (h *httpServer) doSet(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	key := vars.Get("key")
	value := vars.Get("value")

	if h.ctx.StCached.RaftNode.Raft.State() != raft.Leader {
		var leaderRaftTcpAddress = h.ctx.StCached.RaftNode.Raft.Leader()
		var leaderHttpAddress = ""
		for _,server:=range h.ctx.StCached.RaftNode.Raft.GetConfiguration().Configuration().Servers{
			if server.Address == leaderRaftTcpAddress { leaderHttpAddress= string(server.ID) }
		}
		if leaderHttpAddress == "" {
			h.log.Println("invalid leaderHttpAddress")
			fmt.Fprint(w, "invalid leaderHttpAddress\n")
			return
		}
		err := SetRaftCluster(leaderHttpAddress,key,value)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("%s\n",err))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("ok\n"))
		return
	}
	if !h.checkWritePermission() {
		fmt.Fprint(w, "write method not allowed\n")
		return
	}

	res := h.Set(key,value)

	//if key == "" || value == "" {
	//	h.Log.Println("doSet() error, get nil key or nil value")
	//	fmt.Fprint(w, "param error\n")
	//	return
	//}
	//
	//event := logEntryData{Key: key, Value: value}
	//eventBytes, err := json.Marshal(event)
	//if err != nil {
	//	h.Log.Printf("json.Marshal failed, err:%v", err)
	//	fmt.Fprint(w, "internal error\n")
	//	return
	//}
	//
	//applyFuture := h.ctx.StCached.RaftNode.RaftNode.Apply(eventBytes, 5*time.Second)
	//if err := applyFuture.Error(); err != nil {
	//	h.Log.Printf("RaftNode.Apply failed:%v", err)
	//	fmt.Fprint(w, "internal error\n")
	//	return
	//}
	//fmt.Fprintf(h.ctx.StCached.Log.Writer(), "ok\n")

	fmt.Fprintf(w, fmt.Sprintf("%s\n",res))
}

func (h *httpServer) doDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	key := vars.Get("key")
	if h.ctx.StCached.RaftNode.Raft.State() != raft.Leader {
		var leaderRaftTcpAddress = h.ctx.StCached.RaftNode.Raft.Leader()
		var leaderHttpAddress = ""
		for _,server:=range h.ctx.StCached.RaftNode.Raft.GetConfiguration().Configuration().Servers{
			if server.Address == leaderRaftTcpAddress { leaderHttpAddress= string(server.ID) }
		}
		if leaderHttpAddress == "" {
			h.log.Println("invalid leaderHttpAddress")
			fmt.Fprint(w, "invalid leaderHttpAddress\n")
			return
		}
		err := DeleteKeyRaftCluster(leaderHttpAddress,key)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("%s\n",err))
			return
		}
		fmt.Fprintf(w, "ok\n")
	}


	ret := h.Delete(key)

	fmt.Fprintf(w, "%s\n", ret)
}

func (h *httpServer) Get(key string) string {
	if key == "" {
		h.log.Println("doGet() error, get nil key")
		return ""
	}
	ret := h.ctx.StCached.CacheManager.Get(key)
	return fmt.Sprintf(ret)
}

func (h *httpServer) Set (key,value string) string {
	if !h.checkWritePermission() {
		return "write method not allowed"
	}
	if key == "" || value == "" {
		h.log.Println("doSet() error, get nil key or nil value")
		return "param error"
	}

	event := logEntryData{Key: key, Value: value}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		h.log.Printf("json.Marshal failed, err:%v", err)
		return "internal error"
	}

	applyFuture := h.ctx.StCached.RaftNode.Raft.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		h.log.Printf("RaftNode.Apply failed:%v", err)
		return "internal error"
	}

	return "ok"
}

func (h *httpServer) Delete (key string) string {
	if !h.checkWritePermission() {
		return "write method not allowed\n"
	}
	if key == "" {
		h.log.Println("Delete() error, get nil key")
		return ""
	}
	h.ctx.StCached.CacheManager.Delete(key)
	return "ok"
}

// doJoin handles joining cluster request
func (h *httpServer) doJoin(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	peerAddress := vars.Get("peerAddress")
	if peerAddress == "" {
		h.log.Println("invalid PeerAddress")
		fmt.Fprint(w, "invalid peerAddress\n")
		return
	}
	httpAddress := vars.Get("httpAddress")
	if httpAddress == "" && h.ctx.StCached.RaftNode.Raft.State() != raft.Leader {
		h.log.Println("invalid httpAddress")
		fmt.Fprint(w, "invalid httpAddress\n")
		return
	}
	if h.ctx.StCached.RaftNode.Raft.State() != raft.Leader {
		var leaderRaftTcpAddress = h.ctx.StCached.RaftNode.Raft.Leader()
		var leaderHttpAddress = ""
		for _,server:=range h.ctx.StCached.RaftNode.Raft.GetConfiguration().Configuration().Servers{
			if server.Address == leaderRaftTcpAddress { leaderHttpAddress= string(server.ID) }
		}
		if leaderHttpAddress == "" {
			h.log.Println("invalid leaderHttpAddress")
			fmt.Fprint(w, "invalid leaderHttpAddress\n")
			return
		}
		err := JoinRaftCluster(Options{HttpAddress: httpAddress,RaftTCPAddress: peerAddress,JoinAddress: leaderHttpAddress})
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("%s\n",err))
			return
		}
		fmt.Fprint(w, "ok")
		return
	}

	ServerID:=fmt.Sprintf("%s:%s", strings.Split(peerAddress,":")[0],
		strings.Split(httpAddress,":")[1])
	addPeerFuture := h.ctx.StCached.RaftNode.Raft.AddVoter(
		raft.ServerID(ServerID),
		raft.ServerAddress(peerAddress),
		0, 0)
	if err := addPeerFuture.Error(); err != nil {
		h.log.Printf("Error joining peer to RaftNode, peeraddress:%s, err:%v, code:%d", peerAddress, err, http.StatusInternalServerError)
		fmt.Fprint(w, "internal error\n")
		return
	}
	fmt.Fprint(w, "ok")
}

// doRemove handles remove cluster request
func (h *httpServer) doRemove(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	serverID := vars.Get("serverID")
	if serverID == "" {
		h.log.Println("invalid ID")
		fmt.Fprint(w, "invalid ID\n")
		return
	}

	if h.ctx.StCached.RaftNode.Raft.State() != raft.Leader {
		var leaderRaftTcpAddress = h.ctx.StCached.RaftNode.Raft.Leader()
		var leaderHttpAddress = ""
		for _,server:=range h.ctx.StCached.RaftNode.Raft.GetConfiguration().Configuration().Servers{
			if server.Address == leaderRaftTcpAddress { leaderHttpAddress= string(server.ID) }
		}
		if leaderHttpAddress == "" {
			h.log.Println("invalid leaderHttpAddress")
			fmt.Fprint(w, "invalid leaderHttpAddress\n")
			return
		}
		err := RemoveRaftCluster(Options{HttpAddress: serverID,JoinAddress: leaderHttpAddress})
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("%s\n",err))
			return
		}
		fmt.Fprint(w, "ok")
		return
	}

	removeServerFuture := h.ctx.StCached.RaftNode.Raft.RemoveServer(raft.ServerID(serverID),0,0)
	//addPeerFuture := h.ctx.StCached.RaftNode.RaftNode.AddVoter(RaftNode.ServerID(serverAddress), RaftNode.ServerAddress(peerAddress), 0, 0)
	if err := removeServerFuture.Error(); err != nil {
		h.log.Printf("Error remove server to RaftNode, ID:%s, err:%v, code:%d", serverID, err, http.StatusInternalServerError)
		fmt.Fprint(w, "internal error\n")
		return
	}
	fmt.Fprint(w, "ok")
}

// Member handles remove cluster request
func (h *httpServer) Member(w http.ResponseWriter, r *http.Request) {
	type serverList struct {
		ServerID      string `json:"serverID"`
		ServerAddress string `json:"serverAddress"`
		IsLeader      bool   `json:"isLeader"`
	}
	var servers [] serverList
	var LeaderAddress = h.ctx.StCached.RaftNode.Raft.Leader()
	for _,s := range h.ctx.StCached.RaftNode.Raft.GetConfiguration().Configuration().Servers {
		server := serverList{ServerID: string(s.ID),ServerAddress: string(s.Address)}
		if LeaderAddress == s.Address { server.IsLeader = true }
		servers=append(servers,server)
	}
	w.Header().Set("content-type","text/json")
	msg, err := json.Marshal(servers)
	if err!=nil{
		fmt.Fprint(w, fmt.Sprintf("data parse json err: %s", err))
		return
	}
	w.Write(msg)
}

// Stats handles remove cluster request
func (h *httpServer) Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	msg, err := json.Marshal(h.ctx.StCached.RaftNode.Raft.Stats())
	if err!=nil{
		fmt.Fprint(w, fmt.Sprintf("data parse json err: %s", err))
		return
	}
	w.Write(msg)
}