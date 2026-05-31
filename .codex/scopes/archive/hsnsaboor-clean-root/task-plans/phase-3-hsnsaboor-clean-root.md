---
description: Phase 3 task plan for owning local root entrypoint docs.
---

# Tasks: HsnSaboor Clean Root Phase 3

## Input

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`
- `.gitattributes`

## Canonical architecture / Key constraints

- Root docs are local project identity and operator policy.
- Future upstream sync should not blindly overwrite root docs.
- README/AGENTS should be concise and current-code-aware.

## Format

- `[ID] [P?] [Component] Description`
- `[P]` means parallelizable.
- Valid `Component` values: `Backend`, `Frontend`, `Agentic`, `Docs`, `Config`, `QA`, `Security`, `Infra`.
- Every task must include a clear DoD.

## Phase 3: Root Entrypoint Ownership

Goal: replace upstream-like root docs with local maintained Plus fork docs and protect them from accidental upstream overwrites.

Definition of Done: root docs describe this maintained fork and `.gitattributes` records local ownership.

Tasks:

- [x] T010 [Docs] Rewrite `README.md`
  - DoD: README states HsnSaboor-first, router-selective maintenance and links to scope/wiki knowledge.
- [x] T011 [Docs] Rewrite `AGENTS.md`
  - DoD: AGENTS records provider preservation, root-doc ownership, docs/wiki boundaries, and verification commands.
- [x] T012 [Docs] Decide `README_CN.md` / `README_JA.md` policy
  - DoD: both files are intentionally absent unless a future scope reintroduces maintained translations.
- [x] T013 [Docs] Decide and update `CLAUDE.md`
  - DoD: file is either created with local policy or intentionally absent with rationale in checklist.
- [x] T014 [Config] Add `.gitattributes` root-doc merge policy
  - DoD: README/AGENTS/CLAUDE root docs use the selected merge policy.
- [x] T015 [QA] Verify root-doc policy
  - DoD: `git diff -- README.md AGENTS.md CLAUDE.md .gitattributes` shows only intended local identity changes.

Checkpoint: Phase 4 can start after root docs and merge policy are committed.

## Dependencies & Execution Order

- Phase 3 depends on Phase 2.
- T010-T014 may be edited together.
- T015 blocks commit and Phase 4.
