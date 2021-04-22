package cmd

import (
	"github.com/spf13/cobra"
	"jobor/internal"
	"jobor/internal/config"
	"jobor/internal/routers"
)

func Server() *cobra.Command {
	//fmt.Println("hello word")
	var (
		//cfg string
	)
	var c = &config.Configs

	cmdServer := &cobra.Command{
		Use:   "server",
		Short: "Start Run Jobor Server",
		Long:  `Welcome User Jobor Server`,
		Example: `## 启动命令 ./app server -p 5000 -c ./configs/config.toml -f ./logs`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.Run(c.Server.ConfigFile)
			addr := c.Server.IP+":"+c.Server.Port
			routers.InitRouter(c.Server.Mode, addr)
			//if len(cfg) == 0 {
			//	_ = cmd.Help()
			//	os.Exit(0)
			//}
			//config.Init(cfg)
			//mylog.Init()
			//alarm.InitAlarm()
			//err := model.InitDb()
			//if err != nil {
			//	log.Fatal("InitDb failed", zap.Error(err))
			//}
			//model.InitRabc()
			//go version.CheckLatest() // check new version
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			//lis, err := router.GetListen(define.Server)
			//if err != nil {
			//	log.Fatal("listen failed", zap.Error(err))
			//}
			//// init alarm
			//err = schedule.Init2()
			//if err != nil {
			//	log.Fatal("init schedule failed", zap.Error(err))
			//}
			//
			//err = router.Run(define.Server, lis)
			//if err != nil {
			//	log.Error("router.Run error", zap.Error(err))
			//}
			return nil
		},
	}
	DefaultConfig := ""
	DefaultIP := "0.0.0.0"
	DefaultPort := "5000"
	DefaultMode := ""		   // release, debug, test
	DefaultLevel := ""
	DefaultLog := "./logs"
	cmdServer.Flags().StringVarP(&c.Server.ConfigFile,"config","c",DefaultConfig,"config file, example: ./configs/config.toml")
	cmdServer.Flags().StringVarP(&c.Server.IP, "ip", "i", DefaultIP, "服务IP")
	cmdServer.Flags().StringVarP(&c.Server.Port, "port", "p", DefaultPort, "服务启动的端口: 5000")
	cmdServer.Flags().StringVarP(&c.Server.Mode, "mode", "m", DefaultMode, "启动模式(release, debug, test e.g)")
	cmdServer.Flags().StringVarP(&c.Server.LogPath, "log", "f", DefaultLog, "日志目录(/data/logs e.g)")
	cmdServer.Flags().StringVarP(&c.Server.LogLevel, "level", "l", DefaultLevel, "日志级别(DEBUG, INFO, WARNING e.g)")
	//cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	if c.Server.ConfigFile == "" {
		c.Server.ConfigFile = "./configs/config.toml"
		//fmt.Println("请使用\"-c\"指定配置文件")
		//os.Exit(-1)
	}
	return cmdServer

	//return nil
}
