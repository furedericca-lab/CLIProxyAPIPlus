---
description: Repair scope for the GitHub Actions GoReleaser failure on manual release run 26578713077.
---

# Release CI Repair Contract

## Context

Manual release run `26578713077` failed in the `goreleaser` job after all target
binaries built. The failing archive step tried to include `README_CN.md`, but
this maintained fork intentionally removed translated root READMEs.

The same run also emitted GitHub's Node.js 20 action deprecation warning for
the release workflow actions.

## Goals / Non-goals

Goals:

- Remove absent translated root .codex/scopes from GoReleaser archives.
- Keep release archives tied to files that exist in this fork.
- Move the release workflow to Node 24-backed action majors.
- Keep GoReleaser on the v2 line while using the current action wrapper.
- Verify the GoReleaser config and the required server build.

Non-goals:

- Do not restore `README_CN.md` or `README_JA.md`.
- Do not change release artifact naming or target OS/architecture coverage.
- Do not run a real publishing release from local validation.

## Target files / modules

- `.goreleaser.yml`
- `.github/workflows/release.yaml`
- `.codex/wiki/**`

## Verification Plan

```bash
go run github.com/goreleaser/goreleaser/v2@v2.16.0 check
go build -o test-output ./cmd/server && rm test-output
rg -n "README_CN.md|README_JA.md" .goreleaser.yml .github/workflows/release.yaml
```

Expected `rg` result: no matches.

## Implementation Evidence

- `.goreleaser.yml` no longer includes `README_CN.md` in release archives.
- `.goreleaser.yml` uses current GoReleaser v2 archive/snapshot keys
  (`formats`, `format_overrides.formats`, `snapshot.version_template`).
- `.github/workflows/release.yaml` now uses Node 24-backed action majors:
  `actions/checkout@v6`, `actions/setup-go@v6`, and
  `goreleaser/goreleaser-action@v7`.
- The GoReleaser action is pinned to the v2 GoReleaser line with
  `version: '~> v2'` instead of floating to `latest`.

## Verification Evidence

Passed on 2026-05-28:

```bash
go run github.com/goreleaser/goreleaser/v2@v2.16.0 check
go build -o test-output ./cmd/server && rm test-output
go run github.com/goreleaser/goreleaser/v2@v2.16.0 release --snapshot --clean --skip=publish
```

Additional evidence:

- Run `26578713077` failed with:
  `failed to find files to archive: globbing failed for pattern README_CN.md:
  matching "./README_CN.md": file does not exist`.
- `gh api` checks confirmed the chosen action majors use `node24` in
  `action.yml`.
- The local snapshot release built the full target matrix and produced the
  expected Linux, Windows, Darwin, and FreeBSD archives without the missing-file
  archive error.

## Rollback

Revert this scope's implementation commit. That restores the previous release
workflow action majors and archive file list.

## Open Questions

- None.

## Archive Record

- Archived on 2026-05-28 under `.codex/scopes/archive/release-ci-repair/`.
- Archive purpose: preserve the completed release-ci-repair audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `docs/<enhancement-scope>/`.
- Archived .codex/scopes should only change for factual errata or path-maintenance updates.
