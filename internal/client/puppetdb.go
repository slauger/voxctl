package client

import (
	"fmt"
	"net/http"
)

// PuppetDBClient provides access to the PuppetDB API.
type PuppetDBClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewPuppetDBClient creates a new PuppetDB API client.
func NewPuppetDBClient(httpClient *http.Client, baseURL string) *PuppetDBClient {
	return &PuppetDBClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// ListNodes queries PuppetDB for all nodes.
func (c *PuppetDBClient) ListNodes() ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// GetNodeFacts retrieves facts for a specific node.
func (c *PuppetDBClient) GetNodeFacts(certname string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// DeactivateNode deactivates a node in PuppetDB.
func (c *PuppetDBClient) DeactivateNode(certname string) error {
	return fmt.Errorf("not yet implemented")
}

// PurgeNode purges a deactivated node from PuppetDB.
func (c *PuppetDBClient) PurgeNode(certname string) error {
	return fmt.Errorf("not yet implemented")
}

// ListReports retrieves reports, optionally filtered by certname.
func (c *PuppetDBClient) ListReports(certname string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// GetReport retrieves a specific report by hash.
func (c *PuppetDBClient) GetReport(hash string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
