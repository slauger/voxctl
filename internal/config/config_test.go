package config

import (
	"os"
	"path/filepath"
	"testing"
)

func testConfig() *Config {
	return &Config{
		APIVersion:     "v1",
		Kind:           "Config",
		CurrentContext: "prod",
		Servers: []Server{
			{Name: "prod-server", Server: "https://puppet:8140", PuppetDB: "https://puppetdb:8081", CACert: "/tmp/ca.pem"},
			{Name: "staging-server", Server: "https://puppet-staging:8140"},
		},
		Credentials: []Credential{
			{Name: "admin", ClientCert: "/tmp/cert.pem", ClientKey: "/tmp/key.pem"},
		},
		Contexts: []Context{
			{Name: "prod", Server: "prod-server", Credential: "admin"},
			{Name: "staging", Server: "staging-server", Credential: "admin"},
		},
	}
}

func TestSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config")

	cfg := testConfig()
	if err := Save(path, cfg); err != nil {
		t.Fatalf("Save: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if perm := info.Mode().Perm(); perm != 0600 {
		t.Errorf("expected file mode 0600, got %04o", perm)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if loaded.APIVersion != "v1" {
		t.Errorf("expected apiVersion v1, got %q", loaded.APIVersion)
	}
	if loaded.CurrentContext != "prod" {
		t.Errorf("expected current-context prod, got %q", loaded.CurrentContext)
	}
	if len(loaded.Servers) != 2 {
		t.Errorf("expected 2 servers, got %d", len(loaded.Servers))
	}
	if len(loaded.Contexts) != 2 {
		t.Errorf("expected 2 contexts, got %d", len(loaded.Contexts))
	}
}

func TestLoadNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

func TestGetContext(t *testing.T) {
	cfg := testConfig()

	ctx, err := GetContext(cfg, "prod")
	if err != nil {
		t.Fatalf("GetContext: %v", err)
	}
	if ctx.Server != "prod-server" {
		t.Errorf("expected server prod-server, got %q", ctx.Server)
	}

	_, err = GetContext(cfg, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent context")
	}
}

func TestResolveContext(t *testing.T) {
	cfg := testConfig()

	rc, err := ResolveContext(cfg, "prod")
	if err != nil {
		t.Fatalf("ResolveContext: %v", err)
	}
	if rc.Server.Server != "https://puppet:8140" {
		t.Errorf("expected server URL https://puppet:8140, got %q", rc.Server.Server)
	}
	if rc.Credential.ClientCert != "/tmp/cert.pem" {
		t.Errorf("expected client cert /tmp/cert.pem, got %q", rc.Credential.ClientCert)
	}
	if rc.Server.PuppetDB != "https://puppetdb:8081" {
		t.Errorf("expected puppetdb URL https://puppetdb:8081, got %q", rc.Server.PuppetDB)
	}
}

func TestResolveContextMissingServer(t *testing.T) {
	cfg := &Config{
		Contexts:    []Context{{Name: "bad", Server: "missing", Credential: "admin"}},
		Credentials: []Credential{{Name: "admin"}},
	}
	_, err := ResolveContext(cfg, "bad")
	if err == nil {
		t.Fatal("expected error for missing server")
	}
}

func TestResolveContextMissingCredential(t *testing.T) {
	cfg := &Config{
		Servers:  []Server{{Name: "s1"}},
		Contexts: []Context{{Name: "bad", Server: "s1", Credential: "missing"}},
	}
	_, err := ResolveContext(cfg, "bad")
	if err == nil {
		t.Fatal("expected error for missing credential")
	}
}

func TestSetCurrentContext(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config")

	cfg := testConfig()
	if err := Save(path, cfg); err != nil {
		t.Fatalf("Save: %v", err)
	}

	if err := SetCurrentContext(cfg, path, "staging"); err != nil {
		t.Fatalf("SetCurrentContext: %v", err)
	}

	if cfg.CurrentContext != "staging" {
		t.Errorf("expected current-context staging, got %q", cfg.CurrentContext)
	}
	if cfg.PreviousContext != "prod" {
		t.Errorf("expected previous-context prod, got %q", cfg.PreviousContext)
	}

	// Verify persisted.
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.CurrentContext != "staging" {
		t.Errorf("expected persisted current-context staging, got %q", loaded.CurrentContext)
	}
}

func TestSetCurrentContextNotFound(t *testing.T) {
	cfg := testConfig()
	err := SetCurrentContext(cfg, "/dev/null", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent context")
	}
}
