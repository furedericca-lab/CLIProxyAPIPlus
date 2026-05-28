---
title: Model Catalog Maintenance
type: reference
status: current
scope: risk-resolution-from-old-local-audit
related_files:
  - internal/registry/models/models.json
  - internal/registry/model_definitions.go
  - internal/registry/model_updater.go
  - internal/thinking/provider/iflow/apply.go
tags:
  - models
  - iflow
  - kiro
  - registry
last_checked: 2026-05-28
updated: 2026-05-28T14:35:00Z
---

# Model Catalog Maintenance

This page preserves useful conclusions from the old `docs/iflow-kimi-k2-5`, `docs/kiro-model-recovery`, `docs/models-fork-submodule`, and `docs/models-fork-sync` scopes, updated for the current clean-root implementation.

## Current model source policy

Runtime model refresh uses the maintained fork source:

- `https://raw.githubusercontent.com/furedericca-lab/models/refs/heads/main/models.json`

Release and PR build workflows no longer rewrite `internal/registry/models/models.json` from an external repository. Release artifacts use the catalog committed in the release tag. This avoids hidden dirty checkout state during GoReleaser and prevents accidental drift back to `router-for-me/models`.

The old `third_party/models` submodule approach is not restored in current `main`.

## iFlow `kimi-k2.5`

Historical finding:

- `kimi-k2.5` previously appeared under iFlow static data.
- The active runtime catalog source was `internal/registry/models/models.json`.
- `kimi-k2.5` also existed under the standalone `kimi` provider.
- Restoring it under iFlow required both catalog metadata and iFlow Kimi fallback thinking handling.

Current implementation:

- iFlow Kimi fallback thinking is implemented in `internal/thinking/provider/iflow/apply.go`.
- Focused tests are in `internal/thinking/provider/iflow/apply_test.go`.
- `kimi-*` iFlow models with `Thinking` support use `chat_template_kwargs.enable_thinking`; GLM remains the only family that also sets `clear_thinking`.

Current verification commands:

```bash
go test ./internal/thinking/provider/iflow ./internal/thinking
```

## Kiro model recovery

Historical finding:

- Kiro models were defined in `internal/registry/model_definitions.go`, not `internal/registry/models/models.json`.
- After normalizing dotted-to-dashed model IDs, the missing historical model was `kiro-claude-opus-4.5-chat`.
- Current naming style used dashed IDs.

Recorded implementation from the old scope:

- Restored `kiro-claude-opus-4-5-chat` in `GetKiroModels()`.

Verification command from the old scope:

```bash
awk '/func GetKiroModels\(\)/,/^}/' internal/registry/model_definitions.go | rg 'kiro-claude-opus-4-5-chat|kiro-claude-opus-4-5|kiro-claude-opus-4-5-agentic'
```

## Forked models source

Historical decision:

- The release workflow was changed away from `https://github.com/router-for-me/models.git`.
- The intended fork source was `https://github.com/furedericca-lab/models.git`.
- A submodule was added at `third_party/models` on branch `main`.
- The release workflow copied `third_party/models/models.json` into `internal/registry/models/models.json`.

Current implementation differs from the old submodule design:

- `internal/registry/model_updater.go` points directly to `furedericca-lab/models`.
- `.github/workflows/release.yaml` and `.github/workflows/pr-test-build.yml` do not fetch a model catalog during CI.
- `internal/registry/model_updater_test.go` prevents runtime model URLs from drifting back to router sources.

Follow-up historical finding:

- The submodule schema only contained `models.json`.
- iFlow could be synchronized into the submodule.
- Kiro could not be mirrored into the submodule because Kiro definitions lived in Go code, not in the submodule schema.

Verification commands from the old scopes:

```bash
git submodule status
jq -e '.' third_party/models/models.json
diff -u <(jq -r '.iflow[].id' internal/registry/models/models.json) <(jq -r '.iflow[].id' third_party/models/models.json)
```

## Current convergence risk

These model/catalog decisions may conflict with HsnSaboor/router registry changes. During upstream convergence:

- Keep `internal/registry/model_updater.go` on the forked source unless a new scope deliberately changes model catalog ownership.
- Re-check whether Kiro models are still code-defined.
- Re-check whether iFlow thinking behavior is still implemented in the same provider package.
