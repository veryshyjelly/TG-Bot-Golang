package PhotoFilter

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
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

		m, err := Functions.SendMediaByIO(res.Body, filter+".png", Functions.Document, message.Chat.Id, message.ReplyToMessage.MessageId, filter, false)
		if err != nil {
			return nil, err
		}

		_, err = MessageMethods.DeleteMessage(message.Chat.Id, message.MessageId)
		delete(Globals.PhotoFilterQueue, message.MessageId)

		return m, err

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}
}
