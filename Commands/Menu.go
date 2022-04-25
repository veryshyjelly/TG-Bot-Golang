package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func MenuCommand(baseUrl string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	menuText := "Hey there! My name is <b>Daemon-bot</b> \n I'm here to help you manage your groups! and many fun commmands"
	buttons := make([][]TgTypes.InlineKeyboardButtonType, 0)

	row1 := make([]TgTypes.InlineKeyboardButtonType, 0)
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Stickers", CallbackData: "stickerMenu"})
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Filter", CallbackData: "filterMenu"})
	buttons = append(buttons, row1)

	return Functions.SendButtons(baseUrl, menuText, message.Chat.Id, message.MessageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons})
}

func StickerMenu(baseUrl, CallBackId string) error {
	text := "/sticker - add image to sticker pack \n/remove - remove sticker from pack \n/resize - resize image to sticker compatible"
	_, err := Functions.AnswerCallbackQuery(baseUrl, CallBackId, text, true)
	return err
}

func FilterMenu(baseUrl, CallBackId string) error {
	text := "/add - add sticker/document/photo/animation/audio to filters \n/revoke - revoke trigger \n/filters - list all the filters"
	_, err := Functions.AnswerCallbackQuery(baseUrl, CallBackId, text, true)
	return err
}
