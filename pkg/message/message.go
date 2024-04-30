// Package message package provides the interface for the message service.
// It provides the methods to interact with the sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/message/message-overview.
package message

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type Message interface {
	SendMessage(ctx context.Context, channelType ChannelType, channelURL string, sendMessageRequest SendMessageRequest) (*SendMessageResponse, error)
	ListMessages(ctx context.Context, channelType ChannelType, channelURL string, listMessagesRequest ListMessagesRequest) (*ListMessagesResponse, error)
}

type message struct {
	client client.Client
}

func NewMessage(c client.Client) Message {
	return &message{client: c}
}
