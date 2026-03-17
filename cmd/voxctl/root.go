package main

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/slauger/voxctl/internal/client"
	"github.com/slauger/voxctl/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile    string
	ctxName    string
	outputFmt  string
	cfg        *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "voxctl",
	Short: "CLI for managing OpenVox/Puppet Server infrastructure",
	Long:  "voxctl is a command-line tool for managing OpenVox/Puppet Server infrastructure via REST APIs.",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip config loading for commands that don't need it.
		if cmd.Name() == "version" || cmd.Name() == "help" {
			return nil
		}

		path := cfgFile
		if path == "" {
			path = config.DefaultConfigPath()
		}

		var err error
		cfg, err = config.Load(path)
		if err != nil {
			// Config is optional for some commands.
			if errors.Is(err, fs.ErrNotExist) && cmd.Parent() != nil && cmd.Parent().Name() == "config" {
				cfg = &config.Config{
					APIVersion: "v1",
					Kind:       "Config",
				}
				return nil
			}
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.voxctl/config)")
	rootCmd.PersistentFlags().StringVar(&ctxName, "context", "", "context to use (overrides current-context)")
	rootCmd.PersistentFlags().StringVarP(&outputFmt, "output", "o", "table", "output format (table|json|yaml)")
}

// resolveContext returns the resolved context, respecting the --context flag override.
func resolveContext() (*config.ResolvedContext, error) {
	if cfg == nil {
		return nil, fmt.Errorf("no config loaded; create one at %s", config.DefaultConfigPath())
	}

	name := ctxName
	if name == "" {
		name = cfg.CurrentContext
	}
	if name == "" {
		return nil, fmt.Errorf("no context specified and no current-context set in config")
	}

	return config.ResolveContext(cfg, name)
}

// buildHTTPClient creates an mTLS HTTP client from the resolved context.
func buildHTTPClient(rc *config.ResolvedContext) (*client.ClientConfig, error) {
	return &client.ClientConfig{
		CACert:     rc.Server.CACert,
		ClientCert: rc.Credential.ClientCert,
		ClientKey:  rc.Credential.ClientKey,
	}, nil
}

// configPath returns the effective config file path.
func configPath() string {
	if cfgFile != "" {
		return cfgFile
	}
	return config.DefaultConfigPath()
}
