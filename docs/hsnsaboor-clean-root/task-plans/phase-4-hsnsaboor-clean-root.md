---
description: Phase 4 task plan for verifying the clean root and deciding whether to replace main.
---

# Tasks: HsnSaboor Clean Root Phase 4

## Input

- `integrate/hsnsaboor-clean-root`
- `docs/hsnsaboor-clean-root/`
- `.codex/wiki/**`

## Canonical architecture / Key constraints

- Do not replace `main` without explicit user approval.
- Do not force-push without explicit user approval.
- Provider surface must remain at least HsnSaboor-equivalent.

## Format

- `[ID] [P?] [Component] Description`
- `[P]` means parallelizable.
- Valid `Component` values: `Backend`, `Frontend`, `Agentic`, `Docs`, `Config`, `QA`, `Security`, `Infra`.
- Every task must include a clear DoD.

## Phase 4: Verification and Switch Decision

Goal: prove the clean-root branch is buildable, provider-complete, and ready for a user decision on replacing `main`.

Definition of Done: build passes, test status is documented, provider surface is intact, and the user has a concrete switch/no-switch recommendation.

Tasks:

- [x] T016 [QA] Run Go formatting
  - DoD: Not applicable for this batch because no Go files changed.
- [x] T017 [QA] Run required build
  - DoD: `go build -o test-output ./cmd/server && rm test-output` passes.
- [x] T018 [QA] Run full tests or document blockers
  - DoD: `go test ./...` passes, or exact failing packages are recorded with pre-existing/upstream/new classification.
- [x] T019 [QA] Verify provider matrix
  - DoD: auth/executor/login matrices show HsnSaboor provider surface is intact.
- [x] T020 [QA] Check no old local commits were replayed unintentionally
  - DoD: `git log --oneline upstream/main..HEAD` shows only our current maintenance commits after commit.
- [x] T021 [Docs] Update checklist evidence and recommendation
  - DoD: `4phases-checklist.md` contains pass/fail evidence and a final switch recommendation.

Checkpoint: Replacing `main` is blocked until the user explicitly approves.

## Dependencies & Execution Order

- Phase 4 depends on Phases 1-3.
- T017 depends on T016.
- T021 depends on T017-T020.
