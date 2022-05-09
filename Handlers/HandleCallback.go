package Handlers

import (
	"Telegram-Bot/Features"
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MediaFunctions"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func HandleCallback(baseUrl string, query TgTypes.CallbackQueryType) (*TgTypes.MessageType, error) {
	thisChatId, thisMessageId, _, _, _ := Functions.ParseMessage(&query.Message)

	switch query.Data {
	case "stickerMenu":
		err := Features.StickerMenu(baseUrl, query.Id)
		return nil, err
	case "filterMenu":
		err := Features.FilterMenu(baseUrl, query.Id)
		return nil, err
	case "deleteMessage":
		_, err := MessageMethods.DeleteMessage(baseUrl, query.Message.Chat.Id, query.Message.MessageId)
		return nil, err
	case "ytAudio":
		if link, ok := Globals.AudioLinks[thisMessageId]; ok {
			Functions.AnswerCallbackQuery(baseUrl, query.Id, "in progress...", true)
			return MediaFunctions.SendMediaByUrl(baseUrl, link, MediaFunctions.Audio, thisChatId, thisMessageId, "", true)
		} else {
			_, err := Functions.AnswerCallbackQuery(baseUrl, query.Id, "Invalid session", false)
			return nil, err
		}
	case "ytVideo":
		if link, ok := Globals.VideoLinks[thisMessageId]; ok {
			Functions.AnswerCallbackQuery(baseUrl, query.Id, "in progress...", true)
			return MediaFunctions.SendMediaByUrl(baseUrl, link, MediaFunctions.Video, thisChatId, thisMessageId, "", true)
		} else {
			_, err := Functions.AnswerCallbackQuery(baseUrl, query.Id, "Invalid session", false)
			return nil, err
		}
	default:
		_, err := Functions.AnswerCallbackQuery(baseUrl, query.Id, "Answering Query", true)
		return nil, err
	}
}
