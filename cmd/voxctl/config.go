package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/slauger/voxctl/internal/config"
	"github.com/slauger/voxctl/internal/output"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage voxctl configuration",
}

var getContextsCmd = &cobra.Command{
	Use:   "get-contexts",
	Short: "List all configured contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfg == nil {
			return fmt.Errorf("no config loaded")
		}

		columns := []string{"", "NAME", "SERVER", "CREDENTIAL"}
		var rows [][]string
		for _, ctx := range cfg.Contexts {
			marker := ""
			if ctx.Name == cfg.CurrentContext {
				marker = "*"
			}
			rows = append(rows, []string{marker, ctx.Name, ctx.Server, ctx.Credential})
		}

		return output.Print(outputFmt, rows, columns)
	},
}

var currentContextCmd = &cobra.Command{
	Use:   "current-context",
	Short: "Display the current context",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfg == nil || cfg.CurrentContext == "" {
			return fmt.Errorf("no current context set")
		}
		fmt.Println(cfg.CurrentContext)
		return nil
	},
}

var useContextCmd = &cobra.Command{
	Use:   "use-context [name]",
	Short: "Switch the current context",
	Long:  "Switch the current context. Use '-' to switch to the previous context. If no argument is given, an interactive picker (fzf) is launched.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfg == nil {
			return fmt.Errorf("no config loaded")
		}

		var name string
		switch {
		case len(args) == 1 && args[0] == "-":
			if cfg.PreviousContext == "" {
				return fmt.Errorf("no previous context set")
			}
			name = cfg.PreviousContext
		case len(args) == 1:
			name = args[0]
		default:
			picked, err := pickContext(cfg)
			if err != nil {
				return err
			}
			name = picked
		}

		if err := config.SetCurrentContext(cfg, configPath(), name); err != nil {
			return err
		}

		fmt.Printf("Switched to context %q.\n", name)
		return nil
	},
}

// pickContext launches fzf to interactively select a context.
func pickContext(cfg *config.Config) (string, error) {
	if len(cfg.Contexts) == 0 {
		return "", fmt.Errorf("no contexts configured")
	}

	var names []string
	for _, ctx := range cfg.Contexts {
		names = append(names, ctx.Name)
	}

	fzfPath, err := exec.LookPath("fzf")
	if err != nil {
		return "", fmt.Errorf("fzf not found in PATH; install fzf or specify a context name")
	}

	fzf := exec.Command(fzfPath, "--prompt", "Select context: ")
	fzf.Stdin = strings.NewReader(strings.Join(names, "\n"))
	fzf.Stderr = os.Stderr

	out, err := fzf.Output()
	if err != nil {
		return "", fmt.Errorf("fzf selection cancelled")
	}

	return strings.TrimSpace(string(out)), nil
}

func init() {
	configCmd.AddCommand(getContextsCmd)
	configCmd.AddCommand(currentContextCmd)
	configCmd.AddCommand(useContextCmd)
	rootCmd.AddCommand(configCmd)
}
