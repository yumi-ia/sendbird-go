package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSessionToken(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	user := NewUser(client)

	getSessionTokenRequest := GetSessionTokenRequest{
		ExpiresAt: 0,
	}

	getSessionTokenResponse := &GetSessionTokenResponse{
		Token:     "token",
		ExpiresAt: 0,
	}

	client.OnPost("/users/42/token", getSessionTokenRequest, &GetSessionTokenResponse{}).Return(getSessionTokenResponse, nil)

	cur, err := user.GetSessionToken(context.Background(), "42", getSessionTokenRequest)
	require.NoError(t, err)
	assert.Equal(t, getSessionTokenResponse, cur)
}
