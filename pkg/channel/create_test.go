package channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGroupChannel(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	channel := NewChannel(client)

	createChannelRequest := CreateChannelRequest{
		UserIDs:                []string{"42", "43"},
		Name:                   "name",
		ChannelURL:             "channel-url",
		CoverURL:               "cover-url",
		CustomType:             "custom-type",
		Data:                   `{"key": "value"}`,
		IsDistinct:             true,
		IsPublic:               true,
		IsSuper:                true,
		IsEphemeral:            true,
		AccessCode:             "access-code",
		InviterID:              "inviter-id",
		Strict:                 true,
		OperatorIDs:            []string{"44", "45"},
		BlocSDKUserChannelJoin: true,
	}

	createChannelResponse := &CreateChannelResponse{
		Name: "name",
	}

	client.OnPost("/group_channels", createChannelRequest, &CreateChannelResponse{}).Return(createChannelResponse, nil)

	cur, err := channel.CreateGroupChannel(context.Background(), createChannelRequest)
	require.NoError(t, err)
	assert.Equal(t, createChannelResponse, cur)
}
