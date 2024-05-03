package channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestStartTyping(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t).
		OnPost("/group_channels/channel-url/typing", typingRequest{UserIDs: []string{"user-id"}}, nil).TypedReturns(nil, nil).Once().
		Parent
	channel := NewChannel(client)

	err := channel.StartTyping(context.Background(), "channel-url", []string{"user-id"})
	require.NoError(t, err)
}

func TestStopTyping(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t).
		OnDelete("/group_channels/channel-url/typing", typingRequest{UserIDs: []string{"user-id"}}, nil).TypedReturns(nil, nil).Once().
		Parent
	channel := NewChannel(client)

	err := channel.StopTyping(context.Background(), "channel-url", []string{"user-id"})
	require.NoError(t, err)
}
