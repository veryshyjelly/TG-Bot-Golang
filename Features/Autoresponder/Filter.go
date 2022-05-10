package Autoresponder

import (
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func FilterMessage(textBody string, chatId, messageId int64, delay int) error {
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
					var respondId int64
					var sendMode Functions.MediaType
					switch filters.FileType {
					case "sticker":
						sendMode = Functions.Sticker
					case "animation":
						sendMode = Functions.Animation
					case "audio":
						sendMode = Functions.Audio
					case "photo":
						sendMode = Functions.Photo
					case "video":
						sendMode = Functions.Video
					case "document":
						sendMode = Functions.Document
					case "message":
						copyID, _ := strconv.ParseInt(filters.FileId, 10, 64)
						respondId, err = MessageMethods.CopyMessage(chatId, chatId, copyID, messageId, "", true)

					}

					if filters.FileType != "message" {
						repliedMessage, err = Functions.SendMediaByUrl(filters.FileId, sendMode, chatId, messageId, "", true)
						respondId = repliedMessage.MessageId
					}

					if err != nil {
						return err
					}

					switch filters.FileType {
					case "sticker", "animation", "photo":
						go func() {
							_, err := MessageMethods.DelayDelete(delay, respondId, chatId)
							if err != nil {
								log.Println(err)
							}
						}()

					default:
						go func() {
							_, err := MessageMethods.DelayDelete(600, respondId, chatId)
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
