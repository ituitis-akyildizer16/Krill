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
type Renderer struct {
	NoColor bool
	Out     io.Writer
}

// New returns a renderer writing to stdout.
func New(noColor bool) *Renderer {
	return &Renderer{NoColor: noColor, Out: os.Stdout}
}

// Header prints a section header.
func (r *Renderer) Header(text string) {
	if r.NoColor {
		fmt.Fprintf(r.Out, "== %s ==\n", text)
		return
	}
