---
description: Execution checklist for the HsnSaboor clean-root integration scope.
---

# Phases Checklist: HsnSaboor Clean Root

## Input References

- `docs/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md`
- `docs/archive/hsnsaboor-clean-root/task-plans/phase-1-hsnsaboor-clean-root.md`
- `docs/archive/hsnsaboor-clean-root/task-plans/phase-2-hsnsaboor-clean-root.md`
- `docs/archive/hsnsaboor-clean-root/task-plans/phase-3-hsnsaboor-clean-root.md`
- `docs/archive/hsnsaboor-clean-root/task-plans/phase-4-hsnsaboor-clean-root.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`

## Global Status Board

| Phase | State | Completion | Health | Blockers |
| --- | --- | ---: | --- | --- |
| Phase 1: Prepare clean root | Completed | 100% | Passed | None |
| Phase 2: Local docs/wiki policy | Completed | 100% | Passed | None |
| Phase 3: Root entrypoint ownership | Completed | 100% | Passed | None |
| Phase 4: Verification and switch decision | Completed | 100% | Passed | Main replacement pending user approval |

## Phase 1: Prepare Clean Root

- [x] Backup branch exists.
- [x] Integration branch exists from `upstream/main`.
- [x] Provider baseline captured.
- [x] No old local commits replayed.

Evidence commands:

```bash
git branch --list 'backup/main-before-hsnsaboor-clean-root' 'integrate/hsnsaboor-clean-root'
git log -1 --oneline
git log --oneline upstream/main..HEAD
```

Result:

- `backup/main-before-hsnsaboor-clean-root` exists.
- `integrate/hsnsaboor-clean-root` was created from `upstream/main` at `8c93cf68`.
- Initial branch had no commits above `upstream/main` before applying current maintenance docs.

## Phase 2: Local Docs/Wiki Policy

- [x] `.codex/wiki/**` is trackable.
- [x] `.codex/notepad.md` remains ignored.
- [x] Scope docs exist under `docs/archive/hsnsaboor-clean-root/`.
- [x] Wiki lint passes.

Evidence commands:

```bash
git check-ignore -v .codex/wiki/index.md .codex/notepad.md docs/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md || true
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
```

Result:

- `.codex/wiki/**` is added to git.
- `.codex/notepad.md` remains ignored.
- `wiki_note.py lint` passed.

## Phase 3: Root Entrypoint Ownership

- [x] `README.md` is our maintained Plus fork README.
- [x] `AGENTS.md` is our local operator contract.
- [x] `README_CN.md` and `README_JA.md` policy is decided.
- [x] `CLAUDE.md` policy is decided.
- [x] `.gitattributes` protects locally owned root docs.

Evidence commands:

```bash
git diff -- README.md README_CN.md README_JA.md AGENTS.md CLAUDE.md .gitattributes
git config --get merge.ours.driver
```

Result:

- Root docs now describe our maintained Plus fork and scope/wiki boundaries.
- `.gitattributes` sets `merge=ours` for root docs and `docs/**`.
- `git config --get merge.ours.driver` returns `true`.

## Phase 4: Verification and Switch Decision

- [x] Build passes.
- [x] Tests pass or failures are documented as pre-existing/upstream.
- [x] Provider matrix is intact.
- [ ] User explicitly approves or rejects replacing `main`.

Evidence commands:

```bash
go build -o test-output ./cmd/server && rm test-output
go test ./...
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Result:

- `go build -o test-output ./cmd/server && rm test-output` passed.
- `go test ./...` passed.
- Provider matrix includes HsnSaboor Plus additions: `cline`, `xai`, and `ollama`.
- `gofmt -w .` was not run because this batch changed docs/config only and no Go files.
- `git log --oneline upstream/main..HEAD` shows only `ef390d0f docs: establish clean-root maintenance policy`.

## Final Release Gate Summary

Clean-root branch verification passed. Replacing `main` remains blocked until explicit user approval.

## Archive Record

- Archived on 2026-05-28 under `docs/archive/hsnsaboor-clean-root/`.
- Archive purpose: preserve the completed hsnsaboor-clean-root audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `docs/<enhancement-scope>/`.
- Archived docs should only change for factual errata or path-maintenance updates.
