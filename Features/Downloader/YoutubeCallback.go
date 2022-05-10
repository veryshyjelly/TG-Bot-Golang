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
				{Globals.ExitButton},
			}
		} else {
			buttons = [][]TgTypes.InlineKeyboardButtonType{{Globals.ExitButton}}
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
				{Globals.ExitButton},
			}
		} else {
			buttons = [][]TgTypes.InlineKeyboardButtonType{{Globals.ExitButton}}
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
