package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func Client() *cobra.Command {
	var (
		cfg string
	)
	cmdServer := &cobra.Command{
		Use:   "client",
		Short: "Start Run jobor client",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
		},
	}
	cmdServer.Flags().StringVarP(&cfg, "conf", "c", "", "client config [toml]")
	return cmdServer
}

