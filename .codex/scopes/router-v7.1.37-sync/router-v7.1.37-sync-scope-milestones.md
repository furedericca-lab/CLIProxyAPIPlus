---
description: Milestones for the router v7.1.37 upstream sync.
status: completed
date: 2026-06-01
---

# Router v7.1.37 Sync Milestones

## Milestone 1: Scope And Evidence Baseline

Status: Complete

- Confirm clean local branch and remote configuration.
- Confirm `v7.1.37` exists on `upstream`.
- Create integration branch `merge-router-v7.1.37`.
- Open active scope docs under `.codex/scopes/router-v7.1.37-sync/`.

## Milestone 2: Merge And Conflict Resolution

Status: Complete

- Merge tag `v7.1.37` into the integration branch.
- Resolve conflicts according to local provider and documentation ownership
  policy.
- Record conflict categories and any provider-impact decisions.

## Milestone 3: Verification

Status: Complete

- Run provider-surface before/after checks.
- Run targeted Go tests for touched modules.
- Run the required compile check.
- Run broad tests if targeted validation passes.
- Run wiki rebuild/doctor if durable knowledge or scope docs changed.

## Milestone 4: Merge Back, Push, And Archive Readiness

Status: Complete

- Merge the validated integration branch back to `main`.
- Commit and push implementation/scope evidence.
- Archive only if the user requests closeout after validation.
