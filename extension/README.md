# krill VS Code extension

Local-first AI copilot for VS Code: ask git-aware questions about the
file you're editing, review the staged diff, and suggest shell commands.
Runs entirely through the local `krill` CLI + Ollama — no cloud.

## Install

```bash
go install github.com/ituitis-akyildizer16/krill/cmd/krill@latest
code --install-extension ./extension.vsix   # or use F5 with the Extension Host
```

## Commands

| Command | Binding | What it does |
|---|---|---|
| krill: Ask about this file | ctrl+alt+a | Grounded answer via blame/diff |
| krill: Review staged diff | — | Summary in a notification |
| krill: Suggest command | — | NL → command, copied to clipboard |

## Configuration

- `krill.binaryPath` — path to the CLI (default `krill`)
- `krill.model` — Ollama model (default `qwen2.5-coder:7b`)

## Development

```bash
cd extension
npm install
npm run compile
```

Press F5 to launch the Extension Development Host.

## Privacy

All model traffic goes to `http://localhost:11434` (Ollama). The extension
sends file paths and the collected git context to that local server only.