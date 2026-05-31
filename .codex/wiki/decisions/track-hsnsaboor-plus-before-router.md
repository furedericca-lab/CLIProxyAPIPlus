---
title: Track HsnSaboor Plus Before Router Core
type: decision
status: superseded
scope: hsnsaboor-clean-root
related_scopes:
  []
related_files:
  - docs/archive/hsnsaboor-clean-root/hsnsaboor-clean-root-contract.md
  - .codex/wiki/reference/upstream-plus-maintenance.md
tags:
  - upstream
  - providers
  - merge-strategy
decision_date: 2026-05-28
last_checked: 2026-05-28
updated: 2026-05-31T05:13:53Z
---

# Track HsnSaboor Plus Before Router Core

Superseded on 2026-05-31 by
`.codex/wiki/decisions/track-router-main-as-upstream.md`.

## Decision

Use `https://github.com/HsnSaboor/CLIProxyAPIPlus` as the first convergence target for this fork. Treat `https://github.com/router-for-me/CLIProxyAPI` as a later selective patch source, not as the direct merge target.

## Rationale

The router branch is newer but removed many Plus providers. HsnSaboor is behind router after `v7.1.9`, but it preserves the Plus provider intent and adds missing local pieces such as Cline, Ollama executor support, and xAI.

Directly merging router risks deleting provider/auth/executor functionality that this fork exists to keep.

## Consequences

- The completed clean-root scope aligned `main` with HsnSaboor and did not replay old local commits wholesale.
- Later router updates should be cherry-picked or manually ported with provider-deletion hunks excluded.
- Provider surface checks are mandatory before declaring any upstream integration complete.

Refinement on 2026-05-28:

- Use HsnSaboor implementations as the default when local old commits overlap by provider or feature.
- Do not replay local old commits wholesale.
- Build a clean branch from `upstream/main`, then add only our current maintenance commits on top.
- Audit old local commits under `docs/archive/old-local-commit-audit/` before reintroducing any missing behavior.
