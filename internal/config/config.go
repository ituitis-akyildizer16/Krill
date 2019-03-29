// Package config loads and defaults krill's local config.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

// Config is the user-editable krill configuration.
type Config struct {
	Model    string   `toml:"model"`
	Shell    string   `toml:"shell"`
	Theme    string   `toml:"theme"`
	CacheTTL int      `toml:"cache_ttl_seconds"`
	NoColor  bool     `toml:"no_color"`
	Context  Context  `toml:"context"`
	Ollama   Ollama   `toml:"ollama"`
	Suggest  Suggest  `toml:"suggest"`
	Plugins  Plugins  `toml:"plugins"`
	Ignore   []string `toml:"ignore"`
}

// Context controls how much repo context is collected.
type Context struct {
	MaxBlameLines int `toml:"max_blame_lines"`
	MaxDiffLines  int `toml:"max_diff_lines"`
	HistoryDays   int `toml:"history_days"`
}

// Ollama holds the local model server connection.
type Ollama struct {
	URL     string `toml:"url"`
	Timeout int    `toml:"timeout_seconds"`
}

// Suggest tunes the NL-to-command flow.
type Suggest struct {
	MaxTokens     int      `toml:"max_tokens"`
	DenyPatterns  []string `toml:"deny_patterns"`
	WarnPatterns  []string `toml:"warn_patterns"`
	UseLocalModel bool     `toml:"use_local_model"`
}

// Plugins configures the Python plugin SDK bridge.
type Plugins struct {
	Enabled  bool     `toml:"enabled"`
	Dir      string   `toml:"dir"`
	Providers []string `toml:"providers"`
}

// Defaults returns the built-in defaults.
func Defaults() *Config {
	return &Config{
		Model:    "qwen2.5-coder:7b",
		Shell:    "bash",
		Theme:    "dark",
		CacheTTL: 3600,
