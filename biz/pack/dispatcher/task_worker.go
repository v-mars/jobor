package dispatcher

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tidwall/gjson"
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
	request := task2.TaskRequest{TaskId: taskId, TaskLang: lang, TaskData: []byte(data)}
	runner, err := GetDataRun(&request)
	if err != nil {
		hlog.CtxErrorf(ctx, "dispatcher getDataRun err: %s", err.Error())
		return err
	}
	gd := gjson.Parse(data)
	if gd.Get("pre_cmd").String() != "" {
		preRes, exitCode, err := runner.RunPreCmd(ctx)
		if err != nil {
			err = fmt.Errorf("\n%s, Return exitCode:%5d", err, exitCode) // write exitCode,total 5 byte
			resp := pbapi.StreamResponse{Resp: []byte(fmt.Sprintf("预执行结果：\n%s\n预执行错误：%s\n", preRes,
				err.Error()))}
			err = stream.Send(&resp)
			return err
		}
		preResp := pbapi.StreamResponse{Resp: []byte(fmt.Sprintf("预执行结果：\n%s\n预执行成功\n\n", preRes))}
		err = stream.Send(&preResp)
		if err != nil {
			return err
		}
	}
	_ = stream.Send(&pbapi.StreamResponse{Resp: []byte(fmt.Sprintf("任务执行返回：\n"))})
	hlog.CtxDebugf(ctx, "Task=[%d] lang=%s runner run pre cmd is success", taskId, lang)
	out := runner.Run(ctx)
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
