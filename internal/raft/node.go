package raft

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/raft"
	raftBoltdb "github.com/hashicorp/raft-boltdb"
)

const (
	LogStoreFileName = "RaftNode-Log.bolt"
	StableStoreFileName = "RaftNode-stable.bolt"
)

type raftNodeInfo struct {
	raft           *raft.Raft
	fsm            *FSM
	leaderNotifyCh chan bool
	nodeNotifyCh   chan bool
	raftConfig     *raft.Config
	a     *raft.AppendPipeline
}

func newRaftTransport(opts Options) (*raft.NetworkTransport, error) {
	address, err := net.ResolveTCPAddr("tcp", opts.RaftTCPAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}
	return transport, nil
}

func newRaftNode(opts Options, ctx *stCachedContext) (*raftNodeInfo, error) {
	raftConfig := raft.DefaultConfig()
	ServerID:=fmt.Sprintf("%s:%s", strings.Split(opts.RaftTCPAddress,":")[0],
		strings.Split(opts.HttpAddress,":")[1])
	raftConfig.LocalID = raft.ServerID(ServerID)
	//raftConfig.LocalID = raft.ServerID(MD5HashString16(Opts.RaftTCPAddress))
	//raftConfig.LocalID = RaftNode.ServerID(Opts.RaftTCPAddress)
	//raftConfig.Logger = Log.New(os.Stderr, "RaftNode: ", Log.Ldate|Log.Ltime)
	//logOpts :=hclog.DefaultOptions
	//logOpts.Name="RaftNode: "
	//raftConfig.Logger = hclog.New(logOpts)
	//raftConfig.SnapshotInterval = 20 * time.Second
	//raftConfig.SnapshotThreshold = 2
	leaderNotifyCh := make(chan bool, 1)
	raftConfig.NotifyCh = leaderNotifyCh

	transport, err := newRaftTransport(opts)
	if err != nil {
		return nil, err
	}

	if err = os.MkdirAll(opts.DataDir, 0700); err != nil {
		return nil, err
	}

	fsm := &FSM{
		ctx: ctx,
		log: log.New(os.Stderr, "FSM: ", log.Ldate|log.Ltime),
	}
	snapshotStore, err := raft.NewFileSnapshotStore(opts.DataDir, 1, os.Stderr)
	if err != nil {
		return nil, err
	}

	logStore, err := raftBoltdb.NewBoltStore(filepath.Join(opts.DataDir, LogStoreFileName))
	if err != nil {
		return nil, err
	}

	stableStore, err := raftBoltdb.NewBoltStore(filepath.Join(opts.DataDir, StableStoreFileName))
	if err != nil {
		return nil, err
	}

	raftNode, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	if opts.Bootstrap {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		raftNode.BootstrapCluster(configuration)
	}

	return &raftNodeInfo{raft: raftNode, fsm: fsm, leaderNotifyCh: leaderNotifyCh,raftConfig:raftConfig}, nil
}

// JoinRaftCluster joins a node to RaftNode cluster
func JoinRaftCluster(opts Options) error {
	url := fmt.Sprintf("http://%s/join?peerAddress=%s&httpAddress=%s",
		opts.JoinAddress,opts.RaftTCPAddress,opts.HttpAddress)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) != "ok" {
		return errors.New(fmt.Sprintf("Error joining cluster: %s", body))
	}

	return nil
}

// RemoveRaftCluster remove a node to RaftNode cluster
func RemoveRaftCluster(opts Options) error {
	url := fmt.Sprintf("http://%s/remove?serverID=%s",opts.JoinAddress,opts.HttpAddress)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) != "ok" {
		return errors.New(fmt.Sprintf("Error remove cluster: %s", body))
	}
	return nil
}

// SetRaftCluster  a node to RaftNode cluster
func SetRaftCluster(leadHttpAddress,key,value string) error {
	url := fmt.Sprintf("http://%s/set?key=%s&value=%s", leadHttpAddress, key,value)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if strings.TrimRight(string(body),"\n") != "ok" {
		return errors.New(fmt.Sprintf("Error set key cluster: %s", body))
	}

	return nil
}

// DeleteKeyRaftCluster  a node to RaftNode cluster
func DeleteKeyRaftCluster(leadHttpAddress,key string) error {
	url := fmt.Sprintf("http://%s/delete?key=%s", leadHttpAddress, key)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if strings.TrimRight(string(body),"\n") != "ok" {
		return errors.New(fmt.Sprintf("Error set key cluster: %s", body))
	}

	return nil
}


// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值 返回一个32位md5加密后的字符串
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// MD5HashString16 返回一个16位md5加密后的字符串
func MD5HashString16(data string) string{
	return MD5HashString(data)[8:24]
}