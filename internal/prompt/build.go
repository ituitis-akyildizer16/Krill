// Package prompt builds bounded, git-grounded prompts.
package prompt

import (
	"fmt"
	"strings"
)

// System is the base system prompt for coding questions.
const System = `You are krill, a local-first terminal copilot running on the
user's machine with no network access to any model service.

Rules:
- Answer ONLY from the provided context (git blame, diff, log, file).
- When you reference history, cite the short commit SHA.
- If the context does not contain the answer, say so explicitly.
- Keep answers under 150 words unless the user asks for detail.
- Never invent commit SHAs or file content.`

// ContextBundle is the collected repo context handed to the model.
type ContextBundle struct {
	Branch    string
	Status    string
	Blame     string
	Diff      string
	Log       string
	TargetFile string
