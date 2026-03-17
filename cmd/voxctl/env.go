package main

import (
	"fmt"

	"github.com/slauger/voxctl/internal/client"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage Puppet environments",
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
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
		puppet := client.NewPuppetClient(httpClient, rc.Server.Server)
		_, err = puppet.ListEnvironments()
		if err != nil {
			return err
		}
		return nil
	},
}

var envCacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage environment cache",
}

var envCacheClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the environment cache",
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
		puppet := client.NewPuppetClient(httpClient, rc.Server.Server)
		if err := puppet.ClearEnvironmentCache(); err != nil {
			return err
		}
		fmt.Println("Environment cache cleared.")
		return nil
	},
}

func init() {
	envCacheCmd.AddCommand(envCacheClearCmd)
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envCacheCmd)
	rootCmd.AddCommand(envCmd)
}
