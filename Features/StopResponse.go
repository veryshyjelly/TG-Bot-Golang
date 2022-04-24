package features

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func StopResponse(baseUrl, trigger string, chatId, messageId int64) {

	if trigger == "" {
		Functions.SendTextMessage(baseUrl, "What to revoke?", chatId, messageId)
		return
	}

	trigger = strings.ToLower(trigger)

	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		log.Fatalln(err)
	}
	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)

	var done bool
LOOP:
	for key, chatData := range data.Data {
		if chatData.ChatId == chatId {
			for k, stickerData := range chatData.Filters {
				if stickerData.Trigger == trigger {
					copy(chatData.Filters[k:], chatData.Filters[k+1:])
					chatData.Filters = chatData.Filters[:len(chatData.Filters)-1]
					Functions.SendTextMessage(baseUrl, "revoked response for <code>"+trigger+"</code>", chatId, messageId)
					done = true
					data.Data[key] = chatData
					break LOOP
				}
			}
			break LOOP
		}
	}

	if !done {
		Functions.SendTextMessage(baseUrl, "The text was not set to a response.", chatId, messageId)
	}

	byteData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("Data/reactions.json", byteData, 0)
	if err != nil {
		log.Fatalln(err)
	}
}
