// Package cli wires the cobra command tree for krill.
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/ituitis-akyildizer16/krill/internal/config"
	"github.com/ituitis-akyildizer16/krill/internal/runner"
)

// Version is overridden at build time via -ldflags.
var Version = "0.9.2"

