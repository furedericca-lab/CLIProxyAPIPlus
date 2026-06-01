---
description: Merge router upstream v7.1.37 into the maintained Plus fork while preserving local provider behavior.
status: completed
date: 2026-06-01
---

# Router v7.1.37 Sync Contracts

## Context

`upstream` and `router` both point to
`https://github.com/router-for-me/CLIProxyAPI`. Local `main` is a maintained
Plus fork, not a plain router mirror. It has already merged router through
`v7.1.33`; the next requested target is `v7.1.37`.

The upstream delta from `v7.1.33` through `v7.1.37` is small but touches core
Codex request handling, websocket defaults, model registry entries, and config
surface:

- `v7.1.34`: Codex identity obfuscation for requests and responses.
- `v7.1.35`: Codex turn/window metadata handling.
- `v7.1.36`: XAI video model coverage and Codex User-Agent assertion fixes.
- `v7.1.37`: Codex session and conversation header handling.

## Goals

- Merge router tag `v7.1.37` on an integration branch created from local
  `main`.
- Preserve local/HsnSaboor-exclusive Plus providers:
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`,
  `kiro`, `qwen`.
- Prefer router implementations for router-owned providers and shared core
  behavior.
- Keep root docs and `.codex/wiki/**` locally owned.
- Record conflict decisions, provider-surface evidence, and verification
  results before merging back to `main`.

## Non-goals

- Do not retire Plus-only providers.
- Do not restore upstream `README_CN.md`, `README_JA.md`, the upstream docs
  tree, or upstream `.codex/scopes/**`.
- Do not rework unrelated provider architecture outside the `v7.1.37` merge
  boundary.
- Do not push until merge, validation, and scope evidence are coherent.

## Target Files And Modules

- `internal/runtime/executor/codex_executor.go`
- `internal/runtime/executor/codex_websockets_executor.go`
- `internal/runtime/executor/codex_*_test.go`
- `internal/config/config.go`
- `internal/watcher/diff/config_diff.go`
- `config.example.yaml`
- `internal/registry/model_definitions.go`
- `internal/registry/model_definitions_test.go`
- `sdk/api/handlers/openai/*`
- `sdk/cliproxy/auth/selector.go`

## Constraints

- `README.md`, `AGENTS.md`, and `CLAUDE.md` are local ownership surfaces.
- Treat router provider deletion hunks as suspect until provider-surface checks
  prove local Plus providers remain.
- Keep generated runtime secrets, tokens, cookies, and auth material out of
  docs and logs.
- Resolve conflicts against current local policy and wiki guidance, not by
  blindly accepting either side.

## Verification Surface

```bash
git remote -v
git ls-remote --tags upstream v7.1.37
git fetch upstream --tags
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth
go build -o test-output ./cmd/server && rm test-output
go test ./...
python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild --json
python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json
git diff --check
```

## Evidence Log

- 2026-06-01: `git status --short --branch` returned
  `## main...origin/main` before opening this scope.
- 2026-06-01: `git remote -v` showed `upstream` and `router` both pointing to
  `https://github.com/router-for-me/CLIProxyAPI`.
- 2026-06-01: `git ls-remote --tags upstream v7.1.37` returned
  `05b972479aeb6885235e8d363cdc8a15be41fd6f refs/tags/v7.1.37`.
- 2026-06-01: `git fetch upstream --tags` fetched `v7.1.34` through
  `v7.1.37`; it returned non-zero only because old local tags
  `v6.8.44-0`, `v6.8.45-0`, `v6.9.2-0`, and `v6.9.5-0` would be clobbered.
- 2026-06-01: integration branch created:
  `merge-router-v7.1.37`.
- 2026-06-01: `git merge --no-edit v7.1.37` conflicted only in
  `internal/config/config.go` and `internal/registry/model_definitions.go`.
- 2026-06-01: conflict resolution kept local `OllamaKey`, local
  CodeBuddy/Plus model definitions, and router's new provider-wide
  `CodexConfig` and XAI `grok-imagine-video-1.5-preview` built-in.
- 2026-06-01: first targeted test run failed in
  `TestCodexExecutorCacheHelper_IdentityConfuseRemapsBodyAndHeaders` because
  HTTP header key casing changed from `Session_id` to `session_id` during
  identity-confuse rewriting.
- 2026-06-01: fixed `setHeaderCasePreserved` to reuse an existing matching
  header key's spelling before replacing its value. This preserves HTTP
  `Session_id` while keeping websocket lowercase `session_id` where that key
  already exists.
- 2026-06-01: post-merge provider surface still included Plus-only auth dirs,
  executors, and login commands:
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`,
  `kiro`, and `qwen`.
- 2026-06-01: targeted verification passed:
  `go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth`.
- 2026-06-01: compile verification passed:
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-06-01: broad verification passed: `go test ./...`.
- 2026-06-01: wiki validation passed:
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild --json`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json`.
- 2026-06-01: scope doc checks passed:
  `doc_placeholder_scan.sh .codex/scopes/router-v7.1.37-sync` and
  `scope_sync_check.sh router-v7.1.37-sync README.md AGENTS.md`.
- 2026-06-01: `git diff --check` passed.
- 2026-06-01: `main` fast-forwarded to merge commit `5a697b18`
  (`Merge tag 'v7.1.37' into merge-router-v7.1.37`).
- 2026-06-01: closeout evidence commit prepared after merge-back so the
  active scope records the final mainline state.
- 2026-06-01: follow-up review found local Codex OAuth endpoint override and
  success-page URL validation changes relative to router. Per provider
  precedence, Codex is router-owned, so `internal/auth/codex/openai_auth.go`
  and `internal/auth/codex/oauth_server.go` were restored to match
  the upstream main branch.
- 2026-06-01: `internal/auth/codex/openai_auth_test.go` remains locally adapted
  only because this fork wraps HTTP transports with bare-IP TLS bypass handling;
  the OAuth implementation under test follows upstream endpoints and flow.
- 2026-06-01: comparing `internal/auth/codex/openai_auth.go` and
  `internal/auth/codex/oauth_server.go` against the upstream main branch
  returned no diff.
- 2026-06-01: follow-up verification passed:
  `go test ./internal/auth/codex ./internal/api/handlers/management ./internal/runtime/executor`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `go test ./...`.
- 2026-06-01: second follow-up aligned all router-owned OAuth/auth files with
  upstream main for `antigravity`, `claude`, `codex`, `gemini`, `kimi`, `xai`,
  `vertex`, and `empty`. `oauth-endpoint-overrides` was retained only as a
  Plus-only config extension for current consumers `github-copilot` and `kiro`;
  router-owned providers no longer consume it.
- 2026-06-01: second follow-up verification passed:
  `git diff upstream/main -- internal/auth/antigravity/auth.go internal/auth/claude/anthropic_auth.go internal/auth/claude/oauth_server.go internal/auth/claude/token.go internal/auth/gemini/gemini_auth.go internal/auth/kimi/kimi.go internal/auth/codex/openai_auth.go internal/auth/codex/oauth_server.go internal/auth/xai internal/auth/vertex internal/auth/empty`;
  `rg -n 'GetOAuthEndpointOverride\("(antigravity|claude|codex|gemini|kimi|xai|vertex|empty)"' internal sdk` returned no source consumers;
  `go test ./internal/auth/antigravity ./internal/auth/claude ./internal/auth/codex ./internal/auth/gemini ./internal/auth/kimi ./internal/auth/xai ./internal/config ./internal/api/handlers/management ./internal/runtime/executor ./sdk/auth`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `go test ./...`.
- 2026-06-01: runtime follow-up confirmed user-reported Codex OAuth login
  succeeded but `/responses` message sends received Cloudflare managed
  challenge HTML with HTTP 403 from `chatgpt.com/backend-api/codex`. Upstream
  issue `router-for-me/CLIProxyAPI#3626` tracks the same failure and points to
  closed PR `#2900`; the local fix extends the existing uTLS protected-host
  list to `chatgpt.com`, routes default Codex HTTP and image requests through
  that client only for the ChatGPT host, and normalizes Cloudflare challenge
  HTML into a stable `cloudflare_challenge` JSON error instead of returning a
  full HTML page to clients.
- 2026-06-01: Cloudflare follow-up verification passed:
  `go test ./internal/runtime/executor ./internal/runtime/executor/helps`;
  `go build -o test-output ./cmd/server && rm test-output`.

## Escalation Triggers

- A conflict would delete a local/HsnSaboor-exclusive Plus provider.
- Router changes require a new secret, token, or external account mutation.
- Tests reveal behavior changes outside Codex/XAI/config/model-registry merge
  scope.
- The merge cannot build without broad unrelated refactors.

## Rollback

Before merge-back, reset or delete the integration branch. After merge-back,
revert the merge commit and any follow-up conflict-resolution commits; do not
rewrite pushed `main`.

## Outcome

Router `v7.1.37` is merged into `main` with Plus provider surfaces preserved,
router-owned OAuth/auth implementations aligned to router upstream,
`oauth-endpoint-overrides` narrowed to Plus-only consumers, Codex ChatGPT
OAuth requests protected from the known Cloudflare TLS-fingerprint challenge,
and validation passing.
