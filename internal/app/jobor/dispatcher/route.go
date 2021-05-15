package dispatcher

import (
	"jobor/internal/models/tbs"
	"jobor/pkg/logger"
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
type JoborWorker func() *tbs.JoborWorker

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetWorkerByRoutePolicy return a type JoborWorker, it will return a worker
func GetWorkerByRoutePolicy(routingKey string, routePolicy int) JoborWorker {
	switch routePolicy {
	case Random:
		return random(routingKey)
	case RoundRobin:
		return roundRobin(routingKey)
	case Weight:
		return weight(routingKey)
	case LeastTask:
		return leastTask(routingKey)
	default:
		return random(routingKey)
	}
}

// 随机调度
func random(routingKey string) JoborWorker {
	return func() *tbs.JoborWorker {
		workers, err := GetWorkers(routingKey)
		if err != nil {
			logger.Jobor.Errorf("get online worker failed: %s", err)
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}

// 轮询调度
func roundRobin(routingKey string) JoborWorker {
	var i = rand.Int()
	return func() *tbs.JoborWorker {
		workers, err := GetWorkers(routingKey)
		if err != nil {
			logger.Jobor.Errorf("get online worker failed: %s", err)
			return nil
		}
		worker := workers[i%len(workers)]
		i++
		return &worker
	}
}

// 权重调度
func weight(routingKey string) JoborWorker {
	return func() *tbs.JoborWorker {
		workers, err := GetWorkers(routingKey)
		if err != nil {
			logger.Jobor.Errorf("get online worker failed: %s", err)
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
func leastTask(routingKey string) JoborWorker {
	return func() *tbs.JoborWorker {
		workers, err := GetWorkers(routingKey)
		if err != nil {
			logger.Jobor.Errorf("get online worker failed: %s", err)
			return nil
		}
		return &workers[rand.Int()%len(workers)]
	}
}


