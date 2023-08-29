package main

import (
	"flag"
	"jobor/conf"
	"jobor/worker"
)

func init() {

}
func main() {
	flag.StringVar(&conf.FlagWorker, "f", "./worker.yaml", "config path, eg: -c conf/worker.yaml")
	flag.Parse()
	worker.MQWorker()
}
