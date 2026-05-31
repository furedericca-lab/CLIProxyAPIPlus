---
description: Batch-sync router upstream releases into the maintained Plus fork while preserving local/HsnSaboor-exclusive providers.
status: completed
date: 2026-05-31
---

# Router Batch Sync Contract

## Context

`upstream` tracks `https://github.com/router-for-me/CLIProxyAPI`. The local
Plus fork has been batch-merged through router `v7.1.32`, which was
`upstream/main` on 2026-05-31. Router changes landed in tag-bounded batches
with provider precedence checks.

## Goals

- Merge router progress in batches starting from `v7.1.10`.
- Use router code for providers router already owns.
- Preserve local/HsnSaboor-exclusive providers:
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`,
  `kiro`, `qwen`.
- Use HsnSaboor's maintenance line as the update reference for those exclusive
  providers when relevant.
- Keep root docs and `docs/**` locally owned.

## Non-goals

- Do not jump directly to `upstream/main`.
- Do not remove Plus providers without an explicit retirement scope.
- Do not import upstream root docs or upstream `docs/**` as product docs.

## Batch Plan

1. Batch 1: merged `v7.1.10` through `v7.1.15`.
2. Batch 2: merged `v7.1.16` through `v7.1.20`.
3. Batch 3: merged `v7.1.21` through `v7.1.25`.
4. Batch 4: merged `v7.1.26` through `v7.1.32`.

## Verification Surface

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
go test ./internal/api ./internal/config ./internal/runtime/executor ./internal/registry
go build -o test-output ./cmd/server && rm test-output
python3 /root/.codex/skills/repo-task-driven/scripts/repo_task.py --root /root/work/CLIProxyAPIPlus check --scope router-batch-sync --json
git diff --check
```

## Evidence

- 2026-05-31: `git rev-list --left-right --count main...upstream/main`
  returned `4440 86` before batch sync.
- 2026-05-31: merge base with router is
  `9ef99aa76688f1462fab96670f75ab0d2fc3a77c` (`v7.1.9`).
- 2026-05-31: merged `v7.1.10`:
  `feat(api): add support for local management password validation and spoofed IP rejection`.
  Verification: `go test ./internal/api ./internal/config`.
- 2026-05-31: merged `v7.1.11`:
  `refactor(api): remove newTestServerWithOptions and spoofed IP rejection test`.
  Conflict: `internal/api/server_test.go` import block only. Verification:
  `go test ./internal/api ./internal/config`.
- 2026-05-31: merged `v7.1.12`: Home mTLS certificate bootstrap, HTTP CONNECT
  proxy support, Claude attribution stripping, and Claude tool-use stability.
  Conflicts: `cmd/server/main.go`,
  `internal/translator/antigravity/claude/antigravity_claude_request.go`,
  `internal/translator/codex/claude/codex_claude_request.go`. Verification:
  `go test ./internal/home ./sdk/proxyutil ./internal/translator/... ./internal/api ./internal/config`;
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-05-31: merged `v7.1.13`: upstream response header tracking for logging
  and usage. Conflict: `internal/runtime/executor/helps/usage_helpers.go`.
  Verification:
  `go test ./internal/logging ./internal/redisqueue ./internal/runtime/executor/helps ./sdk/api/handlers ./sdk/cliproxy/usage`;
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-05-31: merged `v7.1.14`: xAI reasoning effort support. Verification:
  `go test ./internal/thinking/... ./internal/runtime/executor`;
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-05-31: merged `v7.1.15`: OpenAI-compatible image models, Gemini max
  output cap, Antigravity thought signatures, and Home CA fingerprint handling.
  Conflicts: `internal/runtime/executor/openai_compat_executor.go`,
  `internal/runtime/executor/openai_compat_executor_compact_test.go`,
  `sdk/api/handlers/handlers.go`. Verification:
  `go test ./internal/runtime/executor ./sdk/api/handlers ./sdk/api/handlers/openai ./internal/watcher/diff ./internal/translator/antigravity/gemini ./internal/home ./internal/config`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `go test ./...`.
- 2026-05-31: provider surface after batch 1 still includes:
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`,
  `kiro`, `qwen`.
- 2026-05-31: `git rev-list --left-right --count main...upstream/main`
  returned `4446 64` after merging through `v7.1.15`.
- 2026-05-31: merged `v7.1.16` through `v7.1.20`: Antigravity project ID
  enforcement, Redis subscription failover, Codex context-length stream
  handling, Gemini 3.5 Flash registry entries, reasoning effort usage events,
  Redis protocol handling, Claude system role conversion, and Grok Build 0.1.
  Conflicts included root translated README deletions, `cmd/server/main.go`,
  Antigravity auth/executor tests, `internal/registry/models/models.json`,
  `sdk/api/handlers/handlers.go`, and `sdk/cliproxy/auth/conductor.go`.
  Verification:
  `go test ./internal/auth/antigravity ./internal/runtime/executor ./sdk/api/handlers ./sdk/cliproxy/auth ./internal/api`.
- 2026-05-31: merged `v7.1.21` through `v7.1.25`: file-backed request logging,
  auth-file websocket metadata handling, Gemini CLI schema cleanup, GPT Image 2
  handling, Claude/OpenAI reasoning signatures, and TTFT usage reporting.
  Conflicts included management auth files, response writer cleanup, model
  registry, Claude/Gemini/Aistudio executors, usage helpers, and OpenAI
  compatible executor. Local Plus behavior retained for Antigravity primary
  management, GitLab PAT auth, Claude alias/base URL/max token compatibility,
  and NVIDIA/OpenAI-compatible request normalization. Verification:
  `go test ./internal/api/handlers/management ./internal/api/middleware ./internal/logging ./internal/runtime/executor ./internal/runtime/executor/helps ./internal/translator/claude/openai/responses ./sdk/api/handlers/openai ./sdk/cliproxy/usage`.
- 2026-05-31: merged `v7.1.26` through `v7.1.32`: cache/service-tier usage
  accounting, signature validation extraction and provider compatibility
  checks, Amp tool casing restoration, Home app log forwarding, Claude Opus
  4.8 model registry, OpenAI websocket input ID dedupe, OAuth callback
  validation, Gemini developer role support, and websocket fallback payload
  cleanup. Conflicts included response writer cleanup, Claude executor local
  compatibility, Antigravity Claude signature helpers, and handler metadata.
  Verification:
  `go test ./internal/signature ./internal/runtime/executor ./internal/translator/... ./sdk/api/handlers ./sdk/api/handlers/openai ./sdk/cliproxy/auth ./sdk/cliproxy/usage ./internal/api/middleware ./internal/logging ./internal/api/handlers/management`.
- 2026-05-31: removed stale local Home address parsing helpers after router
  Home switched to JWT-only config; verification:
  `go test ./cmd/server`;
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-05-31: final provider surface check still includes local/HsnSaboor-only
  providers `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`,
  `kilo`, `kiro`, `qwen`.
- 2026-05-31: final `git rev-list --left-right --count main...upstream/main`
  returned `4450 0`, proving no remaining router `upstream/main` commits were
  missing.
- 2026-05-31: final verification passed:
  `go test ./...`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `git diff --check`.

## Rollback

Revert the batch merge commit and any follow-up conflict-resolution commits.
Do not reset the branch after pushing.

## Archive Record

- Archived on 2026-05-31 under `docs/archive/router-batch-sync/`.
- Archive purpose: preserve the completed router-batch-sync audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `docs/<enhancement-scope>/`.
- Archived docs should only change for factual errata or path-maintenance updates.
