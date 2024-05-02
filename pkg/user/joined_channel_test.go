package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestGetGroupChannelCount(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t)
	user := NewUser(client)

	getGroupChannelCountRequest := GetGroupChannelCountRequest{
		CustomTypes: []string{"custom-type"},
		HiddenMode:  ModeAll,
		State:       StateInvited,
		SuperMode:   SuperModeAll,
	}

	getGroupChannelCountResponse := &GetGroupChannelCountResponse{
		GroupChannelCount: 42,
	}

	client.OnGet("/users/42/group_channel_count", getGroupChannelCountRequest, &GetGroupChannelCountResponse{}).Return(getGroupChannelCountResponse, nil)

	cur, err := user.GetGroupChannelCount(context.Background(), "42", getGroupChannelCountRequest)
	require.NoError(t, err)
	assert.Equal(t, getGroupChannelCountResponse, cur)
}
