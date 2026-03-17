# webuild-it

`webuild-it.com` is the public home for infrastructure, CI/CD, automation, and systems work under the `WeBuild-IT` brand.

## Phase 1

This repository starts as a **static business card site**:

- lightweight;
- easy to deploy on the same production Kubernetes cluster as `ra-planet.com`;
- safe to evolve later into an authenticated online dashboard.

## Deployment Model

- separate app from `ra-planet.com`
- separate ingress and TLS
- same production cluster
- static Nginx container now
- richer app later

## Local Usage

```bash
make serve
```

Visit `http://127.0.0.1:8088`.

## Container

```bash
make docker-build
```

## Kubernetes

The repository already includes a simple base manifest set in [k8s/base](/home/ja/repos/webuild-it/k8s/base).

Expected production hostnames:

- `webuild-it.com`
- `www.webuild-it.com`

## Roadmap

See [docs/ARCHITECTURE.md](/home/ja/repos/webuild-it/docs/ARCHITECTURE.md).
