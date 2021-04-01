package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"jobor/cmd"
)

func main(){
	//version.Commit = c
	//version.Version = v
	//version.BuildDate = d
	rootCmd := &cobra.Command{Use: "jobor"}
	rootCmd.AddCommand(cmd.Client())
	rootCmd.AddCommand(cmd.Server())
	rootCmd.AddCommand(cmd.Version())
	//rootCmd.AddCommand(cmd.GeneratePemKey())
	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("rootCmd.Execute failed %s", err.Error())
	}
}