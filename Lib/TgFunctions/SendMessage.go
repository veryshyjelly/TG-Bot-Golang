package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SendMessageResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func SendTextMessage(baseUrl, text string, chatId int64, replyId int64) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendMessageQuery)
	sendQuery.ChatId, sendQuery.Text, sendQuery.ReplyToMessageId, sendQuery.ParseMode = chatId, text, replyId, "HTML"

	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/sendMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(body))
	data := SendMessageResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Result
}
