# 🦊 voxctl

> **🚧 Status: Concept Phase** — This project is in the early concept and design phase. Nothing is implemented yet. We are currently collecting ideas and defining the scope. Feedback, suggestions, and ideas are very welcome — feel free to open an [issue](https://github.com/slauger/voxctl/issues).

CLI tool for managing OpenVox/Puppet Server infrastructure via REST APIs.

voxctl provides a convenient command-line interface for common Puppet infrastructure operations that would otherwise require manual `curl` calls or `kubectl exec` into server pods:

- 🔐 **Certificate Lifecycle** — List, sign, revoke, clean, and inspect certificates via the Puppet CA API
- 🌍 **Environment Management** — List environments and flush the environment cache via the Puppet Server API
- 🖥️ **Node Management** — List, deactivate, and purge nodes, retrieve facts via PuppetDB
- 📊 **Report Viewing** — List and inspect Puppet run reports from PuppetDB
- ⚙️ **Kubeconfig-style Configuration** — Multi-server context switching with named servers, credentials, and contexts
- 🔒 **mTLS Authentication** — Native mutual TLS support for secure API communication
- 📋 **Flexible Output** — Table, JSON, and YAML output formats for scripting and automation
- 🔀 **Interactive Context Switching** — fzf-powered interactive picker and `use-context -` for quick switching

## Quick Start

```bash
# Install from source
go install github.com/slauger/voxctl/cmd/voxctl@latest

# List all certificates
voxctl ca list

# List all nodes from PuppetDB
voxctl node list
```

## Commands

```
voxctl ca          Certificate lifecycle management (list, sign, revoke, clean, show)
voxctl env         Environment management (list, cache clear)
voxctl node        Node management (list, deactivate, purge, facts)
voxctl report      Report viewing (list, show)
voxctl config      Context and configuration management
```

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
