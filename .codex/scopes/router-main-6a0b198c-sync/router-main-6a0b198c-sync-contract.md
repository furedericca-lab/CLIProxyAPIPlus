---
description: Sync maintained Plus fork with router upstream/main at 6a0b198c.
status: complete
date: 2026-06-13
host: hermes
service: CLIProxyAPIPlus
category: upstream-sync
---

# Router Main 6a0b198c Sync Contract

## Context

- Current local branch before integration: `main` at `338220f3` (`Archive router main 58305350 sync scope`).
- Integration branch: `router-main-6a0b198c-sync`.
- Router upstream target: `upstream/main` at `6a0b198c` (`Merge pull request #3823 from router-for-me/fix/cors-expose-plugin-support`).
- Latest fetched router release tag: `v7.1.73`.
- Previous archived sync scope: `.codex/scopes/archive/router-main-58305350-sync/router-main-58305350-sync-contract.md`.
- Durable upstream policy: `.codex/wiki/reference/upstream-plus-maintenance.md`.

## Findings

- `git fetch upstream main --tags` updated `upstream/main` from `58305350` to `6a0b198c`; it returned non-zero only because old local tags `v6.8.44-0`, `v6.8.45-0`, `v6.9.2-0`, and `v6.9.5-0` would be clobbered.
- Upstream advanced through plugin store install/update support, host model callbacks, plugin scheduler updates, CORS plugin-support header exposure, auth singleflight, Antigravity/Codex fixes, and release/docs churn.
- A raw `HEAD..upstream/main` diff proposes deleting local `.codex/**`, root docs policy, and Plus-only providers/executors. Those deletions are upstream mirror differences, not acceptable local sync results.

## Outcome

- Merged router `upstream/main` `6a0b198c` into the maintained Plus fork on integration branch `router-main-6a0b198c-sync`.
- Resolved conflicts according to local provider/documentation policy: local root docs and intentionally absent translated READMEs stayed local, release workflow stayed local, and Plus-only providers/executors/login commands were preserved.
- Runtime/source state now includes upstream pluginstore, host model callback, scheduler, CORS plugin support header, sanitizer, translator, registry, and management API fixes while preserving Plus behavior.
- Durable knowledge updated in `.codex/wiki/reference/upstream-plus-maintenance.md`.

## Goals / Non-goals

Goals:
- Merge router `upstream/main` at `6a0b198c`.
- Preserve local Plus provider surfaces under `internal/auth`, `internal/runtime/executor`, `internal/cmd`, `sdk/auth`, and `sdk/cliproxy`.
- Keep local root docs owned by this fork: `README.md`, `AGENTS.md`, and intentionally absent translated READMEs unless local policy changes.
- Accept upstream pluginhost, pluginstore, translator, registry, CORS, and management API fixes where compatible with Plus behavior.
- Run required Go compile/test checks and wiki/scope hygiene checks before claiming completion.

Non-goals:
- Retire any Plus-only provider.
- Adopt upstream translated README files or replace local root docs with router docs.
- Rewrite release packaging beyond the smallest compatibility changes required by the merge.
- Change external accounts, production services, or host-level Hermes runtime.

## Target files / modules

- `cmd/server/main.go`
- `internal/api/**`
- `internal/auth/**`
- `internal/config/**`
- `internal/pluginhost/**`
- `internal/pluginstore/**`
- `internal/registry/**`
- `internal/runtime/executor/**`
- `internal/translator/**`
- `sdk/**`
- `go.mod`, `go.sum`
- `.github/workflows/**`, `.goreleaser.yml`, `Dockerfile`, `.gitignore`, `.gitattributes`
- `.codex/scopes/router-main-6a0b198c-sync/**`
- `.codex/wiki/reference/upstream-plus-maintenance.md`

## Constraints

- Follow `AGENTS.md`: compile check `go build -o test-output ./cmd/server && rm test-output` is required after Go changes.
- Do not leak credentials or auth material in logs, docs, diffs, or final answers.
- Use router implementation as source of truth for router-owned providers, but preserve Plus-only providers unless a separate scope retires them.
- Treat upstream deletion of local docs/wiki/scopes/provider files as a conflict requiring local preservation.
- Use `gofmt` after Go edits and run `go mod tidy` only if module metadata requires it.

## Boundaries

Allowed changes:
- Merge and adapt upstream source, tests, examples, plugin docs, and build metadata needed for router `6a0b198c`.
- Update scope and wiki documentation with merge evidence and pitfalls.

Forbidden changes:
- Remove Plus-only auth providers, executors, login commands, Kiro/GitLab/Copilot/Cursor/iFlow/Qwen/Kilo/CodeBuddy/Cline surfaces, or local `.codex/**` knowledge without explicit retirement scope.
- Recreate `README_CN.md` or `README_JA.md` as local root docs.
- Replace local `README.md`, `AGENTS.md`, or `.gitattributes` policy with upstream mirror content.

## Decision Summary

| Decision | Evidence Source | Evidence Strength | Conflict | Result | Confidence Reason |
| --- | --- | --- | --- | --- | --- |
| Track `upstream/main` `6a0b198c` instead of only latest tag | `git fetch upstream main --tags`, `git rev-parse upstream/main`, latest tag `v7.1.73` | high | resolved | Merge exact upstream main hash | Repo policy tracks router `main`; tags are release markers but main is the active baseline. |
| Preserve Plus-only providers during merge | `AGENTS.md`, `.codex/wiki/reference/upstream-plus-maintenance.md`, provider-surface command output | high | resolved | Reject upstream deletion hunks for local providers | This fork is explicitly not a plain router mirror. |
| Keep root docs and `.codex/**` local | `AGENTS.md`, wiki source-of-truth policy | high | resolved | Preserve local docs/wiki/scopes | Upstream docs are not authoritative for this maintained fork. |
| Keep local release workflow | prior `58305350` scope, `.github/workflows/release.yaml` conflict, translated README absence policy | high | resolved | Used local release workflow side | Upstream release flow still assumes root translated README files that are absent by policy. |
| Fix Antigravity refresh test transport | failing `go test ./...`, SIGQUIT stack at `TestAntigravityRefresh_DeduplicatesConcurrentRefresh`, `newAntigravityHTTPClient` proxy-aware path | high | resolved | Inject test transport through context `cliproxy.roundtripper` | This matches the current runtime transport priority and removes the test hang. |

## Verification surface

- `git remote -v`
- `git rev-parse --short HEAD upstream/main`
- `git ls-tree -d --name-only HEAD:internal/auth | sort`
- `git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort`
- `git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort`
- `gofmt -w <changed Go files>`
- `go build -o test-output ./cmd/server && rm test-output`
- `go test ./sdk/cliproxy ./sdk/cliproxy/auth`
- `go test ./...`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json`
- `ok-skill run repo-task-driven placeholder-scan .codex/scopes/router-main-6a0b198c-sync`
- `git diff --check`

## Escalation triggers

- Escalate only when code/runtime evidence, authoritative wiki, and scope docs materially conflict and the conflict cannot be resolved from local evidence.
- Escalate for data deletion, permission semantics, production access model, or public API compatibility decisions outside the stated boundaries.
- Escalate when user-specified boundaries cannot be satisfied together.

## Rollback

- Before merging back to `main`, abandon the integration branch and return to `main` at `338220f3`.
- After merge commit, revert the merge commit if validation misses a regression.

## Open questions

- None.

## Execution log / evidence updates

- 2026-06-13: Created scope contract with target `upstream/main` `6a0b198c` and latest fetched tag `v7.1.73`.
- 2026-06-13: `git fetch upstream main --tags` updated `upstream/main`; command returned non-zero due known old tag clobber warnings while the remote-tracking branch advanced.
- 2026-06-13: `git merge --no-edit upstream/main` produced conflicts in `.github/workflows/release.yaml`, `README_CN.md`, `README_JA.md`, `internal/api/handlers/management/handler.go`, `internal/api/server.go`, `internal/api/server_test.go`, `internal/runtime/executor/gemini_cli_executor.go`, `internal/runtime/executor/openai_compat_executor.go`, `sdk/api/handlers/handlers.go`, `sdk/cliproxy/auth/conductor.go`, and `sdk/cliproxy/builder.go`.
- 2026-06-13: Resolved root docs/release conflicts by preserving local policy: deleted root translated READMEs remained absent and `.github/workflows/release.yaml` kept local GoReleaser/Plus packaging.
- 2026-06-13: Resolved code conflicts by combining local Plus behavior with upstream plugin/core changes: management handler kept IP blacklist and config-applied hooks while adding pluginstore reload fields; CORS now exposes plugin support headers; SDK execution paths keep local token/fallback metadata while adding upstream plugin interceptors; auth conductor applies plugin after-auth interception after local model/fallback rewrites; builder sets both fallback chains and plugin scheduler.
- 2026-06-13: Fixed post-merge test failures: `internal/logging/gin_logger_test.go` now calls `GinLogrusLogger(&config.Config{})`; `internal/runtime/executor/antigravity_refresh_test.go` injects test transport through context `cliproxy.roundtripper`.
- 2026-06-13: Validation passed: `go build -o test-output ./cmd/server && rm test-output`; `go test ./sdk/cliproxy ./sdk/cliproxy/auth`; `go test ./internal/logging`; `go test -run TestAntigravityRefresh_DeduplicatesConcurrentRefresh -count=1 -v ./internal/runtime/executor`; `go test ./...`.
- 2026-06-13: Provider surface preserved in the worktree: `internal/auth` includes Plus-only `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, `qwen`; `internal/runtime/executor` includes Plus-only executors; `internal/cmd` includes Plus login/cookie/import commands.
