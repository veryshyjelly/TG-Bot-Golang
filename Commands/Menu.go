package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
)

func MenuCommand(baseUrl string, message *TgTypes.MessageType) {
	menuText := "Hi <b>" + message.From.FirstName + "!</b> \nMy name is <b>AB22TGBot</b>. I am a group management bot, here to help you get around and keep the order in your groups!. \nI have lots of handy features, such as flood control, a warning system, a note keeping system, and even predetermined replies on certain keywords."
	buttons := make([][]TgTypes.InlineKeyboardButtonType, 0)

	row1 := make([]TgTypes.InlineKeyboardButtonType, 0)
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})

	buttons = append(buttons, row1)
	row2 := make([]TgTypes.InlineKeyboardButtonType, 0)
	row2 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})
	row2 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})
	row2 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Sticker Menu", CallbackData: "stickerMenu"})

	buttons = append(buttons, row2)
	Functions.SendButtons(baseUrl, menuText, message.Chat.Id, message.MessageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons})
}
