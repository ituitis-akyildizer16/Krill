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

	branch, _ := r.Git.Exec(ctx, "rev-parse", "--abbrev-ref", "HEAD")
	bundle.Branch = strings.TrimSpace(branch)
	status, _ := r.Git.Status(ctx)
	bundle.Status = status

	type result struct {
		kind string
		val  string
		err  error
	}
	ch := make(chan result, 3)

	go func() {
		target := file
		if target == "" {
			target = "."
		}
		bl, err := r.Git.Blame(ctx, target, r.Cfg.Context.MaxBlameLines)
		if err != nil {
			ch <- result{"blame", "", err}
			return
		}
		ch <- result{"blame", git.FormatBlame(bl), nil}
	}()
	go func() {
		d, err := r.Git.Diff(ctx, true, r.Cfg.Context.MaxDiffLines)
		if err != nil {
			ch <- result{"diff", "", err}
			return
		}
		ch <- result{"diff", d, nil}
	}()
	go func() {
		l, err := r.Git.Log(ctx, r.Cfg.Context.HistoryDays*2)
		if err != nil {
			ch <- result{"log", "", err}
			return
		}
		ch <- result{"log", l, nil}
	}()

	for i := 0; i < 3; i++ {
		res := <-ch
		switch res.kind {
		case "blame":
			bundle.Blame = res.val
		case "diff":
			bundle.Diff = res.val
		case "log":
			bundle.Log = res.val
		}
	}
	return bundle, nil
}

// Ask answers a question grounded in repo context.
func Ask(ctx context.Context, cfg *config.Config, question, file string, contextOnly bool) (string, error) {
	r, err := New(cfg)
