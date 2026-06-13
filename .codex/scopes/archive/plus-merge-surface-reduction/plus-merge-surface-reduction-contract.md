---
description: Reduce recurring upstream merge conflicts by moving local Plus extension blocks out of router-owned files.
---

# plus-merge-surface-reduction Contract

## Context

- Current repo/worktree: `/root/work/CLIProxyAPIPlus` on `main` after router `upstream/main` `6a0b198c` sync.
- User boundary: CI, `.codex/**`, scope/wiki, `.github`, README, and related build/documentation surfaces are local-owned; do not chase upstream there.
- Relevant source paths: `cmd/server/main.go`, `cmd/server/plus_login_flags.go`, `sdk/auth/refresh_registry.go`, `sdk/auth/refresh_registry_plus.go`, `internal/config/config.go`, `internal/config/oauth_plus.go`, `internal/registry/model_definitions.go`, `internal/registry/model_definitions_codebuddy.go`, `internal/registry/model_definitions_codebuddy_intl.go`, `internal/registry/model_definitions_cursor.go`, `internal/registry/model_definitions_github_copilot.go`, `internal/registry/model_definitions_kiro.go`, `internal/api/handlers/management/auth_files.go`, `internal/api/handlers/management/auth_files_antigravity.go`.
- Durable policy source: `.codex/wiki/reference/upstream-plus-maintenance.md`.

## Findings

- `cmd/server/main.go` had router-owned startup flow interleaved with Plus-only login flag variables, flag registration, and command dispatch.
- `sdk/auth/refresh_registry.go` mixed router-owned refresh lead registrations with Plus-only providers.
- `internal/config/config.go` mixed central config loading with Plus-oriented OAuth alias/exclusion/endpoint override helpers.
- `internal/registry/model_definitions.go` carried Plus-only CodeBuddy, CodeBuddy Intl, Cursor, GitHub Copilot, Kiro, and Amazon Q static model definitions inside the central router model definition file.
- `internal/api/handlers/management/auth_files.go` carried Antigravity primary/tier helpers beside generic auth-file handlers; dedicated Antigravity tests made this safe enough for a medium-risk mechanical split.
- Long-term upstream diff remains large; reducing merge pain means isolating local extension blocks where behavior can stay local without changing public API names.

## Outcome

- Plus-only login flags and dispatch moved to `cmd/server/plus_login_flags.go`.
- Plus-only refresh lead registrations moved to `sdk/auth/refresh_registry_plus.go`.
- CodeBuddy static model definitions moved to `internal/registry/model_definitions_codebuddy.go`.
- CodeBuddy Intl static model definitions moved to `internal/registry/model_definitions_codebuddy_intl.go`.
- Cursor static model definitions moved to `internal/registry/model_definitions_cursor.go`.
- GitHub Copilot static model definitions and allowlist moved to `internal/registry/model_definitions_github_copilot.go`.
- Kiro and Amazon Q static model definitions moved to `internal/registry/model_definitions_kiro.go`.
- Plus-oriented OAuth config helpers moved to `internal/config/oauth_plus.go`.
- Antigravity management auth primary/tier helpers moved to `internal/api/handlers/management/auth_files_antigravity.go`.
- Existing public/internal function names stay stable: `GetCodeBuddyModels`, `GetCodeBuddyIntlModels`, `GetCursorModels`, `GetGitHubCopilotModels`, `IsAllowedGitHubCopilotModel`, `GetKiroModels`, `GetAmazonQModels`, `NormalizeOAuthExcludedModels`, `GetOAuthEndpointOverride`, `registerRefreshLead`, and login command functions are unchanged.
- `cmd/server/main.go` now keeps a single `plusLoginFlags` hook for local providers and remains closer to router startup shape.

## Goals / Non-goals

Goals:
- Reduce future conflict scope in router-owned files without dropping Plus behavior.
- Keep local-owned docs, CI, scope, wiki, release, and README surfaces independent from upstream.
- Preserve existing command-line flags and provider behavior.

Non-goals:
- No provider retirement.
- No config schema migration.
- No `.github`, README, AGENTS, CLAUDE, `.codex/wiki`, or `.codex/scopes` alignment to upstream.
- No high-risk conductor, scheduler, routing policy, auth persistence, or public API behavior refactor in this scope.

## Target files / modules

- `cmd/server/main.go`
- `cmd/server/plus_login_flags.go`
- `sdk/auth/refresh_registry.go`
- `sdk/auth/refresh_registry_plus.go`
- `internal/config/config.go`
- `internal/config/oauth_plus.go`
- `internal/registry/model_definitions.go`
- `internal/registry/model_definitions_codebuddy.go`
- `internal/registry/model_definitions_codebuddy_intl.go`
- `internal/registry/model_definitions_cursor.go`
- `internal/registry/model_definitions_github_copilot.go`
- `internal/registry/model_definitions_kiro.go`
- `internal/api/handlers/management/auth_files.go`
- `internal/api/handlers/management/auth_files_antigravity.go`
- `.codex/wiki/reference/upstream-plus-maintenance.md`

## Constraints

- Required Go compile check after Go edits: `go build -o test-output ./cmd/server && rm test-output`.
- Run focused package tests for touched packages before broad `go test ./...`.
- Keep comments in Go code English.
- Do not alter secrets, runtime credentials, or token file handling.
- Preserve user-specified local ownership for CI/docs/scope/wiki/root docs.

## Boundaries

Allowed changes:
- Move Plus-only code into new files in the same package/module.
- Add local registry/helper files that keep existing exported functions and flags stable.
- Update wiki and scope docs with the new maintenance boundary.

Forbidden changes:
- Do not accept or restore upstream versions of `.codex/**`, `.github/**`, README, AGENTS, CLAUDE, release workflow, or local build/documentation policy files.
- Do not rename existing CLI flags.
- Do not change provider auth file formats.
- Do not remove Plus-only providers.
- Do not modify `sdk/cliproxy/auth/conductor.go`, scheduler selection, request routing policy, auth persistence semantics, or management API behavior contracts.

## Decision Summary

| Decision | Evidence Source | Evidence Strength | Conflict | Result | Confidence Reason |
| --- | --- | --- | --- | --- | --- |
| Keep CI/docs/scope/wiki/root docs local-owned | User instruction, `AGENTS.md`, wiki policy | high | resolved | Preserve local files; no upstream alignment work here | User explicitly narrowed these surfaces and repo policy already says root docs are local identity. |
| Extract Plus login flags and dispatch from `cmd/server/main.go` | `cmd/server/main.go` diff against router-owned startup flow | high | resolved | Add `cmd/server/plus_login_flags.go` | Same package extraction preserves access to `setKiroIncognitoMode` and avoids changing command behavior. |
| Split refresh lead registrations by router vs Plus provider ownership | `sdk/auth/refresh_registry.go`, provider policy wiki | high | resolved | Add `sdk/auth/refresh_registry_plus.go` | Init registration order does not matter for distinct provider keys, and the helper remains shared. |
| Move CodeBuddy static models to a provider file | `internal/registry/model_definitions.go`, `GetCodeBuddyModels` callers | high | resolved | Add `internal/registry/model_definitions_codebuddy.go` | Function name stays unchanged, so service and lookup call sites remain stable. |
| Move Cursor static models to a provider file | `internal/registry/model_definitions.go`, `GetCursorModels` callers | high | resolved | Add `internal/registry/model_definitions_cursor.go` | Function name stays unchanged, so registry lookup and service call sites remain stable. |
| Move remaining Plus static model catalogs to provider files | `internal/registry/model_definitions.go`, registry tests | high | resolved | Add GitHub Copilot, Kiro/Amazon Q, and CodeBuddy Intl provider files | This is a same-package mechanical move; exported function names and call sites are unchanged. |
| Move Plus-oriented OAuth config helpers to a local extension file | `internal/config/config.go`, config OAuth tests | high | resolved | Add `internal/config/oauth_plus.go` | Struct fields remain on `Config`; helper methods and normalizers keep the same exported names. |
| Move Antigravity management helpers to a provider-specific file | `internal/api/handlers/management/auth_files.go`, `auth_files_antigravity_test.go` | high | resolved | Add `internal/api/handlers/management/auth_files_antigravity.go` | This is a same-package mechanical move with dedicated management tests covering the affected helper methods. |
| Remove high-risk candidates from this scope | User instruction, scope boundaries | high | resolved | Do not touch `sdk/cliproxy/auth/conductor.go` or broad auth routing behavior | User explicitly limited this scope to low- and medium-risk changes. |

## Verification surface

- `gofmt -w cmd/server/main.go cmd/server/plus_login_flags.go sdk/auth/refresh_registry.go sdk/auth/refresh_registry_plus.go internal/config/config.go internal/config/oauth_plus.go internal/registry/model_definitions.go internal/registry/model_definitions_codebuddy.go internal/registry/model_definitions_codebuddy_intl.go internal/registry/model_definitions_cursor.go internal/registry/model_definitions_github_copilot.go internal/registry/model_definitions_kiro.go`
- `go test ./cmd/server ./sdk/auth ./internal/config ./internal/registry ./internal/api/handlers/management`
- `go build -o test-output ./cmd/server && rm test-output`
- `go test ./...`
- Wiki checks after durable docs update:
  - `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild`
  - `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json`
  - `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint`
  - `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json`
- Scope hygiene:
  - `ok-skill run repo-task-driven placeholder-scan .codex/scopes/archive/plus-merge-surface-reduction`
  - `ok-skill run repo-task-driven text-scan .codex/scopes/archive/plus-merge-surface-reduction README.md`
  - `git diff --check`

## Escalation triggers

- Escalate only if a provider behavior change is required to compile or test.
- Escalate if user-specified local-owned surfaces must be changed to complete code verification.
- Escalate if upstream/router ownership evidence contradicts the current wiki policy and cannot be resolved from repo code.

## Rollback

- Delete the new Go extension files and move their contents back into the original files.
- Re-run `gofmt`, targeted tests, required compile check, and wiki/scope checks.

## Open questions

- None.

## Execution log / evidence updates

- 2026-06-13T12:23:29+08:00: Created scope and implemented file extractions for Plus login flags, Plus refresh lead registration, and CodeBuddy model definitions.
- 2026-06-13T12:23:29+08:00: `gofmt -w ...` completed successfully.
- 2026-06-13T12:23:29+08:00: Initial `go test ./cmd/server ./sdk/auth ./internal/registry` failed because `cmd/server/main.go` still uses `kiro.InitializeAndStart` and `kiro.StopGlobalRefreshManager` in service startup. Restored the `internal/auth/kiro` import.
- 2026-06-13T12:23:29+08:00: `go test ./cmd/server ./sdk/auth ./internal/registry` passed.
- 2026-06-13T12:23:29+08:00: `go build -o test-output ./cmd/server && rm test-output` passed.
- 2026-06-13T12:23:29+08:00: `go test ./...` passed.
- 2026-06-13T12:23:29+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild` passed and rebuilt `.codex/wiki/index.md` plus `.codex/wiki/decision-log.md`.
- 2026-06-13T12:23:29+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json` passed.
- 2026-06-13T12:23:29+08:00: Legacy-form commands `python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint` and `... legacy surface-check --json` returned `unexpected arguments before wiki command: legacy`; current unified equivalents are `lint` and `surface-check`.
- 2026-06-13T12:23:29+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py lint` passed.
- 2026-06-13T12:23:29+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py surface-check --json` passed.
- 2026-06-13T12:23:29+08:00: `ok-skill run repo-task-driven placeholder-scan .codex/scopes/archive/plus-merge-surface-reduction` passed.
- 2026-06-13T12:23:29+08:00: `ok-skill run repo-task-driven text-scan .codex/scopes/archive/plus-merge-surface-reduction README.md` passed.
- 2026-06-13T12:23:29+08:00: `git diff --check` passed.
- 2026-06-13T12:31:20+08:00: Implemented second batch by moving Plus-oriented OAuth config helpers to `internal/config/oauth_plus.go` and Cursor model definitions to `internal/registry/model_definitions_cursor.go`.
- 2026-06-13T12:31:20+08:00: `gofmt -w internal/config/config.go internal/config/oauth_plus.go internal/registry/model_definitions.go internal/registry/model_definitions_cursor.go` completed successfully.
- 2026-06-13T12:31:20+08:00: `go test ./internal/config` passed.
- 2026-06-13T12:31:20+08:00: `go test ./cmd/server ./sdk/auth ./internal/config ./internal/registry` passed.
- 2026-06-13T12:31:20+08:00: `go build -o test-output ./cmd/server && rm test-output` passed.
- 2026-06-13T12:31:20+08:00: `go test ./...` passed.
- 2026-06-13T12:31:20+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild` passed and rebuilt `.codex/wiki/index.md` plus `.codex/wiki/decision-log.md`.
- 2026-06-13T12:31:20+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json` passed.
- 2026-06-13T12:31:20+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py lint` passed.
- 2026-06-13T12:31:20+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py surface-check --json` passed.
- 2026-06-13T12:31:20+08:00: `ok-skill run repo-task-driven placeholder-scan .codex/scopes/archive/plus-merge-surface-reduction` passed.
- 2026-06-13T12:31:20+08:00: `ok-skill run repo-task-driven text-scan .codex/scopes/archive/plus-merge-surface-reduction README.md` passed.
- 2026-06-13T12:31:20+08:00: `git diff --check` passed.
- 2026-06-13T12:39:00+08:00: Implemented final low-risk registry split by moving GitHub Copilot, Kiro/Amazon Q, and CodeBuddy Intl static model catalogs into provider-specific files.
- 2026-06-13T12:39:00+08:00: `gofmt -w internal/registry/model_definitions.go internal/registry/model_definitions_github_copilot.go internal/registry/model_definitions_kiro.go internal/registry/model_definitions_codebuddy_intl.go` completed successfully.
- 2026-06-13T12:39:00+08:00: `go test ./internal/registry` passed.
- 2026-06-13T12:42:51+08:00: Implemented medium-risk same-package split by moving Antigravity management auth helpers to `internal/api/handlers/management/auth_files_antigravity.go`.
- 2026-06-13T12:42:51+08:00: `gofmt -w internal/api/handlers/management/auth_files.go internal/api/handlers/management/auth_files_antigravity.go` completed successfully.
- 2026-06-13T12:42:51+08:00: `go test ./internal/api/handlers/management` passed.
- 2026-06-13T12:42:51+08:00: `go test ./cmd/server ./sdk/auth ./internal/config ./internal/registry ./internal/api/handlers/management` passed.
- 2026-06-13T12:42:51+08:00: `go build -o test-output ./cmd/server && rm test-output` passed.
- 2026-06-13T12:42:51+08:00: `go test ./...` passed.
- 2026-06-13T12:42:51+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild` passed and rebuilt `.codex/wiki/index.md` plus `.codex/wiki/decision-log.md`.
- 2026-06-13T12:42:51+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json` passed.
- 2026-06-13T12:42:51+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py lint` passed.
- 2026-06-13T12:42:51+08:00: `python3 /root/.codex/skills/wiki-note/scripts/wiki.py surface-check --json` passed.
- 2026-06-13T12:42:51+08:00: `ok-skill run repo-task-driven placeholder-scan .codex/scopes/archive/plus-merge-surface-reduction` passed.
- 2026-06-13T12:42:51+08:00: `ok-skill run repo-task-driven text-scan .codex/scopes/archive/plus-merge-surface-reduction README.md` passed.
- 2026-06-13T12:42:51+08:00: `git diff --check` passed.


## Archive Record

- Archived on 2026-06-13 under `.codex/scopes/archive/plus-merge-surface-reduction/`.
- Archive purpose: preserve the completed plus-merge-surface-reduction audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `.codex/scopes/<enhancement-scope>/`.
- Archived docs should only change for factual errata or path-maintenance updates.
