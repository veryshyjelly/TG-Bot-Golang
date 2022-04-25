package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetChatMemberQuery struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type GetChatMemberResult struct {
	Ok          bool                   `json:"ok"`
	Result      TgTypes.ChatMemberType `json:"result"`
	ErrorCode   int                    `json:"error_code"`
	Description string                 `json:"description"`
}

func GetChatMember(baseUrl string, chatId, userId int64) (*TgTypes.ChatMemberType, error) {
	query, err := json.Marshal(GetChatMemberQuery{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/getChatMember", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := GetChatMemberResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	if !data.Ok {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
