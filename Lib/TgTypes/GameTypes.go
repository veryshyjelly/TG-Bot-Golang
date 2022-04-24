package TgTypes

type GameType struct {
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	Photo        []PhotoSizeType     `json:"photo"`
	Text         string              `json:"text,omitempty"`
	TextEntities []MessageEntityType `json:"text_entities,omitempty"`
	Animation    AnimationType       `json:"animation,omitempty"`
}

type CallBackGameType struct {
	// No Information
}

type CallbackQueryType struct {
	Id              string      `json:"id"`
	From            UserType    `json:"from"`
	Message         MessageType `json:"message,omitempty"`
	InlineMessageId string      `json:"inline_message_id,omitempty"`
	ChatInstance    string      `json:"chat_instance"`
	Data            string      `json:"data,omitempty"`
	GameShortName   string      `json:"game_short_name,omitempty"`
}
