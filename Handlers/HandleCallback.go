package Handlers

import (
	"Telegram-Bot/Features"
	"Telegram-Bot/Features/Downloader"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func HandleCallback(query *TgTypes.CallbackQueryType) (*TgTypes.MessageType, error) {

	switch query.Data {

	case "stickerMenu":
		err := Features.StickerMenu(query.Id)
		return nil, err

	case "filterMenu":
		err := Features.FilterMenu(query.Id)
		return nil, err

	case "deleteMessage":
		_, err := MessageMethods.DeleteMessage(query.Message.Chat.Id, query.Message.MessageId)
		return nil, err

	case "ytAudio":
		return Downloader.HandleYoutubeAudio(query.Id, &query.Message)

	case "ytVideo":
		return Downloader.HandleYoutubeVideo(query.Id, &query.Message)

	default:
		_, err := Functions.AnswerCallbackQuery(query.Id, "Answering Query", true)
		return nil, err

	}
}
