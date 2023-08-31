package dispatcher

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"jobor/kitex_gen/pbapi"
	task2 "jobor/kitex_gen/task"
	"time"
)

func TaskWorker(ctx context.Context, data string, taskId int64, lang string, stream pbapi.TaskService_RunTaskServer) error {
	startTime := time.Now()
	defer func() {
		hlog.CtxDebugf(ctx, "Task=[%d] lang=%s worker run Task finish, cost time: %s.",
			taskId, lang, time.Since(startTime).String())
	}()
	//taskCtx, taskCancel := context.WithCancel(ctx)
	//timeoutCtx, timeoutCac := context.WithCancel(taskCtx)
	//if t.Timeout > 0 {
	//	timeoutCtx, timeoutCac = context.WithTimeout(timeoutCtx, time.Second*time.Duration(t.Timeout))
	//}
	//defer taskCancel()
	//defer timeoutCac()
	//_ = timeoutCtx
	request := task2.TaskRequest{TaskId: taskId, TaskLang: lang, TaskData: []byte(data)}
	runner, err := GetDataRun(&request)
	if err != nil {
		hlog.CtxErrorf(ctx, "dispatcher getDataRun err: %s", err.Error())
		return err
	}
	out := runner.Run(context.Background())
	hlog.CtxDebugf(ctx, "Task=[%d] lang=%s runner run start", taskId, lang)
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			hlog.CtxErrorf(ctx, "runner run close failed: %s", err)
		}
	}(out)
	var buf = make([]byte, 1024)

	resp := ""
	for {
		n, err := out.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println(resp)
				hlog.CtxDebugf(ctx, "Task=[%d] lang=%s runner run is finish", taskId, lang)
				return nil
			}
			// if read failed please send default err code -1
			hlog.CtxErrorf(ctx, "read failed from %s", err)
			err = stream.Send(&pbapi.StreamResponse{Resp: []byte(err.Error() + fmt.Sprintf("%3d", DefaultExitCode))})
			if err != nil {
				hlog.Errorf("send failed: %s", err)
			}
			return nil
		}
		if n > 0 {
			//fmt.Println("worker out:", string(buf[:n]))
			resp += string(buf[:n])
			resp := pbapi.StreamResponse{Resp: buf[:n]}
			err = stream.Send(&resp)
			if err != nil {
				hlog.Errorf("stream send failed: %s", err)
				return nil
			}
		}
	}
	//hlog.CtxInfof(taskCtx, "TaskWorker: %s", data)
	//return nil
}
