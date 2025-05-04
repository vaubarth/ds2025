package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(customerCmd)
	customerCmd.AddCommand(customerServicesCmd)
	customerCmd.AddCommand(customersNewCmd)
}

var customerCmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage customers.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Listing all known customers")
		t := table.NewWriter()
		tTemp := table.Table{}
		tTemp.Render()
		t.AppendHeader(table.Row{"ID", "Customer"})
		t.AppendRow(table.Row{1, "Customer A"})
		t.AppendRow(table.Row{1, "Customer B"})
		t.AppendRow(table.Row{1, "Customer C"})
		t.AppendRow(table.Row{1, "Customer D"})
		t.AppendRow(table.Row{1, "Customer E"})

		fmt.Println(t.Render())
	},
}

var customerServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage customers.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Listing all services for customer %s\n", args[0])
		t := table.NewWriter()
		tTemp := table.Table{}
		tTemp.Render()
		t.AppendHeader(table.Row{"ID", "Service", "Repository"})
		t.AppendRow(table.Row{1, "Service A", "github.com/some/repo_a"})
		t.AppendRow(table.Row{1, "Service B", "github.com/some/repo_a"})
		t.AppendRow(table.Row{1, "Service C", "github.com/some/repo_a"})

		fmt.Println(t.Render())
	},
}

var customersNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new service",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Adding new customer %s\n", args[0])
		pw := progress.NewWriter()
		pw.SetUpdateFrequency(time.Millisecond * 100)

		customerTracker := track("Creating customer data")
		namespaceTracker := track("Registering namespaces")

		pw.AppendTracker(customerTracker)
		pw.AppendTracker(namespaceTracker)

		go pw.Render()

		for i := 0; i <= 100; i += 10 {
			time.Sleep(100 * time.Millisecond)
			customerTracker.Increment(10)
			namespaceTracker.Increment(10)
		}

	},
}
