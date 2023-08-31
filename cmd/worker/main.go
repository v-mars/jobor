package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/conf"
	"jobor/worker"
)

func init() {

}
func main() {
	flag.StringVar(&conf.FlagWorker, "f", "./worker.yaml", "config path, eg: -c conf/worker.yaml")
	flag.Parse()
	go func() {
		hlog.Fatal(worker.StartWorkerRpc())
	}()
	worker.MQWorker()
}
