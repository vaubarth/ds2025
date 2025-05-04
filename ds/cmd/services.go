package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"time"
)

var database string
var customer string
var stack string

func init() {
	rootCmd.AddCommand(servicesCmd)

	servicesCmd.AddCommand(servicesNewCmd)
	servicesNewCmd.Flags().StringVarP(&customer, "customer", "c", "", "Customer the service belongs to")
	servicesNewCmd.Flags().StringVarP(&database, "database", "d", "", "Type of database")
	servicesNewCmd.Flags().StringVarP(&database, "replicas", "r", "", "Replicas")
	servicesNewCmd.Flags().StringVarP(&database, "dns", "w", "", "DNS required?")
	servicesNewCmd.Flags().StringVarP(&stack, "stack", "s", "", "Language stack")
	servicesNewCmd.MarkFlagRequired("customer")

	servicesCmd.AddCommand(servicesDetailCmd)

}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Deal with services",
	Long:  `Create, modify, manage services in DS`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Listing all known services")
		t := table.NewWriter()
		tTemp := table.Table{}
		tTemp.Render()
		t.AppendHeader(table.Row{"ID", "Service", "Customer", "Repository"})
		t.AppendRow(table.Row{1, "Service A", "Customer A", "github.com/some/repo_a"})
		t.AppendRow(table.Row{1, "Service B", "Customer A", "github.com/some/repo_b"})
		t.AppendRow(table.Row{1, "Service C", "Customer A", "github.com/some/repo_c"})
		t.AppendRow(table.Row{1, "Service Y", "Customer B", "github.com/some/repo_d"})
		t.AppendRow(table.Row{1, "Service Z", "Customer C", "github.com/some/repo_e"})

		fmt.Println(t.Render())
	},
}

func track(message string) *progress.Tracker {
	return &progress.Tracker{
		Total:   100,
		Message: message,
	}
}

var servicesNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new service",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Adding new service %s\n", args[0])
		pw := progress.NewWriter()
		pw.SetUpdateFrequency(time.Millisecond * 100)

		repoTracker := track("Creating repo")
		manifestsTracker := track("Writing manifests")
		dbTracker := track("Setting up database")
		o11yTracker := track("Setting up o11y")
		deploymentsTracker := track("Setting up deployments")
		cdTracker := track("Creating CI/CD configurations (GH Actions)")
		backstageTracker := track("Configuring Backstage")

		pw.AppendTracker(repoTracker)
		pw.AppendTracker(dbTracker)
		pw.AppendTracker(o11yTracker)
		pw.AppendTracker(manifestsTracker)
		pw.AppendTracker(deploymentsTracker)
		pw.AppendTracker(backstageTracker)
		pw.AppendTracker(cdTracker)

		go pw.Render()

		for i := 0; i <= 100; i += 10 {
			time.Sleep(100 * time.Millisecond)
			repoTracker.Increment(10)
			dbTracker.Increment(10)
			o11yTracker.Increment(10)
			manifestsTracker.Increment(10)
			backstageTracker.Increment(10)
			deploymentsTracker.Increment(10)
			cdTracker.Increment(10)
		}

	},
}

var servicesDetailCmd = &cobra.Command{
	Use:   "details",
	Short: "Get details of a given service",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Details for service %s\n\n", args[0])

		fmt.Println("Repository: github.com/some/repo")
		fmt.Println("Dashboard:  dash.ds.dev/a17b")
		fmt.Println("Pipelines:  argo.ds.dev/bjk7")
		fmt.Println("")
		fmt.Println("Replicas:   4")
		fmt.Println("Endpoints:  some.prod.com | some.dev.com")

	},
}
