---
description: Technical notes for router v7.1.37 merge adaptation.
status: completed
date: 2026-06-01
---

# Router v7.1.37 Sync Technical Documentation

## Expected Runtime Change Areas

### Codex Executor

Router adds identity obfuscation and session/conversation header handling in
the Codex executor and websocket executor. Local merge review should verify
that these changes do not break local Codex-compatible routing or OpenAI image
handling.

### Codex OAuth

Codex is a router-owned provider. The OAuth implementation in
`internal/auth/codex/openai_auth.go` and `internal/auth/codex/oauth_server.go`
should stay aligned with router upstream for authorization URL generation, token
exchange, refresh, and callback success behavior. This fork should not add
Codex OAuth endpoint overrides unless a future explicit decision scope changes
that policy.

`internal/auth/codex/openai_auth_test.go` may differ from upstream only where
needed to test this fork's HTTP transport wrapping, such as bare-IP TLS bypass
behavior.

### Config Surface

Router adds config fields and default/diff handling for Codex websocket header
behavior. Local merge review should keep `config.example.yaml`,
`internal/config/config.go`, and `internal/watcher/diff/config_diff.go`
consistent.

### Model Registry

Router adds `grok-imagine-video-1.5-preview` coverage. Local merge review
should ensure registry tests still match the maintained fork's model-source
policy.

### OpenAI Handlers

Router updates OpenAI video handler behavior and tests. Local merge review
should preserve existing SDK handler compatibility.

## Provider Impact Boundary

No Plus-only provider should be removed or renamed by this scope. Any compile
breakage in Plus providers caused by shared config/executor changes is in
scope to adapt narrowly.

## Local Adaptation

The merge exposed a header key-casing mismatch in identity-confuse handling:
HTTP tests expect `Session_id`, while websocket tests expect lowercase
`session_id` when the websocket path already carries that key. The local
adaptation updates `setHeaderCasePreserved` to preserve the existing matching
key spelling before replacing the value. This keeps both HTTP and websocket
paths stable.

## Validation Plan

Start with touched-module tests, then compile, then broad tests:

```bash
go test ./internal/runtime/executor ./internal/config ./internal/registry ./internal/watcher/diff ./sdk/api/handlers/openai ./sdk/cliproxy/auth
go build -o test-output ./cmd/server && rm test-output
go test ./...
```

All three checks passed on 2026-06-01.

Follow-up Codex OAuth checks also passed on 2026-06-01:

```bash
git diff <upstream-main> -- internal/auth/codex/openai_auth.go internal/auth/codex/oauth_server.go
go test ./internal/auth/codex ./internal/api/handlers/management ./internal/runtime/executor
go build -o test-output ./cmd/server && rm test-output
go test ./...
```
