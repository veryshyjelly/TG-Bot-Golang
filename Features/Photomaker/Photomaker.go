package Photomaker

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HandlePhotoMaker(option, queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if message == nil {
		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err
	}

	if photoId, ok := Globals.PhotoFilterQueue[message.MessageId]; ok {

		imagePath, err := Functions.GetFile(photoId)
		if err != nil {
			return nil, err
		}

		resImageLink := "https://violetics.pw/api/photomaker/" + option + "?apikey=" + Settings.VioKey + "&animation=static&colour=col&image=" + url.QueryEscape("https://api.telegram.org/file/bot"+Settings.ApiToken+"/"+imagePath.FilePath) + "&image2=" + url.QueryEscape("https://api.telegram.org/file/bot"+Settings.ApiToken+"/"+imagePath.FilePath)

		res, err := http.Get(resImageLink)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 200 {
			body, _ := ioutil.ReadAll(res.Body)
			errData := make(map[string]interface{})
			err := json.Unmarshal(body, &errData)
			if err != nil {
				return nil, err
			}

			_, err = Functions.AnswerCallbackQuery(queryId, fmt.Sprint(errData["message"]), true)
			return nil, err
		}

		Functions.AnswerCallbackQuery(queryId, "in progress...", true)

		m, err := Functions.SendMediaByIO(res.Body, option+".png", Functions.Document, message.Chat.Id, message.ReplyToMessage.MessageId, option, false)
		if err != nil {
			return nil, err
		}

		_, err = MessageMethods.DeleteMessage(message.Chat.Id, message.MessageId)

		return m, err

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}

}
