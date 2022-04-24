package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ForwardResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func ForwardMessage(baseUrl string, toChatId, fromChatId int64, messageId int64) *TgTypes.MessageType {
	query, err := json.Marshal(TgTypes.ForwardQuery{ChatId: toChatId, FromChatId: fromChatId, MessageId: messageId})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/forwardMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := ForwardResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if !data.Ok {
		return nil
	}

	return &data.Result
}
