package Autoresponder

import (
	"Telegram-Bot/Lib/MediaFunctions"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
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
					var respondId int64
					var sendMode MediaFunctions.MediaType
					switch filters.FileType {
					case "sticker":
						sendMode = MediaFunctions.Sticker
					case "animation":
						sendMode = MediaFunctions.Animation
					case "audio":
						sendMode = MediaFunctions.Audio
					case "photo":
						sendMode = MediaFunctions.Photo
					case "video":
						sendMode = MediaFunctions.Video
					case "document":
						sendMode = MediaFunctions.Document
					case "message":
						copyID, _ := strconv.ParseInt(filters.FileId, 10, 64)
						respondId, err = MessageMethods.CopyMessage(baseUrl, chatId, chatId, copyID, messageId, "", true)

					}

					if filters.FileType != "message" {
						repliedMessage, err = MediaFunctions.SendMediaByUrl(baseUrl, filters.FileId, sendMode, chatId, messageId, "", true)
						respondId = repliedMessage.MessageId
					}

					if err != nil {
						return err
					}

					switch filters.FileType {
					case "sticker", "animation", "photo":
						go func() {
							_, err := MessageMethods.DelayDelete(baseUrl, delay, respondId, chatId)
							if err != nil {
								log.Println(err)
							}
						}()

					default:
						go func() {
							_, err := MessageMethods.DelayDelete(baseUrl, 600, respondId, chatId)
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
