---
description: Clean-root integration plan to rebuild this maintained Plus line on top of HsnSaboor CLIProxyAPIPlus.
---

# HsnSaboor Clean Root Contract

## Context

This repository currently has old second-development history, historical upstream merge commits, local CI/doc/model changes, and at least one local Copilot patch that overlaps functionally with HsnSaboor's current Plus fork.

The target is to make future history clearer:

```text
upstream/main (HsnSaboor)
  -> our: wiki and docs cleanup policy
  -> our: local README/AGENTS/merge policy
  -> our: selected repo-specific maintenance only after review
```

Current baseline checked on 2026-05-28:

- Current local `main`: `044678b0`
- HsnSaboor `upstream/main`: `8c93cf68`
- router `router/main`: `2bcc7622`
- HsnSaboor/router split point: `9ef99aa7` (`v7.1.9`)

## Findings

- HsnSaboor is the best first base because it preserves the Plus provider direction.
- Local auth providers are all present in HsnSaboor.
- HsnSaboor additionally has `cline`, `xai`, and `ollama` executor support.
- Router is newer but removed many Plus providers, so it must remain a selective patch source, not a direct merge target.
- Local old commits are not patch-identical to HsnSaboor, but some are same-topic. The main example is Copilot Claude routing.

## Goals / Non-goals

Goals:

- Create a non-destructive integration branch from `upstream/main`.
- Preserve old `main` under a backup branch before any final switch.
- Recreate only current maintenance decisions as new commits on top of HsnSaboor.
- Track `.codex/wiki/**`.
- Keep `.codex/notepad.md` and scratch ignored.
- Keep durable scope execution docs under `docs/hsnsaboor-clean-root/`.
- Keep `docs/` available for `repo-task-driven` active scopes and archives; the cleanup only removed old upstream/historical docs.
- Keep long-lived project knowledge under `.codex/wiki/**`.
- Own root entrypoint docs locally:
  - `README.md`
  - `README_CN.md`
  - `README_JA.md`
  - `AGENTS.md`
  - `CLAUDE.md`
- Add merge protection for locally owned root docs.
- Verify provider surface before and after reroot.

Non-goals for the first reroot:

- Do not directly replay old local commits.
- Do not bring old model fork/submodule behavior forward automatically.
- Do not bring old CI/release workflow customizations forward automatically.
- Do not merge router.
- Do not delete Plus providers.
- Do not publish or force-push until the user explicitly approves final branch replacement.

## Local old commit disposition

Old local commits not in HsnSaboor should not be replayed wholesale.

| Commit | Old purpose | Disposition |
| --- | --- | --- |
| `044678b0` | Copilot Claude model routing | Do not replay by default. HsnSaboor has same-topic commits. Validate behavior first. |
| `2dc61ac6` | Manual release workflow trigger | Do not replay initially. Rebuild CI policy later if still needed. |
| `566bfb69` | Point model updater to forked models source | Do not replay initially. Re-decide model source after clean root builds. |
| `1aee8779` | Sync model catalogs and forked models source | Do not replay initially. Mixed model/submodule/docs commit is too broad. |
| `ad7d7999` | Goreleaser dirty-state avoidance | Do not replay initially. Re-evaluate with current workflow. |
| `e3d123f8` | Skip Docker steps without DockerHub secrets | Do not replay initially. Rebuild CI policy later if needed. |
| `dd9f77fc` | Skip docker-image without DockerHub secrets | Do not replay initially. Rebuild CI policy later if needed. |
| `2a1cf2b3` | Revert static model cleanup | Do not replay by default. High risk against newer registry architecture. |

## Target files / modules

- `.gitignore`
- `.gitattributes`
- `.codex/wiki/**`
- `docs/hsnsaboor-clean-root/**`
- `README.md`
- `README_CN.md`
- `README_JA.md`
- `AGENTS.md`
- `CLAUDE.md`

Potential later review areas, only after clean root is stable:

- `.github/workflows/**`
- `internal/registry/**`
- `internal/runtime/executor/github_copilot_executor.go`
- `internal/registry/model_definitions.go`

## Constraints

- HsnSaboor is the Plus baseline.
- If HsnSaboor already has a provider or feature, use HsnSaboor's version by default.
- Local old implementation wins only if a current behavioral gap is proven after testing.
- Root docs are local project identity and should not be blindly synced from upstream.
- Upstream `docs/**` content should not be synced; our `docs/**` is reserved for scopes and scope archives.
- `.gitattributes` records owned-doc merge policy; local clones must set `git config merge.ours.driver true`.
- Router provider deletion hunks are rejected by default.

## Verification plan

Provider surface:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Build and tests:

```bash
gofmt -w .
go build -o test-output ./cmd/server && rm test-output
go test ./...
```

Docs/wiki:

```bash
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
find docs -maxdepth 3 -type f -print | sort
git log --oneline upstream/main..HEAD
```

## Rollback

Before any final switch:

```bash
git branch backup/main-before-hsnsaboor-clean-root main
```

The integration branch is non-destructive. If it is wrong, delete it and keep current `main`:

```bash
git branch -D integrate/hsnsaboor-clean-root
```

Final replacement of `main` and any force push require explicit user approval.

## Open questions

- Should `README_CN.md` and `README_JA.md` become full translations immediately or short pointers until translation is refreshed?
- Should `CLAUDE.md` be created now, or only if a Claude-specific workflow appears?
- After clean root builds, do we still need any part of the old model fork/submodule behavior?

## Execution log / evidence updates

- 2026-05-28: User selected HsnSaboor-first clean-root strategy.
- 2026-05-28: User clarified that scope plans belong under `docs/*`; wiki remains for durable knowledge.
