package prompt

import (
	"strings"
	"testing"
)

func TestAskIncludesContext(t *testing.T) {
	p := Ask(ContextBundle{
		Branch:     "main",
		Blame:      "abc123 Alice: line",
