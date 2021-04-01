package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Version() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Jobor",
		Long:  `This is Mars plan jobor cron`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("jobor version is v1.0")
		},
	}
	return versionCmd
}
