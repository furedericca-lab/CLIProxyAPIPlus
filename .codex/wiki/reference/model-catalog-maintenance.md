---
title: Model Catalog Maintenance
type: reference
status: current
scope: docs-migration
related_files:
  - internal/registry/models/models.json
  - internal/registry/model_definitions.go
  - internal/registry/model_updater.go
  - internal/thinking/provider/iflow/apply.go
  - third_party/models
tags:
  - models
  - iflow
  - kiro
  - registry
last_checked: 2026-05-28
updated: 2026-05-28T13:20:00Z
---

# Model Catalog Maintenance

This page preserves useful conclusions from the old `docs/iflow-kimi-k2-5`, `docs/kiro-model-recovery`, `docs/models-fork-submodule`, and `docs/models-fork-sync` scopes.

## iFlow `kimi-k2.5`

Historical finding:

- `kimi-k2.5` previously appeared under iFlow static data.
- The active runtime catalog source was `internal/registry/models/models.json`.
- `kimi-k2.5` also existed under the standalone `kimi` provider.
- Restoring it under iFlow required both catalog metadata and iFlow Kimi fallback thinking handling.

Recorded implementation from the old scope:

- Restored `kimi-k2.5` to `internal/registry/models/models.json` under `iflow`.
- Restored iFlow Kimi fallback thinking in `internal/thinking/provider/iflow/apply.go`.
- Added focused tests in `internal/thinking/provider/iflow/apply_test.go`.

Verification commands from the old scope:

```bash
jq -e '.iflow | map(.id) | index("kimi-k2.5") != null' internal/registry/models/models.json
jq -e '.iflow[] | select(.id == "kimi-k2.5") | .thinking.levels | index("high") != null' internal/registry/models/models.json
jq -e '.' internal/registry/models/models.json
git diff --check
```

Stale environment note from the old scope:

- At that time, `go test ./internal/thinking/provider/iflow` was blocked because the installed Go toolchain rejected `go 1.26.0` in `go.mod`.

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

- Re-check whether `third_party/models` still exists and is desired.
- Re-check whether `internal/registry/model_updater.go` still supports a forked source.
- Re-check whether Kiro models are still code-defined.
- Re-check whether iFlow thinking behavior is still implemented in the same provider package.
