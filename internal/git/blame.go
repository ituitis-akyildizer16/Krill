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
	if strings.HasPrefix(s, "^") {
		s = s[1:]
	}
	if len(s) < 40 {
		return false
	}
	for _, c := range s[:40] {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}

// FormatBlame renders blame lines as compact text for the prompt.
func FormatBlame(bl []BlameLine) string {
	var b strings.Builder
	for _, l := range bl {
		fmt.Fprintf(&b, "%s %s (%s): %s\n", l.Commit[:8], l.Author, l.Timestamp, l.Content)
	}
	return b.String()
}