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
	fmt.Fprintln(r.Out, color.New(color.Bold, color.FgCyan).Sprintf("== %s ==", text))
}

// Dim prints dimmed text.
func (r *Renderer) Dim(text string) {
	if r.NoColor {
		fmt.Fprintln(r.Out, text)
		return
	}
	fmt.Fprintln(r.Out, color.New(color.FgHiBlack).Sprint(text))
}
