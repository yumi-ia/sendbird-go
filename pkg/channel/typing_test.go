package channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestStartTyping(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t)
	channel := NewChannel(client)

	client.OnPost("/group_channels/channel-url/typing", typingRequest{UserIDs: []string{"user-id"}}, nil).Return(nil, nil)

	err := channel.StartTyping(context.Background(), "channel-url", []string{"user-id"})
	require.NoError(t, err)
}

func TestStopTyping(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t)
	channel := NewChannel(client)

	client.OnDelete("/group_channels/channel-url/typing", typingRequest{UserIDs: []string{"user-id"}}, nil).Return(nil, nil)

	err := channel.StopTyping(context.Background(), "channel-url", []string{"user-id"})
	require.NoError(t, err)
}
