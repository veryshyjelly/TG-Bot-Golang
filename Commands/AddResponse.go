package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type FileFilterType struct {
	FileId   string `json:"file_id"`
	Trigger  string `json:"trigger"`
	FileType string `json:"file_type"`
}

type ChatFileFilter struct {
	ChatId  int64            `json:"chat_id"`
	Filters []FileFilterType `json:"filters"`
}

type FilterDataArray struct {
	Data []ChatFileFilter `json:"data"`
}

func AddResponse(baseUrl, trigger string, chatId, messageId int64, repliedMessage *TgTypes.MessageType) error {
	if trigger == "" {
		_, err := Functions.SendTextMessage(baseUrl, "Add some text", chatId, messageId)
		return err
	}

	trigger, fileId, fileType := strings.ToLower(trigger), "", ""

	if repliedMessage.Sticker.FileId != "" {
		fileId = repliedMessage.Sticker.FileId
		fileType = "sticker"
	} else if repliedMessage.Animation.FileId != "" {
		fileId = repliedMessage.Animation.FileId
		fileType = "animation"
	} else if repliedMessage.Audio.FileId != "" {
		fileId = repliedMessage.Audio.FileId
		fileType = "audio"
	} else if len(repliedMessage.Photo) > 0 {
		fileId = repliedMessage.Photo[0].FileId
		fileType = "photo"
	} else if repliedMessage.Document.FileId != "" {
		fileType = "document"
		fileId = repliedMessage.Document.FileId
	} else {
		fileType = "message"
		fileId = fmt.Sprint(repliedMessage.MessageId)
	}

	storage, _ := ioutil.ReadFile("Data/reactions.json")
	data := FilterDataArray{}
	err := json.Unmarshal(storage, &data)
	if err != nil {
		return err
	}

	var done bool

	for key, chatData := range data.Data {
		if chatData.ChatId == chatId {
			for k, stickerData := range chatData.Filters {
				if stickerData.Trigger == trigger {
					copy(chatData.Filters[k:], chatData.Filters[k+1:])
					chatData.Filters = chatData.Filters[:len(chatData.Filters)-1]
					data.Data[key] = chatData
					break
				}
			}
			stickerData := FileFilterType{FileId: fileId, Trigger: trigger, FileType: fileType}
			chatData.Filters = append(chatData.Filters, stickerData)
			data.Data[key] = chatData
			done = true
			break
		}
	}

	if !done {
		chatData := ChatFileFilter{
			ChatId:  chatId,
			Filters: []FileFilterType{},
		}
		stickerData := FileFilterType{FileId: fileId, Trigger: trigger, FileType: fileType}
		chatData.Filters = append(chatData.Filters, stickerData)
		data.Data = append(data.Data, chatData)
	}

	byteData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("Data/reactions.json", byteData, 0)
	if err != nil {
		return err
	}

	return nil
}
