package git

import (
	"context"
	"fmt"
	"strings"
)

// BlameLine is a single parsed line of `git blame --porcelain`.
type BlameLine struct {
	Commit    string
	Author    string
	Timestamp string
	LineNo    int
	Content   string
}

// Blame returns up to maxLines lines of porcelain blame for a path,
// newest-first is NOT guaranteed; git returns file order. We keep file
// order so line numbers stay meaningful.
func (r *Runner) Blame(ctx context.Context, path string, maxLines int) ([]BlameLine, error) {
	out, err := r.Exec(ctx, "blame", "--porcelain", "--", path)
	if err != nil {
