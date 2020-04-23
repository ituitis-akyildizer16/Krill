// Package suggest turns natural language into shell commands with
// safety checks layered on top.
package suggest

import (
	"regexp"
	"strings"
)

// Result carries the suggested command and any safety warning.
type Result struct {
	Command string
	Warning string
}

// SafetyCheck applies deny/warn patterns to a suggested command.
type SafetyCheck struct {
	Deny []*regexp.Regexp
	Warn []*regexp.Regexp
}

// NewSafetyCheck compiles the given patterns.
