package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yumi-ia/sendbird-go/pkg/client"
)

func TestGetUnreadMessagesCount(t *testing.T) {
	t.Parallel()

	getUnreadMessagesCountRequest := GetUnreadMessagesCountRequest{
		CustomTypes: []string{"custom-type"},
		SuperMode:   SuperModeSuper,
	}

	getUnreadMessagesCountResponse := &GetUnreadMessagesCountResponse{
		UnreadCount: 42,
	}

	client := client.NewClientMock(t).
		OnGet("/users/42/unread_message_count", getUnreadMessagesCountRequest, &GetUnreadMessagesCountResponse{}).TypedReturns(getUnreadMessagesCountResponse, nil).Once().
		Parent
	user := NewUser(client)

	cur, err := user.GetUnreadMessagesCount(context.Background(), "42", getUnreadMessagesCountRequest)
	require.NoError(t, err)
	assert.Equal(t, getUnreadMessagesCountResponse, cur)
}
