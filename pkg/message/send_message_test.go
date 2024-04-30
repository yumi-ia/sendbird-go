package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptr[T any](t T) *T {
	return &t
}

func TestPtr(t *testing.T) {
	t.Parallel()

	assert.True(t, *ptr(true))
	assert.False(t, *ptr(false))
	assert.Zero(t, *ptr(0))
	assert.Equal(t, "foo", *ptr("foo"))
}

func TestValidateSMR(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		smr       SendMessageRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name:      "empty",
			smr:       SendMessageRequest{},
			assertErr: assert.Error,
		},
		{
			name: "Without Message Type",
			smr: SendMessageRequest{
				UserID:  "42",
				Message: "Hello, World!",
			},
			assertErr: assert.Error,
		},
		{
			name: "Without user id",
			smr: SendMessageRequest{
				MessageType: MessageTypeText,
				Message:     "Hello, World!",
			},
			assertErr: assert.Error,
		},
		{
			name: "Without message",
			smr: SendMessageRequest{
				MessageType: MessageTypeText,
				UserID:      "42",
			},
			assertErr: assert.Error,
		},
		{
			name: "Valid",
			smr: SendMessageRequest{
				MessageType: MessageTypeText,
				UserID:      "42",
				Message:     "Hello, World!",
			},
			assertErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.smr.Validate()
			test.assertErr(t, err)
		})
	}
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	message := NewMessage(client)

	sendMessageRequest := SendMessageRequest{
		MessageType:         MessageTypeText,
		UserID:              "42",
		Message:             "Hello, World!",
		CustomType:          "custom-type",
		Data:                `{ "key": "value" }`,
		SendPush:            ptr(true),
		PushMessageTemplate: "push-notification-template",
		MentionType:         MentionTypeChannels,
		MentionUserIDs:      []string{"mention-user-id"},
		IsSilent:            ptr(true),
		MarkAsRead:          ptr(true),
		SortedMetaArray:     []MetaArray{{Key: "key1", Value: []string{"value1"}}, {Key: "key2", Value: []string{"value2"}}},
		CreatedAt:           42,
		PollID:              42,
		IncludePollDetails:  ptr(true),
		DedupID:             "dedup-id",
		ApnsBundleID:        "apns-bundle-id",
		Sound:               "sound",
		Volume:              0.5,
	}

	sendMessageResponse := &SendMessageResponse{
		MessageID: 42,
	}

	client.OnPost("/group_channels/url/messages", sendMessageRequest, &SendMessageResponse{}).Return(sendMessageResponse, nil)

	cur, err := message.SendMessage(context.Background(), ChannelTypeGroup, "url", sendMessageRequest)
	require.NoError(t, err)
	assert.Equal(t, sendMessageResponse, cur)
}
