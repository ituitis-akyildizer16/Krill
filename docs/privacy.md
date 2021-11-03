# Privacy

krill's core promise: **your code and prompts never leave your machine.**
This document is the audit trail for that promise.

## Network calls

There is exactly one network destination in the entire codebase: the
configured Ollama base URL (`http://localhost:11434` by default).

- `POST /api/generate` — model inference (the only payload that contains
  repo context)
- `GET /api/tags` — model listing for `krill models`

Nothing else opens a socket. No update checks, no telemetry, no crash
reporting, no analytics endpoints, no DNS lookups beyond the Ollama host.

## What is collected

When you run `krill ask`, these are sent to the local model:

- branch name and dirty-file count
- up to `max_blame_lines` lines of blame for the target file
- up to `max_diff_lines` lines of your uncommitted diff
- recent commit subjects/authors (up to `history_days`)

You control all of it via config caps. `--context-only` prints the exact
payload without calling the model.

## What is never read

- `.env` files, SSH keys, GitHub tokens, password managers
- Anything outside the current repository
- Files matching `ignore.paths`
