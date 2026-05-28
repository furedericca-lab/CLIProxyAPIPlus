---
title: Upstream Plus Maintenance Strategy
type: reference
status: current
scope: upstream-plus-gap-analysis
related_scopes:
  - upstream-plus-gap-analysis
related_files:
  - docs/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md
  - internal/auth
  - internal/runtime/executor
  - internal/cmd
tags:
  - upstream
  - providers
  - maintenance
last_checked: 2026-05-28
updated: 2026-05-28T12:55:00Z
---

# Upstream Plus Maintenance Strategy

## Current source of truth

Use three remotes when maintaining this fork:

- `origin`: this maintained fork, currently `https://github.com/furedericca-lab/CLIProxyAPIPlus`
- `upstream`: Plus continuation target, currently `https://github.com/HsnSaboor/CLIProxyAPIPlus`
- `router`: original upstream, currently `https://github.com/router-for-me/CLIProxyAPI`

As of the clean-root push on 2026-05-28:

- local `main`: `6325b63a`
- HsnSaboor `upstream/main`: `8c93cf68`
- router `router/main`: `2bcc7622`
- HsnSaboor and router split at `v7.1.9` (`9ef99aa7`)
- pre-clean-root local backup: `backup/main-before-hsnsaboor-clean-root` at `044678b0`

## Maintenance rule

Move to HsnSaboor Plus first, then selectively absorb router changes. Do not directly merge router into the Plus line without a provider-preservation review.

Reason: router removed multiple Plus provider surfaces that HsnSaboor still carries. The provider-deletion set includes `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, and `qwen`, with `cline`/`ollama` also preserved in HsnSaboor but absent from router.

## Locally owned documentation

Treat these root docs as local project identity and operator policy, not upstream-sync files:

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`

Future upstream merges should keep local versions for these files. `README_CN.md` and `README_JA.md` are intentionally absent and should not be restored from upstream. The recommended implementation is a `.gitattributes` merge rule using an `ours` merge driver, plus repo-local merge driver config. Root docs should explain this maintained Plus fork, not mirror upstream sponsor/marketing text unless intentionally reintroduced.

Tracked maintainer knowledge belongs under `.codex/wiki/**`. Scratch notes should stay outside git, for example `.codex/notepad.md`.

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

## First convergence target

HsnSaboor adds provider/runtime pieces missing locally:

- `internal/auth/cline`
- `internal/auth/xai`
- `internal/runtime/executor/cline_executor.go`
- `internal/runtime/executor/ollama_executor.go`
- `internal/runtime/executor/xai_executor.go`
- `internal/cmd/cline_login.go`
- `internal/cmd/xai_login.go`

Pre-clean-root local commits, especially provider routing changes such as `044678b0 fix(copilot): route claude models through native messages`, require a dedicated audit before reintroduction.

Current refined policy: HsnSaboor is the baseline. If HsnSaboor has the same provider or same-topic implementation, prefer HsnSaboor first and validate behavior before reintroducing old local patches.

The clean-root history cleanup scope is archived under `docs/archive/hsnsaboor-clean-root/`: it created a branch from `upstream/main`, then added our current maintenance commit on top.

## Router absorption rule

Router changes after `v7.1.9` are still valuable but should be reviewed as core patches, not as a branch replacement. Candidate areas include logging, image support, Codex model fetching, Redis/home behavior, WebSocket auth parsing, translator fixes, and registry updates.

Provider deletion hunks from router should be excluded unless a separate scope explicitly decides to retire a provider.
