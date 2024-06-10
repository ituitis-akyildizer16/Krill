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

// Warn prints a warning line to stderr.
func (r *Renderer) Warn(text string) {
	if r.NoColor {
		fmt.Fprintln(os.Stderr, text)
		return
	}
	fmt.Fprintln(os.Stderr, color.New(color.FgYellow).Sprintf("! %s", text))
}

// Spinner is a minimal progress indicator.
type Spinner struct {
	stop chan struct{}
	done chan struct{}
}

// Start begins a spinner with the given label.
func (r *Renderer) Start(label string) *Spinner {
	sp := &Spinner{stop: make(chan struct{}), done: make(chan struct{})}
	frames := []string{"|", "/", "-", "\\"}
	go func() {
		i := 0
		for {
			select {
			case <-sp.stop:
				close(sp.done)
				return
			default:
				fmt.Fprintf(os.Stderr, "\r%s %s", frames[i%len(frames)], label)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return sp
}

// Stop clears the spinner line.
func (sp *Spinner) Stop() {
	close(sp.stop)
	<-sp.done
	fmt.Fprint(os.Stderr, "\r\x1b[2K")
}

// Markdown strips common markdown tokens for terminal display.
func Markdown(s string) string {
	s = strings.ReplaceAll(s, "**", "")
	s = strings.ReplaceAll(s, "`", "")
	s = strings.ReplaceAll(s, "### ", "")
	return s
}