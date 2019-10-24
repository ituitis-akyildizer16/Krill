// Package output renders terminal-friendly output.
package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Renderer holds styling state.
