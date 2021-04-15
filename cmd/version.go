package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Version() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of Jobor",
		Long:  `This is Mars plan jobor cron`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version   : %s\n", "v1.0.0")
			fmt.Printf("Commit    : %s\n", "")
			fmt.Printf("BuildDate : %s\n", "2021.4.3")
		},
	}
	return versionCmd
}
