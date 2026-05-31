# CLAUDE.md

This file is locally owned by this maintained CLIProxyAPIPlus fork.

Do not replace it with upstream CLAUDE.md content during synchronization.

Claude-specific agents should follow the same repository policy as `AGENTS.md`:

- Router `main` is the active upstream baseline.
- This fork owns the Plus adaptation layer on top of router.
- Use router provider code for providers router already has.
- Preserve local/HsnSaboor-exclusive providers as Plus extensions, using
  HsnSaboor's maintenance line as their update reference when available.
- Root docs and `.codex/scopes/**` are locally owned.
- Active scope plans live under `.codex/scopes/<scope>/`.
- Durable project knowledge lives under `.codex/wiki/**`.
- Preserve the Plus provider surface unless a scope explicitly decides otherwise.
