---
description: Risk audit for the eight pre-clean-root local commits that were not replayed on top of HsnSaboor.
---

# Old Local Commit Audit Contract

## Context

The clean-root history now uses HsnSaboor `upstream/main` as the Plus baseline and keeps the previous local line under `backup/main-before-hsnsaboor-clean-root`.

This scope audits the eight old local non-merge commits that are not in HsnSaboor:

```bash
git log --oneline --no-merges upstream/main..backup/main-before-hsnsaboor-clean-root
```

Result:

| Commit | Subject |
| --- | --- |
| `044678b0` | `fix(copilot): route claude models through native messages` |
| `2dc61ac6` | `ci(release): add manual workflow trigger` |
| `566bfb69` | `point model updater to forked models source` |
| `1aee8779` | `sync model catalogs and forked models source` |
| `ad7d7999` | `ci(release): avoid dirty git state in goreleaser workflow` |
| `e3d123f8` | `ci: skip docker steps when DockerHub secrets are missing` |
| `dd9f77fc` | `ci: skip docker-image when DockerHub secrets are missing` |
| `2a1cf2b3` | `Revert "fix(registry): clean up outdated model definitions in static data"` |

## Findings

### [High] Fork model catalog source was not carried forward

Evidence:

- Current runtime updater still fetches router model catalogs from `internal/registry/model_updater.go:22`.
- Current release workflow refreshes `internal/registry/models/models.json` from `https://github.com/router-for-me/models.git` at `.github/workflows/release.yaml:19`.
- Current PR build workflow does the same at `.github/workflows/pr-test-build.yml:15`.
- Old commit `566bfb69` changed the runtime source to `https://raw.githubusercontent.com/furedericca-lab/models/refs/heads/main/models.json`.
- Old commit `1aee8779` changed release refresh to a `third_party/models` submodule pinned to `https://github.com/furedericca-lab/models.git`.

Impact:

The current clean-root line may silently drift back to router's model catalog even though this fork's provider policy is HsnSaboor-first and Plus-preserving. This is the highest risk because it affects runtime model availability and release artifacts.

Assessment:

Do not replay the old submodule approach as-is. It added `.gitmodules`, `third_party/models`, and old docs in one broad commit. Rebuild this as a small current-code change if we still want the fork-owned catalog source.

Recommended fix scope:

- Decide whether `furedericca-lab/models` is still the intended model source.
- If yes, update both runtime URLs and CI refresh URLs in one small commit.
- Prefer direct fetch/download over restoring the submodule unless pinning model catalog history is a real release requirement.
- Add a narrow test or script check that detects accidental `router-for-me/models` reintroduction.

### [Medium] iFlow Kimi thinking fallback is missing from current clean-root

Evidence:

- Current `internal/thinking/provider/iflow/apply.go:151` enables `chat_template_kwargs.enable_thinking` only for GLM, `qwen3-max-preview`, `deepseek-v3.2`, and `deepseek-v3.1`.
- Current `internal/thinking/provider/iflow/` has only `apply.go`; the old `apply_test.go` from `1aee8779` is absent.
- Old commit `1aee8779` added `isKimiModel()` and tests proving `kimi-k2.5` maps thinking config to `chat_template_kwargs.enable_thinking`.
- Current model definitions still contain Kimi model surfaces such as `kimi-k2.5` in other providers, and old local intent explicitly covered iFlow Kimi.

Impact:

If iFlow exposes Kimi models with `Thinking` support from the model catalog, current clean-root will advertise thinking but leave the outgoing body unchanged for Kimi IDs. That is a behavioral gap, not just a docs gap.

Assessment:

This is a good candidate to port after confirming current iFlow model catalog entries. The old change is small and testable, but it should be ported in the current v7 module context with fresh tests.

Recommended fix scope:

- Add Kimi prefix handling to iFlow thinking only when `modelInfo.Thinking != nil`.
- Recreate the old `TestApplyKimiFallbackThinking` against current package imports.
- Run `go test ./internal/thinking/provider/iflow ./internal/thinking`.

### [Medium] Release workflow dirty-state avoidance was not retained

Evidence:

- Current release workflow writes `internal/registry/models/models.json` during release at `.github/workflows/release.yaml:19`.
- Old commit `ad7d7999` removed the refresh step to avoid dirty state during GoReleaser.
- Current `goreleaser` invocation uses `release --clean --skip=validate`, which may mask validation but still builds from a modified checkout.

Impact:

Release artifacts may include a catalog that was never committed to the release tag. This can make published binaries harder to reproduce and can hide model catalog changes inside CI state.

Assessment:

Not necessarily a blocker if releases intentionally embed latest catalog at build time, but the current repo needs an explicit release policy. It should not remain accidental.

Recommended fix scope:

- Either remove release-time catalog refresh and require committed catalog updates before tags, or document and enforce "release embeds latest external catalog".
- If external refresh remains, add a post-refresh `git diff -- internal/registry/models/models.json` summary in CI logs.

### [Medium] Manual release trigger is missing

Evidence:

- Current `.github/workflows/release.yaml:3` runs only on tag push.
- Old commit `2dc61ac6` added `workflow_dispatch` with `tag` and optional `ref`.

Impact:

No runtime behavior risk, but maintainer release recovery is weaker. If a tag workflow fails or needs rebuild, manual release requires pushing/re-pushing tags or editing workflow dispatch externally.

Assessment:

Safe to reintroduce after release catalog policy is settled. The old implementation force-created a tag in CI; that should be reviewed because it can publish artifacts for a tag that does not exist on the remote unless the workflow also pushes or clearly treats it as build metadata only.

Recommended fix scope:

- Re-add `workflow_dispatch` with explicit tag/ref inputs.
- Avoid hidden remote tag mutation unless intentionally needed.
- Keep release version metadata deterministic for manual and tag-triggered runs.

### [Low] DockerHub-secret skip commits no longer apply directly

Evidence:

- Current `.github/workflows/` has no `docker-image.yml`.
- Old commits `e3d123f8` and `dd9f77fc` only modified `.github/workflows/docker-image.yml`.

Impact:

No current pipeline gap because the Docker image workflow is absent.

Assessment:

Do not replay. If Docker publishing is reintroduced, carry the lesson forward: secret-dependent publish jobs must skip cleanly when DockerHub secrets are absent.

### [Low] Old static model-data revert should not be replayed

Evidence:

- Old commit `2a1cf2b3` modified `internal/registry/model_definitions_static_data.go`.
- Current `HEAD:internal/registry` does not contain `model_definitions_static_data.go`.
- Current registry uses embedded `internal/registry/models/models.json` plus `internal/registry/model_updater.go`.

Impact:

Replaying this would revive an obsolete static-data path and likely fight the current registry architecture.

Assessment:

Reject by default. Mine it only as historical evidence for model IDs if a current model-catalog audit proves a missing model.

### [Resolved] Copilot Claude native routing appears covered by HsnSaboor/current code

Evidence:

- Current `internal/runtime/executor/github_copilot_executor.go:118` routes Copilot requests through the normal Copilot executor and chooses chat vs responses endpoint via `useGitHubCopilotResponsesEndpoint`.
- Current `internal/registry/model_definitions.go:503` marks Copilot Claude fallback with supported endpoints including `/chat/completions`.
- Current tests include `TestGitHubCopilotExecute_ClaudeModelUsesNativeGateway` and `TestGitHubCopilotExecuteStream_ClaudeModelUsesNativeGateway`.
- Old commit `044678b0` implemented same-topic behavior by delegating Claude models through a native gateway before HsnSaboor's later implementation.

Impact:

No immediate missing behavior was found for this commit. The old implementation should not be replayed because it would conflict with the current executor shape and v7 imports.

Residual risk:

The old test name says "native gateway" while the current implementation no longer uses the old `nativeGateway` helper. The behavior appears covered, but a live Copilot account smoke test would be the only proof against GitHub's actual endpoint behavior.

## Goals / Non-goals

Goals:

- Preserve a durable, risk-ranked audit of the eight old local commits.
- Identify current missing behavior that should become future implementation scopes.
- Prevent accidental replay of obsolete docs, submodules, or static-data architecture.
- Update wiki references so this audit is discoverable after compaction.

Non-goals:

- Do not port code in this audit scope.
- Do not restore old upstream docs.
- Do not restore `README_CN.md` or `README_JA.md`.
- Do not restore `third_party/models` unless a future model-catalog scope explicitly chooses that design.

## Target files / modules

Audit evidence touched:

- `.github/workflows/release.yaml`
- `.github/workflows/pr-test-build.yml`
- `internal/registry/model_updater.go`
- `internal/registry/model_definitions.go`
- `internal/thinking/provider/iflow/apply.go`
- `internal/runtime/executor/github_copilot_executor.go`
- `internal/runtime/executor/github_copilot_executor_test.go`

Potential future implementation files:

- `internal/registry/model_updater.go`
- `.github/workflows/release.yaml`
- `.github/workflows/pr-test-build.yml`
- `internal/thinking/provider/iflow/apply.go`
- `internal/thinking/provider/iflow/apply_test.go`

## Constraints

- HsnSaboor remains the baseline.
- Prefer HsnSaboor/current implementation when an old local commit overlaps by provider or feature.
- Reintroduce old local behavior only as fresh, narrow commits on top of current `main`.
- Keep `docs/` for active scopes and archived scope records, not upstream documentation mirrors.
- Keep `.codex/wiki/**` aligned with current code and scope state.

## Verification plan

Audit-only verification:

```bash
git log --oneline --no-merges upstream/main..backup/main-before-hsnsaboor-clean-root
rg -n "router-for-me/models|furedericca-lab/models|workflow_dispatch|docker-image|kimi-k2\\.5|isEnableThinkingModel|NativeGateway|ClaudeModelUsesNativeGateway" .github internal
bash /root/.codex/skills/repo-task-driven/scripts/scope_inventory.sh
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
```

No Go test is required for this audit-only scope unless code changes are made.

Future implementation verification:

```bash
go test ./internal/thinking/provider/iflow ./internal/thinking
go test ./internal/registry
go build -o test-output ./cmd/server && rm test-output
```

## Rollback

This scope is documentation-only. Roll back by reverting the docs/wiki commit that adds this audit.

## Open questions

- Is `furedericca-lab/models` still the intended maintained model catalog source?
- Should release artifacts embed the latest external catalog at build time, or only the catalog committed in the release tag?
- Does iFlow currently expose Kimi thinking models from the live/forked catalog, or was the old `kimi-k2.5` path only provisional?

## Recommendation

REQUEST CHANGES before porting old local commits wholesale.

Immediate follow-up scopes, in order:

1. `model-catalog-source-policy`: decide and implement fork-owned vs router-owned model catalog source across runtime and CI.
2. `iflow-kimi-thinking-recovery`: port the small Kimi thinking fallback with tests if current model catalog still needs it.
3. `release-workflow-policy`: reintroduce manual release and dirty-state behavior after model catalog policy is settled.
