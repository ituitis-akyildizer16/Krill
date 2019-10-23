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
	System   string `json:"system,omitempty"`
}

// GenerateResponse is a single (non-streaming) generation result.
type GenerateResponse struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	EvalCount int    `json:"eval_count"`
}

// Generate runs a non-streaming generation.
func (c *Client) Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.BaseURL+"/api/generate", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(httpReq)
