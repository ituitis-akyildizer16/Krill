// Command krill is a local-first terminal AI copilot.
//
// It collects git context (blame, diff, history), builds a bounded prompt,
// and streams an answer from a local Ollama model. Nothing leaves the
// machine; the VS Code extension (extension/) and plugin SDK (plugins/)
// talk to the same core over stdio/JSON-RPC.
package main

import (
	"os"

	"github.com/ituitis-akyildizer16/krill/internal/cli"
)

