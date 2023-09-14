package dispatcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"html/template"
	"io"
	"jobor/biz/dal/db"
	"jobor/biz/model"
	"jobor/biz/pack/task_ssh"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/pbapi"
	"jobor/kitex_gen/pbapi/taskservice"
	task2 "jobor/kitex_gen/task"
	"jobor/pkg/convert"
	"jobor/pkg/notify/dingding"
	"jobor/pkg/notify/email"
	"jobor/pkg/notify/lark"
	"jobor/pkg/notify/wechat"
	"jobor/pkg/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Fn(channel, payload string) error {
	defer func() {
		if err := recover(); err != any(nil) {
			//stack := response.Stack(3)
			hlog.Errorf("Jobor 订阅者 panic, channel=%s, payload=%s, 错误信息: %s\n", channel, payload, err)
		}
	}()
	var ed = model.Event{}
	err := json.Unmarshal([]byte(payload), &ed)
	if err != nil {
		hlog.Errorf("Jobor 订阅者 channel=%s, payload=%s, payload解析错误: %s", channel, payload, err)
	}
	t, err := model.GetTaskInfoById(ed.TaskID, false)
	if err != nil {
		hlog.Errorf("Jobor 订阅者 channel=%s, payload=%s, 获取任务错误: %s", channel, payload, err)
	}
	err = EventFunc(ed, t)
	if err != nil {
		hlog.Errorf("Jobor 订阅者 channel=%s, payload=%s, EventFunc错误: %s", channel, payload, err)
	}
	return nil
}

var c = cron.New(cron.WithSeconds())

func EventFunc(ed model.Event, t *model.JoborTask) error {
	tt := *t
	switch ed.TE {
	case model.AddEvent:
		hlog.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() { o.Unlock() }()
		if !ok && t.Status == model.TaskStatusRunning {
			hlog.Debugf("jobor cron taskId=%d add event, dispatcher cron entry is not exist", ed.TaskID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("add", model.TriggerAuto, tt)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: int(ed.TaskID), EntryID: entryId, Entry: entry}
		} else if ok && t.Status == model.TaskStatusRunning {
			c.Remove(o.Entry.ID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("add", model.TriggerAuto, tt)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		hlog.Infof("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] add Task is success", ed.TaskID, t.Name, t.Expr)
	case model.ChangeEvent:
		hlog.Debugf("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() { o.Unlock() }()
		if ok && (t.Status != model.TaskStatusRunning || t.Deleted) {
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
			hlog.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] remove Task is success", ed.TaskID, t.Name, t.Expr)
			return nil
		}
		if !ok && t.Status == model.TaskStatusRunning {
			hlog.Debugf("jobor cron taskId=%d change event, dispatcher cron entry is not exist", ed.TaskID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("change", model.TriggerAuto, tt)
			})
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		} else if ok && t.Status == model.TaskStatusRunning {
			c.Remove(o.Entry.ID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("change", model.TriggerAuto, tt)
			})
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		hlog.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] add Task is success", ed.TaskID, t.Name, t.Expr)
	case model.RunEvent:
		hlog.Debugf("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		if !ok {
			err := fmt.Sprintf("jobor cron taskId=%d run event, dispatcher cron entry is not exist", ed.TaskID)
			hlog.Debugf(err)
			return fmt.Errorf(err)
		} else {
			job := c.Entry(o.EntryID).Job
			job.Run()
			c.Remove(o.Entry.ID)
			hlog.Infof("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name, t.Expr)
		}
	case model.DeleteEvent:
		hlog.Debugf("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		if ok {
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
		} else {
			hlog.Debugf("jobor cron taskId=%d delete event, dispatcher cron entry is not exist", ed.TaskID)
		}
		o.Unlock()
		hlog.Infof("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name, t.Expr)
	case model.KillEvent:
		hlog.Debugf("jobor cron taskId=%d kill event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() { o.Unlock() }()
		if ok {
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
		} else {
			hlog.Debugf("jobor cron taskId=%d kill event, dispatcher cron entry is not exist", ed.TaskID)
		}
		hlog.Infof("jobor cron taskId=%d kill event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name, t.Expr)
	}
	return nil
}

type Task struct {
	sync.RWMutex
	Name    string
	Expr    string
	TaskId  int
	EntryID cron.EntryID
	Entry   cron.Entry
}

type Registry struct {
	sync.RWMutex
	cron map[int]Task
}

var (
	RegistryDispatcher = Registry{cron: make(map[int]Task)}
	TaskConcurrency    = 5
)

func InitCron() {
	ctx := context.TODO()
	if conf.GetConf().Server.CronType == "m" {
		c = cron.New()
		hlog.Info("jobor Task is 'min' cron mode")
	}
	taskList, err := model.GetAllRunningTask()
	if err != nil {
		hlog.CtxFatalf(ctx, err.Error())
	}
	c.Start()

	for _, t := range taskList {
		hlog.CtxDebugf(ctx, "jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", t.Id, t.Name, t.Expr)
		t := t
		entryId, err := c.AddFunc(t.Expr, func() {
			RunTasks("init", model.TriggerAuto, t)
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "jobor cron taskId=%d add func [name=%s, expr=\"%s\"] is err, %s", t.Id, t.Name, t.Expr, err)
		}
		entry := c.Entry(entryId)
		RegistryDispatcher.cron[t.Id] = Task{Name: t.Name, Expr: t.Expr, TaskId: t.Id, EntryID: entryId, Entry: entry}
		hlog.CtxInfof(ctx, "jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is success", t.Id, t.Name, t.Expr)
	}

	//defer c.Stop()
	hlog.Info("jobor Task dispatcher is start")
}

var CacheTask = cacheTask{TaskLogs: make(map[int]*taskSession)}

type cacheTask struct {
	sync.RWMutex
	TaskLogs map[int]*taskSession
}

type taskSession struct {
	sync.RWMutex
	TaskLog    *model.JoborLog
	Abort      chan bool
	Status     string
	Output     byte
	Err        error
	Conn       *grpc.ClientConn
	TaskCancel context.CancelFunc
	TaskCtx    context.Context
	Task       *model.JoborTask
	Stream     taskservice.TaskService_RunTaskClient
	RunTaskIds []int
}

// RunTasks evt 事件, add/change
func RunTasks(evt, trigger string, t model.JoborTask) {
	//if Raft.St.RaftNode.Raft.State() != raft.Leader && trigger == TriggerAuto {
	//	return
	//}
	//var tt = *t
	var ctx = context.Background()
	//if t {
	//}
	RunTasksWithRPC(ctx, evt, trigger, t, nil, "")
	//RunTasksWithBroker(evt, trigger, t)
}

func RunTasksWithRPC(ctx context.Context, evt, trigger string, t model.JoborTask, runbyid *int, parentChild string) {
	//if Raft.St.RaftNode.Raft.State() != raft.Leader && trigger == TriggerAuto {
	//	return
	//}
	var s = taskSession{TaskCtx: context.Background()}
	var tx = db.DB
	jsonTime := db.JSONTime{Time: time.Now()}
	workers := GetWorkerByRoutePolicy(t.RoutingKeys, t.RoutePolicy, t.Lang)
	var taskLog = model.JoborLog{Name: t.Name, Lang: t.Lang, TaskId: t.Id, TriggerMethod: trigger, Expr: t.Expr,
		Data: t.Data, StartTime: jsonTime, Result: model.TaskLogStatusWait, ExpectCode: t.ExpectCode,
		ExpectContent: t.ExpectContent,
		Idempotent:    fmt.Sprintf("%d", RegistryDispatcher.cron[t.Id].Entry.Prev.Unix()),
	}
	if runbyid != nil {
		taskLog.ByTaskLogId = *runbyid
	}
	var startTimeTotal = time.Now()
	defer func() {
		//if Raft.St.RaftNode.Raft.State() != raft.Leader {
		//	hlog.Infof("this dispatcher server is not Leader, Task %s[%d] lang %s is skip run,now time: %s ", t.Name, t.Id, t.Lang, time.Now())
		//	return
		//}
		defer func() {
			if errPanic := recover(); errPanic != any(nil) {
				stack := response.Stack(3)
				s.Err = fmt.Errorf("defer panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
				hlog.Errorf(s.Err.Error())
			}
		}()
		DealTaskResp(&t, &taskLog, &s)
		if errPanic := recover(); errPanic != any(nil) {
			stack := response.Stack(3)
			s.Err = fmt.Errorf("RunTasks panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
			taskLog.Result = model.TaskLogStatusUnknown
		} else if taskLog.Result == model.TaskLogStatusTimeout {
			s.Err = fmt.Errorf("任务 %s[%d] 类型=%s %s", t.Name, t.Id, t.Lang, "已经超时")
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
		} else if taskLog.Result == model.TaskLogStatusAbort {
			s.Err = fmt.Errorf("任务 %s[%d] 类型=%s %s", t.Name, t.Id, t.Lang, "已经终止")
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
		} else if s.Err != nil {
			s.Err = fmt.Errorf("任务 %s[%d] 类型=%s %s", t.Name, t.Id, t.Lang, s.Err)
			hlog.Errorf("%s", s.Err)
			taskLog.ErrMsg = s.Err.Error()
			taskLog.Result = model.TaskLogStatusFailed
		} else {
			taskLog.Result = model.TaskLogStatusSuccess
			hlog.Infof("任务 %s[%d] 类型=%s 执行成功", t.Name, t.Id, t.Lang)
		}
		var totalTime = time.Now().UnixMilli() - startTimeTotal.UnixMilli()
		taskLog.CostTime = db.MillisTime((time.Millisecond * time.Duration(totalTime)).String())
		taskLog.EndTime = db.JSONTime{Time: time.Now()}
		if s.Err = tx.Model(&taskLog).Omit([]string{"CreatedAt", "UpdatedAt"}...).Save(&taskLog).Error; s.Err != nil {
			s.Err = fmt.Errorf("保存tasklog错误: %s", s.Err)
			hlog.Error(s.Err)
			return
		}
		if s.Err = tx.Model(&t).Updates(model.JoborTask{
			Prev: jsonTime, Next: db.JSONTime{Time: RegistryDispatcher.cron[t.Id].Entry.Next}}).Error; s.Err != nil {
			s.Err = fmt.Errorf("update Task pre/next time err: %s", s.Err)
			hlog.Error(s.Err)
		}
		s.Err = s.Alarm()
		s.Delete()
	}()

	if s.Err = tx.Create(&taskLog).Error; s.Err != nil {
		s.Err = fmt.Errorf("create tasklog err: %s", s.Err)
		return
	}
	//s.RunTaskIds = append(s.RunTaskIds, taskLog.TaskId)
	//if runbyid != nil {
	//	CacheTask.TaskLogs[*runbyid].RunTaskIds = append(CacheTask.TaskLogs[*runbyid].RunTaskIds, taskLog.TaskId)
	//}
	if runbyid != nil && len(parentChild) > 0 {
		byTask := model.JoborLog{Model: db.Model{Id: taskLog.Id}}
		if s.Err = db.DB.Model(&model.JoborLog{Model: db.Model{Id: *runbyid}}).Association(parentChild).Append(&byTask); s.Err != nil {
			return
		}
	}

	s.TaskLog = &taskLog
	s.Task = &t
	s.Add()
	hlog.Infof("Task %s[%d] lang %s success start,now time: %s ", t.Name, t.Id, t.Lang, time.Now())
	s.TaskCtx, s.TaskCancel = context.WithCancel(ctx)
	timeoutCtx, timeoutCac := context.WithCancel(ctx)
	if t.Timeout > 0 {
		timeoutCtx, timeoutCac = context.WithTimeout(timeoutCtx, time.Second*time.Duration(t.Timeout))
	}
	defer s.TaskCancel()
	defer timeoutCac()

	// 执行父任务
	s.Err = runMultiTasks(ctx, trigger, t.ParentRunParallel, &taskLog.Id, "Parents", t.ParentTaskIds...)
	if s.Err != nil {
		s.Err = fmt.Errorf("执行父任务[%v]失败，%s", t.ParentTaskIds, s.Err)
		return
	}
	defer func() {
		// 执行子任务
		e := runMultiTasks(ctx, trigger, t.ChildRunParallel, &taskLog.Id, "Childs", t.ChildTaskIds...)
		if e != nil {
			s.Err = fmt.Errorf("执行子任务[%v]失败，%s", t.ParentTaskIds, e)
			return
		}
	}()

	w := workers()
	if w == nil {
		s.Err = fmt.Errorf("can't get valid worker")
		return
	}

	//var executor = DataCode{Lang: Lang(t.Lang),ScriptCode: t.Data.Code}
	var marshal []byte
	marshal, s.Err = json.Marshal(t.Data)
	if s.Err != nil {
		hlog.Errorf("DataCode Marshal  err: %s", s.Err)
		return
	}

	//time.Sleep(1 * time.Second)
	taskLog.Addr = w.Addr
	taskLog.Result = model.TaskLogStatusRunning
	if s.Err = tx.Model(&taskLog).Omit([]string{"CreatedAt", "UpdatedAt"}...).Updates(taskLog).Error; s.Err != nil {
		s.Err = fmt.Errorf("update tasklog addr/status err: %s", s.Err)
		hlog.Error(s.Err)
		return
	}

	if w.Mode == model.WorkerModeSsh {
		ts := task_ssh.SshServer{Cmd: t.Data.Content, Username: w.Username, Password: string(w.Password),
			Host: w.Ip, Port: w.Port,
			AuthMode: w.AuthMode, PrivateKey: string(w.Rsa), PrivateSecret: string(w.Private)}
		s.Err = ts.ExecRemoteSshInit()
		if s.Err != nil {
			s.Err = fmt.Errorf("task worker mode remote ssh init err: %s", s.Err)
			hlog.Errorf(s.Err.Error())
			taskLog.Result = model.TaskLogStatusFailed
			return
		}
		if t.Data.PreCmd != "" {
			_, e := ts.ExecRemoteSshCmd(t.Data.PreCmd)
			if e != nil {
				s.Err = fmt.Errorf("task worker mode ssh run pre cmd err: %s", e)
				hlog.Errorf(s.Err.Error())
				taskLog.Result = model.TaskLogStatusFailed
				return
			}
			hlog.CtxInfof(ctx, "task=[%d] lang=%s worker mode ssh run pre cmd is success", t.Id, t.Lang)
		}

		sshOut, e := ts.ExecRemoteSshCmd(t.Data.Content)
		if e != nil {
			s.Err = fmt.Errorf("task worker mode ssh run task err: %s", e)
			hlog.Errorf(s.Err.Error())
			taskLog.Result = model.TaskLogStatusFailed
			return
		}
		taskLog.Resp = sshOut
		ts.Close()
		hlog.Infof("exec WorkerModeSsh finish")
		return
	}
	tc, err := taskservice.NewClient(conf.AppWorkerName, client.WithHostPorts(w.Addr))
	if err != nil {
		hlog.Errorf("tet task worker rpc client err: %s", s.Err)
		return
	}
	s.Stream, s.Err = tc.RunTask(s.TaskCtx, &task2.TaskRequest{TaskId: int64(t.Id), TaskLang: t.Lang, TaskData: marshal})
	if s.Err != nil {
		hlog.Errorf("run Task err: %s", s.Err)
		return
	}
	defer func(Stream taskservice.TaskService_RunTaskClient) { _ = Stream.Close() }(s.Stream)
	defer func(resStream taskservice.TaskService_RunTaskClient) { _ = resStream.Close() }(s.Stream)
	hlog.Debugf("Task %s[%d] RunTasks success", t.Name, t.Id)
	taskLog.Resp = ""
	streamChan := func() (chan *pbapi.StreamResponse, chan error) {
		errChan := make(chan error, 1)
		Message := make(chan *pbapi.StreamResponse, 1)
		go func() {
			//defer func() { hlog.Debugf("Task %s[%d] res stream.recv finish", t.Name, t.Id) }()
			rec, errRecv := s.Stream.Recv()
			Message <- rec
			errChan <- errRecv
		}()
		return Message, errChan
	}
	for {
		msg, errChan := streamChan()
		//hlog.Debugf("Task %s[%d] res stream.recv start", t.Name, t.Id)
		select {
		//case <-s.Abort:
		case <-s.TaskCtx.Done():
			hlog.Debugf("Task %s[%d] is abort", t.Name, t.Id)
			taskLog.Result = model.TaskLogStatusAbort
			return
		case <-timeoutCtx.Done():
			hlog.Debugf("Task %s[%d] is timeout", t.Name, t.Id)
			taskLog.Result = model.TaskLogStatusTimeout
			return
		case d := <-msg:
			//hlog.Debugf("Task %s[%d] stream recv data: %s", t.Name, t.Id, d.GetResp())
			taskLog.Resp += string(d.GetResp())
		case e := <-errChan:
			if e == io.EOF {
				hlog.Infof("Task %s[%d] stream recv finish", t.Name, t.Id)
				//goto Next
				//DealTaskResp(&t, &taskLog, &s)
				return
			}
			if e != nil {
				hlog.Errorf("Task %s[%d] resStream.Recv err: %s", t.Name, t.Id, e.Error())
				taskLog.Result = "failed"
				s.Err = e
				return
			}
		}
	}
}
func DealTaskResp(t *model.JoborTask, taskLog *model.JoborLog, s *taskSession) {
	if len(taskLog.Resp) > 3000 {
		taskLog.Resp = fmt.Sprintf("%s\n……省略过多内容……\n%s", taskLog.Resp[0:1499], taskLog.Resp[len(taskLog.Resp)-1499:])
	}
	var taskRespCode int
	judges := func() error {
		if s.Err != nil {
			taskLog.Result = model.TaskLogStatusFailed
			taskLog.ErrMsg = s.Err.Error()
			return s.Err
		}
		taskRespCode, s.Err = s.GetRespCode()
		if s.Err != nil {
			s.Err = fmt.Errorf("get response code failed: %s", s.Err)
			hlog.Error(s.Err)
			return s.Err
		}
		fmt.Println("GetRespCode:", t.ExpectCode, taskRespCode)
		s.TaskLog.ErrCode = taskRespCode
		if t.ExpectCode != taskRespCode {
			s.Err = fmt.Errorf("任务 %s[%d] 返回码： %d, 期望返回码：%d", t.Name, taskLog.Id, taskRespCode, t.ExpectCode)
			return s.Err
		}
		if t.ExpectContent != "" {
			if !strings.Contains(taskLog.Resp, t.ExpectContent) {
				s.Err = fmt.Errorf("%s Task %s[%d] resp context not contains expect content: %s", "server", t.Name, taskLog.Id, t.ExpectContent)
				return s.Err
			}
		}
		return nil
	}
	s.Err = judges()
	if s.Err != nil {
		return
	}
}

func RunTasksWithBroker(evt, trigger string, t model.JoborTask) {
	var s = taskSession{TaskCtx: context.Background()}
	defer func() {
		defer func() {
			if errPanic := recover(); errPanic != any(nil) {
				stack := response.Stack(3)
				s.Err = fmt.Errorf("defer panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
				hlog.Error(s.Err)
			}
		}()
	}()
	hlog.Infof("broker Task %s[%d] lang %s routingKey %s success start,now time: %s ",
		t.Name, t.Id, t.Lang, t.RoutingKey, time.Now())
	var marshal []byte
	marshal, s.Err = json.Marshal(t.Data)
	if s.Err != nil {
		hlog.Errorf("DataCode Marshal  err: %s", s.Err)
		return
	}
	//if _, err := q.InitQSrv(&conf.GetConf().Redis, q.Queue); err != nil {
	//	hlog.Fatal(err)
	//	return
	//}
	//err := q.QSrv.RegisterTask("TaskWorker", q.TaskWorker)
	//if err != nil {
	//	hlog.Errorf("register Task  err: %s", err)
	//	return
	//}

	signature := &tasks.Signature{
		Name: RegisterTaskName,
		//RoutingKey:   t.RoutingKey,
		RetryCount:   t.Retry,
		RetryTimeout: t.Timeout,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: string(marshal),
			},
			{
				Type:  "int64",
				Value: t.Id,
			},
			{
				Type:  "string",
				Value: t.Lang,
			},
		},
	}
	asyncResult, err := QSrv.SendTask(signature)
	if err != nil {
		s.Err = err
		return
	}
	//asyncResult.GetState()
	_, err = asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		//return fmt.Errorf("Getting Task result failed with error: %s", err.Error())
	}
	hlog.Infof("broker Task %s[%d] lang %s routingKey %s success start,asyncResult: %s ",
		t.Name, t.Id, t.Lang, t.RoutingKey, asyncResult.GetState())

}

func runMultiTasks(ctx context.Context, trigger string, RunParallel bool, runbyid *int, parentChild string, taskids ...int) error {
	con := make(chan struct{}, TaskConcurrency)
	for _, v := range taskids {
		v := v
		t, err := model.GetTaskInfoById(v, false)
		if err != nil {
			err = fmt.Errorf("jobor runMultiTasks task id %v, 获取任务错误: %s", v, err)
			hlog.Errorf(err.Error())
			return err
		}
		if RunParallel {
			con <- struct{}{}
			go doConnMultiTask(ctx, con, trigger, *t, runbyid, parentChild)
		} else {
			RunTasksWithRPC(ctx, "", trigger, *t, runbyid, parentChild)
		}
	}
	return nil
}

func doConnMultiTask(ctx context.Context, d chan struct{}, trigger string, t model.JoborTask, runbyid *int, parentChild string) {
	RunTasksWithRPC(ctx, "", trigger, t, runbyid, parentChild)
	<-d
}

func (s *taskSession) Alarm() error {
	isFailed, err := utils.In([]string{model.TaskLogStatusFailed, model.TaskLogStatusUnknown, model.TaskLogStatusAbort,
		model.TaskLogStatusCancel}, s.TaskLog.Result)
	if err != nil {
		return err
	}
	switch model.AlarmPolicy(s.Task.AlarmPolicy) {
	case model.Never:
		return nil
	case model.Always:
		err = s.Notify()
		if err != nil {
			return err
		}
	case model.Failed:
		if isFailed {
			err = s.Notify()
			if err != nil {
				return err
			}
		}
	case model.Success:
		if s.TaskLog.Result == model.TaskLogStatusSuccess {
			err = s.Notify()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateRow(row []string, odd int) string {
	var arr []string
	width := ""
	if len(arr) > 1 {
		width = "width:70px"
	}
	for i, x := range row {
		s := "<td style=\"padding:12px 8px\">%s</td>"
		if i == 0 && len(width) > 0 {
			s = "<td style=\"padding:12px 8px;" + width + "\">%s</td>"
		}
		arr = append(arr, fmt.Sprintf(s, strings.ReplaceAll(template.HTMLEscaper(x), "\n", "<br/>")))
	}
	bgColor := "f8f8f9"
	if odd > 0 {
		bgColor = "fff"
	}
	tds := strings.Join(arr, "\n")
	res := fmt.Sprintf(`<tr style="height:25px; background-color:#%s; border:1px solid #e8e8e8">%s
	</tr>`, bgColor, tds)
	return res
}

func generateForm(title string, rows [][]string) string {
	var trs []string
	for i, o := range rows {
		trs = append(trs, generateRow(o, 1-i%2))
	}
	tr := strings.Join(trs, "\n")
	tableTemplate := fmt.Sprintf(`
	<div style="overflow: auto; margin-bottom: 20px;">
	<table border="0" style="width: %s; text-align:left; border-collapse:collapse; font-size:14px; color:rgba(0,0,0,0.75); line-height:1.1; word-break:break-all; white-space:nowrap; border-left:1px solid #e8e8e8; border-right:1px solid #e8e8e8; font-family:'Microsoft YaHei UI','Microsoft YaHei','微软雅黑',SimSun,'宋体',sans-serif,serif; margin:5px;border-top:1px solid #e8e8e8;">
	<caption style="text-align: left; font-size: 16px; font-weight: bold; margin-bottom: 10px">
	%s
	</caption>
	<tbody>
	%s
	</tbody>
	</table>
	</div>`, "80%", title, tr)
	return tableTemplate
}

func (s *taskSession) GenerateHtml() string {
	before := `<div style=" text-align:left; border-collapse:co
llapse; font-size:18px; line-height:1.1;  font-family:'Microsoft YaHei UI','Microsoft YaHei','微软雅黑',SimSun,'
宋体',sans-serif,serif; margin:30px">
	Hello, 定时任务出现问题了:
    </div>`
	taskDetail := [][]string{
		{"任务 Id：", fmt.Sprintf("%d", s.Task.Id)},
		{"任务名称：", fmt.Sprintf("%s", s.Task.Name)},
		{"执行 Id：", fmt.Sprintf("%d", s.TaskLog.Id)},
		{"类   型：", fmt.Sprintf("%s", s.TaskLog.Lang)},
		{"触发方式：", fmt.Sprintf("%s", s.TaskLog.TriggerMethod)},
		{"表达式 ：", fmt.Sprintf("%s", s.TaskLog.Expr)},
		{"Worker：", fmt.Sprintf("%s", s.TaskLog.Addr)},
		{"执行时间：", fmt.Sprintf("[ %s - %s ]", s.TaskLog.StartTime.Format("2006-01-02 15:04:05"), s.TaskLog.EndTime.Format("2006-01-02 15:04:05"))},
		{"耗   时：", fmt.Sprintf("%s", fmt.Sprintf("%s", s.TaskLog.CostTime))},
		{"状   态：", fmt.Sprintf("%s", s.TaskLog.Result)},
	}
	stdOut := [][]string{
		{"返回内容：", fmt.Sprintf("%s", s.TaskLog.Resp)}}
	errMsg := [][]string{
		{"错误状态码：", fmt.Sprintf("%d", s.TaskLog.ErrCode)},
		{"错误信息：", fmt.Sprintf("%s", s.TaskLog.ErrMsg)},
	}
	detail := generateForm("任务执行详情:", taskDetail)
	std := generateForm("任务执行输出:", stdOut)
	err := generateForm("输出错误:", errMsg)
	res := before + detail + std + err
	return res
}

func (s *taskSession) Notify() error {
	var title = fmt.Sprintf("定时任务[%s]记录ID[%d]执行结果", s.Task.Name, s.TaskLog.Id)
	var msg = fmt.Sprintf(
		`
任务 Id：%d
任务名称：%s
执行 Id：%d
类   型：%s
触发方式：%s
表达式 ：%s
Worker：%s
执行时间：[ %s - %s ]
耗   时：%s秒
状   态：%s
返回内容：%s
错误状态码：%d
错误信息：%s`,
		s.Task.Id, s.Task.Name, s.TaskLog.Id, s.TaskLog.Lang, s.TaskLog.TriggerMethod, s.TaskLog.Expr,
		s.TaskLog.Addr, s.TaskLog.StartTime.Format("2006-01-02 15:04:05"),
		s.TaskLog.EndTime.Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%s", s.TaskLog.CostTime), s.TaskLog.Result,
		s.TaskLog.Resp, s.TaskLog.ErrCode, s.TaskLog.ErrMsg)
	var apiData = map[string]interface{}{
		"task_log_id":    s.TaskLog.Id,
		"task_id":        s.Task.Id,
		"task_name":      s.Task.Name,
		"lang":           s.TaskLog.Lang,
		"worker_addr":    s.TaskLog.Addr,
		"trigger_method": s.TaskLog.TriggerMethod,
		"cron_expr":      s.TaskLog.Expr,
		"resp":           s.TaskLog.Resp,
		"start_time":     s.TaskLog.StartTime,
		"end_time":       s.TaskLog.EndTime,
		"err_code":       s.TaskLog.ErrCode,
		"err_msg":        s.TaskLog.ErrMsg,
		"cost_time":      s.TaskLog.CostTime,
		"status":         s.TaskLog.Result,
	}
	if s.Task.Notify.Webhook.Enable {
		hlog.Debugf("Task notify webhook send is success, %s", apiData)
	}
	if s.Task.Notify.Email.Enable {
		mailMsg := s.GenerateHtml()
		notify := email.NewMail(conf.GetConf().Email.Username, conf.GetConf().Email.Password,
			conf.GetConf().Email.SmtpServer, conf.GetConf().Email.From, conf.GetConf().Email.SmtpPort,
			conf.GetConf().Email.Tls)
		err := notify.Send(title, mailMsg, s.Task.Notify.Email.Receivers, []string{})
		if err != nil {
			//return err
		}
		hlog.Debugf("Task notify email send is success")
	}
	if s.Task.Notify.Wechat.Enable {
		notify := wechat.NewWeChat(conf.GetConf().EntWeChat.CorpId,
			convert.ToInt(conf.GetConf().EntWeChat.NotifyAgentId), conf.GetConf().EntWeChat.NotifySecret)
		err := notify.Send(s.Task.Notify.Wechat.Receivers, s.Task.Notify.Wechat.Groups, []string{},
			title, msg, nil)
		if err != nil {
			//return err
		}
		hlog.Debugf("Task notify wechat send is success")
	}
	if s.Task.Notify.Lark.Enable && s.Task.Notify.Lark != nil && s.Task.Notify.Lark.Webhooks != nil {
		whs := (*(s.Task.Notify).Lark).Webhooks
		for _, v := range whs {
			notify := lark.NewLark(v.WebhookUrl, 1, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				//return err
			}
		}
		hlog.Debugf("Task notify lark send is success")
	}
	if s.Task.Notify.Dingding.Enable {
		for _, v := range s.Task.Notify.Dingding.Webhooks {
			notify := dingding.NewDing(v.WebhookUrl, 2, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				//return err
			}
		}
		hlog.Debugf("Task notify dinding send is success")
	}
	hlog.Infof("Task notify send is success")
	return nil
}

func (s *taskSession) GetRespCode() (int, error) {
	if len(s.TaskLog.Resp) > 5 {
		taskRespCode, err := strconv.Atoi(strings.TrimLeft(s.TaskLog.Resp[len(s.TaskLog.Resp)-5:], " "))
		if err != nil {
			hlog.Errorf("change str to int failed: %s", err)
			return DefaultExitCode, err
		}
		return taskRespCode, err
	}

	return DefaultExitCode, fmt.Errorf("response length less than 5")
}

func (s *taskSession) Add() {
	CacheTask.Lock()
	CacheTask.TaskLogs[s.TaskLog.Id] = s
	CacheTask.Unlock()
}

func (s *taskSession) Delete() {
	//s.Close()
	CacheTask.Lock()
	delete(CacheTask.TaskLogs, s.TaskLog.Id)
	CacheTask.Unlock()
}

func (s *taskSession) Close() {
	//if s.Stream.Context()
	//_ = s.Stream.CloseSend()
	_ = s.Conn.Close()
	s.TaskCancel()
}

// 并发控制
/*func con() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// if exist a err Task,will stop all Task
	g := errgroup.WithCancel(ctx)
	g.GOMAXPROCS(1)
	// parent tasks
	g.Go(func(ctx context.Context) error {
		return nil //s.runMultiTasks(ctx, Task.ParentRunParallel, define.ParentTask, Task.Id, Task.ParentTaskIds...)
	})
	// server Task
	g.Go(func(ctx context.Context) error {
		return nil // s.runTask(ctx, Task.Id, Task.Id, define.MasterTask)
	})
	// childs Task
	g.Go(func(ctx context.Context) error {
		return nil // s.runMultiTasks(ctx, Task.ChildRunParallel, define.ChildTask, Task.Id, Task.ChildTaskIds...)
	})
	err := g.Wait()
	if err != nil {
		fmt.Errorf("Task run failed taskId[%d] err: %s", 1, err)
	}
}*/

func ConvertSecond(spec string) string {
	if spec != "" {
		_spec := strings.Split(spec, " ")
		hour, minute, second := time.Now().Clock()

		// second处理
		_second := strconv.FormatInt(int64(second), 10) + "/60"
		spec = strings.Replace(spec, _spec[0], _second, 1)

		// minute处理(类似: */5)
		if strings.HasPrefix(_spec[1], "*/") {
			_minute := strconv.FormatInt(int64(minute), 10) + "/" + strings.Trim(_spec[1], "*/")
			spec = strings.Replace(spec, _spec[1], _minute, 1)
		}

		// hour处理(类似: */5)
		if strings.HasPrefix(_spec[2], "*/") {
			_hour := strconv.FormatInt(int64(hour), 10) + "/" + strings.Trim(_spec[2], "*/")
			spec = strings.Replace(spec, _spec[2], _hour, 1)
		}

		return spec
	}

	return ""
}
