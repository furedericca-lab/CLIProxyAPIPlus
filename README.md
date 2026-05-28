# CLIProxyAPIPlus Maintained Fork

This repository is our maintained Plus line for CLIProxyAPI-compatible local proxy work.

The project goal is to keep the Plus provider surface while using the best maintained upstream base available. The current policy is:

- Use `HsnSaboor/CLIProxyAPIPlus` as the Plus baseline.
- Treat `router-for-me/CLIProxyAPI` as a later selective patch source.
- Do not blindly merge provider deletions from router.
- Do not sync upstream root docs or upstream `docs/` content into this repository.

## Maintenance Model

Root documentation is locally owned:

- `README.md`
- `README_CN.md`
- `README_JA.md`
- `AGENTS.md`
- `CLAUDE.md`

These files describe our fork, not upstream marketing or upstream operator policy. Upstream changes to these paths should be reviewed manually only when they affect real commands, compatibility, or runtime behavior.

Repository task documentation lives under `docs/<scope>/` and completed scopes may move to `docs/archive/<scope>/`. This directory is for our active scope plans, contracts, evidence, and archived scope records. It is not a place to mirror upstream documentation.

Durable maintainer knowledge lives under `.codex/wiki/**`. Use it for decisions, maintenance notes, architecture breadcrumbs, provider inventories, and lessons learned that should survive context compaction.

## Completed Scope

The clean-root maintenance scope has been completed and archived:

- `docs/archive/hsnsaboor-clean-root/`

It created a clean branch rooted at HsnSaboor `upstream/main`, then added only our current maintenance decisions as new commits.

## Upstream Policy

Configured remotes:

- `origin`: this fork
- `upstream`: `https://github.com/HsnSaboor/CLIProxyAPIPlus`
- `router`: `https://github.com/router-for-me/CLIProxyAPI`

Rules:

- HsnSaboor implementations win by default when local old commits overlap by provider or feature.
- Local old commits are not replayed wholesale.
- Router changes are reviewed later as targeted core patches.
- Router provider deletion hunks are rejected by default.

## Provider Preservation

Before and after upstream integration, capture the provider surface:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

The HsnSaboor baseline should preserve Plus providers and currently adds `cline`, `xai`, and `ollama` executor support beyond the old local line.

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

The merge policy for these owned paths is recorded in `.gitattributes`.

On a fresh clone, enable the local merge driver once:

```bash
git config merge.ours.driver true
```
