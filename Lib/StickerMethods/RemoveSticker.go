package StickerMethods

import (
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
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

func RemoveSticker(chatId int64, messageId int64, repliedMessage *TgTypes.MessageType) (bool, error) {
	if repliedMessage == nil || repliedMessage.Sticker.FileId == "" {
		_, err := MessageMethods.SendTextMessage("Reply to the sticker.", chatId, messageId)
		return false, err
	}

	if repliedMessage.Sticker.SetName != "x"+fmt.Sprint(uint64(repliedMessage.Chat.Id-int64(Settings.BotUserId)))+"_by_"+Settings.BotId[1:] {
		_, err := MessageMethods.SendTextMessage("The pack is not of this group or not made by this bot.", chatId, messageId)
		return false, err
	}

	query, err := json.Marshal(RemoveStickerQuery{Sticker: repliedMessage.Sticker.FileId})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/deleteStickerFromSet", "application/json", bytes.NewBuffer(query))
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

	_, err = MessageMethods.SendTextMessage("The sticker was successfully removed.", chatId, messageId)

	return data.Result, err
}
