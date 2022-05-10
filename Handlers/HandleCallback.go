package Handlers

import (
	"Telegram-Bot/Features"
	"Telegram-Bot/Features/Downloader"
	"Telegram-Bot/Features/PhotoFilter"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"strings"
)

func HandleCallback(query *TgTypes.CallbackQueryType) (*TgTypes.MessageType, error) {

	switch query.Data {

	case "stickerMenu":
		return Features.StickerMenu(query.Id, &query.Message)

	case "filterMenu":
		return Features.FilterMenu(query.Id, &query.Message)

	case "deleteMessage":
		_, err := MessageMethods.DeleteMessage(query.Message.Chat.Id, query.Message.MessageId)
		return nil, err

	case "ytAudio":
		return Downloader.HandleYoutubeAudio(query.Id, &query.Message)

	case "ytVideo":
		return Downloader.HandleYoutubeVideo(query.Id, &query.Message)

	case "GoBack":
		return Features.BackMenu(query.Id, &query.Message)

	default:
		x := strings.Split(query.Data, " ")
		if len(x) < 2 {
			_, err := Functions.AnswerCallbackQuery(query.Id, "Answering Query", true)
			return nil, err
		}

		switch x[0] {
		case "filterNext":
			return PhotoFilter.HandleFilterNext(x[1], query.Id, &query.Message)

		case "photoFilter":
			return PhotoFilter.HandlePhotoFilter(x[1], query.Id, &query.Message)

		}

		return nil, nil
	}
}
