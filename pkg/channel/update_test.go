package channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateGroupChannel(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	channel := NewChannel(client)

	updateChannelRequest := UpdateGroupChannelRequest{
		Name:        "channel-name",
		CoverURL:    "cover-url",
		CustomType:  "custom-type",
		Data:        `{"key": "value"}`,
		IsDistinct:  true,
		IsPublic:    true,
		IsSuper:     true,
		AccessCode:  "access-code",
		OperatorIDs: []string{"42", "43"},
	}

	updateChannelResponse := &UpdateGroupChannelResponse{
		Name: "channel-name",
	}

	client.OnPut("/group_channels/channel-url", updateChannelRequest, &UpdateGroupChannelResponse{}).Return(updateChannelResponse, nil)

	cur, err := channel.UpdateGroupChannel(context.Background(), "channel-url", updateChannelRequest)
	require.NoError(t, err)
	assert.Equal(t, updateChannelResponse, cur)
}
