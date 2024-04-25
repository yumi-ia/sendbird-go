package user

import (
	"context"
	"fmt"
)

// GetUnreadMessagesCountRequest is the request to get the number of unread
// messages.
type GetUnreadMessagesCountRequest struct {
	// CustomTypes is the list of one or more custom types to filter group
	// channels with the corresponding types.
	CustomTypes []string `json:"custom_types,omitempty"`
	// SuperMode restricts the search scope to either Supergroup channels or
	// non-Supergroup channels or both. Acceptable values are all, super, and
	// nonsuper. If not specified, the default value is all.
	SuperMode string `json:"super_mode,omitempty"`
}

type GetUnreadMessagesCountResponse struct {
	// UnreadCount is the total number of the user's unread messages.
	UnreadCount int `json:"unread_count"`
}

// GetUnreadMessagesCount retrieves the number of unread messages of a user.
// https://sendbird.com/docs/chat/platform-api/v3/user/managing-unread-count/get-number-of-unread-messages
func (u *user) GetUnreadMessagesCount(ctx context.Context, userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) (*GetUnreadMessagesCountResponse, error) {
	path := fmt.Sprintf("/users/%s/unread_message_count", userID)

	gumcr, err := u.client.Get(ctx, path, getUnreadMessagesCountRequest, &GetUnreadMessagesCountResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to get unread messages count: %w", err)
	}

	getUnreadMessagesCountResponse, ok := gumcr.(*GetUnreadMessagesCountResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to GetUnreadMessagesCountResponse: %+v", gumcr)
	}

	return getUnreadMessagesCountResponse, nil
}
