package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetChatMemberQuery struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type GetChatMemberResult struct {
	Ok     bool                   `json:"ok"`
	Result TgTypes.ChatMemberType `json:"result"`
}

func GetChatMember(baseUrl string, chatId, userId int64) *TgTypes.ChatMemberType {
	query, err := json.Marshal(GetChatMemberQuery{ChatId: chatId, UserId: userId})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/getChatMember", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := GetChatMemberResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	if !data.Ok {
		return nil
	}

	return &data.Result
}
