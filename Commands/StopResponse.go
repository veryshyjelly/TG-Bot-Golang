package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"encoding/json"
	"io/ioutil"
	"strings"
)

func StopResponse(baseUrl, trigger string, chatId, messageId int64) (bool, error) {

	if trigger == "" {
		_, err := Functions.SendTextMessage(baseUrl, "What to revoke?", chatId, messageId)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	trigger = strings.ToLower(trigger)

	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		return false, err
	}

	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)
	if err != nil {
		return false, err
	}

	var done bool
LOOP:
	for key, chatData := range data.Data {
		if chatData.ChatId == chatId {
			for k, stickerData := range chatData.Filters {
				if stickerData.Trigger == trigger {
					copy(chatData.Filters[k:], chatData.Filters[k+1:])
					chatData.Filters = chatData.Filters[:len(chatData.Filters)-1]
					_, err = Functions.SendTextMessage(baseUrl, "revoked response for <code>"+trigger+"</code>", chatId, messageId)
					if err != nil {
						return false, err
					}

					done = true
					data.Data[key] = chatData
					break LOOP
				}
			}
			break LOOP
		}
	}

	if !done {
		_, err = Functions.SendTextMessage(baseUrl, "The text was not set to a response.", chatId, messageId)
		if err != nil {
			return false, err
		}
	}

	byteData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile("Data/reactions.json", byteData, 0)
	if err != nil {
		return false, err
	}

	return true, nil
}
