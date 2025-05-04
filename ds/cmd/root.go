package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ds",
	Short: "ds is our tool to create and manage your services and projects.",
	Long:  `With ds you can create projects and services. You can manage projects, deploy and monitor your services and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		println(cmd.UsageString())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
