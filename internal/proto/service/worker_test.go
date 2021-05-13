package service

import (
	"jobor/internal/config"
	"log"
	"testing"
)

func TestClientGRPCTest(t *testing.T) {
	if err := config.LoadWorkerConfig("../../../configs/worker.toml");err != nil {
		log.Fatal(err)
	}
	if err:=WorkerGRPC();err!=nil{
		log.Fatal(err)
	}
}
