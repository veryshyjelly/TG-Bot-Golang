package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RemoveStickerResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type RemoveStickerQuery struct {
	Sticker string `json:"sticker"`
}

type ChatStickerSet struct {
	ChatId int64  `json:"chat_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}

type CreatedSticker struct {
	Data []ChatStickerSet `json:"data"`
}

func RemoveSticker(baseUrl string, chatId int64, messageId int64, repliedMessage *TgTypes.MessageType) (bool, error) {
	if repliedMessage == nil || repliedMessage.Sticker.FileId == "" {
		_, err := SendTextMessage(baseUrl, "Reply to the sticker.", chatId, messageId)
		return false, err
	}

	//fmt.Println(repliedMessage.Sticker.SetName)
	if repliedMessage.Sticker.SetName != "x"+fmt.Sprint(uint64(repliedMessage.Chat.Id))+"_by_AB22TGBot" {
		_, err := SendTextMessage(baseUrl, "The pack is not of this group.", chatId, messageId)
		return false, err
	}

	query, err := json.Marshal(RemoveStickerQuery{Sticker: repliedMessage.Sticker.FileId})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(baseUrl+"/deleteStickerFromSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := RemoveStickerResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	storage, _ := ioutil.ReadFile("Data/createdStickers.json")
	stickerData := CreatedSticker{}
	err = json.Unmarshal(storage, &stickerData)

	for k, chats := range stickerData.Data {
		if chats.ChatId == chatId {
			chats.Count--
			stickerData.Data[k] = chats
		}
	}

	byteData, err := json.MarshalIndent(stickerData, "", "\t")
	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile("Data/createdStickers.json", byteData, 0)
	if err != nil {
		return false, err
	}

	_, err = SendTextMessage(baseUrl, "The sticker was successfully removed.", chatId, messageId)

	return data.Result, err
}
