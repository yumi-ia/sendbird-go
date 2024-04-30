package channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarkAsRead(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	channel := NewChannel(client)

	client.OnPut("/group_channels/channel-url/messages/mark_as_read", markAsReadRequest{UserID: "user-id"}, nil).Return(nil, nil)

	err := channel.MarkAsRead(context.Background(), "channel-url", "user-id")
	require.NoError(t, err)
}
