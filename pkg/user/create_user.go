// This file contains methods for creating a user.
package user

import (
	"context"
	"errors"
	"fmt"
)

// CreateUserRequest is the request to create a user.
type CreateUserRequest struct {
	UserID                string                 `json:"user_id"`
	Nickname              string                 `json:"nickname"`
	ProfileURL            string                 `json:"profile_url"`
	IssueAccessToken      bool                   `json:"issue_access_token"`
	SessionTokenExpiresAt int64                  `json:"session_token_expires_at"`
	Metadata              map[string]interface{} `json:"metadata"`
}

// CreateUserResponse is the response of the create user request.
type CreateUserResponse struct {
	UserID                     string                 `json:"user_id"`
	Nickname                   string                 `json:"nickname"`
	ProfileURL                 string                 `json:"profile_url"`
	AccessToken                string                 `json:"access_token"`
	IsOnline                   bool                   `json:"is_online"`
	IsActive                   bool                   `json:"is_active"`
	IsCreated                  bool                   `json:"is_created"`
	PhoneNumber                string                 `json:"phone_number"`
	RequireAuthForProfileImage bool                   `json:"require_auth_for_profile_image"`
	SessionTokens              []interface{}          `json:"session_tokens"`
	LastSeenAt                 int                    `json:"last_seen_at"`
	DiscoveryKeys              []string               `json:"discovery_keys"`
	PreferredLanguages         []interface{}          `json:"preferred_languages"`
	HasEverLoggedIn            bool                   `json:"has_ever_logged_in"`
	Metadata                   map[string]interface{} `json:"metadata"`
}

func (u *user) CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (*CreateUserResponse, error) {
	cur, err := u.client.Get(ctx, "/user", createUserRequest, &CreateUserResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	createUserResponse, ok := cur.(*CreateUserResponse)
	if !ok {
		return nil, errors.New("failed to cast body to CreateUserResponse")
	}

	return createUserResponse, nil
}
