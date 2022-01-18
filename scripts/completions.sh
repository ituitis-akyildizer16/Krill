#!/usr/bin/env bash
# Regenerate shell completions into completions/.
set -euo pipefail

mkdir -p completions

go run ./cmd/krill completion bash > completions/krill.bash
go run ./cmd/krill completion zsh  > completions/_krill
