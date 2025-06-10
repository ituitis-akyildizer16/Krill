# Architecture

krill is three components in three languages, one goal: answer questions
about your repo with a model that never leaves your machine.

## Components

```
┌─────────────────┐   stdio/JSON-RPC   ┌──────────────────┐
│  VS Code ext      │ ─────────────────► │   Go core (CLI)   │
│  (TypeScript)     │ ◄───────────────── │   cmd/krill       │
└─────────────────┘                    └─────────┬────────┘
                                                 │ exec git
                                        ┌────────▼────────┐
                                        │  git context      │
                                        │  (blame/diff/log) │
                                        └────────┬────────┘
                                                 │ HTTP /api/*
                                        ┌────────▼────────┐
                                        │  Ollama (local)   │
                                        └──────────────────┘

  Python plugins (plugins/) feed extra context via subprocess protocol.
```

### Go core

`cmd/krill` + `internal/`. Each internal package is small and has one
job:

| Package | Job |
|---|---|
| `cli` | cobra command tree, flags, version |
| `config` | TOML load/save with defaults |
| `git` | blame/diff/log/status wrappers + parsing |
| `ollama` | minimal /api/generate + /api/tags client |
| `prompt` | bounded prompt assembly, system prompts |
| `suggest` | NL→command + safety deny/warn checks |
| `output` | terminal rendering, spinner, markdown strip |
| `cache` | on-disk TTL cache for responses |
| `runner` | orchestration of the above |

### TypeScript extension

`extension/` is a standard VS Code extension: three commands
(ask/review/suggest) that shell out to the CLI binary. It adds no logic of
its own — the CLI is the single source of truth.

### Python plugin SDK

`plugins/sdk` ships `krill_plugins` (ContextProvider,
SuggestPostprocessor, discovery). Providers are loaded from
`plugins.dir` at runtime; each runs as a subprocess with a 2s deadline.

## Data flow (ask)

1. CLI collects branch, status, blame, diff, log **in parallel**.
2. `prompt.Ask` assembles a bounded prompt (context caps enforced).
3. CLI POSTs to Ollama `/api/generate` (stream=false).
4. Response is markdown-stripped and printed; cached by key.

## Design rules

- The CLI is the only component that talks to git or Ollama.
- Context is always bounded before the model is called.
- Caching is keyed on question+file+branch and TTL'd.
- Plugins never block: 2s deadline, failures skipped.