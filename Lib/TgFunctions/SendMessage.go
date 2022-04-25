package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SendMessageQuery struct {
	ChatId                   int64                       `json:"chat_id"`
	Text                     string                      `json:"text"`
	ReplyToMessageId         int64                       `json:"reply_to_message_id,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []TgTypes.MessageEntityType `json:"entities,omitempty"`
	DisableWebPagePreview    bool                        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendMessageResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func SendTextMessage(baseUrl, text string, chatId int64, replyId int64) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(SendMessageQuery{
		ChatId:           chatId,
		Text:             text,
		ReplyToMessageId: replyId,
		ParseMode:        "HTML",
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/sendMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := SendMessageResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
