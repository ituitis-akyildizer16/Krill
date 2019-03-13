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
