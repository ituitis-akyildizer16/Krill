// Package runner orchestrates commands: context collection, prompt
// building, model calls, and output.
package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ituitis-akyildizer16/krill/internal/cache"
	"github.com/ituitis-akyildizer16/krill/internal/config"
	"github.com/ituitis-akyildizer16/krill/internal/git"
	"github.com/ituitis-akyildizer16/krill/internal/ollama"
	"github.com/ituitis-akyildizer16/krill/internal/output"
	"github.com/ituitis-akyildizer16/krill/internal/prompt"
	"github.com/ituitis-akyildizer16/krill/internal/suggest"
)

// Version matches the CLI's build-time version.
var Version = "0.9.2"

// Runner holds the pieces one command needs.
type Runner struct {
	Cfg    *config.Config
	Git    *git.Runner
	Ollama *ollama.Client
	Cache  *cache.Store
	Out    *output.Renderer
}

// New assembles a Runner for the current directory.
func New(cfg *config.Config) (*Runner, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	gr, err := git.New(wd)
	if err != nil {
		return nil, fmt.Errorf("not a git repository: %w", err)
	}
	ol := ollama.New(cfg.Ollama.URL, cfg.Ollama.Timeout)
	cacheDir, err := userCacheDir()
	if err != nil {
		return nil, err
	}
	st, err := cache.New(cacheDir, cfg.CacheTTLDuration())
	if err != nil {
		return nil, err
	}
	return &Runner{
		Cfg:    cfg,
		Git:    gr,
		Ollama: ol,
		Cache:  st,
		Out:    output.New(cfg.NoColor),
	}, nil
}

func userCacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "krill"), nil
}

// collectContext gathers blame/diff/log in parallel.
func (r *Runner) collectContext(ctx context.Context, file string) (prompt.ContextBundle, error) {
	bundle := prompt.ContextBundle{TargetFile: file}

