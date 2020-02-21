// Package runner orchestrates commands: context collection, prompt
// building, model calls, and output.
package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ituitis-akyildizer16/krill/internal/cache"
	"github.com/ituitis-akyildizer16/krill/internal/config"
	"github.com/ituitis-akyildizer16/krill/internal/git"
	"github.com/ituitis-akyildizer16/krill/internal/ollama"
	"github.com/ituitis-akyildizer16/krill/internal/output"
	"github.com/ituitis-akyildizer16/krill/internal/prompt"
	"github.com/ituitis-akyildizer16/krill/internal/suggest"
)

// Version matches the CLI's build-time version.
var Version = "0.9.2"

// Runner holds the pieces one command needs.
type Runner struct {
