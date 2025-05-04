package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with DS",
	Long:  `Opens a browser window that allows you to authenticate with DS.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Pretending to authenticate with DS")

		pw := progress.NewWriter()
		pw.SetUpdateFrequency(time.Millisecond * 100)

		tracker := progress.Tracker{
			Total:   100,
			Message: "Authenticating",
		}
		pw.AppendTracker(&tracker)
		go pw.Render()

		for i := 0; i <= 100; i += 10 {
			time.Sleep(100 * time.Millisecond)
			tracker.Increment(10)
		}

		fmt.Println("\nYou are authenticated!")
	},
}
