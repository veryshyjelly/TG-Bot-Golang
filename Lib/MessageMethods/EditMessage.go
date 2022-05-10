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

type EditTextQuery struct {
	ChatId                int64                            `json:"chat_id,omitempty"`
	MessageId             int64                            `json:"message_id,omitempty"`
	InlineMessageId       string                           `json:"inline_message_id,omitempty"`
	Text                  string                           `json:"text"`
	ParseMode             string                           `json:"parse_mode,omitempty"`
	Entities              []TgTypes.MessageEntityType      `json:"entities,omitempty"`
	DisableWebPagePreview bool                             `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           TgTypes.InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type EditCaptionQuery struct {
	ChatId          int64                            `json:"chat_id"`
	MessageId       int64                            `json:"message_id"`
	InlineMessageId string                           `json:"inline_message_id"`
	Caption         string                           `json:"caption"`
	ParseMode       string                           `json:"parse_mode"`
	CaptionEntities []TgTypes.MessageEntityType      `json:"caption_entities"`
	ReplyMarkup     TgTypes.InlineKeyboardMarkupType `json:"reply_markup"`
}

type EditMarkupQuery struct {
	ChatId          int64                            `json:"chat_id"`
	MessageId       int64                            `json:"message_id"`
	InlineMessageId string                           `json:"inline_message_id"`
	ReplyMarkup     TgTypes.InlineKeyboardMarkupType `json:"reply_markup"`
}

type EditResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func EditMessageText(text string, chatId, messageId int64, inlineMessId string) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(EditTextQuery{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessId,
		Text:            text,
		ParseMode:       "HTML",
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/editMessageText", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := EditResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}

func EditMessageCaption(caption string, chatId, messageId int64, inlineMessId string) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(EditCaptionQuery{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessId,
		Caption:         caption,
		ParseMode:       "HTML",
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/editMessageCaption", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := EditResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}

func EditMessageMarkup(chatId, messageId int64, replyMarkup TgTypes.InlineKeyboardMarkupType, inlineMessId string) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(EditMarkupQuery{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessId,
		ReplyMarkup:     replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/editMessageReplyMarkup", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := EditResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
