#!/usr/bin/env bash
set -euo pipefail

# Build krill from source and install into GOBIN.
VERSION="${VERSION:-$(git describe --tags --always 2>/dev/null || echo dev)}"

