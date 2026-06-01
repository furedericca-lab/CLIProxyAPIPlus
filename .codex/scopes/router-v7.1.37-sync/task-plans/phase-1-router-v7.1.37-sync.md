---
description: Phase 1 plan for scope setup, merge, and conflict resolution.
status: completed
date: 2026-06-01
---

# Phase 1: Scope And Merge

## Objective

Merge router tag `v7.1.37` into `merge-router-v7.1.37` and preserve local Plus
provider behavior.

## Tasks

- Confirm the merge target and current local branch.
- Merge `v7.1.37`.
- Resolve conflicts in Codex/config/registry surfaces.
- Keep local documentation ownership intact.
- Capture post-merge provider surface.

## Done When

- The working tree has no merge conflicts.
- Plus-only providers remain present.
- Scope docs record conflicts and immediate merge evidence.

## Evidence

- `git merge --no-edit v7.1.37` produced two content conflicts:
  `internal/config/config.go` and `internal/registry/model_definitions.go`.
- Conflict resolution preserved local Plus additions and accepted router
  Codex/XAI additions.
- Provider surface still includes local/HsnSaboor-exclusive Plus providers.
