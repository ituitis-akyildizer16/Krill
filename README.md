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
- [Configuration](#configuration)
- [Shell integration](#shell-integration)
- [Security & privacy](#security--privacy)
- [Performance](#performance)
- [Development](#development)
- [Known limitations](#known-limitations)
- [License](#license)

---

## Why local-first

Every agent-era CLI wants to stream your diff to a cloud model. That has
three problems: your code leaves the machine, prompts are rate-limited, and
the answers don't know your repo's history.

krill flips the architecture:

- **Ollama runs the model** — `qwen2.5-coder:7b` and friends run on your
  hardware. The model never sees anything you haven't deliberately asked
  about.
- **Git is the ground truth** — instead of asking a model to guess, krill
  feeds it the actual `git blame`, `git diff`, and recent log for the files
  that matter. Answers cite the commit they came from.
- **The prompt is built locally** — context collection is parallel and
  bounded; the assembled prompt fits your context window before it is sent.

The result is a copilot that is *fast*, *private*, and *specific to your
repo* — the three things cloud copilots trade away.

## Install

Requirements:

- Go 1.21+
- [Ollama](https://ollama.com) running locally (`ollama serve`)

```bash
# build from source
go install github.com/ituitis-akyildizer16/krill/cmd/krill@latest

# or via the release script
curl -fsSL https://raw.githubusercontent.com/ituitis-akyildizer16/krill/main/scripts/install.sh | sh

# pull a coding model (7b is a good size/latency tradeoff)
ollama pull qwen2.5-coder:7b
```

Verify:

```bash
krill status
# krill v0.9.2 · model qwen2.5-coder:7b · shell bash · cache ok
```

## Quick start

```bash
