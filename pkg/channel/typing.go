package channel

import (
	"context"
	"fmt"
	"net/url"
)

// typingRequest is the request to start showing a typing indicator for
// the channel.
type typingRequest struct {
	// UserIDs specifies an array of IDs of users who are to use the typing
	// indicator. You can list up to ten user IDs.
	UserIDs []string `json:"user_ids"`
}

// StartTyping start showing a typing indicator for the channel. This feature
// is only applicable to group channels.
// When there are 100 or more members in a group channel, typing indicator
// events of up to two users are delivered to other users. Once typing
// indicator events of three or more users occur, the events aren't delivered
// to users.
// See https://docs.sendbird.com/docs/chat/platform-api/v3/channel/managing-typing-indicators/start-typing-indicators
func (c *channel) StartTyping(ctx context.Context, channelURL string, userIDs []string) error {
	u, err := url.Parse(fmt.Sprintf("/group_channels/%s/typing", channelURL))
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	req := typingRequest{
		UserIDs: userIDs,
	}

	_, err = c.client.Post(ctx, u.String(), req, nil)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// StopTyping stop showing a typing indicator using this API. To signal that a
// user is no longer typing, you can let the indicator disappear when the user
// sends a message or completely deletes the message text.
// See https://docs.sendbird.com/docs/chat/platform-api/v3/channel/managing-typing-indicators/stop-typing-indicators
func (c *channel) StopTyping(ctx context.Context, channelURL string, userIDs []string) error {
	u, err := url.Parse(fmt.Sprintf("/group_channels/%s/typing", channelURL))
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	req := typingRequest{
		UserIDs: userIDs,
	}

	_, err = c.client.Delete(ctx, u.String(), req, nil)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
