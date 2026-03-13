## Context

The user reported that Kiro also had model deletions. Historical review of `GetKiroModels()` in `internal/registry/model_definitions.go` shows the current registry still contains nearly all historical Kiro models, except for one chat-only variant that disappeared during a Kiro refactor.

## Findings

- Kiro models are defined in `internal/registry/model_definitions.go`, not in `internal/registry/models/models.json`.
- Historical Kiro IDs used both dotted (`4.5`) and normalized dashed (`4-5`) naming forms during refactors.
- After normalizing away dotted-to-dashed renames, the only clearly missing historical Kiro model is `kiro-claude-opus-4.5-chat`.
- The model was removed in commit `204bba9d` during a Kiro refactor and was not replaced with an equivalent chat-only variant.

## Goals / Non-goals

Goals:
- Restore the missing Kiro Claude Opus 4.5 chat-only variant.
- Use the current normalized ID style to fit the existing registry.

Non-goals:
- Do not revert Kiro naming normalization from dotted to dashed IDs.
- Do not change the current Kiro executor behavior or dynamic model conversion logic.
- Do not add speculative Kiro models without clear historical evidence.

## Target files / modules

- `internal/registry/model_definitions.go`

## Constraints

- Keep the diff minimal and consistent with neighboring Kiro model definitions.
- Preserve current Kiro naming conventions (`4-5`, `4-6`, etc.).

## Verification plan

- Confirm the restored model ID is present in `GetKiroModels()`.
- Run a focused text scan for the restored ID.
- Run `git diff --check`.

## Rollback

- Remove the restored `kiro-claude-opus-4-5-chat` entry from `GetKiroModels()`.

## Open questions

- None for this bounded recovery task.

## Execution Status

- Completed: compared historical and current Kiro model ID sets.
- Completed: restored `kiro-claude-opus-4-5-chat` in `GetKiroModels()`.

## Evidence

- `awk '/func GetKiroModels\\(\\)/,/^}/' /root/work/CLIProxyAPIPlus/internal/registry/model_definitions.go | rg 'kiro-claude-opus-4-5-chat|kiro-claude-opus-4-5|kiro-claude-opus-4-5-agentic'`
- `git -C /root/work/CLIProxyAPIPlus diff --check`
- `bash /root/.openclaw/workspace/skills/repo-task-driven/scripts/doc_placeholder_scan.sh /root/work/CLIProxyAPIPlus/docs/kiro-model-recovery`
- `bash /root/.openclaw/workspace/skills/repo-task-driven/scripts/post_refactor_text_scan.sh /root/work/CLIProxyAPIPlus/docs/kiro-model-recovery /root/work/CLIProxyAPIPlus/README.md`
