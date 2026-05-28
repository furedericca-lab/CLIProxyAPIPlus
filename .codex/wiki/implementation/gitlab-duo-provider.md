---
title: GitLab Duo Provider Notes
type: implementation
status: current
scope: docs-migration
related_files:
  - internal/auth/gitlab
  - internal/runtime/executor/gitlab_executor.go
  - internal/api/handlers/management/auth_files_gitlab_test.go
  - internal/cmd/gitlab_login.go
tags:
  - gitlab
  - provider
  - duo
last_checked: 2026-05-28
updated: 2026-05-28T13:20:00Z
---

# GitLab Duo Provider Notes

GitLab Duo is a first-class Plus provider, not just a plain text wrapper.

Supported behavior from the migrated docs:

- OAuth login
- Personal access token login
- Refresh of GitLab `direct_access` metadata
- Dynamic model discovery from GitLab metadata
- Native GitLab AI gateway routing for Anthropic-managed models
- Native GitLab AI gateway routing for OpenAI/Codex-managed models
- Claude-compatible and OpenAI-compatible downstream APIs

Runtime routing rule:

- Anthropic-managed Duo models should route through the GitLab AI gateway Anthropic proxy and reuse the Claude executor path.
- OpenAI-managed Duo models should route through the GitLab AI gateway OpenAI proxy and reuse the Codex/OpenAI executor path.

Client-visible behavior:

- Claude-compatible clients use Duo models through `/v1/messages`.
- OpenAI-compatible clients use Duo models through `/v1/chat/completions`.
- OpenAI Responses clients use Duo models through `/v1/responses`.
- Stable alias: `gitlab-duo`.
- Discovered model names may include names such as `claude-sonnet-4-5` or `gpt-5-codex`.

Operational notes:

- OAuth login command: `./CLIProxyAPI -gitlab-login`.
- PAT login command: `./CLIProxyAPI -gitlab-token-login`.
- Environment knobs existed in the old docs for base URL, OAuth client metadata, and PAT input. Re-check current names in code before documenting them publicly.
- Self-managed GitLab instances use `GITLAB_BASE_URL`.
- PAT scope baseline is `api`.

Current caution:

- The old docs were under `docs/` and may lag the code after the HsnSaboor/router convergence. Re-verify command names and management endpoints before publishing user-facing README instructions.
