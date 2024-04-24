package message

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type Message interface {
	SendMessage(ctx context.Context, channelType, channelURL string, sendMessageRequest SendMessageRequest) (*SendMessageResponse, error)
	ListMessages(ctx context.Context, channelType, channelURL string, listMessagesRequest ListMessagesRequest) (*ListMessagesResponse, error)
}

type message struct {
	client client.Client
}

func NewMessage(c client.Client) Message {
	return &message{client: c}
}
