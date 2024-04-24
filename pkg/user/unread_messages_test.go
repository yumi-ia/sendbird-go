package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUnreadMessagesCount(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	user := NewUser(client)

	getUnreadMessagesCountRequest := GetUnreadMessagesCountRequest{
		CustomTypes: []string{"custom-type"},
		SuperMode:   SuperModeSuper,
	}

	getUnreadMessagesCountResponse := &GetUnreadMessagesCountResponse{
		UnreadCount: 42,
	}

	client.OnGet("/users/42/unread_message_count", getUnreadMessagesCountRequest, &GetUnreadMessagesCountResponse{}).Return(getUnreadMessagesCountResponse, nil)

	cur, err := user.GetUnreadMessagesCount(context.Background(), "42", getUnreadMessagesCountRequest)
	require.NoError(t, err)
	assert.Equal(t, getUnreadMessagesCountResponse, cur)
}
