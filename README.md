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

This repository contains a lightweight multilingual landing page and narrowly scoped contact service:

- lightweight;
- one dependency-free Go binary serving static assets and `/api/contact`;
- EN / RO / RU content from one semantic page;
- server-side SMTP delivery with credentials supplied only at runtime;
- easy to deploy in the existing Kubernetes workload;
- safe to evolve later into an authenticated online dashboard.

## Deployment Model

- separate app from `ra-planet.com`
- separate ingress and TLS
- same production cluster
- single non-root application container
- richer app later

## Local Usage

```bash
make serve
```

Visit `http://127.0.0.1:8080`.

Without SMTP environment variables the site remains available and the contact
API returns a clear `503` response. See [docs/CONTACT.md](docs/CONTACT.md) for
the runtime secret contract and validation flow.

## Validation

```bash
make verify
make docker-build
```

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
