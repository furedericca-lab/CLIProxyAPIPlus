## Context

The release workflow currently refreshes `internal/registry/models/models.json` by fetching from `https://github.com/router-for-me/models.git`. The user requested switching this source to their fork and adding that fork as a git submodule in the repository.

## Findings

- The repository did not previously have a `.gitmodules` file.
- The release workflow refresh step is defined in `.github/workflows/release.yaml`.
- A submodule allows the repository to track the fork explicitly and makes the catalog source auditable.

## Goals / Non-goals

Goals:
- Add the user's forked `models` repository as a submodule.
- Update the release workflow to source `models.json` from that submodule instead of the upstream `router-for-me/models` repository.

Non-goals:
- Do not change unrelated CI behavior.
- Do not rewrite the internal registry format.

## Target files / modules

- `.gitmodules`
- `third_party/models` (git submodule)
- `.github/workflows/release.yaml`

## Constraints

- Keep the release workflow behavior minimal and explicit.
- Use the fork URL provided by the user.

## Verification plan

- Verify `.gitmodules` contains the new fork URL and branch.
- Verify the release workflow references `third_party/models`.
- Verify git recognizes the submodule.

## Rollback

- Remove the submodule entry and revert the workflow refresh step to its previous state.

## Open questions

- None for this bounded task.

## Execution Status

- Completed: added `https://github.com/furedericca-lab/models.git` as the `third_party/models` submodule on branch `main`.
- Completed: updated the release workflow to refresh `internal/registry/models/models.json` from the forked submodule.

## Evidence

- `.gitmodules` contains:
  - path `third_party/models`
  - url `https://github.com/furedericca-lab/models.git`
  - branch `main`
- `git -C /root/work/CLIProxyAPIPlus submodule status`
- `.github/workflows/release.yaml` now initializes submodules and copies `third_party/models/models.json`
