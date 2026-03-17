package main

import (
	"github.com/slauger/voxctl/internal/client"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Manage Puppet reports",
}

var reportNodeFlag string

var reportListCmd = &cobra.Command{
	Use:   "list",
	Short: "List reports",
	RunE: func(cmd *cobra.Command, args []string) error {
		rc, err := resolveContext()
		if err != nil {
			return err
		}
		cc, _ := buildHTTPClient(rc)
		httpClient, err := client.NewHTTPClient(*cc)
		if err != nil {
			return err
		}
		pdb := client.NewPuppetDBClient(httpClient, rc.Server.PuppetDB)
		_, err = pdb.ListReports(reportNodeFlag)
		if err != nil {
			return err
		}
		return nil
	},
}

var reportShowCmd = &cobra.Command{
	Use:   "show <hash>",
	Short: "Show report details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rc, err := resolveContext()
		if err != nil {
			return err
		}
		cc, _ := buildHTTPClient(rc)
		httpClient, err := client.NewHTTPClient(*cc)
		if err != nil {
			return err
		}
		pdb := client.NewPuppetDBClient(httpClient, rc.Server.PuppetDB)
		_, err = pdb.GetReport(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	reportListCmd.Flags().StringVarP(&reportNodeFlag, "node", "n", "", "filter reports by node certname")
	reportCmd.AddCommand(reportListCmd)
	reportCmd.AddCommand(reportShowCmd)
	rootCmd.AddCommand(reportCmd)
}
