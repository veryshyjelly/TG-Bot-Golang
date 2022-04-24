package TgTypes

type MessageType struct {
	MessageId                     int64                             `json:"message_id"`
	From                          UserType                          `json:"from,omitempty"`
	SenderChat                    ChatType                          `json:"sender_chat,omitempty"`
	Date                          int                               `json:"date"`
	Chat                          ChatType                          `json:"chat"`
	ForwardFrom                   UserType                          `json:"forward_from,omitempty"`
	ForwardFromChat               ChatType                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int                               `json:"forward_from_message_id,omitempty,omitempty"`
	ForwardSignature              string                            `json:"forward_signature,omitempty"`
	ForwardSenderName             string                            `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                               `json:"forward_date,omitempty"`
	IsAutomaticForward            bool                              `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *MessageType                      `json:"reply_to_message,omitempty"`
	ViaBot                        UserType                          `json:"via_bot,omitempty"`
	EditDate                      int                               `json:"edit_date,omitempty"`
	HasProtectedContent           bool                              `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                            `json:"media_group_id,omitempty"`
	AuthorSignature               string                            `json:"author_signature,omitempty"`
	Text                          string                            `json:"text,omitempty"`
	Entities                      []MessageEntityType               `json:"entities,omitempty"`
	Animation                     AnimationType                     `json:"animation,omitempty"`
	Audio                         AudioType                         `json:"audio,omitempty"`
	Document                      DocumentType                      `json:"document,omitempty"`
	Photo                         []PhotoSizeType                   `json:"photo,omitempty"`
	Sticker                       StickerType                       `json:"sticker,omitempty"`
	Video                         VideoType                         `json:"video,omitempty"`
	VideoNote                     VideoNoteType                     `json:"video_note,omitempty"`
	Voice                         VoiceType                         `json:"voice,omitempty"`
	Caption                       string                            `json:"caption,omitempty"`
	CaptionEntities               []MessageEntityType               `json:"caption_entities,omitempty"`
	Contact                       ContactType                       `json:"contact,omitempty"`
	Dice                          DiceType                          `json:"dice,omitempty"`
	Game                          GameType                          `json:"game,omitempty"`
	Poll                          PollType                          `json:"poll,omitempty"`
	Venue                         VenueType                         `json:"venue,omitempty"`
	Location                      LocationType                      `json:"location,omitempty"`
	NewChatMembers                []UserType                        `json:"new_chat_members,omitempty"`
	LeftChatMember                UserType                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                            `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSizeType                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                              `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                              `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                              `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                              `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChangedType `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int64                             `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int64                             `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MessageType                      `json:"pinned_message,omitempty"`
	Invoice                       InvoiceType                       `json:"invoice,omitempty"`
	SuccessfulPayment             SuccessfulPaymentType             `json:"successful_payment,omitempty"`
	ConnectedWebsite              string                            `json:"connected_website,omitempty"`
	PassportData                  PassportDataType                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       ProximityAlertTriggeredType       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatScheduled            VoiceChatScheduledType            `json:"voice_chat_scheduled,omitempty"`
	VoiceChatStarted              VoiceChatStartedType              `json:"voice_chat_started,omitempty"`
	VoiceChatEnded                VoiceChatEndedType                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  VoiceChatParticipantsInvitedType  `json:"voice_chat_participants_invited,omitempty"`
	ReplyMarkup                   InlineKeyboardMarkupType          `json:"reply_markup,omitempty"`
}

type MessageEntityType struct {
	Type     string   `json:"type"` // mention, hashtag, cashtag, botcommand, url, email, phonenumber, bold, italic, underline, strikethrough, spoiler, code, pre, text_link, text_mention
	Offset   int      `json:"offset"`
	Length   int      `json:"length"`
	Url      string   `json:"url,omitempty"`
	User     UserType `json:"user,omitempty"`
	Language string   `json:"language,omitempty"`
}

type MessageAutoDeleteTimerChangedType struct {
	MessageAutoDeleteTimer int `json:"message_auto_delete_timer"`
}
