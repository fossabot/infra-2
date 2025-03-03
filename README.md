<p align="center">
  <img alt="logo" src="https://user-images.githubusercontent.com/3325447/162053538-b497fc85-11d8-4fb2-b43e-11db2fd0829a.png" />
</p>

# Infra
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Finfrahq%2Finfra.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Finfrahq%2Finfra?ref=badge_shield)


Infra enables you to **discover and access** infrastructure (e.g. Kubernetes, databases). We help you connect an identity provider such as Okta or Azure active directory, and map users/groups with the permissions you set to your infrastructure.

If you don't have an identity provider, Infra supports local users for you to get started before connecting an identity provider.

## Features

* Single-command to discover & access all your infrastructure (as an example, for Kubernetes, Infra automatically creates and syncs your kubeconfig locally after `infra login` and gets out of your way so you can use your favorite tools to access it)
* No more out-of-sync user configurations no matter where your clusters are hosted
* Support for native RBAC (e.g. support for default Kubernetes cluster roles or mapping to your own existing cluster roles)
* Onboard and offboard users via an identity provider (e.g. Okta)
* Workflow for dynamically requesting & granting access to users (coming soon)
* Audit logs for who did what, when (coming soon)

<p align="center">
  <img alt="product screenshot" src="https://user-images.githubusercontent.com/3325447/162065853-0073e6f2-8094-42f4-b88b-1bf03b2264e0.png"  />
</p>

## Get Started

* [Quickstart](https://infrahq.com/docs/getting-started/quickstart)

## Documentation
* [What is Infra?](https://infrahq.com/docs/getting-started/what-is-infra)
* [Architecture](https://infrahq.com/docs/reference/architecture)
* [Security](https://infrahq.com/docs/reference/security)

## Community
* [Community Forum](https://github.com/infrahq/infra/discussions) Best for: help with building, discussion about infrastructure access best practices.
* [GitHub Issues](https://github.com/infrahq/infra/issues) Best for: bugs and errors you encounter using Infra.


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Finfrahq%2Finfra.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Finfrahq%2Finfra?ref=badge_large)