// Package git wraps the git CLI for context collection.
//
// Everything here shells out to `git` (the source of truth); we only
// parse its output. Commands are bounded by context, never by timeouts
// that would corrupt the output.
package git

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

// Runner executes git commands in the repo root.
type Runner struct {
	dir string
}

// New locates the repo root from startDir and returns a Runner.
func New(startDir string) (*Runner, error) {
	out, err := exec.Command("git", "-C", startDir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, err
	}
	return &Runner{dir: strings.TrimSpace(string(out))}, nil
}

// Dir returns the repo root.
func (r *Runner) Dir() string { return r.dir }

// Exec runs `git <args...>` and returns trimmed stdout.
func (r *Runner) Exec(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", append([]string{"-C", r.dir}, args...)...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
