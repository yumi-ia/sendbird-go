// This package is for the client of the sendbird API
package client

import (
	"log/slog"
)

// Client is the interface for the client of the sendbird API.
type Client interface{}

// NewClient creates a new client for the sendbird API.
func NewClient(opts ...Option) Client {
	cfg := &client{
		logger: slog.Default(),
	}

	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	return cfg
}
