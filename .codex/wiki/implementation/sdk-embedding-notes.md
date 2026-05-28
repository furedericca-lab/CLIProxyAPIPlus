---
title: SDK Embedding Notes
type: implementation
status: current
scope: docs-migration
related_files:
  - sdk/cliproxy
  - sdk/access
  - sdk/translator
  - internal/watcher
tags:
  - sdk
  - embedding
  - watcher
last_checked: 2026-05-28
updated: 2026-05-28T13:20:00Z
---

# SDK Embedding Notes

The old `docs/sdk-*` pages were migrated into this wiki as maintainer notes. Treat examples as conceptually useful but version-sensitive, because the module path and APIs may change during upstream convergence.

## `sdk/cliproxy`

The `sdk/cliproxy` module exposes the proxy as an embeddable Go service. The core shape is:

- Load config through `internal/config`.
- Build with `cliproxy.NewBuilder()`.
- Provide config and config path.
- Run with a cancellable context.
- Service startup wires config/auth watching, token refresh, routes, and graceful shutdown.

Important extension points from the old docs:

- `WithServerOptions`
- `WithMiddleware`
- `WithEngineConfigurator`
- `WithRouterConfigurator`
- `WithRequestLoggerFactory`
- `WithCoreAuthManager`
- lifecycle hooks such as `OnBeforeStart` and `OnAfterStart`

Management API note:

- Management endpoints require the configured remote-management shared credential field.
- Remote management additionally requires the remote-management allow-remote flag.

## Custom executors and translators

Provider executors implement the core provider execution interface and may also implement request preparation for auth header injection.

Conceptual extension path:

1. Implement an executor identifier.
2. Implement non-stream and stream execution.
3. Optionally implement token refresh.
4. Register the executor with the auth manager.
5. Register request/response translators when the upstream wire format differs from built-in protocols.
6. Register models in the global model registry so `/v1/models` and routing hints include them.

Repository rule still applies: avoid standalone changes to `internal/translator/`; prefer SDK translator extension points or broader, justified integration changes.

## `sdk/access`

The access SDK centralizes inbound request authentication. The useful concepts are:

- Providers are registered globally and copied into a `Manager`.
- `Manager.Authenticate` walks providers in order.
- Providers can return success, not-handled, no-credentials, invalid-credential, or internal auth errors.
- Built-in config API-key access validates configured inbound API credentials from supported headers and query aliases.
- Custom access providers should populate metadata for auditing.

Hot reload concept:

- When config changes, refresh config-backed providers and reset the manager provider chain from the registered provider snapshot.

## Watcher queue

The SDK watcher surfaces granular auth updates without forcing full reloads.

Important queue contract:

- `watcher.AuthUpdate` carries `add`, `modify`, or `delete`.
- The queue should be created before the watcher starts.
- Auth updates are coalesced per credential identifier.
- Producers should not block on downstream consumers.
- The service-side queue capacity in the old docs was `256`; re-check current code before relying on that exact value.

## Staleness warning

The old SDK docs referenced the `v6` module path. This repo is already on a later line. Before turning these notes into public README content, verify the current `go.mod`, exported package paths, and builder APIs.
