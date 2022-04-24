package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DeleteResult struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

func DeleteMessage(baseUrl string, ChatId int64, messageId int64) bool {
	query, err := json.Marshal(TgTypes.DeleteMessageQuery{ChatId: ChatId, MessageId: messageId})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/deleteMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := DeleteResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if !data.Ok {
		return false
	}

	return data.Result
}

func DelayDelete(baseUrl string, delay int, messageId, chatId int64) bool {
	time.Sleep(time.Second * time.Duration(delay))
	//fmt.Println("Deleted")
	return DeleteMessage(baseUrl, chatId, messageId)
}
