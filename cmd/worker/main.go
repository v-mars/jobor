package main

import (
	"fmt"
	"github.com/spf13/cobra"
	Cmd "jobor/cmd"
	"jobor/internal/config"
	"jobor/internal/proto/service"
	"jobor/pkg/logger"
	"log"
	"os"
)

var (
	cfg string

	rootCmd = &cobra.Command{
		Use:   "",
		Short: "Start Run Jobor Worker",
		Long:  `This is mars jobor worker`,
		Example: `## 启动命令 ./app -p 20052 -c ./configs/worker.toml -f ./logs`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
			// 加载配置
			err := config.LoadWorkerConfig(cfg)
			if err != nil {
				log.Fatal(err)
			}
			//conf :=config.GetConf()
			logger.Initial()

			if err=service.WorkerGRPC();err!=nil{
				log.Fatal(err)
			}

		},
	}
)

func init() {
	var c = &config.WorkerConfig
	DefaultIP := "0.0.0.0"
	DefaultPort := int32(20052)
	DefaultMode := "release"		   // release, debug, test
	DefaultLevel := "info"
	DefaultLog := "./logs"
	rootCmd.Flags().StringVarP(&cfg, "conf", "c", "", "config file, example: ./configs/config.toml")
	rootCmd.Flags().StringVarP(&c.IP, "ip", "i", DefaultIP, "服务IP")
	rootCmd.Flags().Int32VarP(&c.Port, "port", "p", DefaultPort, "服务启动的端口: 20052")
	rootCmd.Flags().StringVarP(&c.Mode, "mode", "m", DefaultMode, "启动模式(release, debug, test e.g)")
	rootCmd.Flags().StringVarP(&c.LogPath, "log", "f", DefaultLog, "日志目录(/data/logs e.g)")
	rootCmd.Flags().StringVarP(&c.LogLevel, "level", "l", DefaultLevel, "日志级别(DEBUG, INFO, WARNING e.g)")
	if cfg == "" {
		cfg = "./configs/worker.toml"
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