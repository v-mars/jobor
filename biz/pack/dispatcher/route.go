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
func GetWorkerByRoutePolicy(routingKey string, routePolicy int, lang string) JoborWorker {
	switch routePolicy {
	case Random:
		return random(routingKey, lang)
	case RoundRobin:
		return roundRobin(routingKey, lang)
	case Weight:
		return weight(routingKey, lang)
	case LeastTask:
		return leastTask(routingKey, lang)
	default:
		return random(routingKey, lang)
	}
}

// 随机调度
func random(routingKey string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKey, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}

// 轮询调度
func roundRobin(routingKey string, lang string) JoborWorker {
	var i = rand.Int()
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKey, lang)
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
func weight(routingKey string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKey, lang)
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
func leastTask(routingKey string, lang string) JoborWorker {
	return func() *model.JoborWorker {
		workers, err := model.GetWorkers(routingKey, lang)
		if err != nil {
			err = fmt.Errorf("get online worker failed: %s", err)
			hlog.Errorf(err.Error())
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}
