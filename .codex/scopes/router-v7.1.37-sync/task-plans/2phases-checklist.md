---
description: Two-phase checklist for router v7.1.37 sync.
status: completed
date: 2026-06-01
---

# Router v7.1.37 Sync Checklist

## Phase 1: Scope And Merge

Status: Complete

- [x] Confirm local `main` status.
- [x] Confirm `upstream` remote URL.
- [x] Confirm `v7.1.37` tag hash.
- [x] Fetch router tags and record fetch caveat.
- [x] Create integration branch.
- [x] Create active scope docs.
- [x] Merge `v7.1.37`.
- [x] Resolve conflicts and record conflict categories.
- [x] Capture provider surface after merge.

## Phase 2: Verify And Publish

Status: Complete

- [x] Run targeted tests.
- [x] Run compile check.
- [x] Run broad tests.
- [x] Rebuild/lint wiki if scope or durable wiki pages changed.
- [x] Run `git diff --check`.
- [x] Merge back to `main`.
- [x] Push.

## Evidence

- 2026-06-01: `git status --short --branch` returned
  `## main...origin/main`.
- 2026-06-01: `git ls-remote --tags upstream v7.1.37` returned
  `05b972479aeb6885235e8d363cdc8a15be41fd6f refs/tags/v7.1.37`.
- 2026-06-01: `git fetch upstream --tags` fetched `v7.1.34` through
  `v7.1.37`; non-zero exit was caused by old local tag clobber protection, not
  by failure to fetch the requested tag.
- 2026-06-01: `git merge --no-edit v7.1.37` conflicted in
  `internal/config/config.go` and `internal/registry/model_definitions.go`.
- 2026-06-01: conflict resolution kept local Plus additions and accepted
  router Codex/XAI/config additions.
- 2026-06-01: targeted tests initially failed on `Session_id` header casing;
  local adaptation preserves existing header key spelling in
  `setHeaderCasePreserved`.
- 2026-06-01: targeted tests passed:
  `go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth`.
- 2026-06-01: compile check passed:
  `go build -o test-output ./cmd/server && rm test-output`.
- 2026-06-01: broad tests passed: `go test ./...`.
- 2026-06-01: wiki checks passed:
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild --json`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint`;
  `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json`.
- 2026-06-01: scope checks passed:
  `bash /root/.codex/skills/repo-task-driven/scripts/doc_placeholder_scan.sh .codex/scopes/router-v7.1.37-sync`;
  `bash /root/.codex/skills/repo-task-driven/scripts/scope_sync_check.sh router-v7.1.37-sync README.md AGENTS.md`.
- 2026-06-01: `git diff --check` passed.
- 2026-06-01: `main` fast-forwarded to merge commit `5a697b18`.
- 2026-06-01: closeout docs prepared for push after merge-back.
- 2026-06-01: follow-up aligned Codex OAuth implementation with the upstream
  main branch; `openai_auth.go` and `oauth_server.go` now have no diff from
  upstream. `openai_auth_test.go` stays locally adapted for this fork's
  transport wrapper.
- 2026-06-01: follow-up verification passed:
  `go test ./internal/auth/codex ./internal/api/handlers/management ./internal/runtime/executor`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `go test ./...`.
- 2026-06-01: second follow-up aligned all router-owned OAuth/auth files with
  upstream main and narrowed `oauth-endpoint-overrides` documentation/tests to
  Plus-only consumers (`github-copilot`, `kiro`).
- 2026-06-01: second follow-up verification passed:
  `git diff upstream/main -- internal/auth/antigravity/auth.go internal/auth/claude/anthropic_auth.go internal/auth/claude/oauth_server.go internal/auth/claude/token.go internal/auth/gemini/gemini_auth.go internal/auth/kimi/kimi.go internal/auth/codex/openai_auth.go internal/auth/codex/oauth_server.go internal/auth/xai internal/auth/vertex internal/auth/empty`;
  `rg -n 'GetOAuthEndpointOverride\("(antigravity|claude|codex|gemini|kimi|xai|vertex|empty)"' internal sdk` returned no source consumers;
  `go test ./internal/auth/antigravity ./internal/auth/claude ./internal/auth/codex ./internal/auth/gemini ./internal/auth/kimi ./internal/auth/xai ./internal/config ./internal/api/handlers/management ./internal/runtime/executor ./sdk/auth`;
  `go build -o test-output ./cmd/server && rm test-output`;
  `go test ./...`.
