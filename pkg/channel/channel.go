// Package channel package provides the interface for the channel service.
// It provides the methods to interact with the sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/channel/channel-overview.
package channel

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type Channel interface {
	CreateGroupChannel(ctx context.Context, createChannelRequest CreateChannelRequest) (*CreateChannelResponse, error)
	UpdateGroupChannel(ctx context.Context, channelURL string, updateChannelRequest UpdateChannelRequest) (*UpdateChannelResponse, error)
	ListGroupChannels(ctx context.Context, channelURL string, listChannelRequest ListChannelRequest) (*ListChannelResponse, error)
}

type channel struct {
	client client.Client
}

func NewChannel(c client.Client) Channel {
	return &channel{client: c}
}
