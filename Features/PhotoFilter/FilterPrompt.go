package PhotoFilter

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"strconv"
)

func FilterButtons(offset int) TgTypes.InlineKeyboardMarkupType {
	initialOffset := offset
	buttons := make([][]TgTypes.InlineKeyboardButtonType, 0)
	dataList, dataSize := Globals.VioAPIData["photofilter"].Stacks, Globals.VioAPIData["photofilter"].Total

	for i := 0; i < 5 && offset < dataSize; i++ {
		row := make([]TgTypes.InlineKeyboardButtonType, 0)
		for j := 0; j < 3 && offset < dataSize; j++ {
			row = append(row, TgTypes.InlineKeyboardButtonType{Text: dataList[offset].Title, CallbackData: "photoFilter " + dataList[offset].Name})
			offset++
		}
		buttons = append(buttons, row)
	}

	row := make([]TgTypes.InlineKeyboardButtonType, 0)
	if offset > 15 {
		row = append(row, TgTypes.InlineKeyboardButtonType{Text: "Back", CallbackData: "filterNext " + strconv.Itoa(initialOffset-15)})
	}
	row = append(row, Globals.ExitButton)
	if offset < dataSize {
		row = append(row, TgTypes.InlineKeyboardButtonType{Text: "Next", CallbackData: "filterNext " + strconv.Itoa(offset)})
	}
	buttons = append(buttons, row)

	return TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}
}

func HandleFilterPrompt(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if message.ReplyToMessage == nil || message.ReplyToMessage.Document.FileId == "" {
		return MessageMethods.SendTextMessage("reply to an image (uncompressed).", message.Chat.Id, message.MessageId)
	}

	replyMarkup := FilterButtons(0)

	buttonMessage, err := MessageMethods.SendButtons("Choose Filter:", message.Chat.Id, message.MessageId, replyMarkup, true)
	if err != nil {
		return nil, err
	}

	Globals.PhotoFilterQueue[buttonMessage.MessageId] = message.ReplyToMessage.Document.FileId

	return buttonMessage, nil
}

func HandleFilterNext(offset string, queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return nil, err
	}

	replyMarkup := FilterButtons(offsetInt)

	return MessageMethods.EditMessageMarkup(message.Chat.Id, message.MessageId, replyMarkup, "")
}
