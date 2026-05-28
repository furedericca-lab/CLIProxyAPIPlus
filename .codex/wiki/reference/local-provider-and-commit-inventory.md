---
title: Pre-clean-root Provider and Commit Inventory
type: reference
status: current
scope: upstream-plus-gap-analysis
related_scopes:
  - upstream-plus-gap-analysis
related_files:
  - internal/auth
  - internal/runtime/executor
  - internal/cmd
  - internal/registry/model_definitions.go
  - .github/workflows/release.yaml
tags:
  - providers
  - commits
  - upstream
last_checked: 2026-05-28
updated: 2026-05-28T13:20:00Z
---

# Pre-clean-root Provider and Commit Inventory

## Current clean-root provider surface

Current `main` is rooted at HsnSaboor and has the HsnSaboor Plus provider surface:

- Auth dirs: old local providers plus `cline` and `xai`.
- Executors: old local executors plus `cline`, `ollama`, and `xai`.
- Login commands: old local login commands plus `cline_login.go` and `xai_login.go`.

The old local provider directories were not lost during the clean-root switch.

## Pre-clean-root local provider surface

Pre-clean-root `backup/main-before-hsnsaboor-clean-root` provider/auth directories:

- `antigravity`
- `claude`
- `codebuddy`
- `codex`
- `copilot`
- `cursor`
- `empty`
- `gemini`
- `gitlab`
- `iflow`
- `kilo`
- `kimi`
- `kiro`
- `qwen`
- `vertex`

Pre-clean-root executor names:

- `aistudio`
- `antigravity`
- `claude`
- `codebuddy`
- `codex`
- `codex_websockets`
- `cursor`
- `gemini`
- `gemini_cli`
- `gemini_vertex`
- `github_copilot`
- `gitlab`
- `iflow`
- `kilo`
- `kimi`
- `kiro`
- `openai_compat`
- `qwen`

Pre-clean-root login commands:

- `anthropic_login.go`
- `antigravity_login.go`
- `codebuddy_login.go`
- `cursor_login.go`
- `github_copilot_login.go`
- `gitlab_login.go`
- `iflow_cookie.go`
- `iflow_login.go`
- `kilo_login.go`
- `kimi_login.go`
- `kiro_login.go`
- `openai_device_login.go`
- `openai_login.go`
- `qwen_login.go`
- `vertex_import.go`

## Pre-clean-root surface relative to HsnSaboor Plus

HsnSaboor `upstream/main` was a superset of the pre-clean-root local line for provider breadth. The old local line was missing these HsnSaboor pieces:

- Auth dirs: `cline`, `xai`
- Executors: `cline`, `ollama`, `xai`
- Login commands: `cline_login.go`, `xai_login.go`

The clean-root switch now includes those HsnSaboor pieces.

## Pre-clean-root surface relative to router

Router `router/main` removed or lacks multiple Plus providers that both the pre-clean-root line and current clean-root line keep:

- Auth dirs: `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, `qwen`
- Executors: `codebuddy`, `cursor`, `github_copilot`, `gitlab`, `iflow`, `kilo`, `kiro`, `qwen`
- Login commands: `codebuddy_login.go`, `cursor_login.go`, `github_copilot_login.go`, `gitlab_login.go`, `iflow_cookie.go`, `iflow_login.go`, `kilo_login.go`, `kiro_login.go`, `qwen_login.go`

Do not merge router wholesale into the Plus line.

## Pre-clean-root local commits not in HsnSaboor

Non-merge commits from `backup/main-before-hsnsaboor-clean-root` that are not in HsnSaboor `upstream/main`:

- `044678b0` `fix(copilot): route claude models through native messages`
  - Touches `internal/registry/model_definitions.go`, `internal/registry/model_definitions_test.go`, `internal/runtime/executor/github_copilot_executor.go`, and `internal/runtime/executor/github_copilot_executor_test.go`.
  - High-value local provider behavior. Preserve or reconcile deliberately.
- `2dc61ac6` `ci(release): add manual workflow trigger`
  - Touches `.github/workflows/release.yaml`.
  - Low-risk CI convenience.
- `566bfb69` `point model updater to forked models source`
  - Touches `internal/registry/model_updater.go`.
  - Fork-specific model source behavior. Review against the current model-source plan.
- `1aee8779` `sync model catalogs and forked models source`
  - Touches `.github/workflows/release.yaml`, `.gitmodules`, docs, `internal/registry/model_definitions.go`, `internal/registry/models/models.json`, `internal/thinking/provider/iflow/apply.go`, `internal/thinking/provider/iflow/apply_test.go`, and `third_party/models`.
  - High-risk because it mixes docs, catalog data, submodule, and iFlow thinking behavior.
- `ad7d7999` `ci(release): avoid dirty git state in goreleaser workflow`
  - Touches `.github/workflows/release.yaml`.
- `e3d123f8` `ci: skip docker steps when DockerHub secrets are missing`
  - Touches `.github/workflows/docker-image.yml`.
- `dd9f77fc` `ci: skip docker-image when DockerHub secrets are missing`
  - Touches `.github/workflows/docker-image.yml`.
- `2a1cf2b3` `Revert "fix(registry): clean up outdated model definitions in static data"`
  - Touches `internal/registry/model_definitions_static_data.go`.
  - Review carefully because it may reintroduce static data that upstream intentionally removed.

Merge commits not in HsnSaboor are mostly historical upstream/router integration points and should not be replayed as feature patches.

## Merge guidance

When auditing old local commits:

- Do not replay `044678b0` unless HsnSaboor lacks equivalent native Claude routing implementation for Copilot.
- Re-evaluate model fork commits against the current decision to avoid stale `docs/*` and stale submodule behavior.
- Treat CI workflow commits as optional and easy to reapply.
- Inspect `2a1cf2b3` before carrying it forward, because static model data may conflict with newer registry architecture.

Evidence commands:

```bash
git log --oneline --no-merges upstream/main..backup/main-before-hsnsaboor-clean-root
git show --stat --oneline <commit>
git ls-tree -d --name-only backup/main-before-hsnsaboor-clean-root:internal/auth | sort
git ls-tree --name-only backup/main-before-hsnsaboor-clean-root:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only backup/main-before-hsnsaboor-clean-root:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```
