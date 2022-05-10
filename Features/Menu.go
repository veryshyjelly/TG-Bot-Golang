package Features

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func MenuCommand(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	menuText := "Hey there! My name is <b>Daemon-bot</b> \n I'm here to help you manage your groups! and many fun commmands"
	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.StickerButton, Globals.FilterButton},
		{Globals.ExitButton},
	}

	return MessageMethods.SendButtons(menuText, message.Chat.Id, message.MessageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, true)
}

func StickerMenu(CallBackId string) error {
	text := "/sticker - add image to sticker pack \n/remove - remove sticker from pack"
	_, err := Functions.AnswerCallbackQuery(CallBackId, text, true)
	return err
}

func FilterMenu(CallBackId string) error {
	text := "/add - add sticker/document/photo/animation/audio to filters \n/revoke - revoke trigger \n/filters - list all the filters"
	_, err := Functions.AnswerCallbackQuery(CallBackId, text, true)
	return err
}
