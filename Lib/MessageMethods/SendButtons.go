package MessageMethods

import (
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SendButtonsQuery struct {
	ChatId                   int64                            `json:"chat_id"`
	Text                     string                           `json:"text"`
	ReplyToMessageId         int64                            `json:"reply_to_message_id,omitempty"`
	ParseMode                string                           `json:"parse_mode,omitempty"`
	Entities                 []TgTypes.MessageEntityType      `json:"entities,omitempty"`
	DisableWebPagePreview    bool                             `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                             `json:"disable_notification,omitempty"`
	ProtectContent           bool                             `json:"protect_content,omitempty"`
	AllowSendingWithoutReply bool                             `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              TgTypes.InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendButtonResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func SendButtons(text string, chatId, replyId int64, buttons TgTypes.InlineKeyboardMarkupType, isProtected bool) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(SendButtonsQuery{
		ChatId:           chatId,
		Text:             text,
		ReplyToMessageId: replyId,
		ParseMode:        "HTML",
		ReplyMarkup:      buttons,
		ProtectContent:   isProtected,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/sendMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := SendButtonResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}

type SendButtonImageQuery struct {
	ChatId                   int64                            `json:"chat_id"`
	Photo                    string                           `json:"photo"`
	Caption                  string                           `json:"caption"`
	ParseMode                string                           `json:"parse_mode,omitempty"`
	CaptionEntities          []TgTypes.MessageEntityType      `json:"caption_entities"`
	DisableNotification      bool                             `json:"disable_notification,omitempty"`
	ProtectContent           bool                             `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                             `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              TgTypes.InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

func SendButtonImage(photoUrl, caption string, chatId, replyId int64, buttons TgTypes.InlineKeyboardMarkupType, isProtected bool) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(SendButtonImageQuery{
		ChatId:           chatId,
		Photo:            photoUrl,
		Caption:          caption,
		ReplyToMessageId: replyId,
		ParseMode:        "HTML",
		ReplyMarkup:      buttons,
		ProtectContent:   isProtected,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/sendPhoto", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := SendButtonResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
