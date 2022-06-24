<p align="center">
  <img alt="Infra" src="https://user-images.githubusercontent.com/251292/176219932-dd8f84ae-5ad4-44e5-8865-dedb13f3c32b.svg" />
</p>

[![build](https://github.com/infrahq/infra/actions/workflows/release.yml/badge.svg)](https://github.com/infrahq/infra/releases/latest)
[![twitter](https://img.shields.io/twitter/follow/infrahq.svg?style=social&label=Follow)](https://twitter.com/infrahq)

Secure, lightweight and high-performance proxy for accessing infrastructure and data

### Features

- **User** and **group** access control with pluggable authentication & authorization
- **Connectors** for different infrastructure & data services
- **Audit** logs for compliance & visibility
- CLI for users to **discover and generate credentials** on the fly

### Connectors

| Connector          | Status        |
| ------------------ | ------------- |
| Kubernetes         | ✅ Stable     |
| AWS                | _Coming soon_ |
| Container Registry | _Coming soon_ |
| Postgres           | _Coming soon_ |
| Kafka              | _Coming soon_ |
| MySQL              | _Coming soon_ |
| MongoDB            | _Coming soon_ |
| SSH                | _Coming soon_ |
| RDP                | _Coming soon_ |
| Snowflake          | _Coming soon_ |

### Quickstart (Kubernetes)

#### Install Infra CLI

```
brew install infrahq/tap/infra
```

#### Deploy to Kubernetes

Create a configuration file `infra.yaml`:

```yaml
access:
  - user: jeff@acme.com
    role: edit
    namespace: default
```

Next, deploy the proxy to Kubernetes via `infra deploy`:

```bash
infra deploy --kubernetes -f infra.yaml
```

#### Create an account and generate credentials

Sign up for an account and access the cluster:

```
$ infra signup
  Email: jeff@acme.com
  Please verify your email... ✓ verified.
  Choose a team name: acme

  Your credentials have been updated:
    - example-cluster (example-cluster.acme.infrahq.dev)
```

Now access the cluster and verify you have `edit` access:

```
$ kubectl run --image=nginx
```

#### View access logs

```
$ infra logs

DESTINATION        USER              ACTION                         TIME           DURATION
example-cluster    jeff@acme.com     kubectl run --image nginx      just now       204ms
```

### Example configuration

```yaml
destinations:
  - name: cluster
    kind: kubernetes
    url: https://example-cluster.acme.infrahq.dev
    serviceaccount: env:KUBERNETES_TOKEN

  # container registry
  - name: registry
    kind: registry
    host: index.docker.io
    token: env:REGISTRY_TOKEN

  # database
  - name: db
    kind: postgres
    host: postgres.acme.internal
    port: 5432
    user: postgres
    password: env:POSTGRES_PASSWORD

  # aws
  - name: aws
    kind: aws
    region: us-east-1
    accessKeyId: env:AWS_ACCESS_KEY_ID
    secretAccessKey: env:AWS_SECRET_ACCESS_KEY

authentication:
  # additional keys to verify credentials against
  - jwks: https://acme.com/jwks.json

# static authorization
authorization:
  - user: admin@acme.com
    role: edit
    namespace: default

  - group: DatabaseAdmins
    destination: db
    role: view

  - user: Developers
    destination: registry
    role: view

audit:
  - format: json
    fields: [user, action, time, duration]

# Advanced: specify your own self-hosted server
# on by default, but configurable
server:
  # custom infra servers
  url: https://infrahq.com

  # authenticate against the server
  authenticate: true

  # external authorization
  authorize: true

  # send audit logs
  audit: true
```

### Documentation

- [Connect an identity provider](https://infrahq.com/configuration/identity-provider)
- [Invite your team](https://infrahq.com/docs/configuration/invite-team)
- [Security](https://infrahq.com/docs/reference/security)

## Community

- [Community Forum](https://github.com/infrahq/infra/discussions) Best for: help with building, discussion about infrastructure access best practices.
- [GitHub Issues](https://github.com/infrahq/infra/issues) Best for: bugs and errors you encounter using Infra.
