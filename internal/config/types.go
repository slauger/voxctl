package config

// Config represents the top-level voxctl configuration file,
// modeled after kubeconfig's structure.
type Config struct {
	APIVersion      string       `yaml:"apiVersion"`
	Kind            string       `yaml:"kind"`
	CurrentContext  string       `yaml:"current-context"`
	PreviousContext string       `yaml:"previous-context,omitempty"`
	Servers         []Server     `yaml:"servers"`
	Credentials     []Credential `yaml:"credentials"`
	Contexts        []Context    `yaml:"contexts"`
}

// Server defines a Puppet infrastructure endpoint.
type Server struct {
	Name     string `yaml:"name"`
	Server   string `yaml:"server"`
	PuppetDB string `yaml:"puppetdb,omitempty"`
	CACert   string `yaml:"ca-cert,omitempty"`
}

// Credential holds mTLS client certificate paths.
type Credential struct {
	Name       string `yaml:"name"`
	ClientCert string `yaml:"client-cert"`
	ClientKey  string `yaml:"client-key"`
}

// Context binds a server to a credential.
type Context struct {
	Name       string `yaml:"name"`
	Server     string `yaml:"server"`
	Credential string `yaml:"credential"`
}

// ResolvedContext contains fully resolved server and credential information.
type ResolvedContext struct {
	ContextName string
	Server      Server
	Credential  Credential
}
