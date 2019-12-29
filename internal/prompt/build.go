// Package prompt builds bounded, git-grounded prompts.
package prompt

import (
	"fmt"
	"strings"
)

// System is the base system prompt for coding questions.
const System = `You are krill, a local-first terminal copilot running on the
user's machine with no network access to any model service.
