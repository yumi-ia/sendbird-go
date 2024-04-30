package channel

import (
	"context"
	"fmt"
	"net/url"
)

// Oh hi mark!

// markAsReadRequest is the request to mark all messages in a group channel as
// read for a specific user.
type markAsReadRequest struct {
	// UserID specifies the ID of the target user.
	UserID string `json:"user_id"`
}

// MarkAsRead marks all messages in a group channel as read for a specific
// user. This action is only applicable for users in a group channel.
// See https://sendbird.com/docs/chat/platform-api/v3/message/read-receipts/mark-all-messages-as-read-message
func (c *channel) MarkAsRead(ctx context.Context, channelURL, userID string) error {
	u, err := url.Parse(fmt.Sprintf("/group_channels/%s/messages/mark_as_read", channelURL))
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	req := markAsReadRequest{
		UserID: userID,
	}

	_, err = c.client.Put(ctx, u.String(), req, nil)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
