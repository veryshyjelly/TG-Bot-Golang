package PhotoFilter

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/StickerMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func HandlePhotoFilter(filter, queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if photoId, ok := Globals.PhotoFilterQueue[message.MessageId]; ok {
		Functions.AnswerCallbackQuery(queryId, "in progress...", true)

		imagePath, err := Functions.GetFile(photoId)
		if err != nil {
			return nil, err
		}

		resImageLink := "https://violetics.pw/api/photofilter/" + filter + "?apikey=" + Settings.VioKey + "&image=" + url.QueryEscape("https://api.telegram.org/file/bot"+Settings.ApiToken+"/"+imagePath.FilePath)

		res, err := http.Get(resImageLink)
		if err != nil {
			return nil, err
		}

		data := new(bytes.Buffer)
		_, err = io.Copy(data, res.Body)

		if err != nil {
			return nil, err
		}

		upFileLink, err := StickerMethods.UploadStickerFile(Settings.OwnerId, data)

		m, err := Functions.SendMediaByUrl(upFileLink.FileId, Functions.Document, message.Chat.Id, message.ReplyToMessage.MessageId, filter, false)
		if err != nil {
			return nil, err
		}

		MessageMethods.DeleteMessage(message.Chat.Id, message.MessageId)

		return m, nil

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}
}
