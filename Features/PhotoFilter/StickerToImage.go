package PhotoFilter

import (
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"net/url"
)

func StickerToImage(chatId, messageId int64, repliedMessage *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if repliedMessage == nil || repliedMessage.Sticker.FileId == "" {
		return MessageMethods.SendTextMessage("reply to a sticker.", chatId, messageId)
	}
	if repliedMessage.Sticker.IsAnimated {
		return MessageMethods.SendTextMessage("animated sticker not supported.", chatId, messageId)
	}
	imagePath, err := Functions.GetFile(repliedMessage.Sticker.FileId)
	if err != nil {
		return nil, err
	}

	imageLink := "https://api.violetics.pw/api/converter/webp-to-image?apikey=" + Settings.VioKey + "&image=https://api.telegram.org/file/bot" + Settings.ApiToken + "/" + url.QueryEscape(imagePath.FilePath)
	return Functions.SendMediaByUrl(imageLink, Functions.Document, chatId, messageId, "here's the image", false)
}
