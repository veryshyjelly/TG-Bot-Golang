package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func FilterMessage(baseUrl, textBody string, chatId, messageId int64, delay int) error {
	storage, err := ioutil.ReadFile("Data/reactions.json")
	if err != nil {
		return err
	}

	data := FilterDataArray{}
	err = json.Unmarshal(storage, &data)
	if err != nil {
		return err
	}

	textBody = strings.ToLower(textBody)

LOOP:
	for _, v := range data.Data {
		if v.ChatId == chatId {
			for _, filters := range v.Filters {
				if strings.Contains(" "+textBody+" ", " "+filters.Trigger+" ") {
					var repliedMessage *TgTypes.MessageType

					switch filters.FileType {
					case "sticker":
						repliedMessage, err = Functions.SendStickerByUrl(baseUrl, filters.FileId, chatId, messageId, true)

					case "animation":
						repliedMessage, err = Functions.SendAnimationByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)

					case "audio":
						repliedMessage, err = Functions.SendAudioByUrl(baseUrl, filters.FileId, chatId, messageId, filters.Trigger, true)

					case "photo":
						repliedMessage, err = Functions.SendPhotoByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)

					case "video":
						repliedMessage, err = Functions.SendVideoByUrl(baseUrl, filters.FileId, chatId, messageId, "", true)

					case "document":
						repliedMessage, err = Functions.SendDocumentByUrl(baseUrl, filters.FileId, chatId, messageId, filters.Trigger, true)

					}

					if err != nil {
						return err
					}

					switch filters.FileType {
					case "sticker", "animation", "photo":
						go func() {
							_, err := Functions.DelayDelete(baseUrl, delay, repliedMessage.MessageId, repliedMessage.Chat.Id)
							if err != nil {
								log.Println(err)
							}
						}()

					default:
						go func() {
							_, err := Functions.DelayDelete(baseUrl, 600, repliedMessage.MessageId, repliedMessage.Chat.Id)
							if err != nil {
								log.Println(err)
							}
						}()

					}

					break LOOP
				}
			}
		}
	}

	return nil
}
