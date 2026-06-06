---
title: Maintenance Log
type: maintenance-log
status: current
updated: 2026-05-28T12:47:48Z
---

# Maintenance Log

Append-only history for wiki updates caused by scope work, implementation closeout, or knowledge refresh.

## 2026-05-28T12:49:21Z [upstream-plus-gap-analysis]

- Summary: Initialized project wiki and recorded upstream Plus maintenance strategy after comparing local, HsnSaboor Plus, and router-for-me refs.
- Pages: .codex/wiki/reference/upstream-plus-maintenance.md
- Verification: git fetch --prune upstream; git fetch --prune router; git rev-list --left-right --count; git ls-tree provider matrix
- Residual risk: No implementation merge has been attempted yet.

## 2026-05-28T12:57:58Z [docs-migration]

- Summary: Migrated useful docs content into project wiki and cleared .codex/scopes/* per repository maintenance policy.
- Pages: .codex/wiki/index.md
- Verification: find docs -mindepth 1 -maxdepth 1 -exec rm -rf {} +
- Residual risk: Migrated docs are concise maintainer notes, not full verbatim documentation.

## 2026-05-28T13:03:48Z [hsnsaboor-clean-root]

- Summary: Added detailed clean-root integration plan using HsnSaboor upstream as base and replaying only current local maintenance decisions.
- Pages: .codex/scopes/archive/hsnsaboor-clean-root/
- Verification: scope archived after build, full tests, provider matrix, wiki lint, and scope inventory passed
- Residual risk: Plan only; integration branch has not been created yet.

## 2026-05-28T13:37:42Z [old-local-commit-audit]

- Summary: Opened old-local commit audit scope, refreshed HsnSaboor clean-root wiki scope pointers, and linked the current audit from provider/commit inventory.
- Pages: .codex/wiki/reference/local-provider-and-commit-inventory.md
- Verification: scope_inventory.sh; wiki_note.py rebuild; wiki_note.py lint

## 2026-05-28T13:39:25Z [old-local-commit-audit]

- Summary: Archived completed old-local commit audit scope and updated wiki references from active docs path to archive path.
- Pages: .codex/scopes/archive/old-local-commit-audit/old-local-commit-audit-contract.md
- Verification: archive_scope_docs.sh old-local-commit-audit 2026-05-28

## 2026-05-28T13:44:58Z [risk-resolution-from-old-local-audit]

- Summary: Resolved old-local audit model source and iFlow thinking risks: runtime now uses furedericca-lab/models, CI no longer rewrites committed catalog, and iFlow Kimi thinking fallback has tests.
- Pages: .codex/wiki/reference/model-catalog-maintenance.md
- Verification: go test ./...; go build -o test-output ./cmd/server && rm test-output

## 2026-05-28T13:45:49Z [risk-resolution-from-old-local-audit]

- Summary: Archived completed risk-resolution scope after verification and scope inventory showed no active scopes.
- Pages: .codex/scopes/archive/risk-resolution-from-old-local-audit/risk-resolution-from-old-local-audit-contract.md
- Verification: archive_scope_docs.sh risk-resolution-from-old-local-audit 2026-05-28; scope_inventory.sh --archive

## 2026-05-28T16:03:33Z [release-ci-repair]

- Summary: Repaired release CI packaging after GoReleaser failed on absent README_CN.md and moved release actions to Node 24-backed majors.
- Pages: .codex/scopes/archive/release-ci-repair/release-ci-repair-contract.md
- Verification: go run github.com/goreleaser/goreleaser/v2@v2.16.0 check, go build -o test-output ./cmd/server && rm test-output, go run github.com/goreleaser/goreleaser/v2@v2.16.0 release --snapshot --clean --skip=publish
- Residual risk: A real GitHub release publish was not rerun from this local validation.

## 2026-05-28T16:04:24Z [release-ci-repair]

- Summary: Archived completed release CI repair scope after config validation and server build passed.
- Pages: .codex/scopes/archive/release-ci-repair/release-ci-repair-contract.md
- Verification: bash /root/.codex/skills/repo-task-driven/scripts/archive_scope_docs.sh release-ci-repair 2026-05-28, bash /root/.codex/skills/repo-task-driven/scripts/scope_inventory.sh --archive
- Residual risk: Release publishing still needs a new GitHub Actions run.

## 2026-05-28T17:40:03Z [remove-iflow-kimi-fallback]

- Summary: Removed the local iFlow Kimi thinking fallback after user confirmed it is not needed; wiki now records HsnSaboor behavior as intended.
- Pages: .codex/wiki/reference/model-catalog-maintenance.md
- Verification: go test ./internal/thinking/provider/iflow ./internal/thinking, go build -o test-output ./cmd/server && rm test-output
- Residual risk: Archived historical audit docs still mention the earlier rationale as provenance.

## 2026-05-28T17:41:58Z [align-runtime-with-hsnsaboor]

- Summary: Aligned runtime source code with HsnSaboor by removing local model source and iFlow Kimi fallback patches while preserving local CI/release and documentation policy.
- Pages: .codex/wiki/reference/model-catalog-maintenance.md
- Verification: go test ./internal/thinking/provider/iflow ./internal/thinking ./internal/registry, go build -o test-output ./cmd/server && rm test-output
- Residual risk: Archived historical scopes still describe the removed local patches as provenance.

## 2026-05-31T05:16:04Z [router-upstream-rebaseline]

- Summary: Rebased upstream policy to router/main and recorded provider precedence: router-owned providers use router code, local/HsnSaboor-exclusive providers remain Plus extensions.
- Pages: .codex/wiki/reference/upstream-plus-maintenance.md
- Verification: git remote set-url upstream https://github.com/router-for-me/CLIProxyAPI; git fetch upstream --prune; comm provider matrices
- Residual risk: Router integration has not been merged yet; future scopes must validate provider behavior before code changes.

## 2026-05-31T05:19:18Z [router-upstream-rebaseline]

- Summary: Clarified provider precedence: router-owned providers follow router, while local/HsnSaboor-exclusive providers stay as Plus extensions and use HsnSaboor maintenance as their update reference when available.
- Pages: .codex/wiki/reference/upstream-plus-maintenance.md
- Verification: update AGENTS.md README.md CLAUDE.md wiki references and router-upstream-rebaseline contract
- Residual risk: Actual router code integration remains future scoped work.

## 2026-05-31T05:19:52Z [router-upstream-rebaseline]

- Summary: Archived the completed router upstream rebaseline scope and updated entry-point references to .codex/scopes/archive/router-upstream-rebaseline/.
- Pages: .codex/scopes/archive/router-upstream-rebaseline/router-upstream-rebaseline-contract.md
- Verification: repo_task.py --root /root/work/CLIProxyAPIPlus archive --scope router-upstream-rebaseline --archive-date 2026-05-31 --apply --json
- Residual risk: Future router integration still needs separate implementation scopes.

## 2026-05-31T06:51:46Z [bare-ip-tls-skip]

- Summary: Documented automatic TLS verification bypass for HTTPS IP-literal upstream transports
- Pages: .codex/wiki/concepts/codebase-function-map.md

## 2026-05-31T07:01:48Z [bare-ip-tls-skip]

- Summary: Archived bare-IP HTTPS TLS bypass scope after implementation and verification
- Pages: .codex/scopes/archive/bare-ip-tls-skip/bare-ip-tls-skip-contract.md

## 2026-05-31T08:23:43Z [bare-ip-tls-skip]

- Summary: Recorded the management resource transport pitfall: `/ai-providers/*` probes use `apiCallTransport`, separate from executor transports, and must also wrap bare-IP HTTPS requests with `proxyutil.WrapBareIPTLSBypass`.
- Pages: .codex/wiki/concepts/codebase-function-map.md; .codex/scopes/archive/bare-ip-tls-skip/bare-ip-tls-skip-contract.md
- Verification: `rg` confirmed the pitfall was only in the archived scope before this wiki update; current source has `apiCallTransport` wrapping all return paths.
- Residual risk: Running services still need a rebuilt/restarted binary before the management UI observes the transport fix.

## 2026-06-01T03:07:47Z [docs-consolidation]

- Summary: Aligned README and AGENTS with wiki-note documentation structure; added source docs archive map so the wiki is the durable docs surface and root docs remain routing summaries.
- Pages: .codex/wiki/reference/source-docs-archive-map.md, README.md, AGENTS.md
- Verification: wiki.py rebuild; wiki.py nav build --json
- Residual risk: No code behavior changed; Go build and tests were not run for this documentation-only update.

## 2026-06-01T10:09:56Z [docs-consolidation]

- Summary: Removed the obsolete .codex/notepad.md scratch-file example and made the generated wiki navigation index explicitly ignored.
- Pages: .gitignore, AGENTS.md, .codex/wiki/reference/upstream-plus-maintenance.md
- Verification: rg notepad README.md AGENTS.md .codex/wiki /root/.codex/skills/wiki-note; git check-ignore -v .codex/wiki/.index .codex/wiki/.index/pages.json
- Residual risk: No code behavior changed; documentation and ignore-rule update only.

## 2026-06-06T09:29:01Z [router-v7.1.46-sync]

- Summary: Recorded router v7.1.46 merge pitfalls: Plus deletion hunks, tag fetch warnings, file-backed logging helpers, usage reporter API migration, Codex uTLS routing, Cloudflare cooldown split, and auth test helper dependency.
- Pages: .codex/wiki/reference/upstream-plus-maintenance.md
- Verification: go test ./...; go build -o test-output ./cmd/server && rm test-output; wiki.py rebuild/doctor/lint/surface-check
- Residual risk: Future router merges still need provider-surface review before accepting upstream deletions.
