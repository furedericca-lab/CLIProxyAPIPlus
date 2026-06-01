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
  - internal/auth
  - internal/runtime/executor
  - internal/cmd
tags:
  - upstream
  - providers
  - maintenance
last_checked: 2026-05-31
updated: 2026-05-31T05:13:53Z
---

# Upstream Plus Maintenance Strategy

## Current source of truth

Use three remotes when maintaining this fork:

- `origin`: this maintained fork, currently `https://github.com/furedericca-lab/CLIProxyAPIPlus`
- `upstream`: active router baseline, currently `https://github.com/router-for-me/CLIProxyAPI`
- `router`: compatibility alias, currently `https://github.com/router-for-me/CLIProxyAPI`

As of the router rebaseline on 2026-05-31:

- local `main`: `df53e88d` (`v7.1.9-5`)
- router `upstream/main`: `3a54fb7f` (`v7.1.32`)
- local/router split point: `9ef99aa7` (`v7.1.9`)
- divergence count: local has 4439 commits not in router, router has 86 commits not in local
- pre-clean-root local backup: `backup/main-before-hsnsaboor-clean-root` at `044678b0`

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

Future upstream merges should keep local versions for these files. `README_CN.md` and `README_JA.md` are intentionally absent and should not be restored from upstream. The recommended implementation is a `.gitattributes` merge rule using an `ours` merge driver, plus repo-local merge driver config. Root docs should explain this maintained Plus fork, not mirror upstream sponsor/marketing text unless intentionally reintroduced.

Tracked maintainer knowledge belongs under `.codex/wiki/**`. Scratch notes
should stay outside git; promote durable decisions, references, or debugging
breadcrumbs into typed wiki pages instead.

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
