package message

type MessageType string

const (
	MessageTypeText         MessageType = "MESG"
	MessageTypeFile         MessageType = "FILE"
	MessageTypeAdminMessage MessageType = "ADMM"
)

type ChannelType string

const (
	ChannelTypeOpen  ChannelType = "open_channels"
	ChannelTypeGroup ChannelType = "group_channels"
)

type MentionType string

const (
	MentionTypeUsers    MentionType = "users"
	MentionTypeChannels MentionType = "channels"
)

const (
	OperatorFilterAll         string = "all"
	OperatorFilterOperator    string = "operator"
	OperatorFilterNonOperator string = "nonoperator"
)

const (
	ReplyTypeNone               string = "NONE"
	RelyTypeAll                 string = "ALL"
	ReplyTypeOnlyReplyToChannel string = "ONLY_REPLY_TO_CHANNEL"
)

type User struct {
	UserID     string                 `json:"user_id"`
	Nickname   string                 `json:"nickname"`
	ProfileURL string                 `json:"profile_url"`
	Metadata   map[string]interface{} `json:"metadata"`
}

type PollOption struct {
	Text      string `json:"text"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	VoteCount int    `json:"vote_count"`
	PollID    int    `json:"poll_id"`
	ID        int    `json:"id"`
}

type Poll struct {
	ID                  int          `json:"id"`
	Title               string       `json:"title"`
	Status              string       `json:"status"`
	AllowUserSuggestion bool         `json:"allow_user_suggestion"`
	Data                string       `json:"data"`
	AllowMultipleVotes  bool         `json:"allow_multiple_votes"`
	CreatedAt           int          `json:"created_at"`
	UpdatedAt           int          `json:"updated_at"`
	CreatedBy           string       `json:"created_by"`
	VoterCount          int          `json:"voter_count"`
	CloseAt             int          `json:"close_at"`
	Options             []PollOption `json:"options"`
}

type MessageEvents struct {
	SendPushNotification string `json:"send_push_notification"`
	UpdateUnreadCount    bool   `json:"update_unread_count"`
	UpdateMentionCount   bool   `json:"update_mention_count"`
	UpdateLastMessage    bool   `json:"update_last_message"`
}

// MessageResource is the resource of a message.
type MessageResource struct {
	MessageID            int           `json:"message_id"`
	Type                 string        `json:"type"`
	CustomType           string        `json:"custom_type"`
	ChannelURL           string        `json:"channel_url"`
	User                 User          `json:"user"`
	MentionType          string        `json:"mention_type"`
	MentionedUsers       []User        `json:"mentioned_users"`
	IsRemoved            bool          `json:"is_removed"`
	Message              string        `json:"message"`
	Data                 string        `json:"data"`
	Poll                 Poll          `json:"poll"`
	MessageEvents        MessageEvents `json:"message_events"`
	CreatedAt            int64         `json:"created_at"`
	UpdatedAt            int           `json:"updated_at"`
	IsAppleCriticalAlert bool          `json:"is_apple_critical_alert"`
}

type MetaArray struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}
