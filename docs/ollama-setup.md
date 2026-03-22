# Ollama setup

krill needs a local Ollama server with a coding model installed.

## Install Ollama

```bash
# macOS
brew install ollama
# Linux
curl -fsSL https://ollama.com/install.sh | sh
# Windows
# download from https://ollama.com/download
```

Start it:

```bash
ollama serve
```

Verify:

```bash
curl http://localhost:11434/api/tags   # -> {"models":[...]}
```

## Pull a model

```bash
ollama pull qwen2.5-coder:7b
```

| Model | Size | Notes |
|---|---|---|
| `qwen2.5-coder:7b` | ~4.7GB | default; good latency/quality on 16GB RAM |
| `qwen2.5-coder:3b` | ~2GB | faster, weaker |
| `qwen2.5-coder:14b` | ~9GB | better quality, needs 32GB RAM |
| `llama3.1:8b` | ~4.9GB | general-purpose alternative |

## Configure krill

```toml
[ollama]
url = "http://localhost:11434"
timeout_seconds = 120
```

If Ollama runs on another host (e.g., a home server), point `url` at it —
but that sends prompts over your network. Keep it loopback for the
privacy guarantee (see `docs/privacy.md`).

## Troubleshooting

- `ollama unreachable at ...` → server not running or wrong `url`.
- `no models installed` → run `ollama pull qwen2.5-coder:7b`.
- Slow first token → try a smaller model or check CPU/GPU load.
<!-- draft note 8 -->
