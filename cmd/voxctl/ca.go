package main

import (
	"fmt"

	"github.com/slauger/voxctl/internal/client"
	"github.com/spf13/cobra"
)

var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Manage Puppet CA certificates",
}

var caListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all certificates",
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
		ca := client.NewCAClient(httpClient, rc.Server.Server)
		_, err = ca.ListCertificates()
		if err != nil {
			return err
		}
		return nil
	},
}

var caShowCmd = &cobra.Command{
	Use:   "show <certname>",
	Short: "Show certificate details",
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
		ca := client.NewCAClient(httpClient, rc.Server.Server)
		_, err = ca.GetCertificate(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

var caSignCmd = &cobra.Command{
	Use:   "sign <certname>",
	Short: "Sign a pending certificate request",
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
		ca := client.NewCAClient(httpClient, rc.Server.Server)
		if err := ca.SignCertificate(args[0]); err != nil {
			return err
		}
		fmt.Printf("Certificate %q signed.\n", args[0])
		return nil
	},
}

var caRevokeCmd = &cobra.Command{
	Use:   "revoke <certname>",
	Short: "Revoke a signed certificate",
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
		ca := client.NewCAClient(httpClient, rc.Server.Server)
		if err := ca.RevokeCertificate(args[0]); err != nil {
			return err
		}
		fmt.Printf("Certificate %q revoked.\n", args[0])
		return nil
	},
}

var caCleanCmd = &cobra.Command{
	Use:   "clean <certname>",
	Short: "Remove a certificate from the CA",
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
		ca := client.NewCAClient(httpClient, rc.Server.Server)
		if err := ca.CleanCertificate(args[0]); err != nil {
			return err
		}
		fmt.Printf("Certificate %q cleaned.\n", args[0])
		return nil
	},
}

func init() {
	caCmd.AddCommand(caListCmd)
	caCmd.AddCommand(caShowCmd)
	caCmd.AddCommand(caSignCmd)
	caCmd.AddCommand(caRevokeCmd)
	caCmd.AddCommand(caCleanCmd)
	rootCmd.AddCommand(caCmd)
}
