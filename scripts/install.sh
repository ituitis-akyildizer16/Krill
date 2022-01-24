#!/usr/bin/env bash
set -euo pipefail

# Build krill from source and install into GOBIN.
VERSION="${VERSION:-$(git describe --tags --always 2>/dev/null || echo dev)}"

echo ">> building krill v${VERSION}"
go build -ldflags "-X github.com/VGaussaleex/krill/internal/cli.Version=${VERSION}" \
  -o "${GOBIN:-$(go env GOPATH)/bin}/krill" ./cmd/krill
