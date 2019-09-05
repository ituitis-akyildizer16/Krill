// Package git wraps the git CLI for context collection.
//
// Everything here shells out to `git` (the source of truth); we only
// parse its output. Commands are bounded by context, never by timeouts
// that would corrupt the output.
package git

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)
