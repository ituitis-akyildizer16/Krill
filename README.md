# krill

**Local-first terminal AI copilot** — git-blame-aware answers, shell command
suggestions, and repo digests, all through a local Ollama model. No cloud
round-trip, no telemetry, no account, no lock-in.

| | |
|---|---|
| License | MIT |
| Language | Go 1.21+ |
| Model backend | Ollama (local, default `http://localhost:11434`) |
| Shells | bash · zsh · fish · powershell |
| Cache | on-disk, TTL-managed, under `~/.cache/krill/` |

```bash
go install github.com/ituitis-akyildizer16/krill/cmd/krill@latest
krill ask "why is this function slow?"
krill suggest "stage and commit everything"
```

---

## Table of Contents

- [Why local-first](#why-local-first)
- [Install](#install)
- [Quick start](#quick-start)
- [Commands](#commands)
- [How it grounds answers in git](#how-it-grounds-answers-in-git)
