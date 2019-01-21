// Package cli wires the cobra command tree for krill.
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/ituitis-akyildizer16/krill/internal/config"
	"github.com/ituitis-akyildizer16/krill/internal/runner"
)

// Version is overridden at build time via -ldflags.
var Version = "0.9.2"

// NewRootCommand builds the full command tree.
func NewRootCommand() *cobra.Command {
	cfg, err := config.Load(config.DefaultPath())
	if err != nil {
		cfg = config.Defaults()
	}

	root := &cobra.Command{
		Use:   "krill",
		Short: "Local-first terminal AI copilot",
		Long: "krill answers questions about your repo, suggests shell commands,\n" +
			"and summarizes diffs - all through a local Ollama model.",
		SilenceUsage: true,
	}

	root.AddCommand(
		newAskCommand(cfg),
		newSuggestCommand(cfg),
		newReviewCommand(cfg),
		newStatusCommand(cfg),
		newInitCommand(cfg),
		newModelsCommand(cfg),
		newVersionCommand(),
	)

	return root
}

