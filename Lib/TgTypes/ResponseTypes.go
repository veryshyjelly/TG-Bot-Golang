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
