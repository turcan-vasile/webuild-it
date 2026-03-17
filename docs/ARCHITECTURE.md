# Architecture

## Goal

Use `webuild-it.com` as both:

- a live public business card for IT infrastructure and automation services;
- the long-term home for an online DevOps/customer dashboard.

## Why a Separate App

`webuild-it.com` should stay separate from `ra-planet.com` because:

- it is a different brand and audience;
- it will likely become a product/control surface later;
- release cadence will differ from the WordPress site.

## Phase Plan

### Phase 1: Static Presence

- static HTML deployed with Nginx;
- ingress + TLS on the production cluster;
- simple, fast, easy to cache.

### Phase 2: Internal Dashboard

Potential features:

- project overview
- deployment status
- CI/CD health
- notes and task memory
- uptime/security snapshots
- client-specific portals

### Phase 3: Service Platform

Potential split:

- `webuild-it.com` public marketing site
- `app.webuild-it.com` authenticated dashboard
- `status.webuild-it.com` public status page
- `docs.webuild-it.com` playbooks and service documentation

## Cluster Placement

Production cluster:

- namespace: `webuild-it`
- ingress: dedicated hostname
- same Traefik controller as `ra-planet`
- same GitOps model, separate app definition

## Delivery Strategy

1. build and publish the static image
2. deploy with GitOps
3. validate HTTPS and DNS
4. later replace the static site with an app while keeping the same hostname
