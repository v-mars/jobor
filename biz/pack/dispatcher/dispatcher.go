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
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/pbapi"
	"jobor/kitex_gen/pbapi/taskservice"
	task2 "jobor/kitex_gen/task"
	"jobor/pkg/notify/dingding"
	"jobor/pkg/notify/email"
	"jobor/pkg/notify/lark"
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
	switch ed.TE {
	case model.AddEvent:
		hlog.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", ed.TaskID, t.Name, t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() { o.Unlock() }()
		if !ok && t.Status == model.TaskStatusRunning {
			hlog.Debugf("jobor cron taskId=%d add event, dispatcher cron entry is not exist", ed.TaskID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("add", model.TriggerAuto, t)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: int(ed.TaskID), EntryID: entryId, Entry: entry}
		} else if ok && t.Status == model.TaskStatusRunning {
			c.Remove(o.Entry.ID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("add", model.TriggerAuto, t)
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
				RunTasks("change", model.TriggerAuto, t)
			})
			if err != nil {
				return err
			}
			entry := c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = Task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		} else if ok && t.Status == model.TaskStatusRunning {
			c.Remove(o.Entry.ID)
			entryId, err := c.AddFunc(t.Expr, func() {
				RunTasks("change", model.TriggerAuto, t)
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
		hlog.CtxDebugf(ctx, "jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", t.ID, t.Name, t.Expr)
		entryId, err := c.AddFunc(t.Expr, func() {
			RunTasks("init", model.TriggerAuto, &t)
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "jobor cron taskId=%d add func [name=%s, expr=\"%s\"] is err, %s", t.ID, t.Name, t.Expr, err)
		}
		entry := c.Entry(entryId)
		RegistryDispatcher.cron[t.ID] = Task{Name: t.Name, Expr: t.Expr, TaskId: t.ID, EntryID: entryId, Entry: entry}
		hlog.CtxInfof(ctx, "jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is success", t.ID, t.Name, t.Expr)
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
}

// RunTasks evt 事件, add/change
func RunTasks(evt, trigger string, t *model.JoborTask) {
	//if Raft.St.RaftNode.Raft.State() != raft.Leader && trigger == TriggerAuto {
	//	return
	//}

	RunTasksWithRPC(evt, trigger, t)
	//RunTasksWithBroker(evt, trigger, t)
}

func RunTasksWithRPC(evt, trigger string, t *model.JoborTask) {
	//if Raft.St.RaftNode.Raft.State() != raft.Leader && trigger == TriggerAuto {
	//	return
	//}
	var s = taskSession{TaskCtx: context.Background()}
	var tx = db.DB
	jsonTime := db.JSONTime{Time: time.Now()}
	workers := GetWorkerByRoutePolicy(t.RoutingKey, t.RoutePolicy)
	var taskLog = model.JoborLog{Name: t.Name, Lang: t.Lang, TaskId: t.ID, TriggerMethod: trigger, Expr: t.Expr,
		Data: t.Data, StartTime: jsonTime, Result: model.TaskLogStatusWait,
		Idempotent: fmt.Sprintf("%d", RegistryDispatcher.cron[t.ID].Entry.Prev.Unix()),
	}
	var startTimeTotal = time.Now()
	defer func() {
		//if Raft.St.RaftNode.Raft.State() != raft.Leader {
		//	hlog.Infof("this dispatcher server is not Leader, Task %s[%d] lang %s is skip run,now time: %s ", t.Name, t.ID, t.Lang, time.Now())
		//	return
		//}
		defer func() {
			if errPanic := recover(); errPanic != any(nil) {
				stack := response.Stack(3)
				s.Err = fmt.Errorf("defer panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
				hlog.Error(s.Err)
			}
		}()
		if errPanic := recover(); errPanic != any(nil) {
			stack := response.Stack(3)
			s.Err = fmt.Errorf("RunTasks panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
			taskLog.Result = model.TaskLogStatusUnknown
		} else if taskLog.Result == model.TaskLogStatusTimeout {
			s.Err = fmt.Errorf("task %s[%d] lang=%s %s", t.Name, t.ID, t.Lang, "is timeout")
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
		} else if taskLog.Result == model.TaskLogStatusAbort {
			s.Err = fmt.Errorf("task %s[%d] lang=%s %s", t.Name, t.ID, t.Lang, "is aborted")
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
		} else if s.Err != nil {
			s.Err = fmt.Errorf("task %s[%d] lang=%s %s", t.Name, t.ID, t.Lang, s.Err)
			hlog.Error(s.Err)
			taskLog.ErrMsg = s.Err.Error()
			taskLog.Result = model.TaskLogStatusFailed
		} else {
			taskLog.Result = model.TaskLogStatusSuccess
			hlog.Infof("task %s[%d] lang=%s 执行成功", t.Name, t.ID, t.Lang)
		}
		//totalTime := time.Since(startTimeTotal)
		//value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", totalTime.Seconds()), 64)
		var totalTime = time.Now().Unix() - startTimeTotal.Unix()
		taskLog.CostTime = db.SecTime((time.Second * time.Duration(totalTime)).String())
		taskLog.EndTime = db.JSONTime{Time: time.Now()}
		if s.Err = tx.Model(&taskLog).Omit([]string{"CreatedAt", "UpdatedAt"}...).Save(&taskLog).Error; s.Err != nil {
			s.Err = fmt.Errorf("保存tasklog错误: %s", s.Err)
			hlog.Error(s.Err)
			return
		}
		if s.Err = tx.Model(&t).Updates(model.JoborTask{
			Prev: jsonTime, Next: db.JSONTime{Time: RegistryDispatcher.cron[t.ID].Entry.Next}}).Error; s.Err != nil {
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

	s.TaskLog = &taskLog
	s.Task = t
	s.Add()
	hlog.Infof("Task %s[%d] lang %s success start,now time: %s ", t.Name, t.ID, t.Lang, time.Now())

	//var executor = DataCode{Lang: Lang(t.Lang),ScriptCode: t.Data.Code}
	var marshal []byte
	marshal, s.Err = json.Marshal(t.Data)
	if s.Err != nil {
		hlog.Errorf("DataCode Marshal  err: %s", s.Err)
		return
	}

	s.TaskCtx, s.TaskCancel = context.WithCancel(context.Background())
	timeoutCtx, timeoutCac := context.WithCancel(context.TODO())
	if t.Timeout > 0 {
		timeoutCtx, timeoutCac = context.WithTimeout(timeoutCtx, time.Second*time.Duration(t.Timeout))
	}
	defer s.TaskCancel()
	defer timeoutCac()
	time.Sleep(1 * time.Second)

	//conn, w, errConn := TryGetGrpcClientConn(s.TaskCtx, workers)
	//if errConn != nil {
	//	s.Err = errConn
	//	hlog.Errorf("TryGetGrpcClientConn err: %s", errConn)
	//	return
	//}
	//defer func(conn *grpc.ClientConn) { _ = conn.Close() }(conn)
	//s.Conn = conn
	w := workers()
	fmt.Println("w:", w)
	if w == nil {
		s.Err = fmt.Errorf("can't get valid worker")
		return
	}
	taskLog.Addr = w.Addr
	taskLog.Result = model.TaskLogStatusRunning
	if s.Err = tx.Model(&taskLog).Omit([]string{"CreatedAt", "UpdatedAt"}...).Updates(taskLog).Error; s.Err != nil {
		s.Err = fmt.Errorf("update tasklog addr/status err: %s", s.Err)
		hlog.Error(s.Err)
		return
	}
	tc, err := taskservice.NewClient(conf.AppWorkerName, client.WithHostPorts(w.Addr))
	if err != nil {
		hlog.Errorf("tet task worker rpc client err: %s", s.Err)
		return
	}
	s.Stream, s.Err = tc.RunTask(s.TaskCtx, &task2.TaskRequest{TaskId: int64(t.ID), TaskLang: t.Lang, TaskData: marshal})
	if s.Err != nil {
		hlog.Errorf("run Task err: %s", s.Err)
		return
	}
	defer func(Stream taskservice.TaskService_RunTaskClient) { _ = Stream.Close() }(s.Stream)
	defer func(resStream taskservice.TaskService_RunTaskClient) { _ = resStream.Close() }(s.Stream)
	hlog.Debugf("Task %s[%d] RunTasks success", t.Name, t.ID)
	taskLog.Resp = ""
	streamChan := func() (chan *pbapi.StreamResponse, chan error) {
		errChan := make(chan error, 1)
		Message := make(chan *pbapi.StreamResponse, 1)
		go func() {
			defer func() { hlog.Debugf("Task %s[%d] res stream.recv finish", t.Name, t.ID) }()
			rec, errRecv := s.Stream.Recv()
			Message <- rec
			errChan <- errRecv
		}()
		return Message, errChan
	}
	for {
		msg, errChan := streamChan()
		hlog.Debugf("Task %s[%d] res stream.recv start", t.Name, t.ID)
		select {
		//case <-s.Abort:
		case <-s.TaskCtx.Done():
			hlog.Debugf("Task %s[%d] is abort", t.Name, t.ID)
			taskLog.Result = model.TaskLogStatusAbort
			return
		case <-timeoutCtx.Done():
			hlog.Debugf("Task %s[%d] is timeout", t.Name, t.ID)
			taskLog.Result = model.TaskLogStatusTimeout
			return
		case d := <-msg:
			hlog.Debugf("Task %s[%d] stream recv data: %s", t.Name, t.ID, d.GetResp())
			taskLog.Resp += string(d.GetResp())
		case e := <-errChan:
			if e == io.EOF {
				hlog.Infof("Task %s[%d] stream recv finish", t.Name, t.ID)
				goto Next
			}
			if e != nil {
				hlog.Errorf("Task %s[%d] resStream.Recv err: %s", t.Name, t.ID, e.Error())
				taskLog.Result = "failed"
				s.Err = e
				return
			}
		}
	}
Next:
	if len(taskLog.Resp) > 3000 {
		taskLog.Resp = fmt.Sprintf("%s\n……省略过多内容……\n%s", taskLog.Resp[0:1499], taskLog.Resp[len(taskLog.Resp)-1499:])
	}
	var taskRespCode int
	judges := func() error {
		if s.Err != nil {
			return s.Err
		}
		taskRespCode, s.Err = s.GetRespCode()
		if s.Err != nil {
			s.Err = fmt.Errorf("get response code failed: %s", s.Err)
			hlog.Error(s.Err)
			return s.Err
		}
		s.TaskLog.ErrCode = taskRespCode
		if t.ExpectCode != taskRespCode {
			return fmt.Errorf("%s Task %s[%d] resp code is %d, expect code %d", "server", t.Name, taskLog.ID, taskRespCode, t.ExpectCode)
		}
		if t.ExpectContent != "" {
			if !strings.Contains(taskLog.Resp, t.ExpectContent) {
				return fmt.Errorf("%s Task %s[%d] resp context not contains expect content: %s", "server", t.Name, taskLog.ID, t.ExpectContent)
			}
		}
		return nil
	}
	s.Err = judges()
	if s.Err != nil {
		return
	}

}

func RunTasksWithBroker(evt, trigger string, t *model.JoborTask) {
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
		t.Name, t.ID, t.Lang, t.RoutingKey, time.Now())
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
				Value: t.ID,
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
		t.Name, t.ID, t.Lang, t.RoutingKey, asyncResult.GetState())

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
	res := fmt.Sprintf(`<tr style="height:25px; background-color:#%s; border-bottom:1px solid #e8e8e8">%s
	</tr>`, bgColor, tds)
	return res
}

func generateForm(title string, rows [][]string) string {
	var trs []string
	for i, o := range rows {
		trs = append(trs, generateRow(o, 1-i%2))
	}
	tr := strings.Join(trs, "\n")
	tableTemplate := `
	<div style="overflow: auto; margin-bottom: 20px;">
	<table border="0" style="width: 80%%; text-align:left; border-collapse:collapse; font-size:14px; color:rgba(0,0,0,0.75); line-height:1.1; word-break:break-all; white-space:nowrap; border-left:1px solid #e8e8e8; border-right:1px solid #e8e8e8; font-family:'Microsoft YaHei UI','Microsoft YaHei','微软雅黑',SimSun,'宋体',sans-serif,serif; margin:5px;border-top:1px solid #e8e8e8;">
	<caption style="text-align: left; font-size: 16px; font-weight: bold; margin-bottom: 10px">
	%s
	</caption>
	<tbody>
	%s
	</tbody>
	</table>
	</div>`
	res := fmt.Sprintf(tableTemplate, title, tr)
	return res
}

func (s *taskSession) GenerateHtml() string {
	before := `<div style=" text-align:left; border-collapse:co
llapse; font-size:18px; line-height:1.1;  font-family:'Microsoft YaHei UI','Microsoft YaHei','微软雅黑',SimSun,'
宋体',sans-serif,serif; margin:30px">
	Hello, 定时任务出现问题了:
    </div>`
	taskDetail := [][]string{
		{"任务 ID：", fmt.Sprintf("%d", s.Task.ID)},
		{"任务名称：", fmt.Sprintf("%s", s.Task.Name)},
		{"执行 ID：", fmt.Sprintf("%d", s.TaskLog.ID)},
		{"类   型：", fmt.Sprintf("%s", s.TaskLog.Lang)},
		{"触发方式：", fmt.Sprintf("%s", s.TaskLog.TriggerMethod)},
		{"表达式 ：", fmt.Sprintf("%s", s.TaskLog.Expr)},
		{"Worker：", fmt.Sprintf("%s", s.TaskLog.Addr)},
		{"执行时间：", fmt.Sprintf("[ %s - %s ]", s.TaskLog.StartTime.Format("2006-01-02 15:04:05"), s.TaskLog.EndTime.Format("2006-01-02 15:04:05"))},
		{"耗   时：", fmt.Sprintf("%s", fmt.Sprintf("%.3f", s.TaskLog.CostTime))},
		{"状   态：", fmt.Sprintf("%s", s.TaskLog.Result)},
	}
	stdOut := [][]string{
		{"返回内容：", fmt.Sprintf("%s", s.TaskLog.Resp+"test multi<div> &<br>line \n line2\n multi\n")}}
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
	var title = fmt.Sprintf("定时任务[%s]记录ID[%d]执行结果", s.Task.Name, s.TaskLog.ID)
	var msg = fmt.Sprintf(
		`任务 ID：%d
任务名称：%s
执行 ID：%d
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
		s.Task.ID, s.Task.Name, s.TaskLog.ID, s.TaskLog.Lang, s.TaskLog.TriggerMethod, s.TaskLog.Expr,
		s.TaskLog.Addr, s.TaskLog.StartTime.Format("2006-01-02 15:04:05"),
		s.TaskLog.EndTime.Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%.3f", s.TaskLog.CostTime), s.TaskLog.Result,
		s.TaskLog.Resp, s.TaskLog.ErrCode, s.TaskLog.ErrMsg)
	var apiData = map[string]interface{}{
		"task_log_id":    s.TaskLog.ID,
		"task_id":        s.Task.ID,
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
			conf.GetConf().Email.SMTPHost, conf.GetConf().Email.From, conf.GetConf().Email.Port,
			conf.GetConf().Email.Tls)
		err := notify.Send(title, mailMsg, s.Task.Notify.Email.Receivers, []string{})
		if err != nil {
			return err
		}
		hlog.Debugf("Task notify email send is success")
	}
	if s.Task.Notify.Lark.Enable && s.Task.Notify.Lark != nil && s.Task.Notify.Lark.Webhooks != nil {
		whs := (*(s.Task.Notify).Lark).Webhooks
		for _, v := range whs {
			notify := lark.NewLark(v.WebhookUrl, 1, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				return err
			}
		}
		hlog.Debugf("Task notify lark send is success")
	}
	if s.Task.Notify.Dingding.Enable {
		for _, v := range s.Task.Notify.Dingding.Webhooks {
			notify := dingding.NewDing(v.WebhookUrl, 2, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				return err
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
	CacheTask.TaskLogs[s.TaskLog.ID] = s
	CacheTask.Unlock()
}

func (s *taskSession) Delete() {
	//s.Close()
	CacheTask.Lock()
	delete(CacheTask.TaskLogs, s.TaskLog.ID)
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
		return nil //s.runMultiTasks(ctx, Task.ParentRunParallel, define.ParentTask, Task.ID, Task.ParentTaskIds...)
	})
	// server Task
	g.Go(func(ctx context.Context) error {
		return nil // s.runTask(ctx, Task.ID, Task.ID, define.MasterTask)
	})
	// childs Task
	g.Go(func(ctx context.Context) error {
		return nil // s.runMultiTasks(ctx, Task.ChildRunParallel, define.ChildTask, Task.ID, Task.ChildTaskIds...)
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
