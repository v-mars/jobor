package main

import (
	"fmt"
	"github.com/spf13/cobra"
	Cmd "jobor/cmd"
	"jobor/internal"
	"jobor/internal/config"
	"jobor/internal/router"
	"os"
)

var (
	c = &config.Configs
	//cfg string
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Start Run Jobor Server",
		Long:  `Welcome User Jobor Server`,
		Example: `## 启动命令 ./app server -p 5000 -c ./configs/config.toml -f ./logs`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.Run(c.Server.ConfigFile)
			addr := fmt.Sprintf("%s:%s",c.Server.IP,c.Server.Port)
			router.InitRouter(c.Server.Mode, addr)
			//if len(cfg) == 0 {
			//	_ = cmd.Help()
			//	os.Exit(0)
			//}
			//config.Init(cfg)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
)

func init() {
	DefaultConfig := ""
	DefaultIP := "0.0.0.0"
	DefaultPort := "5000"
	DefaultMode := ""		   // release, debug, test
	DefaultLevel := ""
	DefaultLog := "./logs"
	rootCmd.Flags().StringVarP(&c.Server.ConfigFile,"config","c",DefaultConfig,"config file, example: ./configs/config.toml")
	rootCmd.Flags().StringVarP(&c.Server.IP, "ip", "i", DefaultIP, "服务IP")
	rootCmd.Flags().StringVarP(&c.Server.Port, "port", "p", DefaultPort, "服务启动的端口: 5000")
	rootCmd.Flags().StringVarP(&c.Server.Mode, "mode", "m", DefaultMode, "启动模式(release, debug, test e.g)")
	rootCmd.Flags().StringVarP(&c.Server.LogPath, "log", "f", DefaultLog, "日志目录(/data/logs e.g)")
	rootCmd.Flags().StringVarP(&c.Server.LogLevel, "level", "l", DefaultLevel, "日志级别(DEBUG, INFO, WARNING e.g)")
	//cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	if c.Server.ConfigFile == "" {
		c.Server.ConfigFile = "./config.toml"
		//fmt.Println("请使用\"-c\"指定配置文件")
		//os.Exit(-1)
	}
	return
}


func main(){
	rootCmd.AddCommand(Cmd.Version())
	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("rootCmd.Execute failed %s", err.Error())
		os.Exit(-1)
	}
}
