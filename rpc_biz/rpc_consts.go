package rpc_biz

import (
	"google.golang.org/grpc/keepalive"
	"time"
)

const (
	// DefaultRPCTimeout RPC connect timeout
	DefaultRPCTimeout = time.Second * 3
	// DefaultHeartbeatInterval worker send heartbeat ttl
	//DefaultHeartbeatInterval2       = time.Second * 5  //
	DefaultHeartbeatInterval        = time.Second * 5 // maxWorkerTTL int64 = 20
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
	MaxConnectionAge:      0 * time.Second,  // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var Kacp = keepalive.ClientParameters{
	Time:                5 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,     // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,            // send pings even without active streams
}
