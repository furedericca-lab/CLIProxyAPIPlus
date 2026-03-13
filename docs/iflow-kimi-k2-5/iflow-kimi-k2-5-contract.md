## Context

`kimi-k2.5` previously appeared in the repository's iFlow model catalog during the pre-migration static data era, but it is absent from the active network-refreshed catalog in `internal/registry/models/models.json`. The user requested restoring `kimi-k2.5` under `iflow`.

## Findings

- The active runtime catalog source is `internal/registry/models/models.json`.
- Git history shows `kimi-k2.5` previously existed under iFlow static data.
- The current catalog still exposes `kimi-k2.5` under the standalone `kimi` provider.
- The current iFlow thinking applier does not define provider-specific thinking behavior for Kimi models, so restoring the model entry should not claim unsupported thinking semantics.

## Goals / Non-goals

Goals:
- Restore `kimi-k2.5` to the active `iflow` model catalog.
- Keep the change minimal and auditable.

Non-goals:
- Do not change standalone `kimi` provider definitions.
- Do not introduce new iFlow-specific thinking behavior for Kimi models without protocol evidence.
- Do not regenerate or refactor the full catalog pipeline.

## Target files / modules

- `internal/registry/models/models.json`

## Constraints

- Keep JSON valid and preserve existing catalog structure.
- Avoid adding metadata that suggests unsupported runtime behavior.
- Use a minimal diff consistent with neighboring iFlow model entries.

## Verification plan

- Query the `iflow` model IDs from `internal/registry/models/models.json` and confirm `kimi-k2.5` is present.
- Run a JSON validity check on `internal/registry/models/models.json`.
- Run `git diff --check` to catch formatting or whitespace issues.

## Rollback

- Remove the restored `kimi-k2.5` object from the `iflow` array in `internal/registry/models/models.json`.

## Open questions

- Whether iFlow currently supports Kimi-specific thinking controls remains unverified and is intentionally left unchanged in this task.

## Implementation Notes

- Initial implementation restores the iFlow catalog entry only.
- Follow-up implementation restores iFlow Kimi fallback thinking behavior by:
  - re-adding `thinking` metadata to `iflow/kimi-k2.5`
  - routing Kimi model IDs through the iFlow `enable_thinking` fallback branch
  - adding focused unit coverage for the Kimi fallback applier path

## Execution Status

- Completed: restored `kimi-k2.5` to the active `iflow` catalog in `internal/registry/models/models.json`.
- Completed: restored iFlow Kimi fallback thinking handling in `internal/thinking/provider/iflow/apply.go`.
- Completed: added focused Kimi fallback tests in `internal/thinking/provider/iflow/apply_test.go`.
- Completed: verified repository task docs with the repo-task-driven validation scripts.

## Evidence

- `bash /root/.openclaw/workspace/skills/repo-task-driven/scripts/doc_placeholder_scan.sh /root/work/CLIProxyAPIPlus/docs/iflow-kimi-k2-5`
- `bash /root/.openclaw/workspace/skills/repo-task-driven/scripts/post_refactor_text_scan.sh /root/work/CLIProxyAPIPlus/docs/iflow-kimi-k2-5 /root/work/CLIProxyAPIPlus/README.md`
- `jq -e '.iflow | map(.id) | index("kimi-k2.5") != null' /root/work/CLIProxyAPIPlus/internal/registry/models/models.json`
- `jq -e '.iflow[] | select(.id == "kimi-k2.5") | .thinking.levels | index("high") != null' /root/work/CLIProxyAPIPlus/internal/registry/models/models.json`
- `jq -e '.' /root/work/CLIProxyAPIPlus/internal/registry/models/models.json`
- `git -C /root/work/CLIProxyAPIPlus diff --check`

## Verification Gaps

- `go test ./internal/thinking/provider/iflow` is currently blocked in this environment because `/root/work/CLIProxyAPIPlus/go.mod` declares `go 1.26.0`, and the installed Go toolchain rejects that version string during module parsing.
