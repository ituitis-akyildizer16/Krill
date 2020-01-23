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
	Extra     string
}

// Ask builds a prompt for a free-form question.
func Ask(bundle ContextBundle, question string, maxBlame, maxDiff int) string {
	var b strings.Builder
	b.WriteString("Repo context:\n")
	b.WriteString(fmt.Sprintf("branch: %s\n", bundle.Branch))
	if bundle.Status != "" {
		b.WriteString(fmt.Sprintf("status: %s\n", bundle.Status))
	}
	if bundle.TargetFile != "" {
		b.WriteString(fmt.Sprintf("file: %s\n", bundle.TargetFile))
	}
	if bundle.Blame != "" {
		b.WriteString("\n# blame (recent):\n" + bundle.Blame)
	}
	if bundle.Diff != "" {
		b.WriteString("\n# uncommitted diff:\n" + bundle.Diff)
	}
	if bundle.Log != "" {
		b.WriteString("\n# recent history:\n" + bundle.Log)
	}
	if bundle.Extra != "" {
