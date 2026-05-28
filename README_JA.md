# CLIProxyAPIPlus Maintained Fork

This repository is our maintained CLIProxyAPIPlus line. It is not an upstream documentation mirror.

Maintenance policy:

- Use `HsnSaboor/CLIProxyAPIPlus` as the Plus baseline.
- Use `router-for-me/CLIProxyAPI` only as a later selective patch source.
- Do not merge router provider deletions blindly.
- Do not sync upstream `README*`, `AGENTS.md`, `CLAUDE.md`, or `docs/**` content into this repository.

## Documentation Boundaries

- `docs/<scope>/`: active scope contracts, plans, checklists, and evidence.
- `docs/archive/<scope>/`: archived completed scopes.
- `.codex/wiki/**`: durable maintainer knowledge, decisions, and implementation notes.

`docs/` remains available for our repo-task-driven scopes and archives. It should not be used to mirror upstream documentation.

## Completed Scope

- `docs/archive/hsnsaboor-clean-root/`

This completed scope built a clean branch rooted at HsnSaboor `upstream/main`, then added only our current maintenance decisions as new commits.

## Commands

```bash
gofmt -w .
go build -o cli-proxy-api ./cmd/server
go build -o test-output ./cmd/server && rm test-output
go test ./...
```

See also:

- `README.md`
- `AGENTS.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
