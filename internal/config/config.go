package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DefaultConfigPath returns the default config file path (~/.voxctl/config).
func DefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", ".voxctl", "config")
	}
	return filepath.Join(home, ".voxctl", "config")
}

// Load reads and parses a voxctl config file.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	return &cfg, nil
}

// Save writes the config to disk with secure permissions.
func Save(path string, cfg *Config) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("writing config %s: %w", path, err)
	}

	return nil
}

// GetContext returns the named context from the config.
func GetContext(cfg *Config, name string) (*Context, error) {
	for i := range cfg.Contexts {
		if cfg.Contexts[i].Name == name {
			return &cfg.Contexts[i], nil
		}
	}
	return nil, fmt.Errorf("context %q not found", name)
}

// ResolveContext looks up the named context and resolves its server and credential references.
func ResolveContext(cfg *Config, name string) (*ResolvedContext, error) {
	ctx, err := GetContext(cfg, name)
	if err != nil {
		return nil, err
	}

	var server *Server
	for i := range cfg.Servers {
		if cfg.Servers[i].Name == ctx.Server {
			server = &cfg.Servers[i]
			break
		}
	}
	if server == nil {
		return nil, fmt.Errorf("server %q referenced by context %q not found", ctx.Server, name)
	}

	var cred *Credential
	for i := range cfg.Credentials {
		if cfg.Credentials[i].Name == ctx.Credential {
			cred = &cfg.Credentials[i]
			break
		}
	}
	if cred == nil {
		return nil, fmt.Errorf("credential %q referenced by context %q not found", ctx.Credential, name)
	}

	return &ResolvedContext{
		ContextName: name,
		Server:      *server,
		Credential:  *cred,
	}, nil
}

// SetCurrentContext updates the current context and saves the config.
func SetCurrentContext(cfg *Config, path, name string) error {
	// Verify the context exists.
	if _, err := GetContext(cfg, name); err != nil {
		return err
	}

	cfg.PreviousContext = cfg.CurrentContext
	cfg.CurrentContext = name

	return Save(path, cfg)
}
