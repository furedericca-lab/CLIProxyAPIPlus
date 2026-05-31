---
description: Contract for bare-ip-tls-skip.
---

# bare-ip-tls-skip Contract

## Context

- Current repo/worktree: `/root/work/CLIProxyAPIPlus`.
- Relevant source paths:
  - `sdk/proxyutil/proxy.go`
  - `internal/runtime/executor/helps/proxy_helpers.go`
  - `internal/runtime/executor/helps/utls_client.go`
  - `internal/util/proxy.go`
  - `internal/api/handlers/management/api_tools.go`
  - `internal/runtime/executor/helps/proxy_helpers_test.go`
  - `internal/api/handlers/management/api_tools_test.go`
  - `internal/auth/codex/openai_auth_test.go`
  - `internal/auth/kimi/kimi_proxy_test.go`
  - `sdk/proxyutil/proxy_test.go`
- Relevant archived scope references: none required; this is a narrow runtime transport change.

## Findings

- Finding: custom Claude `base-url` requests use the shared proxy-aware HTTP client path and currently inherit Go's default certificate hostname verification.
- Evidence: `internal/runtime/executor/claude_executor.go` calls `newProxyAwareHTTPClient(...).Do(...)` for streaming and raw Claude HTTP paths; `sdk/proxyutil.BuildHTTPTransport` clones the default transport without changing TLS verification.
- Finding: the user's target shape is an HTTPS URL whose hostname is an IP literal, where certificate SAN/CN mismatch is expected.
- Evidence: user request asked for bare-IP HTTPS to skip TLS verification by default, without a user-facing `insecure-skip-tls-verify` parameter.
- Finding: the management resource test path used by `/ai-providers/codex` and
  `/ai-providers/claude` builds its own transport, separate from executor
  runtime HTTP clients.
- Evidence: `internal/api/handlers/management/api_tools.go` returned raw
  transports from `apiCallTransport` for management API calls; those were not
  wrapped by `proxyutil.WrapBareIPTLSBypass` before the follow-up fix.
- Finding: the reported `https://114.119.173.40/v1` endpoint is reachable after
  TLS verification is skipped.
- Evidence: live probe without `-k` failed on certificate SAN mismatch; with
  `-k`, `https://114.119.173.40/v1` returned HTTP 404 and
  `https://114.119.173.40/v1/responses` returned HTTP 401 `Missing API key`,
  proving the remaining 502 boundary was local TLS verification in the
  management transport path.

## Outcome

- Done when: HTTPS requests whose URL hostname is an IPv4 or IPv6 literal use an insecure TLS transport, while domain-name HTTPS requests continue to use normal certificate verification.
- User-visible/runtime state: `https://114.119.173.40/...` style upstreams can be reached through CLIProxyAPIPlus without adding a config flag solely for TLS verification bypass.
- Durable knowledge to preserve: the bypass is automatic only for URL host IP literals, not a global TLS-disable switch.

## Goals / Non-goals

Goals:
- Add default TLS verification bypass for HTTPS upstream URLs with IP-literal hosts.
- Preserve proxy behavior for the insecure IP-literal path.
- Keep domain-name HTTPS certificate verification unchanged.
- Cover the routing behavior with unit tests.

Non-goals:
- Do not add a new config key.
- Do not disable TLS verification globally.
- Do not change provider authentication, headers, model routing, or retry policy.

## Target files / modules

- `sdk/proxyutil/proxy.go`: central round-tripper wrapper for IP-literal HTTPS TLS bypass.
- `internal/runtime/executor/helps/proxy_helpers.go`: apply the wrapper to shared proxy-aware HTTP clients.
- `internal/runtime/executor/helps/utls_client.go`: apply the wrapper to the standard fallback path used outside official Anthropic uTLS hosts.
- `internal/util/proxy.go`: apply the same wrapper to utility-created HTTP clients that use `SetProxy`.
- `internal/api/handlers/management/api_tools.go`: apply the same wrapper to management API calls used by `/ai-providers/*` resource checks.
- `sdk/proxyutil/proxy_test.go`, `internal/runtime/executor/helps/proxy_helpers_test.go`, `internal/api/handlers/management/api_tools_test.go`, `internal/auth/codex/openai_auth_test.go`, and `internal/auth/kimi/kimi_proxy_test.go`: regression tests for routing, proxy-preserving clone behavior, and direct/proxy override semantics after wrapping.

## Constraints

- Keep comments and user-visible strings in English.
- Do not leak API keys or auth material in logs, docs, or tests.
- Keep `proxy-url: direct` semantics intact by cloning the existing direct transport.
- Avoid weakening domain-name HTTPS verification.

## Boundaries

Allowed changes:
- Transport selection and tests needed for automatic bare-IP HTTPS TLS bypass.
- Scope evidence updates under `.codex/scopes/archive/bare-ip-tls-skip/`.

Forbidden changes:
- Config schema additions for a user-facing TLS skip flag.
- Provider-specific behavioral rewrites beyond choosing the outgoing HTTP transport.
- Changes to archived scope contents.

## Decision Summary

| Decision | Evidence Source | Evidence Strength | Conflict | Confidence Reason | Result |
| --- | --- | ---: | --- | --- | --- |
| Implement a central `proxyutil` round-tripper wrapper that routes only HTTPS IP-literal hosts to an insecure clone of the existing transport. | code/user | high | none | This satisfies the user request without introducing a config flag and preserves proxy/direct settings by cloning the selected transport. | Accepted |
| Apply the wrapper in shared executor HTTP client helpers and Claude uTLS fallback transport. | code | high | none | Claude custom `base-url` and most provider executors use these helpers; official `api.anthropic.com` keeps its uTLS path. | Accepted |
| Leave domain-name HTTPS verification enabled. | security/user | high | none | The requested bypass is for bare IP only; disabling all TLS verification would be broader than needed. | Accepted |

## Verification surface

- `gofmt -w sdk/proxyutil/proxy.go sdk/proxyutil/proxy_test.go internal/runtime/executor/helps/proxy_helpers.go internal/runtime/executor/helps/utls_client.go internal/util/proxy.go`
- `go test ./sdk/proxyutil`
- `go test ./internal/runtime/executor/helps`
- `go test ./internal/api/handlers/management`
- `go test ./internal/auth/codex ./internal/auth/kimi`
- `go test ./...`
- `go build -o test-output ./cmd/server && rm test-output`
- `bash /root/.codex/skills/repo-task-driven/scripts/doc_placeholder_scan.sh .codex/scopes/archive/bare-ip-tls-skip`
- `git diff --check`

## Escalation triggers

- Escalate only when code/runtime evidence, authoritative wiki, and scope .codex/scopes materially conflict and the conflict cannot be resolved from local evidence.
- Escalate for data deletion, permission semantics, production access model, or public API compatibility decisions outside the stated boundaries.
- Escalate when user-specified boundaries cannot be satisfied together.

## Rollback

- Revert the wrapper in `sdk/proxyutil/proxy.go`, remove wrapper calls in `internal/runtime/executor/helps` and `internal/util/proxy.go`, and remove the added proxyutil tests.
- Revert the wrapper call in `internal/api/handlers/management/api_tools.go`
  if management API call transport behavior must be restored.

## Open questions

- None.

## Execution log / evidence updates

- 2026-05-31: Created scope contract with single-contract mode for automatic bare-IP HTTPS TLS bypass.
- 2026-05-31: Implemented `proxyutil.WrapBareIPTLSBypass` and applied it to shared executor/client transport construction.
- 2026-05-31: Verification passed:
  - `gofmt -w sdk/proxyutil/proxy.go sdk/proxyutil/proxy_test.go internal/runtime/executor/helps/proxy_helpers.go internal/runtime/executor/helps/utls_client.go internal/util/proxy.go internal/runtime/executor/helps/proxy_helpers_test.go`
  - `go test ./sdk/proxyutil`
  - `go test ./internal/runtime/executor/helps`
  - `go test ./internal/auth/codex ./internal/auth/kimi`
  - `git diff --check`
  - `go test ./...`
  - `go build -o test-output ./cmd/server && rm test-output`
- 2026-05-31: Reopened after `/ai-providers/codex` management resource checks
  still returned frontend 502 for `https://114.119.173.40/v1`.
- 2026-05-31: Confirmed `curl -k` reaches the target endpoint and receives
  upstream HTTP responses, while normal TLS fails on certificate SAN mismatch.
- 2026-05-31: Extended the bypass to management API calls by wrapping all
  `apiCallTransport` return paths with `proxyutil.WrapBareIPTLSBypass`.
- 2026-05-31: Follow-up verification passed:
  - `gofmt -w sdk/proxyutil/proxy.go internal/api/handlers/management/api_tools.go internal/api/handlers/management/api_tools_test.go`
  - `go test ./internal/api/handlers/management`
  - `go test ./sdk/proxyutil`
  - `go test ./internal/runtime/executor/helps`

## Archive Record

- Archived on 2026-05-31 under `.codex/scopes/archive/bare-ip-tls-skip/`.
- Archive purpose: preserve the completed bare-ip-tls-skip audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `docs/<enhancement-scope>/`.
- Archived .codex/scopes should only change for factual errata or path-maintenance updates.
