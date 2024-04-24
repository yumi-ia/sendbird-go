package channel

const (
	DistrincModeAll         string = "all"
	DistrincModeDistinct    string = "distinct"
	DistrincModeNonDistinct string = "nondistinct"
)

const (
	PublicModeAll       string = "all"
	PublicModePublic    string = "public"
	PublicModeNonPublic string = "private"
)

const (
	SuperModeAll      string = "all"
	SuperModeSuper    string = "super"
	SuperModeNonSuper string = "nonsuper"
)

const (
	OrderChronological             string = "chronological"
	OrderLatestLastMessage         string = "latest_last_message"
	OderChannelNameAlphabetical    string = "channel_name_alphabetical"
	OrderMetadataValueAlphabetical string = "metadata_value_alphabetical"
)

const (
	QueryTypeAnd string = "AND"
	QueryTypeOr  string = "OR"
)

type CreatedBy struct {
	UserID                     string `json:"user_id"`
	Nickname                   string `json:"nickname"`
	ProfileURL                 string `json:"profile_url"`
	RequireAuthForProfileImage bool   `json:"require_auth_for_profile_image"`
}

type Member struct {
	UserID     string                 `json:"user_id"`
	Nickname   string                 `json:"nickname"`
	ProfileURL string                 `json:"profile_url"`
	IsActive   bool                   `json:"is_active"`
	IsOnline   bool                   `json:"is_online"`
	LastSeenAt int64                  `json:"last_seen_at"`
	State      string                 `json:"state"`
	Role       string                 `json:"role"`
	Metadata   map[string]interface{} `json:"metadata"`
}

type Operator struct {
	UserID     string                 `json:"user_id"`
	Nickname   string                 `json:"nickname"`
	ProfileURL string                 `json:"profile_url"`
	IsActive   bool                   `json:"is_active"`
	IsOnline   bool                   `json:"is_online"`
	LastSeenAt int                    `json:"last_seen_at"`
	State      string                 `json:"state"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// ChannelResource is the resource of a channel.
type ChannelResource struct {
	Name                   string      `json:"name"`
	ChannelURL             string      `json:"channel_url"`
	CoverURL               string      `json:"cover_url"`
	CustomType             string      `json:"custom_type"`
	UnreadMessageCount     int         `json:"unread_message_count"`
	Data                   string      `json:"data"`
	IsDistinct             bool        `json:"is_distinct"`
	IsPublic               bool        `json:"is_public"`
	IsSuper                bool        `json:"is_super"`
	IsEphemeral            bool        `json:"is_ephemeral"`
	IsAccessCodeRequired   bool        `json:"is_access_code_required"`
	MemberCount            int         `json:"member_count"`
	JoinedMemberCount      int         `json:"joined_member_count"`
	UnreadMentionCount     int         `json:"unread_mention_count"`
	CreatedBy              CreatedBy   `json:"created_by"`
	Members                []Member    `json:"members"`
	Operators              []Operator  `json:"operators"`
	LastMessage            interface{} `json:"last_message"`
	MessageSurvivalSeconds int         `json:"message_survival_seconds"`
	MaxLengthMessage       int         `json:"max_length_message"`
	CreatedAt              int         `json:"created_at"`
	Freeze                 bool        `json:"freeze"`
}
