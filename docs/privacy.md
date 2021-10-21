# Privacy

krill's core promise: **your code and prompts never leave your machine.**
This document is the audit trail for that promise.

## Network calls

There is exactly one network destination in the entire codebase: the
configured Ollama base URL (`http://localhost:11434` by default).

- `POST /api/generate` — model inference (the only payload that contains
