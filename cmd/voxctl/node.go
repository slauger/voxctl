package main

import (
	"fmt"

	"github.com/slauger/voxctl/internal/client"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage Puppet nodes",
}

var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all nodes",
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
		_, err = pdb.ListNodes()
		if err != nil {
			return err
		}
		return nil
	},
}

var nodeFactsCmd = &cobra.Command{
	Use:   "facts <certname>",
	Short: "Show facts for a node",
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
		_, err = pdb.GetNodeFacts(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

var nodeDeactivateCmd = &cobra.Command{
	Use:   "deactivate <certname>",
	Short: "Deactivate a node",
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
		if err := pdb.DeactivateNode(args[0]); err != nil {
			return err
		}
		fmt.Printf("Node %q deactivated.\n", args[0])
		return nil
	},
}

var nodePurgeCmd = &cobra.Command{
	Use:   "purge <certname>",
	Short: "Purge a deactivated node",
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
		if err := pdb.PurgeNode(args[0]); err != nil {
			return err
		}
		fmt.Printf("Node %q purged.\n", args[0])
		return nil
	},
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeFactsCmd)
	nodeCmd.AddCommand(nodeDeactivateCmd)
	nodeCmd.AddCommand(nodePurgeCmd)
	rootCmd.AddCommand(nodeCmd)
}
