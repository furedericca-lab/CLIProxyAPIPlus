# AGENTS.md

This is the local operator contract for our maintained CLIProxyAPIPlus fork.

## Project Identity

This repository is not a plain mirror of upstream CLIProxyAPI or any upstream Plus fork. It is our maintained Plus line.

Current upstream policy:

- `upstream` points to `https://github.com/router-for-me/CLIProxyAPI`.
- `router` also points to `https://github.com/router-for-me/CLIProxyAPI` as an
  explicit compatibility alias for older local commands.
- Router `main` is the active upstream baseline.
- This fork owns the Plus adaptation layer on top of router.

When router changes overlap with local Plus provider behavior, prefer the
current router implementation for core behavior, then reintroduce or adapt Plus
provider support only after current-code validation proves the gap.

Provider precedence:

- For providers that already exist in router, use router's implementation as
  the source of truth.
- For providers that exist only in this fork or the former HsnSaboor Plus line,
  keep them as local Plus providers, use HsnSaboor's maintenance line as their
  update reference, and adapt them to router core changes.

## Documentation Boundaries

Use `docs/` for repo-task-driven scope work:

- Active scopes: `docs/<scope>/`
- Archived scopes: `docs/archive/<scope>/`
- Scope contracts, phase plans, checklists, evidence, and archive records belong here.

Do not put upstream docs in `docs/`. The previous cleanup removed old upstream/user-facing docs so the directory can be used for our scope workflow.

Use `.codex/wiki/**` for durable local project knowledge:

- decisions
- maintenance notes
- provider inventories
- architecture breadcrumbs
- implementation notes
- lessons learned

Keep scratch notes out of git, for example `.codex/notepad.md`.

## Root Docs Are Locally Owned

These paths are local project identity and operator policy:

- `README.md`
- `AGENTS.md`
- `CLAUDE.md`

Do not blindly sync these files from upstream. If upstream changes contain useful command or compatibility information, manually port the relevant facts into our local docs or wiki.

The same rule applies to upstream `docs/**`: do not merge it into this repo as documentation content.

`README_CN.md` and `README_JA.md` are intentionally absent. Do not recreate them from upstream during synchronization.

The merge policy is recorded in `.gitattributes`.

On a fresh clone, enable the local merge driver once:

```bash
git config merge.ours.driver true
```

## Provider Preservation

The Plus provider surface is the main reason this fork exists.

Before and after any upstream integration, capture:

```bash
git ls-tree -d --name-only HEAD:internal/auth | sort
git ls-tree --name-only HEAD:internal/runtime/executor | rg '_executor\.go$' | sed 's/_executor\.go$//' | sort
git ls-tree --name-only HEAD:internal/cmd | rg '(_login|_cookie|vertex_import)\.go$' | sort
```

Expected strategy:

- Preserve the local Plus provider surface unless a scope explicitly retires a
  provider.
- Use router's provider code for providers router already owns.
- Keep local/HsnSaboor-exclusive providers as Plus extensions and update them
  from HsnSaboor's maintenance line when available.
- Treat router provider deletions or API moves as adaptation work, not as
  automatic permission to drop Plus behavior.
- Accept router core fixes and protocol updates after provider-surface review.
- Do not replay old local provider patches unless tests or code review prove they are still needed.

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

## Code Conventions

- Keep changes small and simple.
- Use `gofmt` after Go edits.
- Comments in Go code should be English.
- Keep user-visible strings in the language already used by that area.
- Avoid `log.Fatal` and `log.Fatalf`; return errors or log with logrus.
- Wrap defer errors when cleanup can fail.
- Avoid leaking credentials, tokens, cookies, or auth material in logs, docs, diffs, final answers, and wiki.
- Do not add standalone `internal/translator/` changes unless repository policy and permissions allow it.
- Keep `internal/runtime/executor/` limited to executors and executor tests. Helper files belong under `internal/runtime/executor/helps/`.

## Scope Workflow

For non-trivial repo work:

1. Create or update a scope under `docs/<scope>/`.
2. Keep long-lived knowledge aligned in `.codex/wiki/**`.
3. Record evidence commands and results in the scope checklist.
4. Run the narrowest meaningful verification, then broader checks when risk justifies it.
5. Do not claim completion without matching evidence.

For the completed clean-root effort, use:

- `docs/archive/hsnsaboor-clean-root/`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`

For the completed router-first rebaseline decision, use:

- `docs/archive/router-upstream-rebaseline/`
- `.codex/wiki/decisions/track-router-main-as-upstream.md`

For the completed router batch sync through `v7.1.32`, use:

- `docs/archive/router-batch-sync/`
