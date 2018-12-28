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
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

// Store is a directory-based TTL cache.
type Store struct {
	dir string
	ttl time.Duration
}

// New creates a store rooted at dir (created if missing).
func New(dir string, ttl time.Duration) (*Store, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &Store{dir: dir, ttl: ttl}, nil
