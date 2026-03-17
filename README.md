# voxctl

CLI tool for managing OpenVox/Puppet Server infrastructure via REST APIs.

## Overview

`voxctl` provides a convenient command-line interface for common Puppet infrastructure operations that would otherwise require manual `curl` calls or `kubectl exec` into server pods.

### Supported Operations

| Command | Description | API |
|---------|-------------|-----|
| `voxctl ca` | Certificate lifecycle management (list, sign, revoke, clean, show) | Puppet CA API |
| `voxctl env` | Environment management (list, cache clear) | Puppet Server API |
| `voxctl node` | Node management (list, deactivate, purge, facts) | PuppetDB API |
| `voxctl report` | Report viewing (list, show) | PuppetDB API |

## Installation

### From source

```bash
go install github.com/slauger/voxctl/cmd/voxctl@latest
```

### From releases

Download the latest binary from the [GitHub Releases](https://github.com/slauger/voxctl/releases) page.

## Configuration

`voxctl` uses a kubeconfig-style YAML configuration file at `~/.voxctl/config`.

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

### Context Management

```bash
voxctl config get-contexts       # list all contexts
voxctl config current-context    # show active context
voxctl config use-context prod   # switch context
voxctl config use-context -      # switch to previous context
voxctl config use-context        # interactive picker (requires fzf)
```

## Usage

### Certificate Management

```bash
voxctl ca list
voxctl ca show <certname>
voxctl ca sign <certname>
voxctl ca revoke <certname>
voxctl ca clean <certname>
```

### Environment Management

```bash
voxctl env list
voxctl env cache clear
```

### Node Management

```bash
voxctl node list
voxctl node facts <certname>
voxctl node deactivate <certname>
voxctl node purge <certname>
```

### Report Management

```bash
voxctl report list
voxctl report list --node <certname>
voxctl report show <hash>
```

### Global Flags

```bash
--config string    # config file (default: ~/.voxctl/config)
--context string   # override current-context
-o, --output       # output format: table, json, yaml (default: table)
```

## Building

```bash
make build       # build binary to bin/voxctl
make test        # run tests
make lint        # run golangci-lint
make snapshot    # goreleaser snapshot build
```

## Related Projects

- [openvox-operator](https://github.com/slauger/openvox-operator) - Kubernetes operator for OpenVox/Puppet Server

## License

MIT License - see [LICENSE](LICENSE) for details.
