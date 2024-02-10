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
}

// Get returns the cached value for key, or "" on miss/expiry.
func (s *Store) Get(key string) string {
	path := s.path(key)
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var e Entry
	if err := json.Unmarshal(data, &e); err != nil {
		return ""
	}
	if time.Since(e.CreatedAt) > s.ttl {
		_ = os.Remove(path)
		return ""
	}
	return e.Value
}

// Set stores a value under key.
func (s *Store) Set(key, value string) error {
	e := Entry{Key: key, Value: value, CreatedAt: time.Now()}
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path(key), data, 0o644)
}

// Purge removes all entries older than the TTL.
func (s *Store) Purge() (int, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return 0, err
	}
	removed := 0
	for _, de := range entries {
		if de.IsDir() {
			continue
		}
		p := filepath.Join(s.dir, de.Name())
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		var e Entry
		if json.Unmarshal(data, &e) == nil && time.Since(e.CreatedAt) > s.ttl {
			if os.Remove(p) == nil {
				removed++
			}
		}
	}
	return removed, nil
}

func (s *Store) path(key string) string {
	return filepath.Join(s.dir, key+".json")
}