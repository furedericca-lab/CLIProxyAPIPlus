---
description: Implementation scope for resolving risks found in the archived old-local commit audit.
---

# Risk Resolution From Old Local Audit Contract

## Context

The completed audit at `docs/archive/old-local-commit-audit/old-local-commit-audit-contract.md` identified one high-risk and three medium-risk gaps that should be fixed before more upstream absorption work:

- Model catalog source still pointed at router instead of this fork's model catalog.
- Release and PR workflows rewrote `internal/registry/models/models.json` from router during CI.
- iFlow Kimi thinking fallback from old local work was missing.
- Manual release dispatch from old local CI was missing.

## Goals / Non-goals

Goals:

- Point runtime model refresh at `furedericca-lab/models`.
- Stop CI/release jobs from silently rewriting committed model catalog data.
- Reintroduce manual release dispatch without mutating remote tags.
- Restore iFlow Kimi thinking fallback with tests.
- Add a regression guard that prevents router model source from returning to runtime updater config.
- Archive this scope after verification.

Non-goals:

- Do not restore the old `third_party/models` submodule.
- Do not restore old upstream docs.
- Do not restore `README_CN.md` or `README_JA.md`.
- Do not change provider surfaces outside the audited risk fixes.
- Do not implement Docker publishing; current repo has no Docker image workflow.

## Target files / modules

- `internal/registry/model_updater.go`
- `internal/registry/model_updater_test.go`
- `.github/workflows/release.yaml`
- `.github/workflows/pr-test-build.yml`
- `internal/thinking/provider/iflow/apply.go`
- `internal/thinking/provider/iflow/apply_test.go`
- `.codex/wiki/**`

## Implementation Plan

1. Model source and CI policy:
   - Replace router model URLs in runtime updater with `furedericca-lab/models`.
   - Remove release/PR build steps that rewrite `internal/registry/models/models.json`.
   - Keep release artifacts tied to the committed catalog in the release tag.
2. Release workflow:
   - Add `workflow_dispatch` with required `tag` and optional `ref`.
   - Create only a local tag for manual release metadata.
3. iFlow thinking:
   - Treat `kimi-*` iFlow model IDs as `chat_template_kwargs.enable_thinking` models when `modelInfo.Thinking` is present.
   - Preserve GLM-only `clear_thinking` behavior.
4. Regression tests:
   - Test model updater URLs use the fork source and do not include router sources.
   - Test Kimi thinking enabled/disabled behavior and that `clear_thinking` is not emitted for Kimi.

## Verification Plan

```bash
gofmt -w internal/registry/model_updater_test.go internal/thinking/provider/iflow/apply.go internal/thinking/provider/iflow/apply_test.go
go test ./internal/thinking/provider/iflow ./internal/thinking ./internal/registry
go build -o test-output ./cmd/server && rm test-output
rg -n "router-for-me/models|models.router-for.me" .github internal/registry -g '!**/*_test.go'
bash /root/.codex/skills/repo-task-driven/scripts/scope_inventory.sh --archive
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
```

Expected `rg` result: no matches. Test files may contain those strings as negative regression fixtures.

## Implementation Evidence

- Runtime updater now fetches from `https://raw.githubusercontent.com/furedericca-lab/models/refs/heads/main/models.json`.
- Release and PR build workflows no longer rewrite `internal/registry/models/models.json`.
- Release workflow supports manual dispatch with a required tag and optional ref, using only a local tag for GoReleaser metadata.
- iFlow `kimi-*` models with `Thinking` support now use the generic `chat_template_kwargs.enable_thinking` fallback.
- Regression tests cover fork model source and iFlow Kimi thinking behavior.

## Verification Evidence

Passed on 2026-05-28:

```bash
go test ./internal/thinking/provider/iflow ./internal/thinking ./internal/registry
go build -o test-output ./cmd/server && rm test-output
go test ./...
rg -n "router-for-me/models|models.router-for.me" .github internal/registry -g '!**/*_test.go'
git diff --check
bash /root/.codex/skills/repo-task-driven/scripts/doc_placeholder_scan.sh docs/risk-resolution-from-old-local-audit
bash /root/.codex/skills/repo-task-driven/scripts/post_refactor_text_scan.sh docs/risk-resolution-from-old-local-audit README.md AGENTS.md
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
```

The model-source `rg` check exited with no matches. `actionlint` is not installed in this environment, so GitHub Actions syntax was reviewed by file inspection rather than tool validation.

## Rollback

Revert the implementation commit and this scope archive commit. This will restore router model source behavior and remove the iFlow Kimi fallback.

## Open Questions

- None for this implementation pass. `furedericca-lab/models.git` was verified to have a `main` branch before implementation.
