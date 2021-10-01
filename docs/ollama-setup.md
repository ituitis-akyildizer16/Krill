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
