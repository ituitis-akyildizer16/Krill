// Package cache provides a small on-disk TTL cache for model responses.
package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Entry is a single cached item.
type Entry struct {
	Key       string    `json:"key"`
