# 🦊 voxctl

[![Go](https://github.com/slauger/voxctl/actions/workflows/go.yaml/badge.svg)](https://github.com/slauger/voxctl/actions/workflows/go.yaml)
[![Go Version](https://img.shields.io/github/go-mod-go-version/slauger/voxctl)](go.mod)
[![License](https://img.shields.io/github/license/slauger/voxctl)](LICENSE)

> **🚧 Status: Early Development** — The CLI skeleton is in place: the kubeconfig-style configuration, context switching (`voxctl config`), mTLS client setup and output formatting work. The API commands (`ca`, `env`, `node`, `report`) are defined but not implemented yet and return `not yet implemented`. Feedback, suggestions, and ideas are very welcome — feel free to open an [issue](https://github.com/slauger/voxctl/issues).

CLI tool for managing OpenVox/Puppet Server infrastructure via REST APIs.

voxctl provides a convenient command-line interface for common Puppet infrastructure operations that would otherwise require manual `curl` calls or `kubectl exec` into server pods:

- 🔐 **Certificate Lifecycle** — List, sign, revoke, clean, and inspect certificates via the Puppet CA API
- 🌍 **Environment Management** — List environments and flush the environment cache via the Puppet Server API
- 🖥️ **Node Management** — List, deactivate, and purge nodes, retrieve facts via PuppetDB
- 📊 **Report Viewing** — List and inspect Puppet run reports from PuppetDB
- ⚙️ **Kubeconfig-style Configuration** — Multi-server context switching with named servers, credentials, and contexts
- 🔒 **mTLS Authentication** — Native mutual TLS support for secure API communication
- 📋 **Flexible Output** — Table, JSON, and YAML output formats for scripting and automation
- 🔀 **Interactive Context Switching** — fzf-powered interactive picker and `use-context -` for quick switching (requires [fzf](https://github.com/junegunn/fzf) for the picker)

## Quick Start

```bash
# Install with Go
go install github.com/slauger/voxctl/cmd/voxctl@latest

# Show the configured contexts and switch between them
voxctl config get-contexts
voxctl config use-context production

# List all certificates (planned)
voxctl ca list

# List all nodes from PuppetDB (planned)
voxctl node list
```

## Installation

Once a release is published, prebuilt archives for Linux, macOS and Windows (amd64 and arm64) are attached to each [GitHub release](https://github.com/slauger/voxctl/releases), together with `checksums.txt`.

Build from source (requires Go, see `go.mod`):

```bash
git clone https://github.com/slauger/voxctl.git
cd voxctl
make build    # binary in bin/voxctl
```

## Commands

```
voxctl ca          Certificate lifecycle management (list, sign, revoke, clean, show)
voxctl env         Environment management (list, cache clear)
voxctl node        Node management (list, deactivate, purge, facts)
voxctl report      Report viewing (list, show)
voxctl config      Context and configuration management (current-context, get-contexts, use-context)
voxctl version     Print version information
```

Global flags: `--config` (config file, default `~/.voxctl/config`), `--context` (overrides `current-context`), `-o, --output` (`table`, `json`, `yaml`).

## Configuration

voxctl uses a kubeconfig-style YAML configuration file at `~/.voxctl/config`.

```yaml
apiVersion: v1
kind: Config
current-context: production
servers:
  - name: production
    server: https://puppet.prod.example.com:8140
    puppetdb: https://puppetdb.prod.example.com:8081
    ca-cert: /path/to/ca.pem
credentials:
  - name: admin
    client-cert: /path/to/cert.pem
    client-key: /path/to/key.pem
contexts:
  - name: production
    server: production
    credential: admin
```

## Related Projects

- [openvox-operator](https://github.com/slauger/openvox-operator) — Kubernetes operator for OpenVox/Puppet Server
- [openvox-code](https://github.com/slauger/openvox-code) — Fast, Git-native Puppet environment deployment tool

## License

MIT License — see [LICENSE](LICENSE) for details.
