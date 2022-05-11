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
		{Globals.FilterButton, Globals.PhotoButton, Globals.StickerButton},
		{Globals.YoutubeButton, Globals.ExitButton},
	}

	return MessageMethods.SendButtons(Settings.MenuText, message.Chat.Id, message.MessageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, true)
}

func StickerMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Stickers\n\nCreate or add stickers to chat sticker set.\nI can automatically resize your image to make beautiful stickers.🔥\n\n"
	text += "Commands:\n-<code>/sticker &lt;Emoji (optional)&gt; &lt;reply to image&gt;</code>: Convert image to sticker or add any sticker (must be static) to sticker set (sticker set are made for every chat)."
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton, Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func FilterMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Filters\n\nMake your chat more lively with filters; The bot will reply to certain words!\n\nFilters are case insensitive; every time someone says your trigger words, Bot will reply something else! can be used to create your own commands, if desired."
	text += "\n\nCommands:\n- <code>/add &lt;trigger&gt; &lt;reply to something&gt;</code>: Every time someone says \"trigger\", the bot will reply with \"sentence\". For multiple word filters, quote the trigger.\n- <code>/filters</code>: List all chat filters.\n- <code>/revoke &lt;trigger&gt;</code>: Stop the bot from replying to \"trigger\"."
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton, Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func YoutubePlayMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Youtube-Play\n\nWant to be able to watch your favorite YouTube videos offline?\nDo you ever wish you had a way to save those great videos to watch later? Now use bot command download any YouTube video you want!\n\n"
	text += "You can download any video (size limited by telegram) by pasting a link as well as you can search for a video.\n\n"
	text += "Command:\n<code>/play &lt;query as link or search keywords&gt;</code>: search and download youtube video as mp3 or mp4.\n\n"
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton, Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func PhotoMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	text := "Photo\n\nYou can apply many beautiful filters and effects to your own photos and images. The process is automatic just reply the image and select filter to get modified image.\n\n"
	text += "Commands:\n<code>/photo &lt;reply to sticker&gt;</code>: get image file from sticker.\n<code>/pfilter &lt;reply to image&gt;</code>: apply filters to your images\n\n"
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.BackButton, Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(text, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}

func BackMenu(CallBackId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	Functions.AnswerCallbackQuery(CallBackId, "", false)

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.FilterButton, Globals.PhotoButton, Globals.StickerButton},
		{Globals.YoutubeButton, Globals.ExitButton},
	}

	return MessageMethods.EditMessageText(Settings.MenuText, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, message.Chat.Id, message.MessageId, "")
}
