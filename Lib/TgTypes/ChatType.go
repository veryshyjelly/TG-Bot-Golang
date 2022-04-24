package TgTypes

type ChatType struct {
	Id                    int64               `json:"id"`
	Type                  string              `json:"type"`
	Title                 string              `json:"title,omitempty"`
	Username              string              `json:"username,omitempty"`
	FirstName             string              `json:"first_name,omitempty"`
	LastName              string              `json:"last_name,omitempty"`
	Photo                 ChatPhotoType       `json:"photo,omitempty"`
	Bio                   string              `json:"bio,omitempty"`
	HasPrivateForwards    bool                `json:"has_private_forwards,omitempty"`
	Description           string              `json:"description,omitempty"`
	InviteLink            string              `json:"invite_link,omitempty"`
	PinnedMessage         *MessageType        `json:"pinned_message,omitempty"`
	Permissions           ChatPermissionsType `json:"permissions,omitempty"`
	SlowModeDelay         int                 `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int                 `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   bool                `json:"has_protected_content,omitempty"`
	StickerSetName        string              `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool                `json:"can_set_sticker_set,omitempty"`
	LinkedChatId          int64               `json:"linked_chat_id,omitempty"`
	Location              ChatLocationType    `json:"location,omitempty"`
}

type ChatLocationType struct {
	Location LocationType `json:"location"`
	Address  string       `json:"address"`
}

type ChatPermissionsType struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessage         bool `json:"can_pin_message,omitempty"`
}

type ChatMemberUpdatedType struct {
	Chat          ChatType           `json:"chat"`
	From          UserType           `json:"from"`
	Date          int                `json:"date"`
	OldChatMember ChatMemberType     `json:"old_chat_member"`
	NewChatMember ChatMemberType     `json:"new_chat_member"`
	InviteLink    ChatInviteLinkType `json:"invite_link,omitempty"`
}

type ChatMemberType struct {
	Status                string   `json:"status,omitempty"`
	User                  UserType `json:"user"`
	CanBeEdited           bool     `json:"can_be_edited,omitempty"`
	IsMember              bool     `json:"is_member"`
	IsAnonymous           bool     `json:"is_anonymous,omitempty"`
	CanManageChat         bool     `json:"can_manage_chat,omitempty"`
	CanDeleteMessages     bool     `json:"can_delete_messages,omitempty"`
	CanManageVoiceChats   bool     `json:"can_manage_voice_chats,omitempty"`
	CanRestrictMembers    bool     `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool     `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool     `json:"can_change_info,omitempty"`
	CanInviteUsers        bool     `json:"can_invite_users,omitempty"`
	CanSendMessages       bool     `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool     `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool     `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool     `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool     `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int      `json:"until_date,omitempty"`
	CanPostMessages       bool     `json:"can_post_messages,omitempty"`
	CanEditMessages       bool     `json:"can_edit_messages,omitempty"`
	CanPinMessages        bool     `json:"can_pin_messages,omitempty"`
	CustomTitle           string   `json:"custom_title,omitempty"`
}

type ChatInviteLinkType struct {
	InviteLink              string   `json:"invite_link"`
	Creator                 UserType `json:"creator"`
	CreatesJoinRequest      bool     `json:"creates_join_request"`
	IsPrimary               bool     `json:"is_primary"`
	IsRevoked               bool     `json:"is_revoked"`
	Name                    string   `json:"name,omitempty"`
	ExpireDate              int      `json:"expire_date,omitempty"`
	MemberLimit             int      `json:"member_limit,omitempty"`
	PendingJoinRequestCount int      `json:"pending_join_request_count,omitempty"`
}

type ChatJoinRequestType struct {
	Chat       ChatType           `json:"chat"`
	From       UserType           `json:"from"`
	Date       int                `json:"date"`
	Bio        string             `json:"bio,omitempty"`
	InviteLink ChatInviteLinkType `json:"invite_link,omitempty"`
}
