package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReactionList(baseUrl string, message *TgTypes.MessageType) {
	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		//fmt.Println("File is not found")
		log.Fatalln(err)
	}
	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)

	res := "List of filter in " + message.Chat.Title + ":\n"

	for _, v := range data.Data {
		if v.ChatId == message.Chat.Id {
			for _, filters := range v.Filters {
				res += "-<code>" + filters.Trigger + "</code>\n"
			}
			break
		}
	}

	go Functions.SendTextMessage(baseUrl, res, message.Chat.Id, message.MessageId)
}
