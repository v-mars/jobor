package dispatcher

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/model"
	"math/rand"
	"time"
)

const (
	// Random 1:Random 2:RoundRobin 3:Weight 4:LeastTask
	Random int = iota + 1
	RoundRobin
	Weight
	LeastTask
)

// Select a next run worker

// JoborWorker will return next run worker
// if JoborWorker is nil,because not find valid worker
type JoborWorker func() *model.JoborWorker

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetWorkerByRoutePolicy return a type JoborWorker, it will return a worker
func GetWorkerByRoutePolicy(routingKeys []string, routePolicy int, lang string) JoborWorker {
	switch routePolicy {
	case Random:
		return random(routingKeys, lang)
	case RoundRobin:
		return roundRobin(routingKeys, lang)
	case Weight:
		return weight(routingKeys, lang)
	case LeastTask:
		return leastTask(routingKeys, lang)
	default:
		return random(routingKeys, lang)
	}
}

// 随机调度
func random(routingKeys []string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKeys, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}

// 轮询调度
func roundRobin(routingKeys []string, lang string) JoborWorker {
	var i = rand.Int()
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKeys, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		worker := workers[i%len(workers)]
		i++
		return &worker
	}
}

// 权重调度
func weight(routingKeys []string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKeys, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		allWeight := 0

		for _, w := range workers {
			allWeight += int(w.Weight)
		}
		get := rand.Int() % allWeight
		pre := 0

		for _, w := range workers {
			if pre <= get && get < pre+int(w.Weight) {
				return &w
			}
			pre += int(w.Weight)
		}

		return nil
	}
}

// 最少调度
func leastTask(routingKeys []string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKeys, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}
