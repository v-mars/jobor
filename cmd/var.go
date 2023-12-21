package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	Ver       = "v3.0.5"
	BuildDate = time.Now().Format("2006.01.02")
)

func Version() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of Jobor",
		Long:  `This is mars jobor`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version   : %s\n", Ver)
			fmt.Printf("Commit    : %s\n", "")
			fmt.Printf("BuildDate : %s\n", BuildDate)
		},
	}
	return versionCmd
}
