# krill

Local-first terminal AI copilot — git-aware answers, shell command
suggestions, and diff reviews through a local Ollama model. No cloud.

## Why local-first

Every agent-era CLI streams your diff to a cloud model. krill flips that:
the model runs on your machine (Ollama), the answers are grounded in your
actual `git blame`/`git diff`/history, and nothing leaves the machine.

## Three components, three languages

- **Go core (`cmd/`, `internal/`)** — the CLI. Context collection is
  parallel and bounded; single static binary.
- **TypeScript extension (`extension/`)** — VS Code companion: ask about
  the open file, review the staged diff, suggest commands.
- **Python plugin SDK (`plugins/`)** — context providers and command
  postprocessors as pip-installable plugins.

## Install

```bash
go install github.com/ituitis-akyildizer16/krill/cmd/krill@latest
ollama pull qwen2.5-coder:7b
```

## Usage

```bash
krill ask "why is this function slow?" -f internal/git/blame.go
krill suggest "stage and commit everything"
krill review
krill status
```

## Development

```bash
make build          # go build ./cmd/krill
make test           # go test ./... -race
make lint           # golangci-lint run
```

## License

MIT. See `LICENSE`.