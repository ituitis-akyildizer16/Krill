package git

import (
	"context"
	"strconv"
	"strings"
)

// Diff returns the uncommitted diff (staged or unstaged) bounded by
// maxLines of output.
func (r *Runner) Diff(ctx context.Context, staged bool, maxLines int) (string, error) {
	args := []string{"diff"}
	if staged {
		args = append(args, "--cached")
	}
	args = append(args, "--stat", "--patch", "--no-color")
	out, err := r.Exec(ctx, args...)
	if err != nil {
		return "", err
	}
	if maxLines > 0 {
		lines := strings.Split(out, "\n")
		if len(lines) > maxLines {
			lines = append(lines[:maxLines], "... (truncated)")
		}
		out = strings.Join(lines, "\n")
	}
	return out, nil
}

// Log returns the last n commit subjects with short SHAs.
func (r *Runner) Log(ctx context.Context, n int) (string, error) {
	out, err := r.Exec(ctx,
		"log", "-n", itoa(n), "--pretty=format:%h %ad %an: %s",
		"--date=short")
	if err != nil {
		return "", err
	}
	return out, nil
}

// Status returns a compact status line: branch, dirty counts.
func (r *Runner) Status(ctx context.Context) (string, error) {
	branch, err := r.Exec(ctx, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}
	porcelain, err := r.Exec(ctx, "status", "--porcelain")
	if err != nil {
		return "", err
	}
