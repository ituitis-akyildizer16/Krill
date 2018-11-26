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
# Ask about the file you're working on (blame + diff aware)
krill ask "why is this function slow?" -f internal/git/blame.go

# Get a shell command for what you want to do
krill suggest "stage and commit everything"
# -> git add -A && git commit -m "wip"

# Summarize the staged diff before committing
krill review

# Repo digest: branches, dirty files, recent history, hints
krill status
```

## Commands

| Command | Description |
|---|---|
| `krill ask <question> [-f file]` | Answer grounded in blame/diff/log context |
| `krill suggest <intent>` | Natural language → shell command, safety-checked |
| `krill review` | Summarize the staged diff |
| `krill status` | Repo digest with model-backed hints |
| `krill init` | Write a default config to `~/.config/krill/config.toml` |
| `krill models` | List models available on the local Ollama server |
| `krill version` | Print version and build info |

## How it grounds answers in git

When you `krill ask`, the following context is collected **in parallel**
and assembled into a bounded prompt:

1. **Blame** for the target file (or the files changed in the working
   tree) — up to `context.max_blame_lines` lines.
2. **Diff** of uncommitted changes — up to `context.max_diff_lines`.
3. **Recent history** — commit subjects and authors for the last
   `context.history_days` days.
4. **Repo metadata** — branch, remote URL (name only), language stats.

The prompt template tells the model to answer strictly from this context
and to cite `commit <short-sha>` when it references history. If the
context doesn't contain the answer, it says so instead of hallucinating.

```
$ krill ask "who last touched the retry loop?" -f internal/ollama/client.go

The retry loop in client.go was last changed in commit 4c8a1f3
("ollama: back off on 5xx during generate", 2025-11-02).
The current 3-attempt loop with exponential backoff was added there;
earlier commits (b91e0aa, 2024-06) used a fixed 2-second wait.
```

## Configuration

Config lives at `~/.config/krill/config.toml` (override with
`KRILL_CONFIG`):

```toml
model = "qwen2.5-coder:7b"
shell = "bash"
theme = "dark"
cache_ttl_seconds = 3600
no_color = false

[context]
max_blame_lines = 400
max_diff_lines = 600
history_days = 14

[ollama]
url = "http://localhost:11434"
