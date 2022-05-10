package StickerMaker

import (
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/StickerMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ChatStickerSet struct {
	ChatId int64  `json:"chat_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}

type CreatedSticker struct {
	Data []ChatStickerSet `json:"data"`
}

func MakeSticker(chatId, messageId int64, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if message == nil {
		return MessageMethods.SendTextMessage("Where is the image? reply to an image document (uncompressed) or a sticker (static).", chatId, messageId)
	}

	var imageLink string
	var wasSticker bool

	if message.Document.FileId != "" {
		imagePath, err := Functions.GetFile(message.Document.FileId)
		if err != nil {
			return nil, err
		}

		imageLink = "https://api.telegram.org/file/bot" + Settings.ApiToken + "/" + url.QueryEscape(imagePath.FilePath)
	} else if message.Sticker.FileId != "" && !message.Sticker.IsAnimated {
		wasSticker = true
		imagePath, err := Functions.GetFile(message.Sticker.FileId)
		if err != nil {
			return nil, err
		}

		imageLink = "https://api.violetics.pw/api/converter/webp-to-image?apikey=" + Settings.VioKey + "&image=https://api.telegram.org/file/bot" + Settings.ApiToken + "/" + url.QueryEscape(imagePath.FilePath)
	} else {
		textMessage, err := MessageMethods.SendTextMessage("Where is the image? reply to an image document (uncompressed) or a sticker (static).", message.Chat.Id, message.MessageId)
		if err != nil {
			return textMessage, err
		}
	}

	res, err := http.Get(imageLink)
	if err != nil {
		return nil, err
	}

	data := new(bytes.Buffer)
	_, err = io.Copy(data, res.Body)

	if err != nil {
		return nil, err
	}

	var resImage *bytes.Buffer
	if wasSticker {
		resImage = data
	} else {
		resImage, err = ResizeImage(data)
	}

	if err != nil {
		return nil, err
	}

	upStickerFile, err := StickerMethods.UploadStickerFile(message.From.Id, resImage)
	if err != nil {
		return nil, err
	}

	packName := "x" + fmt.Sprint(uint64(message.Chat.Id)) + "_by_" + Settings.BotId[1:]
	fmt.Println("Packname", packName)
	var stickUserId int64
	var title string
	if message.Chat.Id < 0 {
		title = message.Chat.Title + " " + Settings.BotName
		stickUserId = Settings.OwnerId
	} else {
		title = message.From.FirstName + " " + Settings.BotName + " Pack"
		stickUserId = message.From.Id
	}

	successMessage := "Sticker added to <a href=\"https://t.me/addstickers/" + packName + "\">Pack</a>"
	failMessage := "Adding of Sticker Failed"

	if set, _ := StickerMethods.GetStickerSet(packName); set != nil {
		if ok, _ := StickerMethods.AddStickerToSet(packName, upStickerFile.FileId, "😂", stickUserId); ok {
			return MessageMethods.SendTextMessage(successMessage, message.Chat.Id, message.MessageId)
		} else {
			return MessageMethods.SendTextMessage(failMessage, message.Chat.Id, message.MessageId)
		}

	} else {
		if ok, _ := StickerMethods.CreateStickerSet(packName, title, "😂", upStickerFile.FileId, stickUserId); ok {
			return MessageMethods.SendTextMessage(successMessage, message.Chat.Id, message.MessageId)
		} else {
			return MessageMethods.SendTextMessage(failMessage, message.Chat.Id, message.MessageId)
		}
	}
}
