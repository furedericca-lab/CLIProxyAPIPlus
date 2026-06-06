---
description: Merge router upstream main through v7.1.46 into the maintained Plus fork while preserving local provider and documentation ownership.
status: completed
date: 2026-06-06
---

# Router v7.1.46 Sync Contract

## Context

`upstream` and `router` both point to
`https://github.com/router-for-me/CLIProxyAPI`. Local `main` is a maintained
Plus fork and not a plain router mirror. The previous validated upstream sync
merged router through `v7.1.37`; the current requested target is upstream
`main` at `fca12a26`, tagged `v7.1.46`.

The upstream delta from `v7.1.38` through `v7.1.46` includes usage reporting,
Redis queue auth/error events, Codex reasoning replay/cache handling,
Cloudflare challenge retry handling, runtime auth removal, API logging source
tracking, and safemode example API key warnings.

## Goals

- Merge upstream `main`/`v7.1.46` into local `main`.
- Preserve local/HsnSaboor-exclusive Plus providers and commands.
- Prefer router implementations for router-owned providers and shared core
  behavior when the change is compatible with Plus surfaces.
- Keep local documentation and repo policy surfaces locally owned.
- Record conflict decisions, provider-surface evidence, and validation results.

## Non-goals

- Do not retire Plus-only providers.
- Do not restore upstream `README_CN.md`, `README_JA.md`, upstream `docs/`, or
  upstream `.codex/scopes/**` as local documentation.
- Do not blindly accept upstream deletion of local wiki/scope history.
- Do not introduce or expose secrets.

## Constraints

- Root docs `README.md`, `AGENTS.md`, and `CLAUDE.md` are locally owned policy
  surfaces. `CLAUDE.md` is currently intentionally absent in local `main`.
- `README_CN.md` and `README_JA.md` are intentionally absent.
- `.codex/wiki/**` and `.codex/scopes/**` are the durable local knowledge and
  scope layers.
- Router provider deletions are adaptation work, not permission to drop Plus
  behavior.
- Keep Go comments in English and run `gofmt` after Go edits.

## Verification Surface

```bash
git status --short --branch
git remote -v
git fetch upstream main --tags
git describe --tags --abbrev=0 upstream/main
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth
go build -o test-output ./cmd/server && rm test-output
go test ./...
python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json
git diff --check
```

## Evidence Log

- 2026-06-06: `git status --short --branch` returned
  `## main...origin/main` before opening this scope.
- 2026-06-06: `git remote -v` showed `upstream` and `router` both pointing to
  `https://github.com/router-for-me/CLIProxyAPI`.
- 2026-06-06: `git fetch upstream main --tags` updated `upstream/main` from
  `05b97247` to `fca12a26` and fetched new tags `v7.1.38` through `v7.1.46`.
  It returned non-zero because existing local old tags `v6.8.44-0`,
  `v6.8.45-0`, `v6.9.2-0`, and `v6.9.5-0` would be clobbered.
- 2026-06-06: `git describe --tags --abbrev=0 upstream/main` returned
  `v7.1.46`.
- 2026-06-06: pre-merge provider-surface checks showed Plus provider auth dirs,
  executors, and login commands still present on local `main`.
- 2026-06-06: `git diff --name-status HEAD..upstream/main` showed upstream
  would delete local `.codex/wiki/**`, `.codex/scopes/**`, Plus provider auth
  dirs, Plus executors, Plus login commands, translated README absence policy,
  and other local-only surfaces if accepted blindly.
- 2026-06-06: `git merge --no-edit upstream/main` conflicted in
  `cmd/server/main.go`, `internal/api/middleware/response_writer.go`,
  `internal/api/server.go`, `internal/runtime/executor/claude_executor.go`,
  `internal/runtime/executor/codex_executor.go`,
  `internal/runtime/executor/gemini_cli_executor.go`,
  `internal/runtime/executor/helps/logging_helpers.go`,
  `internal/runtime/executor/helps/utls_client.go`,
  `internal/runtime/executor/openai_compat_executor.go`, `README_CN.md`, and
  `README_JA.md`.
- 2026-06-06: conflict resolution kept local Kiro login/incognito handling,
  local Codex `chatgpt.com`-scoped uTLS routing with proxy-aware fallback for
  other hosts, local quota cooldown policy, local Plus provider surfaces, and
  local docs ownership while accepting upstream safemode, usage reporter,
  file-backed API logging, Redis/auth event, Codex reasoning replay, and auth
  removal changes.
- 2026-06-06: `README_CN.md` and `README_JA.md` remained absent per local
  policy; `git diff --cached --name-status | rg
  'README_CN|README_JA|CLAUDE|\.codex/wiki|\.codex/scopes'` showed only this
  new scope file under `.codex/scopes/router-v7.1.46-sync/`.
- 2026-06-06: the first targeted test run failed because old local executor
  reporter method calls needed the upstream `helps.UsageReporter` method names,
  `sdk/cliproxy/auth/auto_refresh_loop_test.go` was needed for the upstream
  `setRefreshLeadFactory` test helper, and Cloudflare challenge cooldown still
  inherited the local one-minute quota cooldown base.
- 2026-06-06: fixes applied: updated reporter calls to
  `Publish`/`PublishFailure`/`EnsurePublished`, added upstream
  `sdk/cliproxy/auth/auto_refresh_loop_test.go`, and split Cloudflare
  challenge cooldown to a 10-second independent backoff while preserving the
  local one-minute quota cooldown base.
- 2026-06-06: targeted verification passed:
  `go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth`.
- 2026-06-06: compile verification passed:
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-06-06: broad verification passed: `go test ./...`.
- 2026-06-06: post-merge provider surface still included Plus auth dirs
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`,
  `kiro`, and `qwen`; executors `cline`, `codebuddy`, `github_copilot`,
  `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, and `qwen`; and their login or
  cookie commands.

## Escalation Triggers

- A conflict cannot be resolved without deleting a Plus provider.
- A router change needs a new secret, token, or external account mutation.
- Tests reveal behavior changes outside the upstream sync boundary.
- Validation cannot build without broad unrelated refactors.

## Rollback

Before final commit, abort the merge with `git merge --abort`. After commit,
revert the merge commit and any follow-up conflict-resolution commits; do not
rewrite pushed `main` unless explicitly requested.

## Outcome

Router upstream `main` through `v7.1.46` is merged into the maintained Plus
fork with Plus provider surfaces and local documentation ownership preserved.
Validation passed for targeted Go packages, the required server compile check,
and full `go test ./...`.
