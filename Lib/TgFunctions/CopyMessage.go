package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CopyResult struct {
	Ok        bool  `json:"ok"`
	MessageId int64 `json:"message_id"`
}

func CopyMessage(baseUrl string, chatId, fromChatId, messageId, replyId int64, caption string) int64 {
	sendQuery := new(TgTypes.CopyQuery)
	sendQuery.ChatId, sendQuery.FromChatId, sendQuery.MessageId = chatId, fromChatId, messageId
	sendQuery.ReplyToMessageId, sendQuery.Caption = replyId, caption
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/copyMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := CopyResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data.MessageId
}
