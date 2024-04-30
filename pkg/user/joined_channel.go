package user

import (
	"context"
	"fmt"
)

const (
	StateJoined  string = "joined"
	StateInvited string = "invited"
)

// GetGroupChannelCountRequest is the request to get the number of unread
// messages.
type GetGroupChannelCountRequest struct {
	// CustomTypes is the list of one or more custom types to filter group
	// channels with the corresponding types.
	CustomTypes []string `json:"custom_types,omitempty"`
	// HiddenMode restricts the search scope to group channels that match a
	// specific hidden_status and operating behavior. Acceptable values are the
	// following:
	// - ModeUnHiddenOnly (default): Specifies channels which the user joined
	// with the unhidden status.
	// - ModeHiddenOnly: Specifies all channels which the user joined with either
	// the hidden_allow_auto_unhide or hidden_prevent_auto_unhide status.
	// - ModeHiddenAllowAutoUnhide: Specifies channels which the user joined with
	// the hidden_allow_auto_unhide status.
	// - ModeHiddenPreventAutoUnhide: Specifies channels which the user joined
	// with the hidden_prevent_auto_unhide status.
	// - ModeAll: Specifies all channels regardless of their hidden_status.
	HiddenMode Mode `json:"hidden_mode,omitempty"`
	// State determines which join status to use to filter the user's group
	// channels and count the total number. Valid values are the following:
	// - joined: Indicates the number count of the user’s joined channels.
	// - invited: Indicates the number count of channels which the user has been
	// invited to but not joined.
	State string `json:"state,omitempty"`
	// SuperMode restricts the search scope to either Supergroup channels or
	// non-Supergroup channels or both. Acceptable values are all, super, and
	// nonsuper. If not specified, the default value is all.
	SuperMode SuperMode `json:"super_mode,omitempty"`
}

type GetGroupChannelCountResponse struct {
	// GroupChannelCount is the total number of the user's group channels.
	GroupChannelCount int `json:"group_channel_count"`
}

// GetGroupChannelCount retrieves the number of group channels of a user.
// https://sendbird.com/docs/chat/platform-api/v3/user/getting-group-channel-count/get-number-of-channels-by-join-status
func (u *user) GetGroupChannelCount(ctx context.Context, userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) (*GetGroupChannelCountResponse, error) {
	path := fmt.Sprintf("/users/%s/group_channel_count", userID)

	ggccr, err := u.client.Get(ctx, path, getGroupChannelCountRequest, &GetGroupChannelCountResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to get unread messages count: %w", err)
	}

	getGroupChannelCountResponse, ok := ggccr.(*GetGroupChannelCountResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to GetGroupChannelCountResponse: %+v", ggccr)
	}

	return getGroupChannelCountResponse, nil
}
