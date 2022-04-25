package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func FilterMessage(baseUrl, textBody string, chatId, messageId int64, delay int) {
	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		//fmt.Println("File is not found")
		log.Fatalln(err)
	}
	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)

	textBody = strings.ToLower(textBody)

LOOP:
	for _, v := range data.Data {
		if v.ChatId == chatId {
			for _, filters := range v.Filters {
				if strings.Contains(" "+textBody+" ", " "+filters.Trigger+" ") {
					var repliedMessage *TgTypes.MessageType
					switch filters.FileType {
					case "sticker":
						repliedMessage = Functions.SendStickerByUrl(baseUrl, filters.FileId, chatId, messageId, true)
					case "animation":
						repliedMessage = Functions.SendAnimationByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)
					case "audio":
						repliedMessage = Functions.SendAudioByUrl(baseUrl, filters.FileId, chatId, messageId, filters.Trigger, true)
					case "photo":
						repliedMessage = Functions.SendPhotoByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)
					case "video":
						repliedMessage = Functions.SendVideoByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)
					case "document":
						repliedMessage = Functions.SendDocumentByUrl(baseUrl, filters.FileId, chatId, messageId, filters.Trigger, true)
					}
					switch filters.FileType {
					case "sticker", "animation", "photo":
						go Functions.DelayDelete(baseUrl, delay, repliedMessage.MessageId, repliedMessage.Chat.Id)
					default:
						go Functions.DelayDelete(baseUrl, 600, repliedMessage.MessageId, repliedMessage.Chat.Id)
					}
					break LOOP
				}
			}
		}
	}
}
