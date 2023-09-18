package srv_http

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"jobor/biz/dal"
	"jobor/biz/dal/redisStore"
	"jobor/biz/mw"
	"jobor/cmd/srv_rpc"
	"jobor/conf"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	hc "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/requestid"
	"github.com/hertz-contrib/sessions"
)

var (
	//cfg string

	RootCmd = &cobra.Command{
		Use:     "",
		Short:   "Start Run Jobor Server",
		Long:    `This is dolphin jobor server`,
		Example: `## 启动命令 ./app -c ./conf/config.yaml`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(conf.FlagConf) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
			// 加载配置
			conf.GetConf()

			//start grpc server service
			go func() {
				hlog.Fatal(srv_rpc.StartSrvRpc())
			}()
			// start http server service
			Start()
		},
	}
)

func init() {
	//var c = &conf.WorkerConfig{}
	//DefaultIP := "0.0.0.0"
	//DefaultPort := int32(20052)
	RootCmd.Flags().StringVarP(&conf.FlagConf, "conf", "c", "", "config file, example: ./conf/config.yaml")
	//rootCmd.Flags().StringVarP(&c.IP, "ip", "i", DefaultIP, "服务IP")
	//rootCmd.Flags().Int32VarP(&c.Port, "port", "p", DefaultPort, "服务启动的端口: 20052 e.g")
	if conf.FlagConf == "" {
		conf.FlagConf = "./conf/config.yaml"
		//fmt.Println("请使用\"-c\"指定配置文件")
		//os.Exit(-1)
	}
	return
}

const sessionName = "dbs_session"

func Start() {

	hlog.Infof("start http server service")
	//flag.Parse()
	startTime := time.Now()
	tracer, cfg := hertztracing.NewServerTracer()
	var options = []hc.Option{
		server.WithHostPorts(conf.GetConf().Server.HttpAddress),
		server.WithMaxRequestBodySize(20 << 20),
		server.WithTransport(standard.NewTransporter),
		tracer,
		//server.WithTracer(prometheus.NewServerTracer(":9091", "/metric")),
	}
	//discover.RegistryWeb(&options)
	//discover.RegistryRPC(&options)
	h := server.Default(options...)
	dal.Init()
	if conf.GetConf().Authentication.EnableSession {
		h.Use(sessions.New(fmt.Sprintf("%s_session", conf.Dom), redisStore.Store))
	}

	h.Use(hertztracing.ServerMiddleware(cfg))

	dal.H = h

	h.Use(mw.AccessLog())
	h.Use(mw.AuditLogHandler())
	hlog.Debugf("init time：%v", time.Since(startTime).String())

	if len(conf.GetConf().Ext.TelemetryEp) > 0 {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(conf.AppName),
			// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
			provider.WithExportEndpoint(conf.GetConf().Ext.TelemetryEp),
			provider.WithInsecure(),
		)
		defer p.Shutdown(context.Background())
	}

	//gin.DefaultWriter = io.MultiWriter(logg.Gin.Writer())
	h.Use(requestid.New())
	// 跨域请求
	h.Use(mw.Cors())
	// 解密加密请求
	h.Use(mw.ReqAesDec())
	//h.Use(gzip.Gzip(gzip.DefaultCompression))

	// 捕获异常，并返回500
	h.Use(mw.Recovery())

	//许可过滤
	//h.Use(mw.LicMw(mw.AllowPathPrefixSkipper(dal.NoAuthorized...)))
	// 认证过滤
	h.Use(mw.UserAuthMw(mw.AllowPathPrefixSkipper(dal.NoLogin...)))
	// 权限过滤
	h.Use(mw.CasbinMw(mw.AllowPathPrefixSkipper(dal.NoAuthorized...)))

	register(h)
	print(banner)
	hlog.Infof("startup time：%v", time.Since(startTime).String())
	h.Spin()
}

var banner = ``
