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
	if err != nil {
		return "", err
	}
	bundle, err := r.collectContext(ctx, file)
	if err != nil {
		return "", err
	}
	promptText := prompt.Ask(bundle, question,
		cfg.Context.MaxBlameLines, cfg.Context.MaxDiffLines)
	if contextOnly {
		return promptText, nil
	}
	key := cacheKey("ask", question, file, bundle.Branch)
	if cached := r.Cache.Get(key); cached != "" {
		return cached, nil
	}
	sp := r.Out.Start("thinking")
	resp, err := r.Ollama.Generate(ctx, ollama.GenerateRequest{
		Model:  cfg.Model,
		Prompt: promptText,
		System: prompt.System,
	})
	sp.Stop()
	if err != nil {
		return "", err
	}
	_ = r.Cache.Set(key, resp.Response)
	return output.Markdown(resp.Response), nil
}

// Suggest returns a shell command for the intent.
func Suggest(ctx context.Context, cfg *config.Config, want, shell string, inline bool) (string, string, error) {
	r, err := New(cfg)
	if err != nil {
		return "", "", err
	}
	promptText := prompt.Suggest(want, shell, cfg.Suggest.MaxTokens)
	key := cacheKey("suggest", want, shell)
	cmd := r.Cache.Get(key)
	if cmd == "" {
		resp, err := r.Ollama.Generate(ctx, ollama.GenerateRequest{
			Model:  cfg.Model,
			Prompt: promptText,
			System: prompt.SuggestSystem,
		})
		if err != nil {
			return "", "", err
		}
		cmd = suggest.Clean(resp.Response)
		_ = r.Cache.Set(key, cmd)
	}
	sc, err := suggest.NewSafetyCheck(cfg.Suggest.DenyPatterns,
		cfg.Suggest.WarnPatterns)
	if err != nil {
		return "", "", err
	}
	res, err := sc.Evaluate(cmd)
	if err != nil {
		return "", "", err
	}
	return res.Command, res.Warning, nil
}

// Review summarizes the staged diff.
func Review(ctx context.Context, cfg *config.Config, staged bool) (string, error) {
	r, err := New(cfg)
	if err != nil {
		return "", err
	}
	diff, err := r.Git.Diff(ctx, staged, cfg.Context.MaxDiffLines)
	if err != nil {
		return "", err
	}
	promptText := "Diff:\n" + diff
	key := cacheKey("review", diff)
	if cached := r.Cache.Get(key); cached != "" {
		return cached, nil
	}
	resp, err := r.Ollama.Generate(ctx, ollama.GenerateRequest{
		Model:  cfg.Model,
		Prompt: promptText,
		System: prompt.ReviewSystem,
	})
	if err != nil {
		return "", err
	}
	_ = r.Cache.Set(key, resp.Response)
	return output.Markdown(resp.Response), nil
}

// Status prints a repo digest.
func Status(ctx context.Context, cfg *config.Config) (string, error) {
	r, err := New(cfg)
	if err != nil {
		return "", err
	}
	branch, _ := r.Git.Exec(ctx, "rev-parse", "--abbrev-ref", "HEAD")
	log, _ := r.Git.Log(ctx, 8)
	var b strings.Builder
	b.WriteString(fmt.Sprintf("krill v%s\n", Version))
	b.WriteString(fmt.Sprintf("model: %s\n", cfg.Model))
	b.WriteString(fmt.Sprintf("branch: %s\n", strings.TrimSpace(branch)))
	b.WriteString(fmt.Sprintf("shell: %s\n", cfg.Shell))
	b.WriteString("recent history:\n" + log + "\n")
	return b.String(), nil
}

// Models lists models available on the local server.
func Models(ctx context.Context, cfg *config.Config) ([]string, error) {
	r, err := New(cfg)
	if err != nil {
		return nil, err
	}
	return r.Ollama.ListModels(ctx)
}

func cacheKey(parts ...string) string {
	joined := strings.Join(parts, "|")
	// keep keys filesystem-safe
	return strings.Map(func(r rune) rune {
		if r == '|' || r == ' ' || r == '/' || r == '\\' {
			return '_'
		}
		return r
	}, joined)
}