---
description: Phase 2 plan for verification, merge-back, and push.
status: in_progress
date: 2026-06-01
---

# Phase 2: Verify And Publish

## Objective

Validate the merged code, update durable evidence, merge back to `main`, and
push.

## Tasks

- Run targeted Go tests for touched modules.
- Run the required compile check.
- Run broad `go test ./...`.
- Run wiki rebuild/doctor after scope docs are updated.
- Run `git diff --check`.
- Merge the integration branch back to `main`.
- Push `main` to `origin`.

## Done When

- Validation commands have recorded pass/fail results.
- `main` contains the validated merge and scope evidence.
- `origin/main` has the pushed result.

## Evidence

- Targeted tests passed:
  `go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth`.
- Compile check passed:
  `go build -o test-output ./cmd/server && rm test-output`.
- Broad tests passed: `go test ./...`.
- Wiki rebuild, doctor, legacy lint, and surface check passed.
- Scope placeholder scan and sync check passed; sync check emitted only
  non-blocking warnings for intentionally absent upstream `docs/` and planned
  `origin/main` text.
- `git diff --check` passed.
