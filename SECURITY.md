# Security

krill is a local tool: your code and prompts never leave the machine
except to the local Ollama socket. This document describes what that
guarantee covers and where the boundaries are.

## Reporting a vulnerability

Do not open a public issue. Email the maintainers directly or use
GitHub's private vulnerability reporting (Security → Report a
vulnerability). Include the component (Go core / extension / plugins),
version, and a minimal reproduction.

## Design guarantees

- **No telemetry.** krill performs no network calls other than the
  configured Ollama URL (`http://localhost:11434` by default). There is
  no update check, no analytics, no crash reporting.
- **No credential access.** krill never reads `.env`, SSH keys, or token
  files. Git context is collected via the git CLI only.
- **Prompt bounds.** Context is capped by `context.max_blame_lines` and
  `context.max_diff_lines`; oversized files are truncated with a marker,
  never silently dropped.
- **Plugin isolation.** Python plugins run as subprocesses with a 2s
  deadline; failures are skipped, never fatal. Enable plugins only from
  directories you control (`plugins.dir`).
- **Command safety.** `suggest` applies deny patterns (hard block) and
  warn patterns (stderr warning) before printing a command.

## Scope

In scope: the Go core, the extension's CLI invocation, the plugin bridge.
Out of scope: the model itself and the Ollama server configuration — if
Ollama is bound to a non-loopback address, prompts may leave the machine.
Keep Ollama on `127.0.0.1:11434`.