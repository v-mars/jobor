package dispatcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"io"
	"jobor/internal/config"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/proto/pb"
	"jobor/internal/response"
	"jobor/internal/utils/errgroup"
	"jobor/pkg/logger"
	"jobor/pkg/notify/dingding"
	"jobor/pkg/notify/email"
	"jobor/pkg/notify/lark"
	"jobor/pkg/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AlarmPolicy int

type TaskEvent uint8

const (
	// AddEvent recv add event
	AddEvent TaskEvent = iota + 1
	// ChangeEvent recv delete task
	ChangeEvent
	// DeleteEvent recv delete task
	DeleteEvent
	// RunEvent run a task2
	RunEvent
	// KillEvent recv stop task
	KillEvent

	Never   AlarmPolicy = 0
	Always  AlarmPolicy = 1
	Failed  AlarmPolicy = 2
	Success AlarmPolicy = 3

	TriggerAuto = "auto"
	TriggerManual = "manual"
)

const (
	PubSubChannel = "jobor.event"
)


type Event struct {
	TaskID uint      // task id
	TE     TaskEvent // task event: add change delete stop task
}

func (ap AlarmPolicy) String() string {
	switch ap {
	case 0:
		return "从不"
	case 1:
		return "总是"
	case 2:
		return "失败"
	case 3:
		return "成功"
	default:
		return "未知告警策略"
	}
}


func Fn(channel, payload string) error {
	defer func() {
		if err := recover(); err != nil{
			stack := response.Stack(3)
			logger.Jobor.Errorf("Jobor 订阅者 panic, channel=%s, payload=%s, 错误信息: %s\n堆栈信息：%s", channel, payload, err, stack)
		}
	}()
	var ed = Event{}
	err:=json.Unmarshal([]byte(payload), &ed)
	if err!=nil{
		logger.Jobor.Errorf("Jobor 订阅者 channel=%s, payload=%s, payload解析错误: %s", channel, payload, err)
	}
	t, err := GetById(ed.TaskID)
	if err!=nil{
		logger.Jobor.Errorf("Jobor 订阅者 channel=%s, payload=%s, 获取任务错误: %s", channel, payload, err)
	}
	err = EventFunc(ed, t)
	if err!=nil{
		logger.Jobor.Errorf("Jobor 订阅者 channel=%s, payload=%s, EventFunc错误: %s", channel, payload, err)
	}
	return nil
}

var c = cron.New(cron.WithSeconds())

func EventFunc(ed Event,t tbs.JoborTask) error {
	switch ed.TE {
	case AddEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if !ok {
			logger.Jobor.Debugf("jobor cron taskId=%d add event, dispatcher cron entry is not exist", ed.TaskID)
			entryId,err := c.AddFunc(t.Expr, func() {
				RunTasks("add",TriggerAuto, t)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err!=nil{return err}
			entry:= c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}else {
			c.Remove(o.Entry.ID)
			entryId,err := c.AddFunc(t.Expr, func() {
				RunTasks("add",TriggerAuto, t)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err!=nil{return err}
			entry:= c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		logger.Jobor.Infof("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] add task is success", ed.TaskID, t.Name,t.Expr)
	case ChangeEvent:
		fmt.Printf("%++v", t)
		logger.Jobor.Debugf("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if t.Status!="running"{
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
			logger.Jobor.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] remove task is success", ed.TaskID, t.Name,t.Expr)
			return nil
		}
		if !ok {
			logger.Jobor.Debugf("jobor cron taskId=%d change event, dispatcher cron entry is not exist", ed.TaskID)
			entryId,err := c.AddFunc(t.Expr, func() {
				RunTasks("change", TriggerAuto,t)
			})
			if err!=nil{return err}
			entry:= c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}else {
			c.Remove(o.Entry.ID)
			entryId,err := c.AddFunc(t.Expr, func() {
				RunTasks("change", TriggerAuto,t)
			})
			if err!=nil{return err}
			entry:= c.Entry(entryId)
			RegistryDispatcher.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		logger.Jobor.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] add task is success", ed.TaskID, t.Name,t.Expr)
	case RunEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		if !ok{
			err:=fmt.Sprintf("jobor cron taskId=%d run event, dispatcher cron entry is not exist", ed.TaskID)
			logger.Jobor.Debugf(err)
			return fmt.Errorf(err)
		}else {
			job := c.Entry(o.EntryID).Job
			job.Run()
			logger.Jobor.Infof("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name,t.Expr)
		}
	case DeleteEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		if ok {
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
		}else {
			logger.Jobor.Debugf("jobor cron taskId=%d delete event, dispatcher cron entry is not exist", ed.TaskID)
		}
		o.Unlock()
		logger.Jobor.Infof("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name,t.Expr)
	case KillEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d kill event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := RegistryDispatcher.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if ok {
			c.Remove(o.Entry.ID)
			delete(RegistryDispatcher.cron, ed.TaskID)
		}else {
			logger.Jobor.Debugf("jobor cron taskId=%d kill event, dispatcher cron entry is not exist", ed.TaskID)
		}
		logger.Jobor.Infof("jobor cron taskId=%d kill event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name,t.Expr)
	}
	return nil
}

type task struct {
	sync.RWMutex
	Name    string
	Expr    string
	TaskId  uint
	EntryID cron.EntryID
	Entry   cron.Entry
}

type Registry struct {
	sync.RWMutex
	cron map[uint]task
}

var (
	RegistryDispatcher = Registry{cron: make(map[uint]task)}
)

func InitCron()  {
	taskList, err := GetAllRunningTask()
	if err!=nil{}
	c.Start()

	for _,t:=range taskList{
		logger.Jobor.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", t.ID,t.Name,t.Expr)
		entryId,err := c.AddFunc(t.Expr, func() {
			RunTasks("init",TriggerAuto, t)
		})
		if err!=nil{}
		entry:= c.Entry(entryId)
		RegistryDispatcher.cron[t.ID] = task{Name: t.Name, Expr: t.Expr, TaskId: t.ID, EntryID: entryId, Entry: entry}
		logger.Jobor.Infof("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is success", t.ID, t.Name,t.Expr)
	}

	//defer c.Stop()
	logger.Jobor.Info("jobor task dispatcher is start")
}

var CacheTask = cacheTask{TaskLogs: make(map[uint]*taskSession)}

type cacheTask struct {
	sync.RWMutex
	TaskLogs map[uint]*taskSession
}

type taskSession struct {
	sync.RWMutex
	TaskLog    *tbs.JoborLog
	Abort      chan bool
	Status     string
	Output     byte
	Err        error
	Conn       *grpc.ClientConn
	Stream     pb.TaskService_RunTaskClient
	TaskCancel context.CancelFunc
	TaskCtx    context.Context
	Task       *tbs.JoborTask
}

// RunTasks evt 事件, add/change
func RunTasks(evt,trigger string, t tbs.JoborTask)  {
	var s = taskSession{TaskCtx: context.Background()}
	var tx = db.DB
	jsonTime := tbs.JSONTime{Time: time.Now()}
	workers := GetWorkerByRoutePolicy(t.RoutingKey,t.RoutePolicy)
	var taskLog = tbs.JoborLog{Name: t.Name,Lang: t.Lang,TaskID: &t.ID,TriggerMethod: trigger,Expr: t.Expr,
		Data:t.Data,StartTime: jsonTime,Result: tbs.TaskLogStatusWait,
	}
	var startTimeTotal = time.Now()
	defer func() {
		defer func() {
			if errPanic := recover(); errPanic != nil{
				stack := response.Stack(3)
				s.Err = fmt.Errorf("defer panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
				logger.Jobor.Error(s.Err)
			}
		}()
		if errPanic := recover(); errPanic != nil{
			stack := response.Stack(3)
			s.Err = fmt.Errorf("RunTasks panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
			logger.Jobor.Error(s.Err)
			taskLog.ErrMsg=s.Err.Error()
			taskLog.Result=tbs.TaskLogStatusUnknown
		}else if taskLog.Result==tbs.TaskLogStatusTimeout{
			s.Err=fmt.Errorf("task %s[%d] lang=%s %s", t.Name,t.ID,t.Lang, "is timeout")
			logger.Jobor.Error(s.Err)
			taskLog.ErrMsg=s.Err.Error()
		}else if taskLog.Result==tbs.TaskLogStatusAbort{
			s.Err=fmt.Errorf("task %s[%d] lang=%s %s", t.Name,t.ID,t.Lang, "is aborted")
			logger.Jobor.Error(s.Err)
			taskLog.ErrMsg=s.Err.Error()
		}else if s.Err!=nil{
			s.Err=fmt.Errorf("task %s[%d] lang=%s %s", t.Name,t.ID,t.Lang, s.Err)
			logger.Jobor.Error(s.Err)
			taskLog.ErrMsg=s.Err.Error()
			taskLog.Result=tbs.TaskLogStatusFailed
		}else {
			taskLog.Result=tbs.TaskLogStatusSuccess
			logger.Jobor.Infof("task %s[%d] lang=%s 执行成功", t.Name,t.ID,t.Lang)
		}
		totalTime := time.Since(startTimeTotal)
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", totalTime.Seconds()), 64)
		taskLog.CostTime = float32(value)
		taskLog.EndTime = tbs.JSONTime{Time: time.Now()}
		if s.Err=tx.Model(&taskLog).Omit([]string{"Ctime", "Mtime"}...).Save(taskLog).Error;s.Err!=nil{
			s.Err=fmt.Errorf("保存tasklog错误: %s", s.Err)
			logger.Jobor.Error(s.Err)
			return
		}
		if s.Err=tx.Model(&t).Updates(tbs.JoborTask{
			Prev: jsonTime, Next: tbs.JSONTime{Time: RegistryDispatcher.cron[t.ID].Entry.Next}}).Error;s.Err!=nil{
			s.Err=fmt.Errorf("update task pre/next time err: %s", s.Err)
			logger.Jobor.Error(s.Err)
		}
		s.Err=s.Alarm()
		s.Delete()
	}()

	if s.Err= tx.Create(&taskLog).Error;s.Err!=nil{
		s.Err=fmt.Errorf("create tasklog err: %s",s.Err)
		return
	}

	s.TaskLog = &taskLog
	s.Task = &t
	s.Add()
	logger.Jobor.Infof("task %s[%d] lang %s success start,now time: %s ", t.Name,t.ID, t.Lang, time.Now())

	//var executor = DataCode{Lang: Lang(t.Lang),ScriptCode: t.Data.Code}
	var marshal []byte
	marshal, s.Err = json.Marshal(t.Data)
	if s.Err != nil {
		logger.Jobor.Errorf("DataCode Marshal  err: %s", s.Err)
		return
	}

	s.TaskCtx, s.TaskCancel = context.WithCancel(context.Background())
	timeoutCtx, timeoutCac := context.WithCancel(context.TODO())
	if t.Timeout > 0 { timeoutCtx, timeoutCac = context.WithTimeout(timeoutCtx, time.Second*time.Duration(t.Timeout)) }
	defer s.TaskCancel()
	defer timeoutCac()

	conn, w,errConn := TryGetGrpcClientConn(s.TaskCtx, workers)
	if errConn != nil{
		s.Err = errConn
		logger.Jobor.Errorf("TryGetGrpcClientConn err: %s",errConn)
		return
	}
	defer func(conn *grpc.ClientConn) {_ = conn.Close()}(conn)
	s.Conn = conn
	taskLog.Addr = w.Addr
	taskLog.Result = tbs.TaskLogStatusRunning
	if s.Err= tx.Model(&taskLog).Omit([]string{"Ctime", "Mtime"}...).Updates(taskLog).Error;s.Err!=nil{
		s.Err=fmt.Errorf("update tasklog addr/status err: %s", s.Err)
		logger.Error(s.Err)
		return
	}
	tc := pb.NewTaskServiceClient(conn)
	s.Stream, s.Err = tc.RunTask(s.TaskCtx,&pb.TaskRequest{TaskId: int32(t.ID),TaskLang: t.Lang,TaskData: marshal})
	if s.Err != nil {
		logger.Jobor.Errorf("run task err: %s", s.Err)
		return
	}
	defer func(Stream pb.TaskService_RunTaskClient) {_ = Stream.CloseSend()}(s.Stream)
	defer func(resStream pb.TaskService_RunTaskClient) {_ = resStream.CloseSend()}(s.Stream)
	logger.Jobor.Debugf("task %s[%d] RunTasks success", t.Name,t.ID)
	taskLog.Resp = ""
	streamChan := func() (chan *pb.StreamResponse, chan error) {
		errChan := make(chan error, 1)
		Message := make(chan *pb.StreamResponse, 1)
		go func() {
			defer func() {logger.Jobor.Debugf("task %s[%d] resStream.Recv finish", t.Name,t.ID)}()
			rec,errRecv := s.Stream.Recv()
			Message <- rec
			errChan <- errRecv
		}()
		return Message, errChan
	}
	for {
		msg,errChan:=streamChan()
		logger.Jobor.Debugf("task %s[%d] resStream.Recv start", t.Name,t.ID)
		select {
		//case <-s.Abort:
		case <-s.TaskCtx.Done():
			logger.Jobor.Debugf("task %s[%d] is abort", t.Name,t.ID)
			taskLog.Result=tbs.TaskLogStatusAbort
			return
		case <-timeoutCtx.Done():
			logger.Jobor.Debugf("task %s[%d] is timeout", t.Name,t.ID)
			taskLog.Result=tbs.TaskLogStatusTimeout
			return
		case d:=<-msg:
			logger.Jobor.Debugf("task %s[%d] stream recv data: %s", t.Name,t.ID,d.GetResp())
			taskLog.Resp += string(d.GetResp())
		case e:=<-errChan:
			if e == io.EOF {
				logger.Jobor.Infof("task %s[%d] stream recv finish", t.Name,t.ID)
				goto Next
			}
			if e != nil {
				logger.Jobor.Errorf("task %s[%d] resStream.Recv err: %s", t.Name,t.ID,e.Error())
				taskLog.Result="failed"
				s.Err = e
				return
			}
		}
	}
Next:
	if len(taskLog.Resp)>3000{
			taskLog.Resp = fmt.Sprintf("%s\n……省略过多内容……\n%s", taskLog.Resp[0:1499],taskLog.Resp[len(taskLog.Resp)-1499:])
		}
	var taskRespCode int
	judges := func() error {
		if s.Err != nil { return s.Err }
		taskRespCode, s.Err = s.GetRespCode()
		if s.Err != nil {
			s.Err = fmt.Errorf("get response code failed: %s",s.Err)
			logger.Jobor.Error(s.Err)
			return s.Err
		}
		s.TaskLog.ErrCode = taskRespCode
		if t.ExpectCode != taskRespCode {
			return fmt.Errorf("%s task %s[%d] resp code is %d, expect code %d", "server", t.Name, taskLog.ID, taskRespCode, t.ExpectCode)
		}
		if t.ExpectContent != "" {
			if !strings.Contains(taskLog.Resp, t.ExpectContent) {
				return fmt.Errorf("%s task %s[%d] resp context not contains expect content: %s", "server", t.Name, taskLog.ID, t.ExpectContent)
			}
		}
		return nil
	}
	s.Err = judges()
	if s.Err != nil {return}

}

func (s *taskSession)Alarm() error {
	isFailed,err:=utils.In([]string{tbs.TaskLogStatusFailed,tbs.TaskLogStatusUnknown,tbs.TaskLogStatusAbort,
		tbs.TaskLogStatusCancel}, s.TaskLog.Result)
	if err!=nil{return err}
	switch AlarmPolicy(s.Task.AlarmPolicy) {
	case Never:
		return nil
	case Always:
		err = s.Notify()
		if err != nil {
			return err
		}
	case Failed:
		if isFailed {
			err = s.Notify()
			if err != nil {
				return err
			}
		}
	case Success:
		if s.TaskLog.Result == tbs.TaskLogStatusSuccess{
			err = s.Notify()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *taskSession)Notify() error {
	var title = fmt.Sprintf("定时任务[%s]记录ID[%d]执行结果",s.Task.Name,s.TaskLog.ID)
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
		s.Task.ID,s.Task.Name,s.TaskLog.ID,s.TaskLog.Lang,s.TaskLog.TriggerMethod,s.TaskLog.Expr,
		s.TaskLog.Addr,s.TaskLog.StartTime.Format("2006-01-02 15:04:05"),
		s.TaskLog.EndTime.Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%.3f", s.TaskLog.CostTime),s.TaskLog.Result,
		s.TaskLog.Resp,s.TaskLog.ErrCode,s.TaskLog.ErrMsg)
	var apiData = map[string]interface{}{
		"task_log_id": s.TaskLog.ID,
		"task_id": s.Task.ID,
		"task_name": s.Task.Name,
		"lang": s.TaskLog.Lang,
		"worker_addr": s.TaskLog.Addr,
		"trigger_method": s.TaskLog.TriggerMethod,
		"cron_expr": s.TaskLog.Expr,
		"resp": s.TaskLog.Resp,
		"start_time": s.TaskLog.StartTime,
		"end_time": s.TaskLog.EndTime,
		"err_code": s.TaskLog.ErrCode,
		"err_msg": s.TaskLog.ErrMsg,
		"cost_time": s.TaskLog.CostTime,
		"status": s.TaskLog.Result,
	}
	if s.Task.Notify.Webhook.Enable  {
		logger.Jobor.Debugf("task notify webhook send is success, %s", apiData)
	}
	if s.Task.Notify.Email.Enable  {
		notify := email.NewMail(config.Configs.Email.Username,config.Configs.Email.Password,
			config.Configs.Email.SMTPHost,config.Configs.Email.From,config.Configs.Email.Port,
			config.Configs.Email.Tls)
		err := notify.Send(title, msg, s.Task.Notify.Email.Receivers, []string{})
		if err != nil {
			return err
		}
		logger.Jobor.Debugf("task notify email send is success")
	}
	if s.Task.Notify.Lark.Enable  {
		for _,v:=range s.Task.Notify.Lark.Webhooks{
			notify := lark.NewLark(v.WebhookUrl,1, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				return err
			}
		}
		logger.Jobor.Debugf("task notify lark send is success")
	}
	if s.Task.Notify.DingDing.Enable  {
		for _,v:=range s.Task.Notify.DingDing.Webhooks{
			notify := dingding.NewDing(v.WebhookUrl,1, v.Secret)
			err := notify.Send([]string{}, title, msg)
			if err != nil {
				return err
			}
		}
		logger.Jobor.Debugf("task notify dinding send is success")
	}
	logger.Jobor.Infof("task notify send is success")
	return nil
}

func (s *taskSession) GetRespCode () (int,error) {
	if len(s.TaskLog.Resp)>5{
		taskRespCode, err := strconv.Atoi(strings.TrimLeft(s.TaskLog.Resp[len(s.TaskLog.Resp)-5:], " "))
		if err != nil {
			logger.Jobor.Errorf("change str to int failed: %s",err)
			return DefaultExitCode,err
		}
		return taskRespCode,err
	}

	return DefaultExitCode,fmt.Errorf("response length less than 5")
}

func (s *taskSession)Add()  {
	CacheTask.Lock()
	CacheTask.TaskLogs[s.TaskLog.ID] = s
	CacheTask.Unlock()
}

func (s *taskSession)Delete()  {
	//s.Close()
	CacheTask.Lock()
	delete(CacheTask.TaskLogs, s.TaskLog.ID)
	CacheTask.Unlock()
}

func (s *taskSession)Close()  {
	//if s.Stream.Context()
	_ = s.Stream.CloseSend()
	_ = s.Conn.Close()
	s.TaskCancel()
}


// 并发控制
func con() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// if exist a err task,will stop all task
	g := errgroup.WithCancel(ctx)
	g.GOMAXPROCS(1)
	// parent tasks
	g.Go(func(ctx context.Context) error {
		return nil//s.runMultiTasks(ctx, task.ParentRunParallel, define.ParentTask, task.ID, task.ParentTaskIds...)
	})
	// server task
	g.Go(func(ctx context.Context) error {
		return nil// s.runTask(ctx, task.ID, task.ID, define.MasterTask)
	})
	// childs task
	g.Go(func(ctx context.Context) error {
		return nil// s.runMultiTasks(ctx, task.ChildRunParallel, define.ChildTask, task.ID, task.ChildTaskIds...)
	})
	err := g.Wait()
	if err != nil {
		fmt.Errorf("task run failed taskId[%d] err: %s", 1, err)
	}
}


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
