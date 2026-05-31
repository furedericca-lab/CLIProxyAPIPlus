---
description: Phase 2 task plan for applying local .codex/scopes and wiki policy on top of HsnSaboor.
---

# Tasks: HsnSaboor Clean Root Phase 2

## Input

- `.codex/scopes/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md`
- `.codex/wiki/**`
- `.gitignore`

## Canonical architecture / Key constraints

- Scope execution .codex/scopes live under `.codex/scopes/archive/hsnsaboor-clean-root/`.
- Durable project knowledge lives under `.codex/wiki/**`.
- Scratch remains ignored.

## Format

- `[ID] [P?] [Component] Description`
- `[P]` means parallelizable.
- Valid `Component` values: `Backend`, `Frontend`, `Agentic`, `Docs`, `Config`, `QA`, `Security`, `Infra`.
- Every task must include a clear DoD.

## Phase 2: Local Docs/Wiki Policy

Goal: apply the local knowledge-management policy as a clean commit on top of HsnSaboor.

Definition of Done: wiki is tracked and valid, scope .codex/scopes are under `.codex/scopes/archive/hsnsaboor-clean-root/`, and old `.codex/scopes/*` content is not restored.

Tasks:

- [x] T005 [Config] Update `.gitignore` for tracked wiki and current scope docs
  - DoD: `.codex/wiki/index.md` and `.codex/scopes/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md` are not ignored; `.codex/notepad.md` remains ignored.
- [x] T006 [Docs] Add or refresh wiki maintenance pages
  - DoD: `wiki_note.py list` shows upstream maintenance, local provider inventory, and model/catalog maintenance pages.
- [x] T007 [Docs] Add current scope .codex/scopes under `.codex/scopes/archive/hsnsaboor-clean-root/`
  - DoD: contract, phase checklist, and phase files exist under `.codex/scopes/archive/hsnsaboor-clean-root/`.
- [x] T008 [QA] Rebuild and lint wiki
  - DoD: `wiki_note.py rebuild` and `wiki_note.py lint` pass.
- [x] T009 [Docs] Commit docs/wiki policy
  - DoD: `git log --oneline upstream/main..HEAD` includes `docs: establish local wiki maintenance policy`.

Checkpoint: Phase 3 can start after the docs/wiki policy commit is clean and wiki lint passes.

## Dependencies & Execution Order

- Phase 2 depends on Phase 1.
- T005 blocks T006-T009.
- T008 blocks T009.
