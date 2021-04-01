package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Server() *cobra.Command {
	fmt.Println("hello word")
	var (
		cfg string
	)
	cmdServer := &cobra.Command{
		Use:   "server",
		Short: "Start Run jobor server",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
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
	cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	return cmdServer

	//return nil
}
