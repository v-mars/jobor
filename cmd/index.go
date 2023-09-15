package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func Server() *cobra.Command {
	//var c = conf.GetConf()

	cmdServer := &cobra.Command{
		Use:     "server",
		Short:   "Start Run Jobor Server",
		Long:    `Welcome User Jobor Server`,
		Example: `## 启动命令 ./app server -p 5000 -c ./configs/config.toml -f ./logs`,
		Run: func(cmd *cobra.Command, args []string) {
			//if len(cfg) == 0 {
			//	_ = cmd.Help()
			//	os.Exit(0)
			//}
			//internal.Run(c.Server.ConfigFile)
			//addr := fmt.Sprintf("%s:%s", c.Server.IP, c.Server.Port)
			//router.InitRouter(c.Server.Mode, addr)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	//DefaultConfig := ""
	//DefaultIP := "0.0.0.0"
	//DefaultPort := "5000"
	//DefaultMode := "" // release, debug, test
	//DefaultLevel := ""
	//DefaultLog := "./logs"
	//cmdServer.Flags().StringVarP(&c.Server.ConfigFile, "config", "c", DefaultConfig, "config file, example: ./configs/config.toml")
	//cmdServer.Flags().StringVarP(&c.Server.IP, "ip", "i", DefaultIP, "服务IP")
	//cmdServer.Flags().StringVarP(&c.Server.Port, "port", "p", DefaultPort, "服务启动的端口: 5000")
	//cmdServer.Flags().StringVarP(&c.Server.Mode, "mode", "m", DefaultMode, "启动模式(release, debug, test e.g)")
	//cmdServer.Flags().StringVarP(&c.Server.LogPath, "log", "f", DefaultLog, "日志目录(/data/logs e.g)")
	//cmdServer.Flags().StringVarP(&c.Server.LogLevel, "level", "l", DefaultLevel, "日志级别(DEBUG, INFO, WARNING e.g)")
	//cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	//if c.Server.ConfigFile == "" {
	//	c.Server.ConfigFile = "./configs/config.toml"
	//	//fmt.Println("请使用\"-c\"指定配置文件")
	//	//os.Exit(-1)
	//}
	return cmdServer
}

func Worker() *cobra.Command {
	var (
		cfg string
	)
	//var c = conf.GetConf()
	cmdServer := &cobra.Command{
		Use:   "worker",
		Short: "Start Run jobor client",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
			// 加载配置
			//err := config.LoadWorkerConfig(cfg)
			//if err != nil {
			//	log.Fatal(err)
			//}
			////conf :=config.GetConf()
			//logger.Initial()
			//
			//if err = service.WorkerGRPC(); err != nil {
			//	log.Fatal(err)
			//}
		},
	}
	//DefaultIP := "0.0.0.0"
	//DefaultPort := int32(20052)
	//DefaultMode := "release" // release, debug, test
	//DefaultLevel := "info"
	//DefaultLog := "./logs"
	//cmdServer.Flags().StringVarP(&cfg, "config", "c", "", "config file, example: ./configs/config.toml")
	//cmdServer.Flags().StringVarP(&c.IP, "ip", "i", DefaultIP, "服务IP")
	//cmdServer.Flags().Int32VarP(&c.Port, "port", "p", DefaultPort, "服务启动的端口: 20052")
	//cmdServer.Flags().StringVarP(&c.Mode, "mode", "m", DefaultMode, "启动模式(release, debug, test e.g)")
	//cmdServer.Flags().StringVarP(&c.LogPath, "log", "f", DefaultLog, "日志目录(/data/logs e.g)")
	//cmdServer.Flags().StringVarP(&c.LogLevel, "level", "l", DefaultLevel, "日志级别(DEBUG, INFO, WARNING e.g)")
	////cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	//if cfg == "" {
	//	cfg = "./configs/worker.toml"
	//	//fmt.Println("请使用\"-c\"指定配置文件")
	//	//os.Exit(-1)
	//}
	return cmdServer
}
