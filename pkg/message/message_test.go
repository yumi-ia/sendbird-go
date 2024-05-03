package message_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/tomMoulard/sendbird-go/pkg/message"
)

// iencli wraps the message.Message interface.
type iencli struct {
	message message.Message
}

// SendMessage wraps the sends message method.
func (c *iencli) SendMessage(ctx context.Context, channelType message.ChannelType, channelURL string, sendMessageRequest message.SendMessageRequest) (*message.SendMessageResponse, error) {
	got, err := c.message.SendMessage(ctx, channelType, channelURL, sendMessageRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return got, nil
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	req := message.SendMessageRequest{
		Message: "hello",
	}
	messageMock := message.NewMessageMock(t).
		OnSendMessage(message.ChannelTypeGroup, "channelURL", req).TypedReturns(&message.SendMessageResponse{}, nil).Once().
		Parent

	c := &iencli{
		message: messageMock,
	}

	_, err := c.SendMessage(context.Background(), message.ChannelTypeGroup, "channelURL", req)
	if err != nil {
		t.Fatalf("failed to send message: %v", err)
	}
}
