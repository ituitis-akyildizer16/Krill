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
