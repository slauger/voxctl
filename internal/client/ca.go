package client

import (
	"fmt"
	"net/http"
)

// CAClient provides access to the Puppet CA API.
type CAClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewCAClient creates a new Puppet CA API client.
func NewCAClient(httpClient *http.Client, baseURL string) *CAClient {
	return &CAClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// ListCertificates retrieves all certificate statuses from the CA.
func (c *CAClient) ListCertificates() ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// GetCertificate retrieves a single certificate status by certname.
func (c *CAClient) GetCertificate(certname string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// SignCertificate signs a pending certificate request.
func (c *CAClient) SignCertificate(certname string) error {
	return fmt.Errorf("not yet implemented")
}

// RevokeCertificate revokes a signed certificate.
func (c *CAClient) RevokeCertificate(certname string) error {
	return fmt.Errorf("not yet implemented")
}

// CleanCertificate removes a certificate from the CA.
func (c *CAClient) CleanCertificate(certname string) error {
	return fmt.Errorf("not yet implemented")
}
