---
title: Track Router Main As Upstream
type: decision
status: accepted
scope: router-upstream-rebaseline
related_scopes: []
related_files:
  - AGENTS.md
  - README.md
  - .codex/wiki/reference/upstream-plus-maintenance.md
source_docs: []
tags:
  - upstream
  - router
  - providers
last_checked: 2026-05-31
updated: 2026-05-31T05:13:53Z
decision_date: 2026-05-31
---

# Track Router Main As Upstream

## Decision

Use `https://github.com/router-for-me/CLIProxyAPI` as the primary `upstream`
remote for this maintained Plus fork. HsnSaboor is no longer treated as the
active maintenance baseline.

## Rationale

HsnSaboor has not shown recent maintenance, while router continues moving. This
fork will track router core changes directly and carry the Plus compatibility
work locally.

## Consequences

- `upstream/main` now resolves to router `main`.
- The existing `router` remote remains as a compatibility alias for older local
  commands and archive evidence.
- Future integration work should adapt router updates into this fork while
  preserving Plus provider behavior through explicit review and local patches.
- Provider deletion hunks are rejected by default unless a scope explicitly
  decides to retire the provider.
- Providers that already exist in router should use router's implementation as
  the baseline.
- Providers that exist only in this fork or the former HsnSaboor Plus line
  should remain local Plus extensions, use HsnSaboor's maintenance line as their
  update reference when available, and be adapted to router core changes.
- Historical HsnSaboor clean-root and audit scopes remain archived as
  provenance, not current policy.
