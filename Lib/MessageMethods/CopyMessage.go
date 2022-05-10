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

type CopyResult struct {
	Ok          bool   `json:"ok"`
	Result      int64  `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type CopyQuery struct {
	ChatId                   int64                       `json:"chat_id"`
	FromChatId               int64                       `json:"from_chat_id"`
	MessageId                int64                       `json:"message_id"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []TgTypes.MessageEntityType `json:"caption_entities,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

func CopyMessage(chatId, fromChatId, messageId, replyId int64, caption string, isProtected bool) (int64, error) {
	query, err := json.Marshal(CopyQuery{
		ChatId:           chatId,
		FromChatId:       fromChatId,
		MessageId:        messageId,
		Caption:          caption,
		ParseMode:        "HTML",
		ProtectContent:   isProtected,
		ReplyToMessageId: replyId,
	})
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/copyMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	data := CopyResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	if !data.Ok {
		return 0, errors.New(data.Description)
	}

	return data.Result, nil
}
