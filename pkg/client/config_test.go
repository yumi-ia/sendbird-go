package client

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLogger(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&b, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	client := &client{}

	client = WithLogger(logger).apply(client)
	assert.Equal(t, logger, client.logger)
}
