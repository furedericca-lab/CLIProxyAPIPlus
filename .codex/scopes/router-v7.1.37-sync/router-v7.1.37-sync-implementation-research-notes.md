---
description: Research notes for the router v7.1.37 upstream sync.
status: in_progress
date: 2026-06-01
---

# Router v7.1.37 Sync Research Notes

## Current Upstream Boundary

- Local baseline before scope: `25c0d9a6 docs: align wiki documentation surfaces`.
- Previous router merge tag in history: `v7.1.33` at `33983b6f`.
- Requested upstream target: `v7.1.37` at `05b97247`.
- Upstream commits after `v7.1.33`:
  - `0f24cafb feat(executor): implement identity obfuscation for Codex requests and responses`
  - `bbcdaab7 feat(executor): enhance Codex identity obfuscation with turn and window metadata handling`
  - `ac1360f4 feat(models): add support for grok-imagine-video-1.5-preview model`
  - `fb4f39d3 test(models, executor): add XAI video model test and fix Codex User-Agent assertions`
  - `05b97247 feat(executor): refine session and conversation header handling for Codex`

## Upstream Diff Shape

`git diff --stat v7.1.33..v7.1.37` showed 17 changed files and about 577
insertions / 64 deletions. The change set is concentrated in:

- Codex executor request/header/session handling.
- Codex websocket header defaults and websocket tests.
- XAI model registry and OpenAI video handler coverage.
- Config schema/example and config diff reporting.
- Auth selector support for new routing/config fields.

## Provider Preservation Risk

The upstream diff does not directly edit Plus-only provider directories, but
the merge still needs provider-surface checks because core executor/config
changes can indirectly affect local providers.

Provider surfaces that must survive:

- Auth dirs: `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`,
  `kilo`, `kiro`, `qwen`.
- Executors: `cline`, `codebuddy`, `cursor`, `github_copilot`, `gitlab`,
  `iflow`, `kilo`, `kiro`, `qwen`.
- Login commands: `codebuddy_login.go`, `cursor_login.go`,
  `github_copilot_login.go`, `gitlab_login.go`, `iflow_cookie.go`,
  `iflow_login.go`, `kilo_login.go`, `kiro_login.go`, `qwen_login.go`.

## Merge Strategy

Use a normal merge from tag `v7.1.37` into `merge-router-v7.1.37`, then resolve
conflicts in place:

- Accept router changes for Codex executor and websocket behavior unless they
  conflict with local Plus compatibility.
- Keep local root docs, wiki, and scope docs.
- Keep Plus provider directories and commands.
- Re-run targeted tests before broad `go test ./...`.

## Decision Summary

Proceed with a direct tag merge to `v7.1.37` rather than four separate tag
merges. The upstream delta is five commits, tightly related to Codex/XAI/config
surfaces, and already tag-bounded by the requested release. Provider-surface
checks and targeted tests provide the main safety net.

