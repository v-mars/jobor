package main

import (
	pbapi "jobor/kitex_gen/pbapi"
	task "jobor/kitex_gen/task"
)

// TaskServiceImpl implements the last service interface defined in the IDL.
type TaskServiceImpl struct{}

func (s *TaskServiceImpl) RunTask(req *task.TaskRequest, stream pbapi.TaskService_RunTaskServer) (err error) {
	println("RunTask called")
	return
}
