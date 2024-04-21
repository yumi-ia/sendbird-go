package client

import (
	"log/slog"
)

// client is the implementation of the Client interface.
type client struct {
	logger *slog.Logger
}

// Option is the interface for the options of the client.
type Option interface {
	apply(client *client) *client
}

// WithLogger is the option for the logger of the client.
func WithLogger(l *slog.Logger) Option {
	return &withLogger{l}
}

type withLogger struct {
	logger *slog.Logger
}

func (w *withLogger) apply(c *client) *client {
	c.logger = w.logger

	return c
}
