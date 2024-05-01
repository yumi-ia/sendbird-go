// Package user package provides the interface for the user service.
// It provides the methods to interact with the sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/user/user-overview.
package user

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type User interface {
	// CreateUser creates a user.
	// See https://sendbird.com/docs/chat/platform-api/v3/user/creating-users/create-a-user
	CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (*CreateUserResponse, error)
	// UpdateUserRequest is the request to update a user.
	// See https://sendbird.com/docs/chat/platform-api/v3/user/managing-users/update-a-user
	UpdateUser(ctx context.Context, userID string, updateUserRequest UpdateUserRequest) (*UpdateUserResponse, error)

	// GetSessionToken retrieves a session token for a user.
	// https://sendbird.com/docs/chat/platform-api/v3/user/managing-session-tokens/issue-a-session-token
	GetSessionToken(ctx context.Context, userID string, getSessionTokenRequest GetSessionTokenRequest) (*GetSessionTokenResponse, error)
	// GetUnreadMessagesCount retrieves the number of unread messages of a user.
	// https://sendbird.com/docs/chat/platform-api/v3/user/managing-unread-count/get-number-of-unread-messages
	GetUnreadMessagesCount(ctx context.Context, userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) (*GetUnreadMessagesCountResponse, error)
	// GetGroupChannelCount retrieves the number of group channels of a user.
	// https://sendbird.com/docs/chat/platform-api/v3/user/getting-group-channel-count/get-number-of-channels-by-join-status
	GetGroupChannelCount(ctx context.Context, userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) (*GetGroupChannelCountResponse, error)
}

type user struct {
	client client.Client
}

func NewUser(c client.Client) User {
	return &user{client: c}
}
