package dal

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/alog"
	"jobor/biz/dal/bind"
	"jobor/biz/dal/casbin"
	"jobor/biz/dal/db"
	"jobor/biz/dal/migrate"
	"jobor/biz/dal/mysql"
	"jobor/biz/dal/redis"
	"jobor/biz/dal/redisStore"
	"jobor/biz/model"
	"jobor/biz/mw"
	"jobor/biz/pack/dispatcher"
	"jobor/biz/pack/oidc_callback"
	"jobor/biz/response"
	"jobor/conf"
)

var (
	H *server.Hertz

	// NoLogin 登录验证
	NoLogin = []string{
		"/login", "/404", "/static", "/favicon.ico", "/air", "/dashboard", "/hostid", "/license",
		"/jobor", "/iconfont", "/index.html", "/vite.svg", "/assets",
		"/favicon.ico", "/ping", "/swagger/", "/api/v1/swagger/*", "/debug/pprof",
		"/metrics", "/routes", "/reverse", "/air", "/api/v1/login", "/fs/", "/api/v1/login-init ",
		oidc_callback.CallbackPath,
		oidc_callback.GotoRedirect,
	}

	// NoAuthorized 权限验证
	NoAuthorized = []string{
		"/api/v1/mfa/", "/api/v1/jobor/enum", "/api/v1/jobor/task", "/api/v1/jobor/task-log", "/api/v1/jobor/dashboard",
		"/api/v1/sys/user-self", "/api/v1/user-info",
		"/api/v1/sys/user/profile", "/api/v1/sys/user/password",
		"/api/v1/logout", "/api/v1/sys/login-history",
		"/api/v1/sys/user-profile/", "/api/v1/otp"}
)

func init() {
	NoAuthorized = append(NoAuthorized, NoLogin...)
}

func Init() {
	//startTime := time.Now()
	//InitZeroLog(config.GetConf().Server.LogFileName, config.LogLevel())
	alog.InitHzLog(conf.GetConf().Server.LogFileName, conf.LogLevel())
	db.DefaultAesKey = conf.GetConf().Ext.DataAesKey
	response.DefaultAesKey = conf.GetConf().Ext.WebAesKey
	// Casbin连接初始化
	if _, err := casbin.NewCasbin(conf.GetConf()); err != nil {
		hlog.Fatal(err)
	}
	//hlog.Debug(time.Since(startTime).String())
	// 参数绑定报错自定义初始化
	bind.Init()
	// Redis连接初始化
	redis.Init(conf.GetConf())

	// MySQL连接初始化
	mysql.Init()

	// init redis store
	redisStore.Init()

	// DB连接初始化
	//migrate.Migrate()

	// DB 数据初始化
	install, err := migrate.QueryIsInstall(context.TODO(), db.DB)
	if err != nil {
		hlog.Fatal(err)
		return
	}
	if !install {
		hlog.Debugf("jobor db init is start ")
		err = migrate.StartInstall(context.TODO(), db.DB, "jobor3.sql")
		if err != nil {
			hlog.Fatal(err)
			return
		}
	}

	// JWT连接初始化
	mw.InitJwt()
	//hlog.Debug(time.Since(startTime).String())

	go func() {
		hlog.Fatal(redis.Subscribe(context.TODO(), dispatcher.Fn, model.PubSubChannel))
	}()

	if _, err := dispatcher.InitQSrv(&conf.GetConf().Redis, dispatcher.Queue); err != nil {
		hlog.Fatal(err)
		return
	}

	dispatcher.InitCron()

}
