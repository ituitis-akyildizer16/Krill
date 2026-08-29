#!/usr/bin/env bash
# Regenerate shell completions into completions/.
set -euo pipefail

mkdir -p completions

go run ./cmd/krill completion bash > completions/krill.bash
go run ./cmd/krill completion zsh  > completions/_krill
go run ./cmd/krill completion fish > completions/krill.fish

echo ">> completions regenerated"
