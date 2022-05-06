package MessageMethods

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ForwardQuery struct {
	ChatId     int64 `json:"chat_id"`
	FromChatId int64 `json:"from_chat_id"`
	MessageId  int64 `json:"message_id"`
}

type ForwardResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func ForwardMessage(baseUrl string, toChatId, fromChatId int64, messageId int64) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(ForwardQuery{
		ChatId:     toChatId,
		FromChatId: fromChatId,
		MessageId:  messageId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/forwardMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := ForwardResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, err
}
