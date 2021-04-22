package routers

import (
	"fmt"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"jobor/internal/app/sys/permission"
	"jobor/internal/middleware"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/pkg/logger"
	"jobor/pkg/utils"
	"log"
	"net/http"
	"os"
	"strings"
)


var (
	Engine *gin.Engine
)

func InitRouter(RunMode string, addr string)  {
	//RunMode := gin.DebugMode
	gin.SetMode(RunMode) //调试模式
	Engine = gin.New()

	// Prom
	Engine.Use(ginprom.PromMiddleware(&ginprom.PromOpts{}))

	if strings.ToLower(RunMode) == "debug" {
		Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}


	gin.DefaultWriter = io.MultiWriter(logger.Gin.Writer()) // os.Stdout, logger.Gin.Writer()

	//错误日志审计
	Engine.Use(middleware.ErrorLogHandler())

	// 不存在方法
	Engine.NoMethod(middleware.NoMethodHandler())
	// 不存在路由
	Engine.NoRoute(middleware.NoRouteHandler())

	// 崩溃恢复
	//Engine.Use(gin.Recovery())
	Engine.Use(middleware.RecoveryMiddleware())
	Engine.Use(middleware.CORSMiddleware())
	// 日志
	Engine.Use(gin.Logger())


	// session
	//store, _ := redis.NewStore(10, "tcp", "192.168.21.138:6379", "", []byte("secret"))
	//store, _ := redis.NewStoreWithDB(200, "tcp", "192.168.21.138:6379", "", "1", []byte(""))
	//store.Options(sessions.Options{MaxAge:60, Secure:false})
	//_ = redis.SetKeyPrefix(store, "")

	//Engine.Use(sessions.Sessions("sessionid", store))

	//var blockArr = []string{"/api", "/v1"}
	// 登录验证 及信息提取
	var notCheckLoginUrlArr = []string{
		"/static","/favicon.ico","/ping","/swager/*","/debug/pprof","/metrics","/api/code",
		"/gin/routes",
		"/api/v1/login","/api/v1/token","/api/v1/refresh-token","/api/login","/api/user/logout",
		}
	Engine.Use(middleware.UserAuthMiddleware(middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...)))

	// 权限验证
	var notCheckPermissionUrlArr = []string{"/api/v1/user-info", "/api/menu/menubuttonlist"}
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	Engine.Use(middleware.CasbinMiddleware(middleware.AllowPathPrefixSkipper(notCheckPermissionUrlArr...)))

	Engine.LoadHTMLGlob("./static/dist/*.html")
	Engine.Static("/static", "./static/dist/static")
	Engine.Static("/resource", "./static/resource")
	Engine.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	Engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 注册路由
	if err := RegisterRouter(Engine); err != nil{
		log.Fatal(err)
	}
	// Prometheus
	prometheusMonitor(Engine)
	//profile(Engine)

	//pprof.Register(Engine)
	//_ = Engine.RunV1(addr)
	go func() {
		log.Fatal(Engine.Run(addr))
	}()
	msg := fmt.Sprintf("服务启动成功，运行模式：%s，版本号：%s，进程号：%d", RunMode, "release", os.Getpid())
	fmt.Println(utils.Green(msg))
	fmt.Println(utils.Green("访问地址 http://"+addr))
	//log.Printf(utils.Green(msg, 1))
	//InitUpdatePermissionByGinRoutes()
	fmt.Println(utils.Green("[*] Waiting for messages. To exit press CTRL+C"))
	select {}
}


func InitUpdatePermissionByGinRoutes()  {
	perms := permission.NewService(db.DB)
	var res []tbs.Permission
	routes := Engine.Routes()
	for _,v := range routes{
		var name string
		pathArray:=strings.Split(strings.Trim(v.Path, "/"), "/")
		if len(pathArray)>=4{
			name = fmt.Sprintf("%s:%s:%s", pathArray[2],pathArray[3], strings.ToLower(v.Method))
		} else if len(pathArray)==3{
			name = fmt.Sprintf("%s:%s", pathArray[2], strings.ToLower(v.Method))
		}else if len(pathArray)==2{
			name = fmt.Sprintf("%s:%s", pathArray[1], strings.ToLower(v.Method))
		} else if len(pathArray)==1{
			name = fmt.Sprintf("%s:%s", pathArray[0], strings.ToLower(v.Method))
		}
		res = append(res, tbs.Permission{Name: name,Method: v.Method, Path: v.Path})
	}
	err := perms.UpdatePermission(res)
	if err!=nil{
		logger.Errorf("InitUpdatePermissionByGinRoutes: %s", err)
	}
}