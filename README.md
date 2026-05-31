# CLIProxyAPIPlus Maintained Fork

This repository is our maintained Plus line for CLIProxyAPI-compatible local proxy work.

The project goal is to keep the Plus provider surface while tracking the
actively maintained router core. The current policy is:

- Use `router-for-me/CLIProxyAPI` as `upstream`.
- Maintain our Plus provider adaptation layer locally.
- For providers router already has, use router's provider implementation as the
  baseline.
- For providers only present in this fork or the former HsnSaboor Plus line,
  keep them, use HsnSaboor's maintenance line as their update reference, and
  adapt them locally.
- Do not blindly accept provider deletions or incompatible router API moves.
- Do not sync upstream root docs or upstream `docs/` content into this repository.

## Maintenance Model

Root documentation is locally owned:

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`

These files describe our fork, not upstream marketing or upstream operator policy. Upstream changes to these paths should be reviewed manually only when they affect real commands, compatibility, or runtime behavior.

Repository task documentation lives under `docs/<scope>/` and completed scopes may move to `docs/archive/<scope>/`. This directory is for our active scope plans, contracts, evidence, and archived scope records. It is not a place to mirror upstream documentation.

Durable maintainer knowledge lives under `.codex/wiki/**`. Use it for decisions, maintenance notes, architecture breadcrumbs, provider inventories, and lessons learned that should survive context compaction.

## Completed Scope

The clean-root maintenance scope has been completed and archived:

- `docs/archive/hsnsaboor-clean-root/`

It created a clean branch rooted at HsnSaboor `upstream/main`, then added only our current maintenance decisions as new commits.

The upstream rebaseline scope has been completed and archived:

- `docs/archive/router-upstream-rebaseline/`

The router batch sync through `v7.1.32` has been completed and archived:

- `docs/archive/router-batch-sync/`

## Upstream Policy

Configured remotes:

- `origin`: this fork
- `upstream`: `https://github.com/router-for-me/CLIProxyAPI`
- `router`: `https://github.com/router-for-me/CLIProxyAPI` as a compatibility alias

Rules:

- Router implementations win by default for core behavior and protocol changes.
- Router provider code wins for providers that exist in router.
- Plus provider behavior is preserved through explicit local adaptation.
- Local/HsnSaboor-exclusive providers remain part of this fork unless a scope
  explicitly retires them, and HsnSaboor remains their update reference when it
  has relevant maintenance.
- Local old commits are not replayed wholesale.
- Provider deletion hunks are rejected by default unless a scope explicitly
  retires the provider.

## Provider Preservation

Before and after upstream integration, capture the provider surface:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Current `internal/auth` split:

- Router-owned providers: `antigravity`, `claude`, `codex`, `empty`, `gemini`,
  `kimi`, `vertex`, `xai`.
- Local/HsnSaboor-exclusive Plus providers to preserve: `cline`, `codebuddy`,
  `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`, `qwen`.

## Development Commands

```bash
gofmt -w .
go build -o cli-proxy-api ./cmd/server
go build -o test-output ./cmd/server && rm test-output
go test ./...
go test -v -run TestName ./path/to/pkg
```

Common server flags:

```bash
--config <path>
--tui
--standalone
--local-model
--no-browser
--oauth-callback-port <port>
```

## Documentation Rules

- Put active scope contracts and task plans in `docs/<scope>/`.
- Put completed scope records in `docs/archive/<scope>/`.
- Put durable maintainer knowledge in `.codex/wiki/**`.
- Do not restore upstream `docs/*` during sync.
- Do not restore upstream `README*`, `AGENTS.md`, or `CLAUDE.md` during sync.
- `README_CN.md` and `README_JA.md` are intentionally absent.

The merge policy for these owned paths is recorded in `.gitattributes`.

On a fresh clone, enable the local merge driver once:

```bash
git config merge.ours.driver true
```
