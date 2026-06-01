---
title: Source Docs Archive Map
type: reference
status: current
scope: docs-consolidation
related_scopes: []
related_files:
  - .codex/wiki
  - README.md
  - AGENTS.md
source_docs:
  - docs
tags:
  - docs
  - wiki
  - consolidation
last_checked: 2026-06-01
updated: 2026-06-01T03:05:26Z
---

# Source Docs Archive Map

## Scope

This page maps documentation surfaces that were intentionally consolidated into
the project wiki for this maintained CLIProxyAPIPlus fork.

The current repository intentionally does not maintain a user-facing `docs/`
tree. Long-lived project knowledge belongs in `.codex/wiki/**`, while active
and archived task evidence belongs in `.codex/scopes/**`.

## Current Surfaces

- `README.md`: short operator entry point for repository identity, commands,
  and links to current wiki pages.
- `AGENTS.md`: coding-agent contract, provider preservation rules, validation
  defaults, and source-of-truth precedence.
- `.codex/wiki/**`: durable detailed knowledge for upstream strategy,
  provider inventory, architecture notes, implementation notes, and
  maintenance history.
- `.codex/scopes/<scope>/`: active scope contracts, plans, and evidence.
- `.codex/scopes/archive/<scope>/`: completed scope records.

## Consolidation Map

- Former upstream or generated user-facing `docs/` content: do not restore as a
  local docs tree. Port useful facts into typed wiki pages.
- Upstream root docs such as `README*`, `AGENTS.md`, and `CLAUDE.md`: do not
  sync blindly. Manually port only facts that affect current commands,
  compatibility, or runtime behavior.
- Upstream `.codex/scopes/**`: do not merge as local documentation. If an
  upstream scope reveals a useful maintenance rule, summarize it in the wiki.
- Local scope records: keep in `.codex/scopes/**`; do not duplicate their full
  evidence in README or AGENTS.

## Primary Wiki Pages

- `.codex/wiki/decisions/track-router-main-as-upstream.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/reference/local-provider-and-commit-inventory.md`
- `.codex/wiki/concepts/codebase-function-map.md`

## Maintenance Rule

When README, AGENTS, scope records, or upstream strategy changes affect durable
knowledge, update the relevant wiki page in the same pass. Then run:

```bash
python3 /root/.codex/skills/wiki-note/scripts/wiki.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki.py doctor --json
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy lint
python3 /root/.codex/skills/wiki-note/scripts/wiki.py legacy surface-check --json
```
