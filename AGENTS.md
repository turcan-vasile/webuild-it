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

## Skills First
For repeatable and risky workflows, run project skills before ad-hoc operations.

## Beads Memory
Use Beads (`bd`) for long-horizon task graph tracking and dependencies.
Primary workspace: `/home/ja/beads-workspace`.
Sync exports via `/home/ja/repos/ja-ops-home/scripts/beads-sync.sh`.
