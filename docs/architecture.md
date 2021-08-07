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
