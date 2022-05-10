package Features

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
)

func MenuCommand(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.StickerButton, Globals.FilterButton},
		{Globals.ExitButton},
	}

	return MessageMethods.SendButtons(Settings.MenuText, message.Chat.Id, message.MessageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, true)
}

func StickerMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Stickers\n\nCreate or add stickers to chat sticker set.\nI can automatically resize your image to make beautiful stickers.🔥\n\nCommands:\n-<code>/sticker &lt;reply&gt;</code>: Convert image to sticker or add any sticker (must be static) to sticker set (sticker set are made for every chat)."
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func FilterMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Filters\n\nMake your chat more lively with filters; The bot will reply to certain words!\n\nFilters are case insensitive; every time someone says your trigger words, Rose will reply something else! can be used to create your own commands, if desired.\n\nCommands:\n- <code>/add &lt;trigger&gt; &lt;reply&gt;</code>: Every time someone says \"trigger\", the bot will reply with \"sentence\". For multiple word filters, quote the trigger.\n- <code>/filters</code>: List all chat filters.\n- <code>/revoke &lt;trigger&gt;</code>: Stop the bot from replying to \"trigger\"."
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func BackMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.StickerButton, Globals.FilterButton},
		{Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(Settings.MenuText, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}
