package schedule

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"jobor/internal/app/jobor/define"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/pkg/logger"
	"net/http"
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

var co = cron.New(cron.WithSeconds())

func EventFunc(ed Event,t tbs.JoborTask) error {
	switch ed.TE {
	case AddEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := DispatcherRegistry.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if !ok {
			logger.Jobor.Debugf("jobor cron taskId=%d add event, dispatcher cron entry is not exist", ed.TaskID)
			entryId,err := co.AddFunc(t.Expr, func() {
				CronAddFuncCommon("add","auto", t)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err!=nil{return err}
			entry:= co.Entry(entryId)
			DispatcherRegistry.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}else {
			co.Remove(o.Entry.ID)
			entryId,err := co.AddFunc(t.Expr, func() {
				CronAddFuncCommon("add","auto", t)
			}) //2 * * * * *, 2 表示每分钟的第2s执行一次
			if err!=nil{return err}
			entry:= co.Entry(entryId)
			DispatcherRegistry.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		logger.Jobor.Infof("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] add task is success", ed.TaskID, t.Name,t.Expr)
	case ChangeEvent:
		fmt.Printf("%++v", t)
		logger.Jobor.Debugf("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := DispatcherRegistry.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if t.Status!="running"{
			co.Remove(o.Entry.ID)
			delete(DispatcherRegistry.cron, ed.TaskID)
			logger.Jobor.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] remove task is success", ed.TaskID, t.Name,t.Expr)
			return nil
		}
		if !ok {
			logger.Jobor.Debugf("jobor cron taskId=%d change event, dispatcher cron entry is not exist", ed.TaskID)
			entryId,err := co.AddFunc(t.Expr, func() {
				CronAddFuncCommon("change", "auto",t)
			})
			if err!=nil{return err}
			entry:= co.Entry(entryId)
			DispatcherRegistry.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}else {
			co.Remove(o.Entry.ID)
			entryId,err := co.AddFunc(t.Expr, func() {
				CronAddFuncCommon("change", "auto",t)
			})
			if err!=nil{return err}
			entry:= co.Entry(entryId)
			DispatcherRegistry.cron[ed.TaskID] = task{Name: t.Name, Expr: t.Expr, TaskId: ed.TaskID, EntryID: entryId, Entry: entry}
		}
		logger.Jobor.Infof("jobor cron taskId=%d change event [name=%s, expr=\"%s\"] add task is success", ed.TaskID, t.Name,t.Expr)
	case RunEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := DispatcherRegistry.cron[ed.TaskID]
		if !ok{
			err:=fmt.Sprintf("jobor cron taskId=%d run event, dispatcher cron entry is not exist", ed.TaskID)
			logger.Jobor.Debugf(err)
			return fmt.Errorf(err)
		}else {
			job := co.Entry(o.EntryID).Job
			job.Run()
			logger.Jobor.Infof("jobor cron taskId=%d run event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name,t.Expr)
		}
	case DeleteEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := DispatcherRegistry.cron[ed.TaskID]
		o.Lock()
		if ok {
			co.Remove(o.Entry.ID)
			delete(DispatcherRegistry.cron, ed.TaskID)
		}else {
			logger.Jobor.Debugf("jobor cron taskId=%d delete event, dispatcher cron entry is not exist", ed.TaskID)
		}
		o.Unlock()
		logger.Jobor.Infof("jobor cron taskId=%d delete event [name=%s, expr=\"%s\"] is success", ed.TaskID, t.Name,t.Expr)
	case KillEvent:
		logger.Jobor.Debugf("jobor cron taskId=%d kill event [name=%s, expr=\"%s\"] is start", ed.TaskID,t.Name,t.Expr)
		o, ok := DispatcherRegistry.cron[ed.TaskID]
		o.Lock()
		defer func() {o.Unlock()}()
		if ok {
			co.Remove(o.Entry.ID)
			delete(DispatcherRegistry.cron, ed.TaskID)
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
	DispatcherRegistry = Registry{cron: make(map[uint]task)}
)

func InitCron()  {
	taskList, err := GetAllRunningTask()
	if err!=nil{}
	co.Start()

	for _,t:=range taskList{
		logger.Jobor.Debugf("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is start", t.ID,t.Name,t.Expr)
		entryId,err := co.AddFunc(t.Expr, func() {
			CronAddFuncCommon("init","auto", t)
		})
		if err!=nil{}
		entry:= co.Entry(entryId)
		DispatcherRegistry.cron[t.ID] = task{Name: t.Name, Expr: t.Expr, TaskId: t.ID, EntryID: entryId, Entry: entry}
		logger.Jobor.Infof("jobor cron taskId=%d add event [name=%s, expr=\"%s\"] is success", t.ID, t.Name,t.Expr)
	}

	//defer co.Stop()
	logger.Jobor.Info("jobor task dispatcher is start")
}

func CronAddFuncCommon(evt,trigger string, t tbs.JoborTask)  {
	var err error
	jsonTime := tbs.JSONTime{Time: time.Now()}
	var taskLog = tbs.JoborLog{Name: t.Name,Lang: t.Lang,TaskID: &t.ID,TriggerMethod: trigger,Expr: t.Expr,
		Data:t.Data,StartTime: jsonTime,
	}
	var startTimeTotal = time.Now()
	defer func() {
		taskLog.Result="failed"
		if errPanic := recover(); errPanic != nil{
			stack := response.Stack(3)
			err = fmt.Errorf("CronAddFuncCommon panic, 错误信息: %s\n堆栈信息：%s", errPanic, stack)
			logger.Jobor.Error(err)
			taskLog.ErrMsg=err.Error()
		}else if err!=nil{
			err=fmt.Errorf("taskId=%d, name=%s, lang=%s 错误信息: %s\n", t.ID, t.Name, t.Lang, err)
			logger.Jobor.Error(err)
			taskLog.ErrMsg=err.Error()
		}else {
			taskLog.Result="success"
			logger.Jobor.Infof("taskId=%d, name=%s, lang=%s 执行成功", t.ID, t.Name, t.Lang)
		}
		totalTime := time.Since(startTimeTotal)
		//float32(totalTime.Seconds())
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", totalTime.Seconds()), 64)
		taskLog.CostTime = float32(value)
		taskLog.EndTime = tbs.JSONTime{Time: time.Now()}
		//var MapData map[string]interface{}
		//if MapData,err=convert.StructToMap(taskLog); err!=nil{
		//	err=fmt.Errorf("结构体转化为map错误: %s", err)
		//	return
		//}
		if err= db.DB.Model(&taskLog).Omit([]string{"Ctime", "Mtime"}...).Save(taskLog).Error;err!=nil{
			err=fmt.Errorf("保存tasklog错误: %s", err)
			logger.Error(err)
			return
		}
	}()

	if err= db.DB.Create(&taskLog).Error;err!=nil{
		err=fmt.Errorf("创建tasklog错误: %s", err)
		return
	}

	fmt.Printf("taskId=%d name=%s cron %s success %s, data: %s\n", t.ID,t.Name, evt, time.Now(), t.Data)
	var executor = DataCode{Lang: Lang(t.Lang),ScriptCode: t.Data.Code}
	out:=executor.Run(context.TODO())
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(out)
	//fmt.Println("stdout: ",buf.String())
	taskLog.Result="success"
	taskLog.Resp=buf.String()
	//if len(buf.String())>3000{
	//	taskLog.Resp = fmt.Sprintf("%s\n……省略过多内容……\n%s", buf.String()[0:1499],buf.String()[len(buf.String())-1499:])
	//}else {
	//	taskLog.Resp=buf.String()
	//}
	//fmt.Println(out)
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


var (
	// Cron2 schedule loop
	Cron2 *cacheSchedule2
)

// task2 running status
type task2 struct {
	sync.RWMutex
	name        string
	cronexpr    string
	close       chan struct{}        // stop schedule
	running     bool                 // task2 is running
	stop        bool                 // stop run task2
	ctxcancel   context.CancelFunc   // store cancelfunc could cancel all task2 by this cancel
	starttime   int64                // run task2 time(ms)
	endtime     int64                // end run task2 time(ms)
	//logcaches   map[string]LogCacher // task2 runing logcaches
	taskids     []string             // save tasks。parent taskids、mainid、childids
	status      int                  // task2 run res fail: -1 success:1
	//next        Next                 // it save a func Next by route policy
	//Trigger     define.Trigger       // how to trigger task2
	errTaskID   string               // run fail task2's id
	errTask     string               // run fail task2's id
	errCode     int                  // failed task2 return code
	errMsg      string               // task2 run failed errmsg
	//errTasktype define.TaskRespType  // failed task2 type
}

type cacheSchedule2 struct {
	sync.RWMutex
	sch map[string]*task2
}

// Init start run already exists task2 from db
/*func Init() error {
	Cron2 = &cacheSchedule2{
		sch: make(map[string]*task2),
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	isinstalll, err := model.QueryIsInstall(ctx)
	if err != nil {
		log.Error("model.QueryIsInstall failed", zap.Error(err))
		return fmt.Errorf("model.QueryIsInstall failed: %w", err)
	}
	if !isinstalll {
		log.Debug("Crocodile is Not Install")
		return nil
	}
	eps, _, err := model.GetTasks(ctx, 0, 0, "", "", "")
	if err != nil {
		log.Error("GetTasks failed", zap.Error(err))
		return err
	}

	for _, task := range eps {
		Cron2.Add(task.ID, task.Name, task.Cronexpr, GetRoutePolicy(task.HostGroupID, task.RoutePolicy))
	}
	//log.Info("init task2 success", zap.Int("Total", len(eps)))
	return nil
}*/

// Add task2 to schedule
//func (s *cacheSchedule2) Add(taskID, taskName string, cronExpr string, next Next) {
//	//log.Debug("Start Add task2 ", zap.String("Name", taskName))
//
//	task := task2{
//		Name:     taskName,
//		cronexpr: cronExpr,
//		close:    make(chan struct{}),
//		next:     next,
//	}
//	s.Lock()
//	// 如果多个用户同时修改 确保不会冲突
//	oldtask, exist := s.sch[taskID]
//	if exist {
//		close(oldtask.close)
//		if oldtask.ctxcancel != nil {
//			oldtask.ctxcancel()
//		}
//	}
//	s.sch[taskID] = &task
//	s.Unlock()
//	//log.Info("Add Task success", zap.String("taskid", taskID), zap.String("Name", taskName))
//	go s.runSchedule(taskID)
//}

// Del schedule task2
// if delete taskid,this taskid must be remove from other task2's child or parent
/*func (s *cacheSchedule2) Del(id string) {
	log.Info("start delete task2", zap.String("taskid", id))
	task2, exist := s.gettask(id)
	if exist {
		log.Debug("start clean ", zap.String("id", id))
		task2.Lock()
		delete(s.sch, id)
		task2.Unlock()
		go s.Clean(task2)

	}
}*/

// Clean task2
/*func (s *cacheSchedule2) Clean(task *task2) {
	log.Info("start clean task2", zap.String("task2", task.Name))
	// if task.ctxcancel != nil {
	// 	task.ctxcancel()
	// }
	close(task.close)
	log.Info("clean task2 success", zap.String("task2", task.Name))
	return
}*/

// kill task2
/*func (s *cacheSchedule2) KillTask(taskid string) {
	task, exist := s.gettask(taskid)
	if !exist {
		log.Error("stoptask failed,task2 is not exist", zap.String("taskid", taskid))
		return
	}
	if task.ctxcancel != nil {
		task.ctxcancel()
	}
	return
}*/

/*func (s *cacheSchedule2) gettask(id string) (*task2, bool) {
	s.RLock()
	task, ok := s.sch[id]
	s.RUnlock()
	if ok && task.logcaches == nil {
		task.logcaches = make(map[string]LogCacher)
	}
	return task, ok
}*/

// saveLog save task2 resp log
/*func (s *cacheSchedule2) storelog(runbyid string, task *task2) error {
	// read all log
	// put logcache to locachepool
	tasklog := &define.Log{
		Name:         task.Name,
		RunByTaskID:  runbyid,
		StartTime:    task.starttime,
		EndTime:      task.endtime,
		TotalRunTime: int(task.endtime - task.starttime),
		Status:       task.status,
		//Trigger:      task.Trigger,
		ErrCode:      task.errCode,
		ErrMsg:       task.errMsg,
		//ErrTasktype:  task.errTasktype,
		ErrTaskID:    task.errTaskID,
		ErrTask:      task.errTask,
		//TaskResps:    make([]*define.TaskResp, 0, len(task.logcaches)),
	}

	for Name, logcache := range task.logcaches {
		var (
			taskresp define.TaskResp
			ok       bool
		)
		// if ok,code runnhost tasktype valid
		taskresp, ok = logcache.Get().(define.TaskResp)

		id, tasktype, _ := splitname(Name)
		// taskresp.Task = task.Name
		if ok {
			taskresp.TaskID = id
			taskresp.TaskTypeStr = taskresp.TaskType.String()
			taskresp.LogData = logcache.ReadAll()

		} else {
			taskresp = define.TaskResp{
				TaskType:    tasktype,
				TaskTypeStr: tasktype.String(),
				TaskID:      id,
				LogData:     logcache.ReadAll(),
			}
		}
		if logcache.GetTaskStatus() == define.TsWait {
			taskresp.Status = define.TsCancel.String()
		} else {
			taskresp.Status = logcache.GetTaskStatus().String()
		}

		// taskresp
		logcache.Close()
		// Put coolpool after logcache close

		cachepool.Put(logcache)
		tasklog.TaskResps = append(tasklog.TaskResps, &taskresp)
	}

	go alarm.JudgeNotify(tasklog)
	// save log
	err := model.SaveLog(context.Background(), tasklog)

	return err
}*/

// runSchedule start run cronexpr schedule
/*func (s *cacheSchedule2) runSchedule(taskid string) {
	task, exist := s.gettask(taskid)
	if !exist {
		return
	}
	//log.Info("start run cronexpr", zap.Any("task2", task), zap.String("id", taskid))

	Expr, err := cronexpr.Parse(task.cronexpr)
	if err != nil {
		//log.Error("cronexpr parse failed", zap.Error(err))
		return
	}
	var (
		last time.Time
		next time.Time
	)
	last = time.Now()
	for {
		//log.Debug("all task2", zap.Any("tasks", s.sch))
		next = Expr.Next(last)
		select {
		case <-task.close:
			//log.Info("close Schedule", zap.String("taskID", taskid), zap.Any("task2", task))
			return
		case <-time.After(next.Sub(last)):
			last = next
			if task.stop {
				//log.Error("task2 is stop run", zap.String("task2", task.Name))
			} else {
				//go s.RunTask(taskid, define.Auto)
			}
		}
	}
}*/

func generatename(id string, tasktype define.TaskRespType) string {
	return id + "_" + strconv.Itoa(int(tasktype))
}

func splitname(taskname string) (string, define.TaskRespType, error) {
	res := strings.Split(taskname, "_")
	if len(res) != 2 {
		return "", 0, fmt.Errorf("split %s failed", taskname)
	}
	id := res[0]
	tasktype, err := strconv.Atoi(res[1])
	if err != nil {
		return "", 0, nil
	}
	return id, define.TaskRespType(tasktype), nil
}

// RunTask start run a task2
/*func (s *cacheSchedule2) RunTask(taskid string, trigger define.Trigger) {
	var (
		masterlogcache LogCacher
		ctx            context.Context
		cancel         context.CancelFunc
		err            error
		task2           *define.GetTask
		g              *errgroup.Group
	)
	log.Info("start run task2", zap.String("taskid", taskid))

	task, exist := s.gettask(taskid)
	if !exist {
		log.Error("this is bug, taskid not exist", zap.String("taskid", taskid), zap.Any("sch", s.sch))
		return
	}
	// if master task2 is running,will do not run this time
	if task.running {
		log.Warn("task2 is running,so not run now", zap.String("task2", task.Name))
		return
	}

	masterlogcache = cachepool.Get().(LogCacher) // this log cache is main
	masterlogcache.SetTaskStatus(define.TsWait)
	masterlogcache.Clean()

	task.errTaskID = ""
	task.errTask = ""
	task.errCode = 0
	task.errMsg = ""
	task.errTasktype = 0
	task.Trigger = trigger
	mastername := generatename(taskid, define.MasterTask)
	task.Lock()
	task.logcaches[mastername] = masterlogcache
	task.Unlock()

	task.running = true
	task.starttime = time.Now().UnixNano() / 1e6

	ctx, cancel = context.WithCancel(context.Background())
	// save control ctx
	task.ctxcancel = cancel
	defer cancel()
	task2, err = model.GetTaskByID(context.Background(), taskid)
	if err != nil {
		log.Error("can'task get main taskID from dataBase", zap.String("task2", task2.Name))
		cachepool.Put(masterlogcache)
		return
	}
	// 只能手动调用运行
	if !task2.RunV1 && trigger == define.Auto {
		log.Error("task2 is forbid auto trigger run by schedule", zap.String("task2", task2.Name))
		cachepool.Put(masterlogcache)
		return
	}

	task.taskids = make([]string, 0, 1+len(task2.ParentTaskIds)+len(task2.ChildTaskIds))

	for _, id := range task2.ParentTaskIds {
		Name := generatename(id, define.ParentTask)
		logcache := cachepool.Get().(LogCacher)
		logcache.SetTaskStatus(define.TsWait)
		logcache.Clean()
		task.Lock()
		task.logcaches[Name] = logcache
		task.Unlock()
		task.taskids = append(task.taskids, Name)
	}

	task.taskids = append(task.taskids, mastername)

	for _, id := range task2.ChildTaskIds {
		Name := generatename(id, define.ChildTask)
		logcache := cachepool.Get().(LogCacher)
		logcache.SetTaskStatus(define.TsWait)
		logcache.Clean()
		task.Lock()
		task.logcaches[Name] = logcache
		task.Unlock()
		task.taskids = append(task.taskids, Name)
	}

	// if exist a err task2,will stop all task2
	g = errgroup.WithCancel(ctx)
	g.GOMAXPROCS(1)
	// parent tasks
	g.Go(func(ctx context.Context) error {
		return s.runMultiTasks(ctx, task2.ParentRunParallel, define.ParentTask, task2.ID, task2.ParentTaskIds...)
	})
	// master task2
	g.Go(func(ctx context.Context) error {
		return s.runTask(ctx, task2.ID, task2.ID, define.MasterTask)
	})
	// childs task2
	g.Go(func(ctx context.Context) error {
		return s.runMultiTasks(ctx, task2.ChildRunParallel, define.ChildTask, task2.ID, task2.ChildTaskIds...)
	})
	err = g.Wait()
	if err != nil {
		log.Error("task2 run failed", zap.String("taskid", taskid), zap.Error(err))
	}
	goto Over
Over:
	task.running = false
	task.endtime = time.Now().UnixNano() / 1e6

	if task.errTaskID == "" {
		task.status = 1
	}
	err = s.storelog(taskid, task)
	if err != nil {
		log.Error("save task2 log failed", zap.Error(err))
	}
}*/

// run multi tasks
// if hash one task2 err, will exit all task2
// TODO: task2 run err whether influence  other task2
// TODO: multi task2 set RunParallel total
/*func (s *cacheSchedule2) runMultiTasks(ctx context.Context, RunParallel bool,
	tasktype define.TaskRespType, runbyid string, taskids ...string) error {
	if len(taskids) == 0 {
		return nil
	}
	log.Info("Start RunV1 Task", zap.Strings("taskids", taskids), zap.String("tasktype", tasktype.String()))
	var maxproc int
	if RunParallel {
		maxproc = len(taskids)
	} else {
		maxproc = 1
	}
	g := errgroup.WithCancel(ctx)
	g.GOMAXPROCS(maxproc)
	for _, id := range taskids {
		taskid := id
		g.Go(func(ctx context.Context) error {
			return s.runTask(ctx, taskid, runbyid, tasktype)
		})
	}
	err := g.Wait()
	return err
}*/

// runTask start run task2,log will store
/*func (s *cacheSchedule2) runTask(ctx context.Context, id, //real run task2 id
	runbyid  string, taskresptype define.TaskRespType) error {
	var (
		// real need task2 status
		// realtask *task2
		logcache LogCacher
		// recverr      error
		taskrespcode = tasktype.DefaultExitCode
		// recv grpc stream
		taskresp *pb.TaskResp
		// error
		err error
		// task2 data
		taskdata *define.GetTask
		// worker conn
		conn *grpc.ClientConn
		// task2 run data
		tdata []byte
		// recv grpc stream
		taskrespstream pb.Task_RunTaskClient
		// grpc client
		taskclient pb.TaskClient
		taskreq    *pb.TaskReq
		ctxcancel  context.CancelFunc
		taskctx    context.Context
		realtask   *task2
		output     []byte
	)

	// when func exit,check res and judge whatever alarm

	runbytask, exist := s.gettask(runbyid)
	if !exist {
		// if happen,this is a bug,
		log.Error("this is a bug,task2 should exist,but can not get task2,", zap.String("taskid", runbyid), zap.Any("allschedule", s.sch))
		err = fmt.Errorf("[bug] can not get taskid %s from schuedle: %v", id, s.sch)
		return err
	}

	if taskresptype == define.MasterTask {
		realtask = runbytask
	} else {
		realtask, exist = s.gettask(id)
		if !exist {
			log.Error("this is a bug,task2 should exist,but can not get task2,", zap.String("taskid", runbyid), zap.Any("allschedule", s.sch))
			err = fmt.Errorf("[bug] can not get taskid %s from schuedle: %v", id, s.sch)
			return err
		}
	}
	Name := generatename(id, taskresptype)
	runbytask.RLock()
	logcache, exist = runbytask.logcaches[Name]
	runbytask.RUnlock()
	if !exist {
		// it happen,it is a bug
		warnbug := fmt.Sprintf("[bug] can get master task2's %s logcache from logcaches: %v", id, runbytask.logcaches)
		log.Error(warnbug)
		logcache = cachepool.Get().(LogCacher)
		logcache.WriteString(warnbug)
		runbytask.Lock()
		runbytask.logcaches[Name] = logcache
		runbytask.Unlock()
	}

	logcache.SetTaskStatus(define.TsRun)

	queryctx, querycancel := context.WithTimeout(ctx,
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer querycancel()

	taskdata, err = model.GetTaskByID(queryctx, id)
	if err != nil {
		log.Error("model.GetTaskByID failed", zap.String("taskid", id), zap.Error(err))
		logcache.WriteStringf("Get Task id %s failed: %v", id, err)
		goto Check
	}
	logcache.WriteStringf("Start Prepare Task %s[%s]", taskdata.Name, id)
	logcache.WriteStringf("Start Conn Worker Host For Task %s[%s]", taskdata.Name, id)

	conn, err = tryGetRCCConn(ctx, realtask.next)
	if err != nil {
		log.Error("tryGetRpcConn failed", zap.Error(err))
		err = fmt.Errorf("Get Rpc Conn Failed From Hostgroup %s[%s] Err: %v",
			taskdata.HostGroup, taskdata.HostGroupID, err)
		logcache.WriteStringf(err.Error())
		goto Check
	}

	logcache.WriteStringf("Success Conn Worker Host[%s]", conn.Target())
	logcache.WriteStringf("Start Get Task %s[%s] RunV1 Data", taskdata.Name, id)
	// Marshal task2 data
	tdata, err = json.Marshal(taskdata.TaskData)
	if err != nil {
		log.Error("json.Marshal", zap.Error(err))
		logcache.WriteStringf("Marshal task2 %s[%s]'s RunData [%v] failed: %v", taskdata.Name, id, taskdata.TaskData, err)
		goto Check
	}

	// task2 run data
	taskreq = &pb.TaskReq{
		TaskId:   id,
		TaskType: int32(taskdata.TaskType),
		TaskData: tdata,
	}

	logcache.WriteStringf("Success Get Task %s[%s] RunV1 Data", taskdata.Name, id)

	logcache.WriteStringf("Start RunV1 Task %s[%s] On Host[%s]", taskdata.Name, id, conn.Target())

	// taskctx only use RunTask
	if taskdata.Timeout > 0 {
		taskctx, ctxcancel = context.WithTimeout(ctx, time.Second*time.Duration(taskdata.Timeout))
	} else {
		taskctx, ctxcancel = context.WithCancel(ctx)
	}

	defer ctxcancel()
	taskclient = pb.NewTaskClient(conn)

	taskrespstream, err = taskclient.RunTask(taskctx, taskreq)
	if err != nil {
		log.Error("RunV1 task2 failed", zap.Error(err))
		logcache.WriteStringf("RunV1 Task %s[%s] TaskData [%v] failed:%v", taskdata.Name, id, taskreq, err)
		goto Check
	}

	// RunTask default resp code

	logcache.WriteStringf("Task %s[%s]  Output----------------", taskdata.Name, id)
	for {
		// Recv return err is nil or io.EOF
		// the last lastrecv must be return code 3 byte
		taskresp, err = taskrespstream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
				taskrespcode = logcache.GetCode()
				goto Check
			}
			err = DealRPCErr(err)
			// taskrespcode =
			logcache.WriteStringf("Task %s[%s] RunV1 Fail: %v", taskdata.Name, id, err.Error())
			// Alarm
			log.Error("Recv failed", zap.Error(err))
			// err = resp.GetMsgErr(taskrespcode)
			goto Check
		}
		logcache.Write(taskresp.GetResp())
		output = append(output, taskresp.GetResp()...)
	}
	// logcache.WriteStringf("Task %s[%s] End Output-------------------\n", taskdata.Name, id)

Check:
	// logcache.WriteStringf("Task %s[%s] RunV1 Over\n", taskdata.Name, id)
	// 记录任务的状态

	// save task2 returncode,tasktype,if task2 could find run host,run host will be save hear
	tmptaskresp := define.TaskResp{
		Task:     realtask.Name,
		ScriptCode:     taskrespcode,
		TaskType: taskresptype,
	}
	if conn != nil {
		// if conn worker failed,can not get worker host
		tmptaskresp.RunHost = conn.Target()
	}
	logcache.Save(tmptaskresp)
	// 当终止任务时，第一个任务取消的任务不经过这里处理，后续的任务才会经过这里处理
	// 所以需要判断t.errTaskId 为空时才经过这里处理
	if err != nil && runbytask.errTaskID != "" {
		select {
		case <-ctx.Done():
			log.Error("task2 is cancel", zap.String("task2", realtask.Name))
			logcache.WriteStringf("task2 %s[%s] is canceled", realtask.Name, id)
			logcache.SetTaskStatus(define.TsCancel)
			return nil
		default:
		}
	}
	var alarmerr error
	// if a task2 fail other task2 will return context.Canceled,but it can not alarm
	// because the first err task2 always alarm,so other task2 do not alarm
	// and the first err task2's errmsg will save tasking

	// check task2 resp code and resp content
	judgeres := func() error {
		if err != nil {
			return err
		}
		if taskdata.ExpectCode != taskrespcode {
			return fmt.Errorf("%s task2 %s[%s] resp code is %d,want resp code %d", taskresptype.String(), id, taskdata.Name, taskrespcode, taskdata.ExpectCode)
		}
		if taskdata.ExpectContent != "" {
			if !strings.Contains(string(output), taskdata.ExpectContent) {
				return fmt.Errorf("%s task2 %s[%s] resp context not contains expect content: %s", taskresptype.String(), id, taskdata.Name, taskdata.ExpectContent)
			}
		}
		return nil
	}
	alarmerr = judgeres()
	if alarmerr != nil {
		//  task2 run err
		// 只运行到这里一次
		// runbytask.status = -1 // task2 fail
		log.Error("task2 run fail", zap.String("task2", realtask.Name), zap.Error(err))
		if runbytask.errTaskID == "" {
			runbytask.status = -1
			runbytask.errTaskID = id
			runbytask.errTask = realtask.Name
			runbytask.errCode = taskrespcode
			runbytask.errMsg = alarmerr.Error()
			runbytask.errTasktype = taskresptype
			// log.Error("-----------------task2 run fail", zap.String("task2", realtask.Name), zap.Error(err))
			logcache.SetTaskStatus(define.TsFail)
		}

	} else {
		log.Error("task2 run success", zap.String("task2", realtask.Name))
		logcache.SetTaskStatus(define.TsFinish)
		// 如有任务失败，那么还未运行的任务可以标记为取消
		// 如果失败的任务还存在并行任务，那么
	}
	return alarmerr
}*/

//  sort running task2
type runningTask []*define.RunTask

func (rt runningTask) Len() int { return len(rt) }
func (rt runningTask) Less(i, j int) bool {
	ii, err := strconv.Atoi(rt[i].ID)
	if err != nil {
		//log.Error("change ID to int failed", zap.String("id", rt[i].ID))
	}
	jj, err := strconv.Atoi(rt[j].ID)
	if err != nil {
		//log.Error("change ID to int failed", zap.String("id", rt[j].ID))
	}
	return ii < jj
}
func (rt runningTask) Swap(i, j int) { rt[i], rt[j] = rt[j], rt[i] }

// GetRunningtask return all running task2
/*func (s *cacheSchedule2) GetRunningtask() []*define.RunTask {
	runtasks := runningTask{}
	s.RLock()
	defer s.RUnlock()
	for taskid, task2 := range s.sch {
		if !task2.running {
			continue
		}
		// task2.running
		runtask := define.RunTask{
			ID:           taskid,
			Name:         task2.Name,
			StartTimeStr: utils.UnixToStr(task2.starttime / 1e3),
			StartTime:    task2.starttime,
			RunTime:      int(time.Now().Unix() - task2.starttime/1e3),
			TriggerStr:   task2.Trigger.String(),
			Cronexpr:     task2.cronexpr,
		}
		runtasks = append(runtasks, &runtask)
	}
	sort.Sort(runtasks)
	return runtasks
}*/

// GetRunTaskStaus return
/*func (s *cacheSchedule2) GetRunTaskStaus(taskid string) []*define.TaskStatusTree {
	retTasksStatus := define.GetTasksTreeStatus()

	parentTasksStatus := retTasksStatus[0]

	taskinfo, exist := s.gettask(taskid)
	if !exist {
		return nil
	}
	mainTaskStatus := retTasksStatus[1]
	mainTaskStatus.Name = taskinfo.Name
	mainTaskStatus.ID = taskid

	childTasksStatus := retTasksStatus[2]

	task2, exist := s.gettask(taskid)
	if !exist {
		return nil
	}
	var status = define.TsNoData
	var isSet = false
	var change = false
	//log.Debug("start get task2 run status", zap.Strings("ids", task2.taskids))
	for _, Name := range task2.taskids {
		// if !task2.running {
		// 	log.Error("task2 is not run", zap.String("Name", Name))
		// 	return nil
		// }
		task2.RLock()
		//logcache := task2.logcaches[Name]
		task2.RUnlock()
		id, _, _ := splitname(Name)

		taskinfo, _ := s.gettask(id)

		if Name == generatename(taskid, define.MasterTask) {
			parentTasksStatus.Status = status.String()
			isSet = false
			status = define.TsNoData
			change = true
			// main task2
			// log.Debug("start get main status" + id + ":" + logcache.GetTaskStatus().String())
			//mainTaskStatus.Status = logcache.GetTaskStatus().String()
			mainTaskStatus.TaskType = define.MasterTask
		} else if change == false {
			// parent taskids
			// 如果全部为wait就设置为wait
			// 如果全部成功那么就设置为finish

			// 如果任务存在run那么就设置为run
			// 如果任务有失败那么就设置为fail
			if !isSet {
				if logcache.GetTaskStatus() == define.TsRun || logcache.GetTaskStatus() == define.TsFail {
					isSet = true
					status = logcache.GetTaskStatus()
				} else {
					status = logcache.GetTaskStatus()
				}

			}
			// log.Debug("start get parent status" + id + ":" + status.String())
			parentaskStatus := &define.TaskStatusTree{
				Name:     taskinfo.Name,
				ID:       id,
				TaskType: define.ParentTask,
				//Status:   logcache.GetTaskStatus().String(),
			}
			parentTasksStatus.TaskType = define.ParentTask
			parentTasksStatus.Children = append(parentTasksStatus.Children, parentaskStatus)

		} else {
			if !isSet {
				if logcache.GetTaskStatus() == define.TsRun || logcache.GetTaskStatus() == define.TsFail {
					isSet = true
					status = logcache.GetTaskStatus()
				} else {
					status = logcache.GetTaskStatus()
				}
			}
			// log.Debug("start get parent status" + id + ":" + status.String())

			// child taskids
			childaskStatus := &define.TaskStatusTree{
				Name:     taskinfo.Name,
				ID:       id,
				TaskType: define.ChildTask,
				Status:   logcache.GetTaskStatus().String(),
			}
			childTasksStatus.TaskType = define.ChildTask
			childTasksStatus.Children = append(childTasksStatus.Children, childaskStatus)
		}
	}
	childTasksStatus.Status = status.String()

	//log.Debug("TaskRunStatus", zap.Any("status", retTasksStatus))
	return retTasksStatus
}*/

// GetLogCache get log cache
/*func (s *cacheSchedule2) GetRunTaskLogCache(taskid, realtaskid string, tasktype define.TaskRespType) (LogCacher, error) {
	task, exist := s.gettask(taskid)
	if !exist {
		return nil, fmt.Errorf("can get task2 id %s", taskid)
	}
	if !task.running {
		return nil, fmt.Errorf("main task2 %s[%s] is not running,can'task get running log", task.Name, taskid)
	}
	Name := generatename(realtaskid, tasktype)
	task.RLock()
	logcache, exist := task.logcaches[Name]
	task.RUnlock()
	if !exist {
		return nil, fmt.Errorf("can get task2's logcache id %s", realtaskid)
	}
	return logcache, nil
}*/

func api(){
	url := "www.baidu.com"
	method := "POST"

	payload := strings.NewReader(`{
    "key1": "v1"
}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("h1", "value1")
	req.Header.Add("h2", "value2")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "BAIDUID=865E9EA2BD4B781410C139B8CA1E3198:FG=1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}