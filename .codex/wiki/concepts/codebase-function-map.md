---
title: Codebase Function Map
type: concept
status: current
scope: docs-migration
related_files:
  - cmd/server/main.go
  - internal/cmd/run.go
  - internal/api/server.go
  - internal/api/modules
  - internal/runtime/executor
  - internal/registry
  - sdk/cliproxy
tags:
  - architecture
  - function-map
last_checked: 2026-05-28
updated: 2026-05-28T13:20:00Z
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

## Maintainer warning

This map is a navigation aid, not a current API contract. Before editing a subsystem, inspect the current source and tests because HsnSaboor/router convergence may move or rename entry points.
