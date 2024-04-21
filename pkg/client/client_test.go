package client

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient()
	cClient, ok := c.(*client)
	require.True(t, ok)
	assert.Equal(t, slog.Default(), cClient.logger)
}
