package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendButtons(baseUrl, text string, chatId, replyId int64, buttons TgTypes.InlineKeyboardMarkupType) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendButtonsQuery)
	sendQuery.ChatId, sendQuery.Text, sendQuery.ReplyToMessageId = chatId, text, replyId
	sendQuery.ParseMode, sendQuery.ReplyMarkup = "HTML", buttons

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

	data := SendMessageResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Result
}
