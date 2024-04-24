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
