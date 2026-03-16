# voxctl

CLI tool for managing OpenVox/Puppet Server infrastructure via REST APIs.

## Overview

`voxctl` provides a convenient command-line interface for common Puppet infrastructure operations that would otherwise require manual `curl` calls or `kubectl exec` into server pods.

### Supported Operations

| Command | Description | API |
|---------|-------------|-----|
| `voxctl cert` | Certificate lifecycle management (list, sign, revoke, clean, show) | Puppet CA API |
| `voxctl env` | Environment management (list, cache clear) | Puppet Server API |
| `voxctl node` | Node management (list, deactivate, purge, facts) | PuppetDB API |
| `voxctl report` | Report viewing (list, show) | PuppetDB API |

## Installation

### From source

```bash
go install github.com/slauger/voxctl@latest
```

### From releases

Download the latest binary from the [GitHub Releases](https://github.com/slauger/voxctl/releases) page.

## Usage

All commands authenticate via mTLS using a client certificate and key.

```bash
voxctl --ca-cert /path/to/ca.pem \
       --client-cert /path/to/cert.pem \
       --client-key /path/to/key.pem \
       --server https://puppet:8140 \
       cert list
```

### Certificate Management

```bash
voxctl cert list
voxctl cert sign <certname>
voxctl cert revoke <certname>
voxctl cert clean <certname>
```

### Environment Management

```bash
voxctl env list
voxctl env cache clear
```

### Node Management

```bash
voxctl node list --puppetdb-server https://puppetdb:8081
voxctl node facts <certname>
```

## Configuration

`voxctl` can be configured via flags, environment variables, or a config file (`~/.voxctl.yaml`).

```yaml
server: https://puppet:8140
puppetdb-server: https://puppetdb:8081
ca-cert: /etc/puppetlabs/puppet/ssl/certs/ca.pem
client-cert: /etc/puppetlabs/puppet/ssl/certs/admin.pem
client-key: /etc/puppetlabs/puppet/ssl/private_keys/admin.pem
```

## Related Projects

- [openvox-operator](https://github.com/slauger/openvox-operator) - Kubernetes operator for OpenVox/Puppet Server

## License

MIT License - see [LICENSE](LICENSE) for details.
