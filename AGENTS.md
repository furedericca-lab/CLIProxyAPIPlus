# AGENTS.md

This is the local coding-agent contract for our maintained CLIProxyAPIPlus
fork. This repository is not a plain mirror of upstream CLIProxyAPI or any
upstream Plus fork.

## Source Of Truth

Use these surfaces in this order:

1. Newest user instruction in the current conversation.
2. This `AGENTS.md` for coding-agent boundaries and validation defaults.
3. `.codex/wiki/index.md` for durable project knowledge.
4. `.codex/scopes/<scope>/` and `.codex/scopes/archive/<scope>/` for scoped
   plans, evidence, and historical implementation records.
5. Current source code and runtime checks.

Root docs are locally owned project identity and policy:

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`

Do not blindly sync these files from upstream. If upstream changes contain
useful command, compatibility, or runtime facts, manually port the facts into
the local README, this contract, or `.codex/wiki/**`.

`README_CN.md` and `README_JA.md` are intentionally absent. Do not recreate
them from upstream during synchronization.

## Upstream And Provider Policy

Current upstream policy:

- `upstream` points to `https://github.com/router-for-me/CLIProxyAPI`.
- `router` points to the same URL as a compatibility alias.
- Router `main` is the active upstream baseline.
- This fork owns the Plus adaptation layer on top of router.

Provider precedence:

- Router-owned providers use router's implementation as the source of truth.
- Local or HsnSaboor-exclusive providers stay as Plus extensions unless a scope
  explicitly retires them.
- HsnSaboor remains the update reference for Plus-only provider behavior when
  it has relevant maintenance.
- Router provider deletions or API moves are adaptation work, not automatic
  permission to drop Plus behavior.

Current detailed policy lives in:

- `.codex/wiki/decisions/track-router-main-as-upstream.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`

## Documentation And Wiki Policy

`.codex/wiki/**` is the durable knowledge layer. Use it for decisions,
maintenance notes, provider inventories, architecture breadcrumbs,
implementation notes, and lessons learned.

`README.md` and `AGENTS.md` should summarize and route; they should not become
competing manuals for the same facts. When long-lived details change, update
the wiki first or in the same pass.

Do not restore upstream `.codex/scopes/**` or upstream `docs/` content as
local documentation. If an upstream doc contains useful facts, absorb those
facts into the wiki and rebuild the wiki index. The docs consolidation map is:

- `.codex/wiki/reference/source-docs-archive-map.md`

Keep scratch notes out of git. Promote durable decisions, references, or
debugging breadcrumbs into typed wiki pages instead.

## Provider Preservation

Before and after any upstream integration, capture:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Accept router core fixes and protocol updates only after provider-surface
review. Do not replay old local provider patches unless tests or code review
prove they are still needed.

## Commands

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

Wiki checks after documentation or scope changes:

```bash
python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json
```

## Code Conventions

- Keep changes small and simple.
- Use `gofmt` after Go edits.
- Comments in Go code should be English.
- Keep user-visible strings in the language already used by that area.
- Avoid `log.Fatal` and `log.Fatalf`; return errors or log with logrus.
- Wrap defer errors when cleanup can fail.
- Avoid leaking credentials, tokens, cookies, or auth material in logs, docs,
  diffs, final answers, and wiki.
- Do not add standalone `internal/translator/` changes unless repository
  policy and permissions allow it.
- Keep `internal/runtime/executor/` limited to executors and executor tests.
  Helper files belong under `internal/runtime/executor/helps/`.

## Scope Workflow

For non-trivial repo work:

1. Create or update a scope under `.codex/scopes/<scope>/`.
2. Keep long-lived knowledge aligned in `.codex/wiki/**`.
3. Record evidence commands and results in the scope checklist.
4. Run the narrowest meaningful verification, then broader checks when risk
   justifies it.
5. Rebuild and lint the wiki after scope or durable-doc changes.
6. Do not claim completion without matching evidence.

Completed scope anchors:

- `.codex/scopes/archive/hsnsaboor-clean-root/`
- `.codex/scopes/archive/router-upstream-rebaseline/`
- `.codex/scopes/archive/router-batch-sync/`
