package worker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/q"
	"jobor/conf"
)

func MQWorker() {
	wc := conf.GetWorkerConf()
	srv, err := q.InitQSrv(&wc.Redis, q.Queue)
	if err != nil {
		hlog.Error(err)
		return
	}
	//worker := srv.NewCustomQueueWorker("worker_name", 10, "default")
	worker := srv.NewWorker("worker_name", 10)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		hlog.Errorf("I am an error handler:", err)
	}

	pretaskhandler := func(signature *tasks.Signature) {
		hlog.Infof("I am a start of task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		hlog.Infof("I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)
	err = worker.Launch()
	fmt.Println("worker is start")
	if err != nil {
		// do something with the error
		hlog.Error(err)
	}
}
