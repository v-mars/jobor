package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"jobor/cmd"
	"os"
)

func main(){
	//version.Commit = c
	//version.Version = v
	//version.BuildDate = d
	rootCmd := &cobra.Command{Use: ""}
	//rootCmd := &cobra.Command{Use: "jobor"}
	rootCmd.AddCommand(cmd.Server())
	rootCmd.AddCommand(cmd.Client())
	rootCmd.AddCommand(cmd.Version())
	//rootCmd.AddCommand(cmd.GeneratePemKey())
	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("rootCmd.Execute failed %s", err.Error())
		os.Exit(-1)
	}
}