---
title: Codebase Function Map
type: concept
status: current
scope: docs-migration
related_files:
  - cmd/server/main.go
  - internal/cmd/run.go
  - internal/api/server.go
  - internal/api/handlers/management/api_tools.go
  - internal/api/modules
  - internal/runtime/executor
  - internal/registry
  - sdk/cliproxy
  - sdk/proxyutil/proxy.go
tags:
  - architecture
  - function-map
last_checked: 2026-05-28
updated: 2026-05-31T08:23:43Z
---

# Codebase Function Map

This page preserves the useful maintainer map from the old `docs/function_map.md` in a shorter wiki-native form.

## Server lifecycle

Entry point:

- `cmd/server/main.go`
  - Defines login and server flags.
  - Loads `.env`.
  - Selects token/config store backends.
  - Registers built-in access providers.
  - Dispatches login flows, TUI mode, or server mode.

Service startup:

- `internal/cmd/run.go`
  - `StartService`
  - `StartServiceBackground`
  - `WaitForCloudDeploy`
  - Builds the service through `sdk/cliproxy`.
  - Runs the service with signal-aware graceful shutdown.

## HTTP server

Core server:

- `internal/api/server.go`
  - `NewServer`
  - `setupRoutes`
  - `Start`
  - `Stop`
  - `AuthMiddleware`

Middleware order from the old function map:

1. logrus request logger
2. recovery
3. extra middleware
4. request logging when enabled
5. CORS

Auth is route-group scoped, especially for `/v1` and `/v1beta`.

Core route families:

- `/v1/models`
- `/v1/chat/completions`
- `/v1/completions`
- `/v1/messages`
- `/v1/messages/count_tokens`
- `/v1/responses`
- `/v1/responses/compact`
- `/v1beta/models`
- OAuth callback routes
- `/v0/management` when management is enabled

Management routes are enabled only when local or remote management credentials are configured.

## Module system

Module registration lives under `internal/api/modules`.

Important types:

- `modules.Context`
- `RouteModule`
- `RouteModuleV2`
- `RegisterModule`

The Amp integration lives under `internal/api/modules/amp`.

## Provider execution

Runtime executors live under `internal/runtime/executor`.

Provider additions usually require changes across:

- auth storage and refresh
- login command
- executor
- model registry
- config examples
- management auth files / config views
- tests

Shared outbound HTTP transport helpers live under `internal/runtime/executor/helps`
and `sdk/proxyutil`. HTTPS requests whose URL host is an IP literal are routed
through an insecure TLS clone of the selected transport by default; domain-name
HTTPS requests keep normal certificate verification. This preserves proxy and
`proxy-url: direct` behavior while supporting bare-IP upstreams.

Management resource probes under `/ai-providers/*` use a separate transport
builder, `internal/api/handlers/management/api_tools.go:apiCallTransport`,
instead of the executor helper path. Keep that transport wrapped with
`proxyutil.WrapBareIPTLSBypass` too; otherwise Codex/Claude resource checks for
bare-IP HTTPS base URLs can still fail with a frontend 502 even though runtime
executor requests already work.

## Maintainer warning

This map is a navigation aid, not a current API contract. Before editing a subsystem, inspect the current source and tests because HsnSaboor/router convergence may move or rename entry points.
