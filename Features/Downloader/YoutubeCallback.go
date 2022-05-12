package Downloader

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func HandleYoutubeAudio(queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if link, ok := Globals.AudioLinks[message.MessageId]; ok {
		Functions.AnswerCallbackQuery(queryId, "in progress...", true)

		var buttons [][]TgTypes.InlineKeyboardButtonType

		if len(message.ReplyMarkup.InlineKeyboard[0]) == 2 {
			buttons = [][]TgTypes.InlineKeyboardButtonType{
				{Globals.VideoButton},
				{Globals.YtLinkButton, Globals.ExitButton},
			}
		} else {
			buttons = [][]TgTypes.InlineKeyboardButtonType{{Globals.YtLinkButton, Globals.ExitButton}}
		}

		_, err := MessageMethods.EditMessageMarkup(message.Chat.Id, message.MessageId,
			TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, "")

		if err != nil {
			return nil, err
		}

		return Functions.SendMediaByUrl(link, Functions.Audio, message.Chat.Id, message.MessageId, "", true)

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}
}

func HandleYoutubeVideo(queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if link, ok := Globals.VideoLinks[message.MessageId]; ok {
		Functions.AnswerCallbackQuery(queryId, "in progress...", true)

		var buttons [][]TgTypes.InlineKeyboardButtonType

		if len(message.ReplyMarkup.InlineKeyboard[0]) == 2 {
			buttons = [][]TgTypes.InlineKeyboardButtonType{
				{Globals.AudioButton},
				{Globals.YtLinkButton, Globals.ExitButton},
			}
		} else {
			buttons = [][]TgTypes.InlineKeyboardButtonType{{Globals.YtLinkButton, Globals.ExitButton}}
		}

		_, err := MessageMethods.EditMessageMarkup(message.Chat.Id, message.MessageId,
			TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, "")

		if err != nil {
			return nil, err
		}

		return Functions.SendMediaByUrl(link, Functions.Video, message.Chat.Id, message.MessageId, "", true)

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}
}

func HandleYoutubeLinks(queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if _, ok := Globals.AudioLinks[message.MessageId]; ok {
		Functions.AnswerCallbackQuery(queryId, "", true)

		myAudioButton, myVideoButton := Globals.AudioButton, Globals.VideoButton
		myAudioButton.Url, myVideoButton.Url = Globals.AudioLinks[message.MessageId], Globals.VideoLinks[message.MessageId]

		buttons := [][]TgTypes.InlineKeyboardButtonType{
			{myAudioButton, myVideoButton},
			{Globals.YtLinkButton, Globals.ExitButton},
		}

		return MessageMethods.EditMessageMarkup(message.Chat.Id, message.MessageId,
			TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, "")

	} else {

		_, err := Functions.AnswerCallbackQuery(queryId, "Invalid session", false)
		return nil, err

	}
}
