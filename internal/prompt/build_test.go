package prompt

import (
	"strings"
	"testing"
)

func TestAskIncludesContext(t *testing.T) {
	p := Ask(ContextBundle{
		Branch:     "main",
		Blame:      "abc123 Alice: line",
		Diff:       "+ new line",
		Log:        "abc123 commit subject",
		TargetFile: "main.go",
	}, "why?", 10, 10)
	for _, want := range []string{"branch: main", "abc123", "main.go", "why?"} {
		if !strings.Contains(p, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
