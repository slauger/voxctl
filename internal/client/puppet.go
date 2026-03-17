package client

import (
	"fmt"
	"net/http"
)

// PuppetClient provides access to the Puppet Server API.
type PuppetClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewPuppetClient creates a new Puppet Server API client.
func NewPuppetClient(httpClient *http.Client, baseURL string) *PuppetClient {
	return &PuppetClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// ListEnvironments retrieves the list of environments from the Puppet Server.
func (c *PuppetClient) ListEnvironments() ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// ClearEnvironmentCache triggers an environment cache clear on the Puppet Server.
func (c *PuppetClient) ClearEnvironmentCache() error {
	return fmt.Errorf("not yet implemented")
}
