---
title: Upstream Plus Maintenance Strategy
type: reference
status: current
scope: router-upstream-rebaseline
related_scopes:
  []
related_files:
  - .codex/scopes/archive/router-upstream-rebaseline/router-upstream-rebaseline-contract.md
  - .codex/scopes/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md
  - .codex/scopes/router-v7.1.46-sync/router-v7.1.46-sync-contract.md
  - internal/auth
  - internal/runtime/executor
  - internal/cmd
tags:
  - upstream
  - providers
  - maintenance
last_checked: 2026-06-13
updated: 2026-06-13T12:42:51+08:00
---

# Upstream Plus Maintenance Strategy

## Current source of truth

Use three remotes when maintaining this fork:

- `origin`: this maintained fork, currently `https://github.com/furedericca-lab/CLIProxyAPIPlus`
- `upstream`: active router baseline, currently `https://github.com/router-for-me/CLIProxyAPI`
- `router`: compatibility alias, currently `https://github.com/router-for-me/CLIProxyAPI`

As of the router `upstream/main` `6a0b198c` sync on 2026-06-13:

- local merge target before commit: `338220f3` (`Archive router main 58305350 sync scope`)
- router `upstream/main`: `6a0b198c` (`Merge pull request #3823 from router-for-me/fix/cors-expose-plugin-support`)
- local/router split point: `9ef99aa7` (`v7.1.9`)
- latest active sync scope: `.codex/scopes/router-main-6a0b198c-sync/`
- latest archived sync scope before this run: `.codex/scopes/archive/router-main-58305350-sync/`
- pre-clean-root local backup: `backup/main-before-hsnsaboor-clean-root` at `044678b0`

The previous validated sync scopes for `v7.1.37` and `v7.1.46` remain under
`.codex/scopes/router-v7.1.37-sync/` and
`.codex/scopes/router-v7.1.46-sync/` as historical evidence.

## Maintenance rule

Track router `main` as the active upstream baseline, then adapt this maintained
Plus fork on top of it. Do not directly merge router into the Plus line without
a provider-preservation review.

Reason: router is the actively moving core, but it removed or lacks multiple
Plus provider surfaces this fork exists to keep. Treat those differences as
adaptation work, not as permission to drop providers automatically. For
providers router already has, use router's implementation as the baseline.

## Locally owned documentation

Treat these root docs as local project identity and operator policy, not upstream-sync files:

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`

Also treat these build and documentation-adjacent surfaces as local-owned unless a future scope explicitly redesigns them:

- `.codex/**`
- `.github/**`
- `.goreleaser.yml`
- release/build documentation linked from the local fork identity

Future upstream merges should keep local versions for these files. `README_CN.md` and `README_JA.md` are intentionally absent and should not be restored from upstream. The recommended implementation is a `.gitattributes` merge rule using an `ours` merge driver, plus repo-local merge driver config. Root docs should explain this maintained Plus fork, not mirror upstream sponsor/marketing text unless intentionally reintroduced.

Tracked maintainer knowledge belongs under `.codex/wiki/**`. Scratch notes
should stay outside git; promote durable decisions, references, or debugging
breadcrumbs into typed wiki pages instead.

## Merge-surface reduction rule

The 2026-06-13 `plus-merge-surface-reduction` scope established a low-risk
pattern for local extensions that repeatedly collide with router-owned files:
keep upstream-shaped core files thin and move Plus-only registrations or
provider catalogs into same-package extension files.

Current split points:

- `cmd/server/main.go` keeps router startup flow plus one `plusLoginFlags`
  hook; Plus-only login flag registration and dispatch live in
  `cmd/server/plus_login_flags.go`.
- `sdk/auth/refresh_registry.go` keeps router-owned refresh lead providers;
  Plus-only refresh lead providers live in
  `sdk/auth/refresh_registry_plus.go`.
- `internal/registry/model_definitions.go` keeps shared/router static model
  lookup and built-ins. Plus provider catalogs now live in provider-specific
  files:
  `internal/registry/model_definitions_codebuddy.go`,
  `internal/registry/model_definitions_codebuddy_intl.go`,
  `internal/registry/model_definitions_cursor.go`,
  `internal/registry/model_definitions_github_copilot.go`, and
  `internal/registry/model_definitions_kiro.go`.
- `internal/config/config.go` keeps central config loading and shared config
  types; Plus-oriented OAuth alias/exclusion/endpoint override helpers live in
  `internal/config/oauth_plus.go`.
- `internal/api/handlers/management/auth_files.go` keeps generic auth-file
  handler flow; Antigravity primary/tier management helpers live in
  `internal/api/handlers/management/auth_files_antigravity.go`.

This scope intentionally excludes high-risk routing and scheduler internals.
Do not fold `sdk/cliproxy/auth/conductor.go`, auth scheduling, request routing
policy, or auth persistence behavior changes into this scope. Those require a
dedicated future scope with behavior-level tests.

Future low/medium-risk candidates for the same treatment, when touched by real merge conflicts:

- Introduce provider-specific management auth helpers before expanding
  `internal/api/handlers/management/auth_files.go` further.

## Provider preservation checklist

Before and after each upstream integration, compare:

```bash
git ls-tree -d --name-only <ref>:internal/auth | sort
git ls-tree --name-only <ref>:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only <ref>:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Must review these areas if router patches are involved:

- `internal/auth/**`
- `internal/runtime/executor/**`
- `internal/cmd/*_login.go`
- `sdk/auth/**`
- `internal/registry/**`
- `internal/config/**`
- `config.example.yaml`

Current `internal/auth` split as of 2026-05-31:

- Router-owned providers: `antigravity`, `claude`, `codex`, `empty`, `gemini`,
  `kimi`, `vertex`, `xai`.
- Local/HsnSaboor-exclusive Plus providers to preserve: `cline`, `codebuddy`,
  `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, `qwen`.

Provider rule:

- If a provider exists in router, converge to router's provider code.
- If a provider exists only in this fork or the former HsnSaboor Plus line,
  preserve it, use HsnSaboor's maintenance line as its update reference when
  available, and adapt it to router core APIs.
- Router-owned OAuth implementations stay aligned with router upstream unless a
  future explicit decision scope approves a local divergence. This currently
  covers `antigravity`, `claude`, `codex`, `gemini`, `kimi`, `xai`, `vertex`,
  and `empty` where OAuth or auth code exists.
- Local tests may adapt around this fork's transport wrappers, but
  authorization URL generation, token exchange, refresh, callback behavior, and
  provider-owned endpoint constants should remain upstream-compatible for
  router-owned providers.
- `oauth-endpoint-overrides` is retained only for Plus-only OAuth providers
  that still consume it, currently `github-copilot` and `kiro`. Do not wire
  router-owned providers back to this generic override map without a new
  explicit local policy decision.

## Historical convergence target

HsnSaboor adds provider/runtime pieces missing locally:

- `internal/auth/cline`
- `internal/auth/xai`
- `internal/runtime/executor/cline_executor.go`
- `internal/runtime/executor/ollama_executor.go`
- `internal/runtime/executor/xai_executor.go`
- `internal/cmd/cline_login.go`
- `internal/cmd/xai_login.go`

Pre-clean-root local commits, especially provider routing changes such as `044678b0 fix(copilot): route claude models through native messages`, require a dedicated audit before reintroduction. The completed audit is archived at `.codex/scopes/archive/old-local-commit-audit/old-local-commit-audit-contract.md`.

Historical refined policy: HsnSaboor was the baseline. That decision is now
superseded by `.codex/wiki/decisions/track-router-main-as-upstream.md`.

The clean-root history cleanup scope is archived under `.codex/scopes/archive/hsnsaboor-clean-root/`: it created a branch from `upstream/main`, then added our current maintenance commit on top.

## Router adaptation rule

Router changes after `v7.1.9` are now the primary upstream stream, but they
still require provider-preservation review before landing on this maintained
Plus line. Candidate areas include logging, image support, Codex model fetching,
Redis/home behavior, WebSocket auth parsing, translator fixes, and registry
updates.

Provider deletion hunks for local/HsnSaboor-exclusive providers should be
excluded unless a separate scope explicitly decides to retire a provider.

## Router v7.1.46 merge pitfalls

The `v7.1.46` sync exposed recurring merge hazards that future upstream work
should check early:

- Router `main` still deletes Plus-only provider surfaces. Do not accept
  upstream deletion hunks for local auth dirs, executors, login commands,
  `.codex/wiki/**`, `.codex/scopes/**`, or the intentionally absent translated
  READMEs without a separate retirement decision.
- `git fetch upstream main --tags` can update `upstream/main` while returning a
  non-zero status because old local tags such as `v6.8.44-0`, `v6.8.45-0`,
  `v6.9.2-0`, and `v6.9.5-0` would be clobbered. Treat that as a tag hygiene
  warning, not proof that `upstream/main` failed to update.
- Upstream file-backed API logging changes require the whole helper set:
  `extractWebsocketTimelineSource`, `extractAPIWebsocketTimelineSource`, and
  `extractFileBodySource` must be present with the `logRequest` source-aware
  path. Package-targeted tests may miss this; the server compile check caught
  the missing helpers.
- Upstream executor usage reporting now prefers
  `helps.NewExecutorUsageReporter` and exported methods
  `Publish`, `PublishFailure`, `EnsurePublished`, and `TrackFailure`.
  Old local wrapper-style calls such as `reporter.publish` compile only while
  using the legacy compat wrapper.
- Codex HTTP routing is intentionally local-adapted: keep `chatgpt.com`
  requests on the uTLS path, but keep other Codex requests proxy-aware with
  bare-IP TLS bypass behavior.
- Local quota cooldown policy remains one minute based on prior fork behavior,
  while upstream Cloudflare challenge tests expect an initial retry window of
  about 10 seconds. Keep Cloudflare challenge cooldown independent from the
  local quota cooldown base.
- Upstream auth removal tests depend on `sdk/cliproxy/auth/auto_refresh_loop_test.go`
  for `setRefreshLeadFactory`. If `conductor_remove_test.go` is added, bring
  that test helper file too.

The validated closeout sequence for this sync was: targeted Go tests, required
server compile check, full `go test ./...`, provider-surface checks, wiki
rebuild/doctor/lint/surface-check, and `git diff --check`.

## Router main 58305350 merge pitfalls

The 2026-06-09 sync to router `58305350` had no exact release tag. Treat the
commit hash as the target baseline unless a later tag appears on the same
commit.

New pitfalls from this sync:

- Upstream pluginhost support is cross-cutting. It touches server startup,
  management API routes, SDK handler interceptors, watcher auth parsing,
  model registration, translator hooks, and service runtime sync. Accepting
  only the new `internal/pluginhost/**` package is insufficient.
- Local Plus runtime hooks still matter. Preserve local Kiro refresh startup,
  `GetWatcher`, GitLab auth/model registration behavior, management IP
  blacklist handling, and config-applied reload callbacks while adding
  pluginhost hooks.
- Release/build changes need local adaptation. Upstream moved plugin-capable
  builds toward CGO/Debian, but its release workflow still assumes translated
  README files. Keep local release packaging unless it is deliberately
  redesigned.
- `.gitignore` patterns can block conflict resolution. Bare `cliproxy` and
  `server` patterns ignore `sdk/cliproxy` and `cmd/server`; keep them
  anchored as `/cliproxy` and `/server`.
- Tests may need signature adaptation after upstream adds context-aware model
  registration. Existing local tests should call `registerModelsForAuth` with
  `context.Background()`.

Validated closeout for this sync: server compile, targeted `sdk/cliproxy`
tests, full `go test ./...`, provider-surface checks, wiki rebuild/doctor/
lint/surface-check, and `git diff --check`.

## Router main 6a0b198c merge pitfalls

The 2026-06-13 sync moved from router `58305350` to `6a0b198c`; the latest
fetched tag in the range was `v7.1.73`.

New pitfalls from this sync:

- Upstream plugin work moved beyond the initial pluginhost stack into
  `internal/pluginstore/**`, latest-release resolution, plugin config get,
  host model callbacks, scheduler support, interceptor skipping, HTML/JSON
  sanitization, and the `X-CPA-SUPPORT-PLUGIN` response header. Future syncs
  should treat pluginhost, pluginstore, management handlers, SDK handlers, and
  CORS exposure as one integration surface.
- Preserve local management behavior while adding upstream plugin store state:
  keep management IP blacklist handling, config-applied callbacks, and local
  `SetOnConfigApplied` behavior together with upstream `SetConfigReloadHook`,
  plugin release cache, and plugin registry HTTP client fields.
- SDK handler merges need both sides of the execution metadata path. Keep local
  estimated input token metadata and fallback-route hints, while accepting
  upstream model execution source metadata, explicit response format, query
  cloning, and request-after-auth interceptors.
- `sdk/cliproxy/auth` merges must combine local fallback model context with
  upstream plugin request-after-auth interception. Apply the interceptor after
  any local OAuth execution-model and fallback-model rewrite so plugins see the
  actual upstream request.
- Executor translator merges should use upstream `ResponseFormatOrSource`
  for response translation, but preserve local OpenAI-compatible stream
  normalization with `normalizeDeltaContentArray` before passing SSE data to
  downstream translators.
- Antigravity refresh tests can hang if they rely on global
  `antigravityTransport` injection. Current runtime builds HTTP clients through
  proxy-aware helpers, so tests should inject the test transport through
  context key `cliproxy.roundtripper`.
- `GinLogrusLogger` now requires a `*config.Config`; older tests that still
  call it with no arguments must pass at least `&config.Config{}`.

Validated closeout for this sync: required server compile, targeted
`sdk/cliproxy` and `sdk/cliproxy/auth` tests, targeted logging and Antigravity
refresh tests during repair, full `go test ./...`, provider-surface checks,
wiki rebuild/doctor/legacy lint/surface-check, scope placeholder scan, and
`git diff --check`.
