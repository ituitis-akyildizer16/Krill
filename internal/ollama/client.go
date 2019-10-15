// Package ollama is a minimal client for the local Ollama server.
package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client talks to a local Ollama instance.
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// New creates a client with the given base URL and timeout.
func New(baseURL string, timeoutSeconds int) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTP: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
	}
}

// GenerateRequest mirrors the Ollama /api/generate payload we use.
type GenerateRequest struct {
	Model    string `json:"model"`
	Prompt   string `json:"prompt"`
	Stream   bool   `json:"stream"`
	Options  any    `json:"options,omitempty"`
	Format   string `json:"format,omitempty"`
