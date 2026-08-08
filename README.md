# webuild-it

`webuild-it.com` is the public home for infrastructure, CI/CD, automation, AI-assisted development, and systems work under the **WeBuildit** brand.

> **We build useful things together.**

## Brand & Storytelling

The living source of truth for the WeBuildit story, community direction, editorial lanes, and proof-of-work content is kept separately from engineering backlog work:

- [Brand & Storytelling Hub](docs/brand/README.md)
- [Origin Story](docs/brand/ORIGIN-STORY.md)
- [Community Vision](docs/brand/COMMUNITY-VISION.md)
- [Content Seeds](docs/brand/CONTENT-SEEDS.md)
- [Case Study Framework](docs/brand/CASE-STUDY-FRAMEWORK.md)

These documents are **brand/storytelling memory**, not implementation issues or engineering authorization.

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

The repository already includes a simple base manifest set in [k8s/base](k8s/base).

Expected production hostnames:

- `webuild-it.com`
- `www.webuild-it.com`

## Roadmap

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).
