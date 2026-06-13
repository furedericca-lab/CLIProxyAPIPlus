---
description: Post-merge risk scan and low-risk hygiene cleanup after the Plus merge-surface reduction closeout.
---

# post-merge-hygiene-scan-2026-06-13 Contract

## Context

- Current repo/worktree: `/root/work/CLIProxyAPIPlus` on `main`.
- Baseline before this scope: `HEAD == origin/main == 3b1109db` (`Refactor Plus extension merge surfaces`).
- Active router baseline: `upstream/main == 6a0b198c`.
- User request: run another latest risk scan and hygiene cleanup.
- Boundary: keep this pass to low/medium-risk cleanup; do not enter high-risk routing, scheduler, auth persistence, provider retirement, CI, release, or upstream documentation alignment work.

## Findings

- `go vet ./...` found three actionable low/medium-risk issues:
  - `internal/logging/request_logger.go`: `FileBodySource.WriteTo` used a `WriteTo(io.Writer) error` signature that looks like `io.WriterTo` but does not implement the interface.
  - `sdk/api/handlers/handlers.go`: unreachable code remained in the stream handler header initialization path.
  - `internal/pluginhost/host_callbacks.go`: stream callback context cancel functions were not used on all return paths.
- `gofmt -l $(git ls-files '*.go')` found tracked Go files with formatting drift from prior work.
- No high-risk provider behavior or routing/scheduler change was required to clear the scan.

## Outcome

- Updated `FileBodySource.WriteTo` to implement `io.WriterTo`-compatible `(int64, error)` semantics and adjusted callers to ignore the byte count where it is not needed.
- Made pluginhost stream callback cancel ownership explicit: cancel on error/unavailable bridge paths, transfer cancel ownership to the stream bridge on success.
- Moved stream header initialization to the synchronous pre-return boundary only when the pre-read stream state has payload or normal close; error-before-payload paths still skip stream interceptors.
- Applied `gofmt` to tracked Go files that were out of format.
- Recorded this maintenance pass in the wiki maintenance log and kept the scope archived as a completed hygiene record.

## Goals / Non-goals

Goals:
- Clear current `go vet` findings.
- Remove formatting drift.
- Preserve stream interceptor behavior covered by existing tests.
- Keep evidence in scope/wiki for future upstream merge hygiene.

Non-goals:
- No provider behavior redesign.
- No scheduler, conductor, auth persistence, routing policy, CI, release, or root README/AGENTS upstream alignment.
- No broad dead-code deletion beyond mechanically proven vet/format cleanup.

## Target files / modules

- `internal/logging/request_logger.go`
- `internal/api/middleware/response_writer.go`
- `internal/pluginhost/host_callbacks.go`
- `sdk/api/handlers/handlers.go`
- Go files listed by `gofmt -l $(git ls-files '*.go')`
- `.codex/wiki/maintenance-log.md`

## Constraints

- Required Go compile check after Go changes: `go build -o test-output ./cmd/server && rm test-output`.
- Keep local-owned documentation and scope/wiki policy intact.
- Do not print or persist secrets.
- Do not make high-risk changes in this scope.

## Verification surface

- `go vet ./...`
- `go test ./internal/logging ./internal/api/middleware ./internal/pluginhost ./sdk/api/handlers`
- `go build -o test-output ./cmd/server && rm test-output`
- `go test ./...`
- `git diff --check`
- `gofmt -l $(git ls-files '*.go')`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py lint`
- `python3 /root/.codex/skills/wiki-note/scripts/wiki.py surface-check --json`

## Execution log / evidence updates

- 2026-06-13T12:54:28+08:00: Initial `go vet ./...` failed on `FileBodySource.WriteTo`, unreachable stream handler code, and pluginhost cancel ownership.
- 2026-06-13T12:54:28+08:00: Initial targeted tests failed after a too-broad stream header init move; `TestHandlerStreamChunkErrorBeforePayloadSkipsResponseInterceptors` proved error-before-payload must not initialize stream interceptors.
- 2026-06-13T12:56:42+08:00: Corrected stream header init to run before return only for pre-read payload or normal close.
- 2026-06-13T12:56:42+08:00: `go vet ./...` passed.
- 2026-06-13T12:56:42+08:00: `go test ./internal/logging ./internal/api/middleware ./internal/pluginhost ./sdk/api/handlers` passed.
- 2026-06-13T12:56:42+08:00: `go build -o test-output ./cmd/server && rm test-output` passed.
- 2026-06-13T12:56:42+08:00: `go test ./...` passed.
- 2026-06-13T12:56:42+08:00: `git diff --check` passed.
- 2026-06-13T12:56:42+08:00: `gofmt -l $(git ls-files '*.go')` returned no files.

## Escalation triggers

- Escalate if future vet findings require changing provider behavior, auth state, routing policy, scheduler selection, public APIs, or persisted config formats.
- Open a new dedicated scope for any high-risk conductor/scheduler/auth-persistence work.

## Rollback

- Revert the cleanup commit and re-run `go vet ./...`, targeted tests, the required server build, and full `go test ./...`.

## Archive Record

- Archived immediately on 2026-06-13 under `.codex/scopes/archive/post-merge-hygiene-scan-2026-06-13/`.
- Future hygiene passes should use a new dated scope instead of reopening this archived record, except for factual errata.
