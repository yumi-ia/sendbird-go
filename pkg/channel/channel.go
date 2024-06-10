// Package channel package provides the interface for the channel service.
// It provides the methods to interact with the sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/channel/channel-overview.
package channel

import (
	"context"

	"github.com/yumi-ia/sendbird-go/pkg/client"
)

type Channel interface {
	// CreateGroupChannel creates a group/super channel.
	// https://sendbird.com/docs/chat/platform-api/v3/channel/creating-a-channel/create-a-group-channel
	CreateGroupChannel(ctx context.Context, createChannelRequest CreateGroupChannelRequest) (*CreateGroupChannelResponse, error)
	// UpdateGroupChannel updates a group channel.
	// See https://sendbird.com/docs/chat/platform-api/v3/channel/managing-a-channel/update-a-group-channel
	UpdateGroupChannel(ctx context.Context, channelURL string, updateChannelRequest UpdateGroupChannelRequest) (*UpdateGroupChannelResponse, error)
	// ListGroupChannels lists group channels.
	// See https://sendbird.com/docs/chat/platform-api/v3/channel/listing-channels-in-an-application/list-group-channels
	ListGroupChannels(ctx context.Context, listChannelRequest ListGroupChannelRequest) (*ListGroupChannelResponse, error)
	// MarkAsRead marks all messages in a group channel as read for a specific
	// user. This action is only applicable for users in a group channel.
	// See https://sendbird.com/docs/chat/platform-api/v3/message/read-receipts/mark-all-messages-as-read-message
	MarkAsRead(ctx context.Context, channelURL, userID string) error
	// StartTyping start showing a typing indicator for the channel. This feature
	// is only applicable to group channels.
	// When there are 100 or more members in a group channel, typing indicator
	// events of up to two users are delivered to other users. Once typing
	// indicator events of three or more users occur, the events aren't delivered
	// to users.
	// See https://docs.sendbird.com/docs/chat/platform-api/v3/channel/managing-typing-indicators/start-typing-indicators
	StartTyping(ctx context.Context, channelURL string, userIDs []string) error
	// StopTyping stop showing a typing indicator using this API. To signal that a
	// user is no longer typing, you can let the indicator disappear when the user
	// sends a message or completely deletes the message text.
	// See https://docs.sendbird.com/docs/chat/platform-api/v3/channel/managing-typing-indicators/stop-typing-indicators
	StopTyping(ctx context.Context, channelURL string, userIDs []string) error
}

type channel struct {
	client client.Client
}

func NewChannel(c client.Client) Channel {
	return &channel{client: c}
}
