---
description: Phase 1 task plan for creating the non-destructive HsnSaboor clean-root branch.
---

# Tasks: HsnSaboor Clean Root Phase 1

## Input

- `.codex/scopes/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md`
- Current `main`
- `upstream/main`
- `router/main`

## Canonical architecture / Key constraints

- HsnSaboor is the base.
- Current `main` must be preserved before final replacement.
- No old local commits are replayed in this phase.

## Format

- `[ID] [P?] [Component] Description`
- `[P]` means parallelizable.
- Valid `Component` values: `Backend`, `Frontend`, `Agentic`, `Docs`, `Config`, `QA`, `Security`, `Infra`.
- Every task must include a clear DoD.

## Phase 1: Prepare Clean Root

Goal: create a safe integration branch from `upstream/main` without changing current `main`.

Definition of Done: backup branch exists, integration branch exists from HsnSaboor, and initial branch/provider evidence is recorded.

Tasks:

- [x] T001 [Infra] Create backup branch for current `main`
  - DoD: `git branch --list backup/main-before-hsnsaboor-clean-root` shows the backup branch.
- [x] T002 [Infra] Fetch current upstream refs
  - DoD: `git fetch --prune upstream` and `git fetch --prune router` complete without errors.
- [x] T003 [Infra] Create `integrate/hsnsaboor-clean-root` from `upstream/main`
  - DoD: `git log -1 --oneline` on the integration branch shows HsnSaboor head.
- [x] T004 [QA] Capture provider baseline on clean root
  - DoD: auth, executor, and login command matrices are saved in the checklist evidence section.

Checkpoint: Phase 2 cannot start until the integration branch is confirmed to have no old local commits on top.

## Dependencies & Execution Order

- Phase 1 blocks all later phases.
- T001 and T002 must run before T003.
- T004 depends on T003.
