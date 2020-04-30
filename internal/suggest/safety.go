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
func NewSafetyCheck(deny, warn []string) (*SafetyCheck, error) {
	sc := &SafetyCheck{}
	for _, p := range deny {
		re, err := regexp.Compile("(?i)" + p)
		if err != nil {
			return nil, err
		}
		sc.Deny = append(sc.Deny, re)
	}
	for _, p := range warn {
		re, err := regexp.Compile("(?i)" + p)
		if err != nil {
			return nil, err
		}
		sc.Warn = append(sc.Warn, re)
	}
	return sc, nil
}

// Evaluate returns a Result with a warning, or an error if the command is
// denied outright.
func (sc *SafetyCheck) Evaluate(cmd string) (Result, error) {
	for _, re := range sc.Deny {
		if re.MatchString(cmd) {
			return Result{}, &DeniedError{Command: cmd, Pattern: re.String()}
		}
	}
	for _, re := range sc.Warn {
		if re.MatchString(cmd) {
			return Result{Command: cmd,
				Warning: "warning: this command may be destructive"},
			nil
		}
	}
	return Result{Command: cmd}, nil
}

// DeniedError marks a command blocked by a deny pattern.
