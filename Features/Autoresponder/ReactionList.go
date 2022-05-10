package Autoresponder

import (
	Functions "Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
)

func ReactionList(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		return nil, err
	}

	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)
	if err != nil {
		return nil, err
	}

	res := "List of filter in " + message.Chat.Title + ":\n"

	for _, v := range data.Data {
		if v.ChatId == message.Chat.Id {
			for _, filters := range v.Filters {
				res += " -<code>" + filters.Trigger + "</code>\n"
			}
			break
		}
	}

	return Functions.SendTextMessage(res, message.Chat.Id, message.MessageId)
}
