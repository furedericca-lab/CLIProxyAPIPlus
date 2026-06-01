# CLIProxyAPIPlus Maintained Fork

This repository is our maintained Plus line for CLIProxyAPI-compatible local
proxy work. It tracks the active router core while preserving the Plus provider
surface that is unique to this fork or the former HsnSaboor Plus line.

Current maintenance policy:

- `upstream` and `router` point to `https://github.com/router-for-me/CLIProxyAPI`.
- Router `main` is the active upstream baseline for core behavior.
- Providers that exist in router follow router's implementation.
- Local or HsnSaboor-exclusive Plus providers remain local extensions unless a
  dedicated scope retires them.
- HsnSaboor remains the reference line for Plus-only provider behavior when it
  has relevant maintenance.

For the detailed policy, use the project wiki:

- `.codex/wiki/index.md`
- `.codex/wiki/decisions/track-router-main-as-upstream.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`

## Repository Workflow

Use the root docs as entry points, not as full manuals:

- `README.md`: this operator entry point.
- `AGENTS.md`: coding-agent contract and validation rules.
- `CLAUDE.md`: local assistant compatibility policy.
- `.codex/wiki/**`: durable decisions, maintenance notes, architecture
  breadcrumbs, provider inventory, and lessons learned.
- `.codex/scopes/<scope>/`: active task contracts, plans, checklists, and
  evidence.
- `.codex/scopes/archive/<scope>/`: completed scope records.

Do not restore upstream root docs, upstream `.codex/scopes/**`, or upstream
`docs/` content by default. Useful upstream facts should be manually ported into
the local README, agent contract, or wiki.

`README_CN.md` and `README_JA.md` are intentionally absent.

## Provider Surface Check

Before and after upstream integration, capture the provider surface:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Treat router provider deletions, API moves, or behavior changes as adaptation
work. Do not drop Plus providers as an incidental merge side effect.

## Development Commands

```bash
gofmt -w .
go build -o cli-proxy-api ./cmd/server
go build -o test-output ./cmd/server && rm test-output
go test ./...
go test -v -run TestName ./path/to/pkg
```

The compile check is required after Go changes:

```bash
go build -o test-output ./cmd/server && rm test-output
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

## Local Merge Policy

The merge policy for locally owned docs is recorded in `.gitattributes`.

On a fresh clone, enable the local merge driver once:

```bash
git config merge.ours.driver true
```
