---
description: Sync maintained Plus fork with router upstream/main at 58305350.
status: complete
date: 2026-06-09
host: hermes
service: CLIProxyAPIPlus
category: upstream-sync
---

# Router Main 58305350 Sync Contract

## Context

The maintained Plus fork was previously synced through router `v7.1.46`
(`fca12a26`). On 2026-06-09, `upstream/main` advanced to `58305350`
(`feat(jshandler): add new plugin providing JavaScript-based interceptors and
capabilities`). There is no exact tag on `58305350`.

Major upstream changes in this range:

- dynamic `pluginhost` support, plugin API/ABI packages, and plugin management
  routes
- JavaScript interceptor example plugin
- Codex service-tier request-normalizer plugin
- Codex generated-image extraction performance changes
- translator fixes for Gemini and Antigravity signature/chunk handling
- Docker runtime move to Debian with `ca-certificates`
- upstream release workflow churn and sponsorship README changes

## Goals

- Merge router `upstream/main` at `58305350` into local `main`.
- Preserve local Plus-only providers, executors, login commands, root docs, and
  `.codex/**` knowledge.
- Accept upstream pluginhost/runtime/translator/core fixes and adapt them into
  the local Plus service path.
- Keep local release packaging from reintroducing intentionally absent
  `README_CN.md`, `README_JA.md`, or `CLAUDE.md`.
- Record merge hazards and validation evidence.

## Non-goals

- Retire any Plus-only provider.
- Replace local root `README.md`/`AGENTS.md` identity with router upstream docs.
- Adopt upstream release workflow wholesale while it assumes translated README
  files that are intentionally absent here.

## Outcome

Merged `upstream/main` `58305350` and resolved conflicts by combining upstream
pluginhost integration with local Plus runtime behavior.

Notable adaptations:

- Kept local GoReleaser release entrypoint and `.goreleaser.yml`, while
  accepting pluginhost-related source/runtime changes.
- Updated `Dockerfile` to use Debian/CGO/`ca-certificates` for dynamic plugin
  support while preserving the local `CLIProxyAPIPlus` binary name and `-plus`
  version suffix.
- Preserved local management API IP blacklist/config-apply behavior while
  adding plugin management routes and plugin auth persistence hooks.
- Preserved Kiro background refresh startup while routing service startup
  through `StartServiceWithPluginHost`.
- Preserved local OAuth alias channels for `github-copilot` and `kiro` while
  adding plugin OAuth provider channel support.
- Narrowed over-broad `.gitignore` patterns from `cliproxy`/`server` to
  `/cliproxy`/`/server` so tracked source paths such as `sdk/cliproxy` and
  `cmd/server` are not ignored.

## Provider Preservation

Post-merge provider surface remained present:

- `internal/auth`: `antigravity`, `claude`, `cline`, `codebuddy`, `codex`,
  `copilot`, `cursor`, `empty`, `gemini`, `gitlab`, `iflow`, `kilo`, `kimi`,
  `kiro`, `qwen`, `vertex`, `xai`
- `internal/runtime/executor`: `aistudio`, `antigravity`, `claude`, `cline`,
  `codebuddy`, `codex`, `codex_websockets`, `cursor`, `gemini`, `gemini_cli`,
  `gemini_vertex`, `github_copilot`, `gitlab`, `iflow`, `kilo`, `kimi`, `kiro`,
  `ollama`, `openai_compat`, `qwen`, `xai`
- `internal/cmd`: `anthropic_login.go`, `antigravity_login.go`,
  `cline_login.go`, `codebuddy_login.go`, `cursor_login.go`,
  `github_copilot_login.go`, `gitlab_login.go`, `iflow_cookie.go`,
  `iflow_login.go`, `kilo_login.go`, `kimi_login.go`, `kiro_login.go`,
  `openai_device_login.go`, `openai_login.go`, `qwen_login.go`,
  `vertex_import.go`, `xai_login.go`

## Verification

Evidence commands run:

```bash
git fetch upstream main
git fetch router main
git fetch origin main
git log --oneline fca12a26..upstream/main
git tag --points-at upstream/main
git describe --tags --exact-match upstream/main || true
gofmt -w ...
go mod tidy
go build -o test-output ./cmd/server && rm test-output
go test ./sdk/cliproxy ./sdk/cliproxy/auth
go test ./...
```

Results:

- `upstream/main` advanced `fca12a26..58305350`.
- `58305350` has no exact tag.
- Required server compile passed.
- Targeted `sdk/cliproxy` and `sdk/cliproxy/auth` tests passed.
- Full `go test ./...` passed.

Pending closeout checks:

- wiki rebuild/doctor/lint/surface-check
- `git diff --check`
- commit and push to `origin/main`
