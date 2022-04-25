package TgTypes

type UpdateType struct {
	UpdateId          int64                  `json:"update_id"`
	Message           MessageType            `json:"message"`
	EditedMessage     MessageType            `json:"edited_message,omitempty"`
	ChannelPost       MessageType            `json:"channel_post,omitempty"`
	EditedChannelPost MessageType            `json:"edited_channel_post"`
	InlineQuery       InlineQueryType        `json:"inline_query,omitempty"`
	Chosen            ChosenInlineResultType `json:"chosen,omitempty"`
	CallbackQuery     CallbackQueryType      `json:"callback_query,omitempty"`
	ShippingQuery     ShippingQueryType      `json:"shipping_query,omitempty"`
	PreCheckoutQuery  PreCheckoutQueryType   `json:"pre_checkout_query"`
	Poll              PollType               `json:"poll,omitempty"`
	MyChatMember      ChatMemberUpdatedType  `json:"my_chat_member,omitempty"`
	ChatMember        ChatMemberUpdatedType  `json:"chat_member,omitempty"`
	ChatJoinRequest   ChatJoinRequestType    `json:"chat_join_request,omitempty"`
}

type SendMessageQuery struct {
	ChatId                   int64               `json:"chat_id"`
	Text                     string              `json:"text"`
	ReplyToMessageId         int64               `json:"reply_to_message_id,omitempty"`
	ParseMode                string              `json:"parse_mode,omitempty"`
	Entities                 []MessageEntityType `json:"entities,omitempty"`
	DisableWebPagePreview    bool                `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendButtonsQuery struct {
	ChatId                   int64                    `json:"chat_id"`
	Text                     string                   `json:"text"`
	ReplyToMessageId         int64                    `json:"reply_to_message_id,omitempty"`
	ParseMode                string                   `json:"parse_mode,omitempty"`
	Entities                 []MessageEntityType      `json:"entities,omitempty"`
	DisableWebPagePreview    bool                     `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                     `json:"disable_notification,omitempty"`
	ProtectContent           bool                     `json:"protect_content,omitempty"`
	AllowSendingWithoutReply bool                     `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type ForwardQuery struct {
	ChatId     int64 `json:"chat_id"`
	FromChatId int64 `json:"from_chat_id"`
	MessageId  int64 `json:"message_id"`
}

type CopyQuery struct {
	ChatId                   int64               `json:"chat_id"`
	FromChatId               int64               `json:"from_chat_id"`
	MessageId                int64               `json:"message_id"`
	Caption                  string              `json:"caption,omitempty"`
	ParseMode                string              `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntityType `json:"caption_entities,omitempty"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageId         int64               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendPhotoQuery struct {
	ChatId                   int64               `json:"chat_id"`
	Photo                    string              `json:"photo"` // Or multipart file
	Caption                  string              `json:"caption,omitempty"`
	ParseMode                string              `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntityType `json:"caption_entities,omitempty"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageId         int64               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendAudioQuery struct {
	ChatId              int64               `json:"chat_id"`
	Audio               string              `json:"audio"` // Or multipart file
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntityType `json:"caption_entities"`
	Duration            int                 `json:"duration,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	Title               string              `json:"title,omitempty"`
	Thumb               string              `json:"thumb,omitempty"` // Or multipart file
	DisableNotification bool                `json:"disable_notification,omitempty"`
	ProtectContent      bool                `json:"protect_content,omitempty"`
	ReplyToMessageId    int64               `json:"reply_to_message_id,omitempty"`
	//ReplyMarkup         InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendDocumentQuery struct {
	ChatId                      int64               `json:"chat_id"`
	Document                    string              `json:"document"`        // Or multipart file
	Thumb                       string              `json:"thumb,omitempty"` // Or multipart file
	Caption                     string              `json:"caption,omitempty"`
	ParseMode                   string              `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntityType `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                `json:"disable_notification,omitempty"`
	ProtectContent              bool                `json:"protect_content,omitempty"`
	ReplyToMessageId            int64               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup                 InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendVideoQuery struct {
	ChatId                   int64               `json:"chat_id"`
	Video                    string              `json:"video"` // Or multipart file
	Duration                 int                 `json:"duration,omitempty"`
	Width                    int                 `json:"width,omitempty"`
	Height                   int                 `json:"height,omitempty"`
	Thumb                    string              `json:"thumb,omitempty"` // Or multipart file
	Caption                  string              `json:"caption,omitempty"`
	ParseMode                string              `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntityType `json:"caption_entities,omitempty"`
	SupportsStreaming        bool                `json:"supports_streaming,omitempty"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageId         int64               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup                 InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendAnimationQuery struct {
	ChatId                   int64               `json:"chat_id"`
	Animation                string              `json:"animation"` // Or multipart file
	Duration                 int                 `json:"duration,omitempty"`
	Width                    int                 `json:"width,omitempty"`
	Height                   int                 `json:"height,omitempty"`
	Thumb                    string              `json:"thumb,omitempty"` // Or multipart file
	Caption                  string              `json:"caption,omitempty"`
	ParseMode                string              `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntityType `json:"caption_entities,omitempty"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageId         int64               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup                 InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendStickerQuery struct {
	ChatId                   int64  `json:"chat_id"`
	Sticker                  string `json:"sticker"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type DeleteMessageQuery struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
}

type GetFileQuery struct {
	FileId string `json:"file_id"`
}

type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
}

type AddStickerQuery struct {
	UserId     int64  `json:"user_id"`
	Name       string `json:"name"`
	PngSticker string `json:"png_sticker,omitempty"`
	Emojis     string `json:"emojis"`
}

type CreateStickerSetQuery struct {
	UserId     int64  `json:"user_id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	PngSticker string `json:"png_sticker"`
	Emojis     string `json:"emojis"`
}
