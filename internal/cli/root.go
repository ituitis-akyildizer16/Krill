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

func newAskCommand(cfg *config.Config) *cobra.Command {
	var file string
	var contextOnly bool
	cmd := &cobra.Command{
		Use:   "ask [question]",
		Short: "Ask a question about the current repo",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			q := joinArgs(args)
			out, err := runner.Ask(context.Background(), cfg, q, file, contextOnly)
			if err != nil {
				return fmt.Errorf("ask: %w", err)
			}
			fmt.Println(out)
			return nil
		},
	}
	cmd.Flags().StringVarP(&file, "file", "f", "", "limit context to this file")
	cmd.Flags().BoolVar(&contextOnly, "context-only", false,
		"print the assembled context without calling the model")
	return cmd
}

func newSuggestCommand(cfg *config.Config) *cobra.Command {
	var inline bool
	var shell string
	cmd := &cobra.Command{
		Use:   "suggest [what you want to do]",
		Short: "Suggest a shell command for what you want",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			want := joinArgs(args)
			sh := shell
			if sh == "" {
				sh = cfg.Shell
			}
			cmdline, warning, err := runner.Suggest(context.Background(), cfg, want, sh, inline)
			if err != nil {
