---
description: Rebaseline this maintained Plus fork so upstream tracks router/main while local scopes preserve Plus provider behavior.
status: complete
date: 2026-05-31
---

# Router Upstream Rebaseline Contract

## Context

The previous maintenance policy used `HsnSaboor/CLIProxyAPIPlus` as the first
Plus convergence baseline and treated `router-for-me/CLIProxyAPI` as a selective
patch source. The user changed direction on 2026-05-31: HsnSaboor appears
unlikely to continue maintenance, so this fork should track router directly and
perform Plus adaptation locally.

## Findings

- `upstream` was changed to `https://github.com/router-for-me/CLIProxyAPI`.
- The `router` remote remains as a compatibility alias for older commands and
  archived evidence.
- `upstream/main` now points at router `v7.1.32` commit `3a54fb7f`.
- Local `main` is `v7.1.9-5` commit `df53e88d`.
- `git rev-list --left-right --count main...upstream/main` returned `4439 86`.
- The merge base is `9ef99aa7` (`v7.1.9`).

## Outcome

The repository policy now treats router `main` as the active upstream baseline.
For providers router already owns, router's implementation is the source of
truth. Providers that are local or former-HsnSaboor exclusive remain locally
owned adaptation work and should not be dropped only because router lacks them.
When HsnSaboor has relevant maintenance for those exclusive providers, use
HsnSaboor's line as their update reference before adapting to router core APIs.

## Goals

- Point `upstream` at router.
- Update local operator .codex/scopes and wiki so future agents do not keep applying the
  old HsnSaboor-first rule.
- Preserve archived HsnSaboor records as historical provenance.
- Record provider precedence so future router integration does not preserve old
  local code where router already has the provider.
- Record that local/HsnSaboor-exclusive provider updates continue to follow
  HsnSaboor maintenance when available.

## Non-goals

- Do not merge router in this scope.
- Do not remove providers in this scope.
- Do not rewrite archived clean-root, audit, or release-repair scope evidence.

## Target Files

- `AGENTS.md`
- `README.md`
- `CLAUDE.md`
- `.codex/wiki/decisions/track-router-main-as-upstream.md`
- `.codex/wiki/decisions/track-hsnsaboor-plus-before-router.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
- `.codex/wiki/maintenance-log.md`

## Constraints

- Root .codex/scopes remain locally owned.
- `docs/` remains reserved for scope contracts and archives.
- `.codex/wiki/**` remains the durable maintainer knowledge layer.
- Provider surface checks are mandatory before future router integration work.

## Boundaries

Archived .codex/scopes under `.codex/scopes/archive/**` may still describe HsnSaboor-first history.
Those records should stay stable unless a future task finds factual stale links
or path references that affect current operation.

## Decision Summary

Router `main` is now primary. This fork is responsible for carrying Plus
provider compatibility on top of router.

## Verification Surface

```bash
git remote -v
git fetch upstream --prune
git rev-list --left-right --count main...upstream/main
git merge-base main upstream/main
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py rebuild
python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint
rg -n "HsnSaboor|selective patch source|first convergence" AGENTS.md README.md CLAUDE.md .codex/wiki .codex/scopes -g '!.codex/scopes/archive/**'
git diff --check
```

## Evidence

- `git remote -v`: `upstream` fetch/push URLs are
  `https://github.com/router-for-me/CLIProxyAPI`.
- `git fetch upstream --prune`: fetched router branches and force-updated
  `upstream/main` from `8c93cf68` to `3a54fb7f`.
- `git rev-list --left-right --count main...upstream/main`: `4439 86`.
- `git merge-base main upstream/main`: `9ef99aa76688f1462fab96670f75ab0d2fc3a77c`.
- `comm -12` on `main:internal/auth` and `upstream/main:internal/auth`:
  `antigravity`, `claude`, `codex`, `empty`, `gemini`, `kimi`, `vertex`, `xai`.
- `comm -23` on `main:internal/auth` and `upstream/main:internal/auth`:
  `cline`, `codebuddy`, `copilot`, `cursor`, `gitlab`, `iflow`, `kilo`, `kiro`,
  `qwen`.
- `python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py rebuild && python3 /root/.codex/skills/wiki-note/scripts/wiki_note.py lint`:
  passed.
- `python3 /root/.codex/skills/repo-task-driven/scripts/repo_task.py --root /root/work/CLIProxyAPIPlus check --scope router-upstream-rebaseline --json`:
  passed with only expected remote/ref string warnings.
- `git diff --check`: passed.

## Escalation Triggers

- Router integration deletes or disables Plus providers without an explicit
  retirement decision.
- Root .codex/scopes or upstream `.codex/scopes/**` overwrite local operator policy.
- Model catalog or auth behavior changes without runtime validation.

## Rollback

If this strategy is reversed, set `upstream` back to
`https://github.com/HsnSaboor/CLIProxyAPIPlus`, restore this scope's policy
edits, and record a new decision page explaining the reversal.

## Open Questions

- Which router feature tranche should be adapted first?
- Should the duplicate `router` remote be kept permanently or removed after
  scripts stop referencing it?

## Archive Record

- Archived on 2026-05-31 under `.codex/scopes/archive/router-upstream-rebaseline/`.
- Archive purpose: preserve the completed router-upstream-rebaseline audit trail.
- Future enhancements should use a new `repo-task-driven` scope under `docs/<enhancement-scope>/`.
- Archived .codex/scopes should only change for factual errata or path-maintenance updates.
