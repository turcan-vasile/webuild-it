# AGENTS.md (webuild-it)

## Repo Scope
Public static site workload and deployment manifests.

## Critical Constraints
- Keep deployment manifests minimal and explicit.
- Validate ingress/TLS changes with safety-first defaults.
- Avoid coupling with `ra-planet` runtime assumptions.

## Verification Defaults
- Run static checks/build for changed frontend assets.
- Validate k8s base manifests when modified.

## Task Memory
Use matrix source-of-truth:
- `/home/ja/repos/ja-ops-home/snapshot/docs/ai/tasks.yaml`
