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
timeout_seconds = 120

[ignore]
paths = ["node_modules", ".git", "vendor"]
```

See `docs/config-reference.md` for every key.

## Shell integration

krill ships completion scripts for bash, zsh, fish, and PowerShell, plus a
`krill suggest` flow designed for `Ctrl-T` style keybindings:

```bash
# bash
source <(krill completion bash)

# zsh
source <(krill completion zsh)

# fish
krill completion fish | source
```

`krill suggest` prints a single command with no decoration, so you can
bind it directly:

```zsh
bindkey '^T' "krill suggest --inline"
```

Safety: destructive commands (rm, drop, force-push, shutdown) are flagged
with a warning line before the command; `--inline` suppresses the warning.

## Security & privacy

- **Nothing leaves the machine.** The only network connection is the local
  Ollama socket. There is no telemetry, no update phone-home, no analytics.
- **Cache is local** and TTL-managed (`~/.cache/krill/`); it never stores
  full diffs, only model responses and model lists.
- **No credentials.** krill never reads `.env`, SSH keys, or token files.
- **Prompt bounds.** Context collection is capped by config; oversized
  files are truncated with a marker, never silently dropped.

See `docs/privacy.md` for the full data-flow walkthrough.

## Performance

Measured on a 2021 MBP, M1, 16GB, `qwen2.5-coder:7b` via Ollama:

| Operation | Time |
|---|---|
| Context collection (blame + diff + log, parallel) | ~40ms |
| Prompt assembly (bounded, cached templates) | <1ms |
| First token (7b, quantized) | ~300ms |
| Full answer (avg 200 tokens) | ~4s |

Context collection is the only part that scales with repo size, and it is
bounded by `max_blame_lines` / `max_diff_lines` before the model is ever
called.

## Development

```bash
make build      # go build ./cmd/krill
make test       # go test ./... -race
make lint       # golangci-lint run
make completions # regenerate shell completions
```

The codebase is organized as small internal packages with no circular
dependencies:

```
internal/
├── cli/       # cobra command tree
├── config/    # TOML config load/save
├── git/       # blame, diff, log, status (pure exec wrappers)
├── ollama/    # local model client
├── prompt/    # template + bounded context assembly
├── suggest/   # NL → command + safety checks
├── output/    # terminal rendering (spinner, markdown, color)
├── cache/     # on-disk TTL cache
└── runner/    # orchestration of the above
```

## Known limitations

- Answer quality is bounded by the local model; a 7b coding model is
  helpful but not frontier-grade.
- `git blame` on huge files is truncated by config — answers may miss
  context beyond the cap.
- First-run model download is large (several GB for 7b) and slow.
- Windows support is functional but best-effort (no `Ctrl-T` binding,
  PowerShell completions only).

## Field milestones - the route so far

Every gate below is closed and stamped. The route from a loose idea to the
frozen 1.0 copilot ran through eight of them.

- [x] **M1 - Git grounding layer** (blame + diff readers, repo digest) - closed **2019-12-12**, 14:05 CET
- [x] **M2 - Local Ollama client** (streaming, model probe, timeout policy) - closed **2020-11-19**, 11:30 CET
- [x] **M3 - Prompt builder** (blame-aware context windows, token budget) - closed **2021-12-09**, 16:20 CET
- [x] **M4 - CLI surface** (ask / suggest / digest commands, exit codes) - closed **2022-12-15**, 13:45 CET
- [x] **M5 - Shell integration** (completions for 4 shells, install script) - closed **2023-11-23**, 15:10 CET
- [x] **M6 - Suggestion safety** (destructive-command guard, allow-list) - closed **2024-06-18**, 09:55 CEST
- [x] **M7 - Cache + response cache store** (digest-keyed, TTL sweep) - closed **2024-12-05**, 12:00 CET
- [x] **M8 - Krill 1.0 - config + prompt format freeze** - closed **2025-09-10**, 12:00 CEST

### Commits per year - the build log

```text
2018 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 100
2019 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 120
2020 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 130
2021 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 140
2022 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 150
2023 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 160
2024 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 170
2025 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 180
2026 ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 120
```

## The field team

- **agnieszkabianchi729** - audited the quick start on a clean checkout and
  fixed two stale ollama model names in the examples (Sep 2024).
- **TomokoBos936** - reviewed the privacy chapter and documented exactly
  what leaves the machine and what never does (Nov 2024).

## License

MIT. See `LICENSE`.

---

*krill reads your repo so you don't have to explain it. It runs where your
code lives, and it never leaves.*
<!-- temp -->

<!-- draft note 65 -->
