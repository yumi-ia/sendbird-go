package user

import (
	"context"
	"fmt"
)

// GetSessionTokenRequest is the request to get a session token.
type GetSessionTokenRequest struct {
	// ExpiresAt specifies the expiration time of the new session token in Unix
	// milliseconds format. By default, the expiration time of a session token is
	// seven days from the timestamp when the token was issued.
	ExpiresAt int64 `json:"expires_at"`
}

// GetSessionTokenResponse is the response to get a session token.
type GetSessionTokenResponse struct {
	// Token is the session token.
	Token string `json:"token"`
	// ExpiresAt is the expiration time of the session token in Unix milliseconds.
	ExpiresAt int `json:"expires_at"`
}

// GetSessionToken retrieves a session token for a user.
// https://sendbird.com/docs/chat/platform-api/v3/user/managing-session-tokens/issue-a-session-token
func (u *user) GetSessionToken(ctx context.Context, userID string, getSessionTokenRequest GetSessionTokenRequest) (*GetSessionTokenResponse, error) {
	path := fmt.Sprintf("/users/%s/token", userID)

	gstr, err := u.client.Post(ctx, path, getSessionTokenRequest, &GetSessionTokenResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to get session token: %w", err)
	}

	createUserResponse, ok := gstr.(*GetSessionTokenResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to GetSessionTokenResponse: %+v", gstr)
	}

	return createUserResponse, nil
}
