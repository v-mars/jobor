package raft

import (
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {
	var opts = Options{Bootstrap: true, HttpAddress: "0.0.0.0:2869",
		RaftTCPAddress: "127.0.0.1:2889", DataDir: "./node1", JoinAddress: ""}
	Server(opts)
	fmt.Println(".....")
}
func TestB(t *testing.T) {
	var opts = Options{Bootstrap: false, HttpAddress: "0.0.0.0:2868",
		RaftTCPAddress: "127.0.0.1:2888", DataDir: "./node2", JoinAddress: "127.0.0.1:2869"}
	Server(opts)
	fmt.Println(".....")
}
func TestC(t *testing.T) {
	var opts = Options{Bootstrap: false, HttpAddress: "0.0.0.0:2867",
		RaftTCPAddress: "127.0.0.1:2887", DataDir: "./node3", JoinAddress: "127.0.0.1:2868"}
	Server(opts)
	fmt.Println(".....")
}


func TestE(t *testing.T) {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:2869")
	fmt.Println(err)
	fmt.Println(tcpaddr)
}


