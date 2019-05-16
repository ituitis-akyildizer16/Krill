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
		return nil, err
	}
	return parseBlame(out, maxLines)
}

func parseBlame(out string, maxLines int) ([]BlameLine, error) {
	lines := strings.Split(out, "\n")
	var result []BlameLine
	var commit, author, ts string
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "\t") {
			result = append(result, BlameLine{
				Commit:    commit,
				Author:    author,
				Timestamp: ts,
				LineNo:    len(result) + 1,
				Content:   strings.TrimPrefix(line, "\t"),
			})
			if maxLines > 0 && len(result) >= maxLines {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "author ") {
			author = strings.TrimPrefix(line, "author ")
		} else if strings.HasPrefix(line, "author-time ") {
			ts = strings.TrimPrefix(line, "author-time ")
		} else if isCommitLine(line) {
			commit = strings.TrimPrefix(line, "^")[:40]
		}
	}
	return result, nil
}

// isCommitLine reports whether line is a porcelain block header
// (40-hex SHA, optionally boundary-marked with ^).
func isCommitLine(line string) bool {
	s := line
