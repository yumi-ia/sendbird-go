package user

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type User interface {
	CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, userID string, updateUserRequest UpdateUserRequest) (*UpdateUserResponse, error)
}

type user struct {
	client client.Client
}

func NewUser(c client.Client) User {
	return &user{client: c}
}
