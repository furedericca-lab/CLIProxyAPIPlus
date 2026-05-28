# AGENTS.md

This is the local operator contract for our maintained CLIProxyAPIPlus fork.

## Project Identity

This repository is not a plain mirror of upstream CLIProxyAPI or any upstream Plus fork. It is our maintained Plus line.

Current upstream policy:

- `upstream` points to `https://github.com/HsnSaboor/CLIProxyAPIPlus`.
- `router` points to `https://github.com/router-for-me/CLIProxyAPI`.
- HsnSaboor is the first convergence baseline.
- Router is a later selective patch source, not a direct merge target.

When local old commits overlap with HsnSaboor by provider or feature, prefer HsnSaboor first. Reintroduce local behavior only after current-code validation proves a gap.

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
- `README_CN.md`
- `README_JA.md`
- `AGENTS.md`
- `CLAUDE.md`

Do not blindly sync these files from upstream. If upstream changes contain useful command or compatibility information, manually port the relevant facts into our local docs or wiki.

The same rule applies to upstream `docs/**`: do not merge it into this repo as documentation content.

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

- Preserve HsnSaboor's Plus provider surface.
- Accept HsnSaboor additions such as `cline`, `xai`, and `ollama` executor support.
- Reject router provider deletion hunks by default.
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

For the current clean-root effort, use:

- `docs/hsnsaboor-clean-root/`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`
