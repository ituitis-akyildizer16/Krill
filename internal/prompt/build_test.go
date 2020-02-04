package prompt

import (
	"strings"
	"testing"
)

func TestAskIncludesContext(t *testing.T) {
	p := Ask(ContextBundle{
